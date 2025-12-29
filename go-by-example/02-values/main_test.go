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

	expected := "golang\n1+1 = 2\n7.0/3.0 = 2.3333333333333335\nfalse\ntrue\nfalse\n"

	got := string(output)
	if got != expected {
		t.Errorf("Unexpected output.\nExpected: %q\nGot: %q", expected, got)
	}
}
