package clog_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func w() (res int, err error) {
	res, _, err = term.GetSize(0)
	return
}

func TestClog(t *testing.T) {
	// Define the strings
	left := "Left"
	center := "Center"
	right := "RRightRightRightight"

	if term.IsTerminal(0) {
		println("in a term")
	} else {
		println("not in a term")
	}

	// Define the total width of the terminal or area where you want to align the strings
	totalWidth, err := w()
	if err != nil {
		totalWidth = 110
	}

	// Calculate the spacing
	leftWidth := lipgloss.Width(left)
	centerWidth := lipgloss.Width(center)
	rightWidth := lipgloss.Width(right)

	// Calculate the space between left and center, and center and right
	// We subtract the widths of the strings and divide the remaining space
	totalSpace := totalWidth - (leftWidth + centerWidth + rightWidth)
	spaceLeftCenter := totalSpace / 2 // Dividing by 3 to distribute space evenly between and around the strings
	spaceCenterRight := totalSpace / 2

	// Construct the final string with calculated spacing
	finalString := fmt.Sprintf("%s%s%s%s%s",
		left,
		strings.Repeat(" ", spaceLeftCenter),
		center,
		strings.Repeat(" ", spaceCenterRight),
		right,
	)

	// Print the final string
	fmt.Println(finalString)
}
