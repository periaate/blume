package blumefmt

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/periaate/blume/types/str"
	"github.com/periaate/blume/yap"
)

// This code has been generated with ChatGPT o1 and fixed by hand.

// regionToInline holds information about a region in the source
// that we want to rewrite: the opening brace, the single statement,
// and the closing brace, all on consecutive lines/positions.
type regionToInline struct {
	openBracePos  token.Pos
	stmtStartPos  token.Pos
	stmtEndPos    token.Pos
	closeBracePos token.Pos
}

// inlineVisitor implements ast.Visitor. It inspects if/switch nodes
// to detect single-statement blocks we want to inline.
type inlineVisitor struct {
	fset     *token.FileSet
	regions  []regionToInline
	srcBytes []byte // the full source code as bytes
	out      bool
}

func (iv *inlineVisitor) Visit(node ast.Node) ast.Visitor {
	switch iv.out {
	case true:
		return iv.visitOut(node)
	default:
		return iv.visit(node)
	}
}

// Visit is called for each node in the AST.
func (iv *inlineVisitor) visitOut(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		if n.Body != nil && len(n.Body.List) == 1 {
			// Grab positions
			openPos := n.Body.Lbrace
			closePos := n.Body.Rbrace
			onlyStmt := n.Body.List[0]

			iv.addRegionIfValid(openPos, onlyStmt, closePos)
		}
	}
	return iv
}

// Visit is called for each node in the AST.
func (iv *inlineVisitor) visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.IfStmt:
		if n.Body != nil && len(n.Body.List) == 1 && n.Else == nil {
			// Grab positions
			openPos := n.Body.Lbrace
			closePos := n.Body.Rbrace
			onlyStmt := n.Body.List[0]

			iv.addRegionIfValid(openPos, onlyStmt, closePos)
		}

	case *ast.CaseClause:
		if len(n.Body) != 1 {
			return iv
		}
		openPos, closePos := iv.findCaseBraces(n)
		// If we found braces in the same line region, store the region
		if openPos.IsValid() && closePos.IsValid() {
			onlyStmt := n.Body[0]
			iv.addRegionIfValid(openPos, onlyStmt, closePos)
		}

	case *ast.CommClause:
		if len(n.Body) == 1 {
			openPos, closePos := iv.findCommBraces(n)
			if openPos.IsValid() && closePos.IsValid() {
				onlyStmt := n.Body[0]
				iv.addRegionIfValid(openPos, onlyStmt, closePos)
			}
		}
	}

	return iv
}

// addRegionIfValid adds a region if we successfully map all
// positions to actual offsets in the original file.
func (iv *inlineVisitor) addRegionIfValid(open token.Pos, stmt ast.Stmt, close token.Pos) {
	stmtStart := stmt.Pos()
	stmtEnd := stmt.End()

	if !open.IsValid() || !close.IsValid() || !stmtStart.IsValid() || !stmtEnd.IsValid() {
		return
	}

	// In principle, we could do more checks (e.g., single-line statement).
	// For simplicity, let's just record the offsets.
	iv.regions = append(iv.regions, regionToInline{
		openBracePos:  open,
		stmtStartPos:  stmtStart,
		stmtEndPos:    stmtEnd,
		closeBracePos: close,
	})
}

// findCaseBraces tries to approximate the positions of the opening/closing brace
// around a case clause. In the Go AST, the "case" clause doesn't explicitly
// store braces, because the braces belong to the switch body. We do a best-effort
// textual search here or simply skip if not feasible.
func (iv *inlineVisitor) findCaseBraces(cc *ast.CaseClause) (token.Pos, token.Pos) {
	// By default in the AST, a CaseClause does not hold separate braces
	// the way an IfStmt's block does. Typically, the block is simply part
	// of the SwitchStmt body. We might do a more advanced approach or skip.
	//
	// However, to produce the exact "case X: stmt" style, we can just pretend
	// we have "virtual braces" hugging each statement for the case. Then
	// the rewrite tries to inline it on the same line. Let's do that:

	// We'll just mark them invalid if we can't do it.
	// We'll define that "openPos" is where the case line ends (the ':'),
	// and "closePos" is the end of the single statement, for the textual rewrite.
	if len(cc.Body) != 1 {
		return token.NoPos, token.NoPos
	}
	// 'case' or 'default' ends at cc.Colon
	openPos := cc.Colon
	closePos := cc.Body[0].End()
	yap.Debug("case body", "open", openPos, "close", closePos, "open valid", openPos.IsValid(), "close valid", closePos.IsValid())
	return openPos, closePos
}

// findCommBraces is analogous to findCaseBraces but for CommClauses in a select.
func (iv *inlineVisitor) findCommBraces(comm *ast.CommClause) (token.Pos, token.Pos) {
	if len(comm.Body) != 1 {
		return token.NoPos, token.NoPos
	}
	openPos := comm.Colon
	closePos := comm.Body[0].End()
	yap.Debug("comm body", "open", openPos, "close", closePos, "open valid", openPos.IsValid(), "close valid", closePos.IsValid())
	return openPos, closePos
}

// Rewrite takes the original source and all the discovered "regions to inline"
// and produces a new source with the single-line blocks inlined.
func (iv *inlineVisitor) Rewrite() []byte {
	// We'll do a simple approach: build a new buffer by scanning from
	// start to end. Whenever we meet a region, we *skip* the braces
	// and the newline, and inline everything onto one line.

	type offsetRegion struct {
		openBraceOffset  int
		stmtStartOffset  int
		stmtEndOffset    int
		closeBraceOffset int
	}

	// Convert token.Pos to absolute offsets.
	regs := make([]offsetRegion, 0, len(iv.regions))
	for _, r := range iv.regions {
		o1 := iv.fset.Position(r.openBracePos).Offset
		o2 := iv.fset.Position(r.stmtStartPos).Offset
		o3 := iv.fset.Position(r.stmtEndPos).Offset
		o4 := iv.fset.Position(r.closeBracePos).Offset
		regs = append(regs, offsetRegion{o1, o2, o3, o4})
	}

	// Sort them by openBraceOffset so we rewrite in ascending order.
	// (This is important if we do multiple inlines in one file.)
	// A simple bubble or insertion sort is enough for a small set.
	for i := 0; i < len(regs); i++ {
		for j := i + 1; j < len(regs); j++ {
			if regs[j].openBraceOffset < regs[i].openBraceOffset {
				regs[i], regs[j] = regs[j], regs[i]
			}
		}
	}

	var out bytes.Buffer
	src := iv.srcBytes

	cursor := 0
	// var offset int
	for _, region := range regs {
		if region.openBraceOffset < cursor {
			yap.Debug("out of bounds", region.openBraceOffset, cursor)
			// Already past it, skip (overlapping or nested inlines).
			continue
		}
		// Write everything up to the opening brace.
		wr := src[cursor:region.openBraceOffset]
		yap.Debug("writing region precontent", string(wr))
		out.Write(wr)
		// Skip the '{' (or ':') and any whitespace/newline until the statement start:
		// We'll be naive and just jump straight to region.stmtStartOffset.
		var wrong bool
		cnt := src[region.stmtStartOffset:region.stmtEndOffset]
		if len(str.Split(string(cnt), false, "\n")) != 1 {
			wrong = true
		}

		// Insert " { " or ": " if it is a switch-case. We do a quick peek:
		ch := src[region.openBraceOffset]
		switch wrong {
		case true:
			switch ch {
			case '{':
				out.WriteString("{\n\t")
			case ':':
				out.WriteString(":\n\t")
			default:
				out.WriteByte(ch)
			}
		default:
			switch ch {
			case '{':
				out.WriteString("{ ")
			case ':':
				out.WriteString(": ")
			default:
				out.WriteByte(ch)
			}
		}
		// Now write from the statement start to the statement end.
		out.Write(cnt)

		// Next, skip the closing brace offset (region.closeBraceOffset)
		// and any whitespace/newlines right after it.
		cursor = region.closeBraceOffset
		// We add " }" if it was originally a '{' block.
		// If it was ':' (case clause), we omit curly braces, just inline the statement.
		if ch == '{' && !wrong {
			out.WriteString(" ")
		}
		if iv.out && wrong {
			out.WriteString("\n")
		}
	}
	// Finally, write the remainder of the file.
	if cursor < len(src) {
		out.Write(src[cursor:])
	}

	return out.Bytes()
}

func basic(input []byte) (res []byte, err error) {
	// Parse the source into an AST.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "stdin.go", input, parser.AllErrors)
	if err != nil {
		return
	}

	// Walk the AST and find single-statement if/case blocks.
	iv := &inlineVisitor{
		fset:     fset,
		srcBytes: []byte(input),
	}
	ast.Walk(iv, file)

	// Perform the textual rewrite.
	res = iv.Rewrite()

	s := string(res)
	var sar []string
	var lastMatch int
	// var lastCaseMatch int
	offset := 1
	matcher_if_fn := str.MatchRegex(`.*(if|func).*{.*}$`)
	// matcher_case := MatchRegex(`.*(default:|case).*$`)
	for i, line := range str.Split(s, true, "\n") {
		switch {
		case matcher_if_fn(line):
			vi := i - offset
			// fmt.Println("foud match", vi, lastMatch, sar[lastMatch], sar[len(sar)-1] == "\n", lastMatch == vi-2)
			if lastMatch == vi-2 && sar[len(sar)-1] == "\n" && lastMatch != 0 {
				sar[len(sar)-1] = line
				lastMatch = vi
				offset += 1
				continue
			}
			lastMatch = i
		}

		sar = append(sar, line)
	}

	s = strings.Join(sar, "")
	res = []byte(s)
	return
}

func basicOut(input []byte) (res []byte, err error) {
	// Parse the source into an AST.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "stdin.go", input, parser.AllErrors)
	if err != nil {
		return
	}

	// Walk the AST and find single-statement if/case blocks.
	iv := &inlineVisitor{
		fset:     fset,
		srcBytes: []byte(input),
		out:      true,
	}
	ast.Walk(iv, file)

	// Perform the textual rewrite.
	res = iv.Rewrite()

	s := string(res)
	var sar []string
	var lastMatch int
	// var lastCaseMatch int
	offset := 1
	matcher_if_fn := str.MatchRegex(`.*(if|func).*{.*}$`)
	// matcher_case := MatchRegex(`.*(default:|case).*$`)
	for i, line := range str.Split(s, true, "\n") {
		switch {
		case matcher_if_fn(line):
			vi := i - offset
			// fmt.Println("foud match", vi, lastMatch, sar[lastMatch], sar[len(sar)-1] == "\n", lastMatch == vi-2)
			if lastMatch == vi-2 && sar[len(sar)-1] == "\n" && lastMatch != 0 {
				sar[len(sar)-1] = line
				lastMatch = vi
				offset += 1
				continue
			}
			lastMatch = i
		}

		sar = append(sar, line)
	}

	s = strings.Join(sar, "")
	res = []byte(s)
	return
}

func Fmt(input []byte) (res []byte, err error) {
	output, err := basic(input)
	if err != nil {
		return
	}
	res, err = basicOut(output)
	return
}
