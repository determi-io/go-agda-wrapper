package go_agda_wrapper

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	// create temporary directory
	dname, err := os.MkdirTemp("", "go_agda_wrapper_compilation_")
	if err != nil {
		t.Errorf("Could not create temp dir: %s", err)
	}

	// run compilation
	_, err = CompileAndRun("myfile.agda", dname)
	if err != nil {
		t.Errorf("Compiling and running failed, error: %s", err)
	}
}
