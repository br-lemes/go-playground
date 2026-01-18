package main

import (
	"os/exec"
	"testing"
)

func TestMainOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute program: %v", err)
	}

	expected := "emp: [0 0 0 0 0]\n" +
		"set: [0 0 0 0 100]\n" +
		"get: 100\n" +
		"len: 5\n" +
		"dcl: [1 2 3 4 5]\n" +
		"dcl: [1 2 3 4 5]\n" +
		"idx: [100 0 0 400 500]\n" +
		"2d:  [[0 1 2] [1 2 3]]\n" +
		"2d:  [[1 2 3] [1 2 3]]\n"

	got := string(output)
	if got != expected {
		t.Errorf("Unexpected output.\nExpected: %q\nGot: %q", expected, got)
	}
}
