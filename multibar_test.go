package multibar

import (
	"testing"

	"github.com/sethgrid/curse"
)

func TestNewLineDetection_EmptyPrintln(t *testing.T) {
	bc := &BarContainer{}
	bc.Println()
	expected := 1
	ResetLines(expected)

	if bc.totalNewlines != expected {
		t.Errorf("empty Println - got %d, want %d newline", bc.totalNewlines, expected)
	}
}

func TestNewLineDetection_Println(t *testing.T) {
	bc := &BarContainer{}
	bc.Println("Look, ma!\n I printed somethin'\n")
	expected := 3
	ResetLines(expected)

	if bc.totalNewlines != expected {
		t.Errorf("Println - got %d, want %d newline", bc.totalNewlines, expected)
	}
}

func TestNewLineDetection_Printf(t *testing.T) {
	bc := &BarContainer{}
	bc.Printf("Look, ma!\n I printed somethin'\n %s %s", "and\n", "\nI can keep printing")
	expected := 4
	ResetLines(expected)

	if bc.totalNewlines != expected {
		t.Errorf("Printf - got %d, want %d newline", bc.totalNewlines, expected)
	}
}

func TestNewLineDetection_Print(t *testing.T) {
	bc := &BarContainer{}
	bc.Print("Look, ma!\n I printed somethin'\n", "and\n", "\nI can keep printing")
	expected := 4
	ResetLines(expected)

	if bc.totalNewlines != expected {
		t.Errorf("Print - got %d, want %d newline", bc.totalNewlines, expected)
	}
}

func TestNewLineDetection_ManyPrints(t *testing.T) {
	bc := &BarContainer{}
	bc.Print("Look, ma!\n I printed somethin'\n", "and\n", "\nI can keep printing")
	bc.Println("and more!")
	expected := 5
	ResetLines(expected)

	if bc.totalNewlines != expected {
		t.Errorf("Print - got %d, want %d newline", bc.totalNewlines, expected)
	}
}
