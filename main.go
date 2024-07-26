package go_agda_wrapper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello flake")
}

func createAndWrite(filename string, content string) (*os.File, error) {
	// create file
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	// write content to file
	_, err = file.WriteString(content)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
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

func CompileAndRun(sourceRoot string, sourceFile string, compileDir string, topLevelValue string) (string, error) {
	// we check that all prereqs exist
	err := checkPrerequisites()
	if err != nil {
		return "", err
	}

	// call agda
	{
		cmd := exec.Command("agda", "--js", "--compile-dir", compileDir, sourceFile)
		cmd.Dir = sourceRoot
		output, err := cmd.Output()
		if err != nil {
			return "", errors.New(fmt.Sprintf("Running agda not successful\noutput:\n%s, error: %s", output, err))
		}
	}

	// Compute module name
	modulename := strings.ReplaceAll(fileNameWithoutExtension(sourceFile), "/", ".")

	// write top level file
	toplevelFile := filepath.Join(compileDir, "main.js")
	toplevelContent := fmt.Sprintf("var main = require('%s')\nprocess.stdout.write(main['%s'])", "jAgda."+modulename, topLevelValue)
	createAndWrite(toplevelFile, toplevelContent)

	// call node
	cmd := exec.Command("node", "main.js")
	cmd.Env = append(cmd.Env, "NODE_PATH=.")
	cmd.Dir = compileDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Running node not successful\noutput:\n%s, error: %s", output, err))
	}

	// done
	return string(output[:]), nil
}
