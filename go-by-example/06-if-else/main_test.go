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

	expected := "7 is odd\n8 is divisible by 4\neither 8 or 7 are even\n9 has 1 digit\n"

	got := string(output)
	if got != expected {
		t.Errorf("Unexpected output.\nExpected: %q\nGot: %q", expected, got)
	}
}
