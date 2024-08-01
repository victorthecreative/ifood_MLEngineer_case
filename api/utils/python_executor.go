package utils

import (
	"fmt"
	"os/exec"
	"strconv"
)

func CallPythonScript(prompt string, maxLen int) (string, error) {
	cmd := exec.Command("python3", "text_generator.py", prompt, strconv.Itoa(maxLen))
	cmd.Dir = "/app"
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error calling Python script: %v, output: %s", err, string(output))
	}
	return string(output), nil
}
