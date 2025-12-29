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

	expected := "initial\n1 2\ntrue\n0\napple\n"

	got := string(output)
	if got != expected {
		t.Errorf("Unexpected output.\nExpected: %q\nGot: %q", expected, got)
	}
}
