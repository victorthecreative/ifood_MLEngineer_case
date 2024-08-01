package runner

import (
	"bufio"
	"fmt"
	"os/exec"

	"ifood_case/cli/config"

	"github.com/schollz/progressbar/v3"
)

func RunPythonScript(cfg *config.Config) error {
	fmt.Println("Executando o script Python...")
	cmd := exec.Command("python3", "./fine_tuned_model/src/main.py")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("falha ao obter o pipe de saída padrão: %v", err)
	}

	bar := progressbar.DefaultBytes(
		-1,
		"executando",
	)

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("falha ao iniciar o script Python: %v", err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		bar.Add(len(scanner.Bytes()))
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("o script Python terminou com erro: %v", err)
	}

	return nil
}
