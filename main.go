package go_agda_wrapper

import (
	"errors"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello flake")
}

func checkPrerequisites() error {
	// check if agda exists
	{
		cmd := exec.Command("agda", "--version")
		_, err := cmd.Output()
		if err != nil {
			return err
		}
	}

	// check if node exists
	{
		cmd := exec.Command("node", "--version")
		_, err := cmd.Output()
		if err != nil {
			return err
		}
	}

	return nil
}

func CompileAndRun(file string, compileDir string) (string, error) {
	// we check that all prereqs exist
	err := checkPrerequisites()
	if err != nil {
		return "", err
	}

	// call agda
	{
		output, err := exec.Command("agda", "--js", "--compile-dir", compileDir, file).Output()
		if err != nil {
			return "", errors.New(fmt.Sprintf("Running agda not successful\noutput:\n%s, error: %s", output, err))
		}
	}

	// We compile
	return "", errors.New("Could not load file. Not implemented.")
}
