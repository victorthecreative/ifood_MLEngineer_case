package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ModelID                 string  `yaml:"model_id"`
	TrainFilePathRaw        string  `yaml:"train_file_path_raw"`
	TrainFilePathCurated    string  `yaml:"train_file_path_curated"`
	PerDeviceTrainBatchSize int     `yaml:"per_device_train_batch_size"`
	OutputDir               string  `yaml:"output_dir"`
	OverwriteOutputDir      bool    `yaml:"overwrite_output_dir"`
	NumTrainEpochs          float64 `yaml:"num_train_epochs"`
	SaveSteps               int     `yaml:"save_steps"`
}

func LoadConfig(configPath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("falha ao ler o arquivo de configuração: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("falha ao analisar a configuração: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	getInput := func(prompt string, defaultValue string) string {
		fmt.Printf("Enter %s (default: %s): ", prompt, defaultValue)
		input, _ := reader.ReadString('\n')
		if input != "\n" {
			return input[:len(input)-1]
		}
		return defaultValue
	}

	getIntInput := func(prompt string, defaultValue int) int {
		for {
			fmt.Printf("Enter %s (default: %d): ", prompt, defaultValue)
			var input int
			_, err := fmt.Scanf("%d\n", &input)
			if err != nil {
				fmt.Println("Entrada inválida. Por favor, insira um número inteiro.")
				reader.ReadString('\n')
				continue
			}
			return input
		}
	}

	getBoolInput := func(prompt string, defaultValue bool) bool {
		for {
			fmt.Printf("Enter %s (default: %t): ", prompt, defaultValue)
			var input string
			_, err := fmt.Scanf("%s\n", &input)
			if err == nil {
				if input == "true" || input == "True" || input == "1" {
					return true
				} else if input == "false" || input == "False" || input == "0" {
					return false
				}
			}
			fmt.Println("Entrada inválida. Por favor, insira 'true' ou 'false'.")
			reader.ReadString('\n')
		}
	}

	getFloatInput := func(prompt string, defaultValue float64) float64 {
		for {
			fmt.Printf("Enter %s (default: %f): ", prompt, defaultValue)
			var input float64
			_, err := fmt.Scanf("%f\n", &input)
			if err != nil {
				fmt.Println("Entrada inválida. Por favor, insira um número decimal.")
				reader.ReadString('\n') // Limpa o buffer de entrada
				continue
			}
			return input
		}
	}

	config.ModelID = getInput("model ID", config.ModelID)
	config.TrainFilePathRaw = getInput("train file path (raw)", config.TrainFilePathRaw)
	config.TrainFilePathCurated = getInput("train file path (curated)", config.TrainFilePathCurated)
	config.PerDeviceTrainBatchSize = getIntInput("per device train batch size", config.PerDeviceTrainBatchSize)
	config.OutputDir = getInput("output directory", config.OutputDir)
	config.OverwriteOutputDir = getBoolInput("overwrite output directory", config.OverwriteOutputDir)
	config.NumTrainEpochs = getFloatInput("number of training epochs", config.NumTrainEpochs)
	config.SaveSteps = getIntInput("save steps", config.SaveSteps)

	// Escreve a configuração atualizada de volta para o arquivo
	updatedConfig, err := yaml.Marshal(&config)
	if err != nil {
		return nil, fmt.Errorf("falha ao serializar a configuração: %v", err)
	}

	err = ioutil.WriteFile(configPath, updatedConfig, 0644)
	if err != nil {
		return nil, fmt.Errorf("falha ao escrever o arquivo de configuração: %v", err)
	}

	return &config, nil
}
