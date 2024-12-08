package main_test

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func TestStaticCheck(t *testing.T) {
	rungo(t, "run", "honnef.co/go/tools/cmd/staticcheck@latest", "./...")
}

func TestUnparam(t *testing.T) {
	rungo(t, "run", "mvdan.cc/unparam@latest", "./...")
}

func TestVet(t *testing.T) {
	rungo(t, "vet", "-all", "./...")
}

func TestGoModTidy(t *testing.T) {
	rungo(t, "mod", "tidy", "-diff")
}

func TestGovulncheck(t *testing.T) {
	rungo(t, "run", "golang.org/x/vuln/cmd/govulncheck@latest", "./...")
}

func TestGofmt(t *testing.T) {
	cmd := exec.Command("gofmt", "-l", ".")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to run gofmt: %v", err)
	}

	// Get the output and check if there are any non-empty lines
	unformattedFiles := strings.TrimSpace(out.String())
	if unformattedFiles != "" {
		t.Errorf("The following files are not properly formatted:\n%s", unformattedFiles)
	}
}

func rungo(t *testing.T, args ...string) {
	t.Helper()

	cmd := exec.Command("go", args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		if ee := (*exec.ExitError)(nil); errors.As(err, &ee) && len(ee.Stderr) > 0 {
			t.Fatalf("%v: %v\n%s", cmd, err, ee.Stderr)
		}
		t.Fatalf("%v: %v\n%s", cmd, err, output)
	}
}
