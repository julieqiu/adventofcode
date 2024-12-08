package main_test

import (
	"errors"
	"os/exec"
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
