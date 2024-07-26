package go_agda_wrapper

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

const testmodule = `
module Test where

open import Agda.Builtin.String

output : String
output = "hello"
`

func TestAll(t *testing.T) {
	// create temporary directory (source)
	sourcedir, err := os.MkdirTemp("", "go_agda_wrapper_source_")
	if err != nil {
		t.Errorf("Could not create temp dir: %s", err)
		return
	}

	// create temporary directory (target)
	targetdir, err := os.MkdirTemp("", "go_agda_wrapper_target_")
	if err != nil {
		t.Errorf("Could not create temp dir: %s", err)
		return
	}

	// create source file
	_, err = createAndWrite(filepath.Join(sourcedir, "Test.agda"), testmodule)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// run compilation
	result, err := CompileAndRun(sourcedir, "Test.agda", targetdir, "output")
	if err != nil {
		t.Errorf("Compiling and running failed, source: %s, target: %s, error: %s", sourcedir, targetdir, err)
		return
	}

	if result != "hello" {
		t.Errorf("Result of execution was %s", strconv.Quote(result))
		return
	}
}
