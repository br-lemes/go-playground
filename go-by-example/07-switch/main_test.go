package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute program: %v", err)
	}

	prefix := "Write 2 as two\nIt's"
	suffix := " noon\nI'm a bool\nI'm an int\nDon't know type string\n"

	got := string(output)
	if !strings.HasPrefix(got, prefix) {
		t.Errorf("Unexpected output.\nExpected to start with: %q\nGot: %q", prefix, got)
	}
	if !strings.HasSuffix(got, suffix) {
		t.Errorf("Unexpected output.\nExpected to end with: %q\nGot: %q", suffix, got)
	}
}
