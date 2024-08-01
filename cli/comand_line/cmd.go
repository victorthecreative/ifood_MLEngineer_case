package main

import (
	"fmt"
	"log"

	"ifood_case/cli/config"
	"ifood_case/cli/runner"
)

func main() {

	cfg, err := config.LoadConfig("./fine_tuned_model/src/config.yml")
	if err != nil {
		log.Fatalf("Falha ao carregar a configuração: %v", err)
	}

	err = runner.RunPythonScript(cfg)
	if err != nil {
		log.Fatalf("O script Python terminou com erro: %v", err)
	}

	fmt.Println("Script Python finalizado com sucesso.")
}
