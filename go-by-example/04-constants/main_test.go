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

	expected := "constant\n6e+11\n600000000000\n-0.28470407323754404\n"

	got := string(output)
	if got != expected {
		t.Errorf("Unexpected output.\nExpected: %q\nGot: %q", expected, got)
	}
}
