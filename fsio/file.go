package fsio

import (
	"bufio"
	"fmt"
	"io"
	"os"

	. "github.com/periaate/blume"
)

// Constants for common file permissions
const (
	// User permissions
	UserRead      = 0o400 // Owner read permission
	UserWrite     = 0o200 // Owner write permission
	UserExecute   = 0o100 // Owner execute permission
	UserReadWrite = UserRead | UserWrite

	// Group permissions
	GroupRead      = 0o040 // Group read permission
	GroupWrite     = 0o020 // Group write permission
	GroupExecute   = 0o010 // Group execute permission
	GroupReadWrite = GroupRead | GroupWrite

	// Other permissions
	OtherRead      = 0o004 // Others read permission
	OtherWrite     = 0o002 // Others write permission
	OtherExecute   = 0o001 // Others execute permission
	OtherReadWrite = OtherRead | OtherWrite

	// Combined permissions
	AllRead      = UserRead | GroupRead | OtherRead          // Read permission for all
	AllWrite     = UserWrite | GroupWrite | OtherWrite       // Write permission for all
	AllExecute   = UserExecute | GroupExecute | OtherExecute // Execute permission for all
	AllReadWrite = AllRead | AllWrite

	// Common file modes
	ReadOnly      = UserRead | GroupRead | OtherRead
	ReadWrite     = UserReadWrite | GroupRead | OtherRead
	ReadWriteExec = ReadWrite | UserExecute | GroupExecute | OtherExecute

	// Directory modes
	DirReadOnly      = ReadOnly | os.ModeDir
	DirReadWrite     = ReadWrite | os.ModeDir
	DirReadWriteExec = ReadWriteExec | os.ModeDir
)

func Copy[DST, SRC ~string](dst DST, src SRC, force bool) error {
	f, err := os.Open(string(src))
	if err != nil {
		return fmt.Errorf("failed to copy from [%s] to [%s] with error: [%s]", src, dst, err)
	}
	defer f.Close()

	switch force {
		case true: err = WriteAll(string(dst), f)
		case false: err = WriteNew(string(dst), f)
	}
	return err
}

// func Read[S ~string](fp S) Result[[]byte] { return AsRes(os.ReadFile(string(fp))) }

// WriteAll writes the contents of the reader to the file, overwriting existing files.
func WriteAll(f string, r io.Reader) (err error) {
	file, err := os.Create(f)
	if err != nil { return err }
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

// WriteNew writes the contents of the reader to a new file, will not overwrite existing files.
func WriteNew(f string, r io.Reader) (err error) {
	if Exists(f) { return fmt.Errorf("file %s already exists", f) }
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil { return err }
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

// AppendTo appends the contents of the reader to the file.
func AppendTo(f string, r io.Reader) (err error) {
	// Open the file in append mode, create if not exists
	file, err := os.OpenFile(f, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil { return err }
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func Open(f string) (rc io.ReadCloser, err error) {
	file, err := os.Open(f)
	if err != nil { return }
	rc = file
	return
}

func Remove(f string) (err error) { return os.Remove(f) }

func ReadTo(f string, r io.Reader) (n int64, err error) {
	file, err := os.Create(f)
	if err != nil { return }
	defer file.Close()

	n, err = io.Copy(file, r)
	return
}

// ReadPipe reads from stdin and returns a slice of lines.
func ReadPipe() (res []string) {
	if HasPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}
	}
	return
}

// HasPipe evaluates whether stdin is being piped in to.
func HasPipe() bool {
	a, err := os.Stdin.Stat()
	if err != nil { return false }
	return (a.Mode() & os.ModeCharDevice) == 0
}

// HasPipe evaluates whether stdin is being piped in to.
func HasOutPipe() bool {
	a, err := os.Stdout.Stat()
	if err != nil { return false }
	return (a.Mode() & os.ModeCharDevice) == 0
}

func Getenv[S ~string](key S) Option[S] {
	val, ok := os.LookupEnv(string(key))
	if !ok { return None[S]() }
	return Some(S(val))
}

// Args returns the command-line arguments without the program name, and including any piped inputs.
func Args[S ~string](opts ...func([]string) bool) Option[Array[S]] {
	var iargs []string
	if len(os.Args) >= 1 { iargs = os.Args[1:] }
	return args[S](append(iargs, ReadPipe()...), opts...)
}

func args[S ~string](arr []string, opts ...func([]string) bool) Option[Array[S]] {
	if len(opts) == 0 { return Some(ToArray(Map[string, S](StoS)(arr))) }
	if !PredAnd(opts...)(arr) { return None[Array[S]]() }
	return Some(ToArray(Map[string, S](StoS)(arr)))
}

// Args returns the command-line arguments without the program name, and including any piped inputs.
func IArgs[S ~string](opts ...func([]string) bool) Option[Array[S]] {
	if len(os.Args) < 1 { return None[Array[S]]() }
	return args[S](os.Args[1:], opts...)
}

// PArgs (piped arguments), returns the piped input as newline separated Array of S typed strings.
func PArgs[S ~string](opts ...func([]string) bool) Option[Array[S]] {
	return args[S](ReadPipe(), opts...)
}

// Args returns the command-line arguments without the program name, and including any piped inputs.
func SepArgs() (res [2][]string) { return [2][]string{os.Args[1:], ReadPipe()} }
