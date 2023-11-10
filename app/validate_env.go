package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {

	testEnvs := getTestEnv()
	iacEnv := getIacEnv(os.Args[1], os.Args[2])
	secrets := getSecret(os.Args[1], os.Args[2])

	error := compare(testEnvs, iacEnv, secrets)
	if error {
		panic("Revise as variaveis de configuração no projeto iac-cluster1")
	}
}

func compare(testEnvs, iacEnv, secrets []string) bool {
	var error bool
	for _, key := range testEnvs {

		if contains(secrets, key) || contains(iacEnv, key) {
			continue
		}
		_, message := fmt.Printf("Chave: %s - Faltando no arquivo env", key)
		fmt.Println(message)
		error = true
	}
	return error
}

func getIacEnv(software, namespace string) []string {
	var keys []string
	arquivo, scanner := readFileToScanner("iac-cluster1/cluster2/argocd-apps-templates/" + namespace + "/" + software + "/env")

	// Garante que o arquivo seja fechado no final do programa
	defer arquivo.Close()

	// Lê o arquivo linha a linha
	for scanner.Scan() {
		linha := scanner.Text()
		if strings.Contains(linha, "=") {
			//obtém o valor antes do "="
			key := strings.TrimSpace(strings.SplitN(linha, "=", 2)[0])
			keys = append(keys, key)
		}
	}
	return keys
}

func getSecret(software, namespace string) []string {
	var keys []string
	secret := readFileToYaml("iac-cluster1/cluster2/argocd-apps-templates/" + namespace + "/" + software + "/SealedSecret.ms-platform-go.yaml")

	// Imprime as chaves do arquivo YAML
	for key := range secret.Spec.EncryptedData {
		keys = append(keys, key)
	}
	return keys
}

func getTestEnv() []string {
	var keys []string
	// Abre o arquivo para leitura
	arquivo, scanner := readFileToScanner("scripts/test.env")

	// Garante que o arquivo seja fechado no final do programa
	defer arquivo.Close()
	// Lê o arquivo linha a linha
	for scanner.Scan() {
		linha := scanner.Text()
		if strings.HasPrefix(linha, "export") && strings.Contains(linha, "=") {
			// Remove "export " e obtém o valor antes do "="
			lineKey := strings.TrimSpace(strings.SplitN(linha, "=", 2)[0])
			key := strings.TrimSpace(strings.SplitN(lineKey, "export", 2)[1])
			keys = append(keys, key)
		}
	}

	// Verifica por erros durante o scan. O scanner devolve false se encontrar um erro
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
	return keys
}

func readFileToYaml(path string) Secret {

	arquivo, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo YAML: %v \n", err)
	}
	var secret Secret

	err = yaml.Unmarshal(arquivo, &secret)
	return secret

}

func readFileToScanner(path string) (*os.File, *bufio.Scanner) {

	arquivo, err := os.Open(path)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
	}

	// Cria um scanner para ler o arquivo linha a linha
	return arquivo, bufio.NewScanner(arquivo)

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Secret struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Template struct {
			Metadata struct {
				Name      string `yaml:"name"`
				Namespace string `yaml:"namespace"`
			} `yaml:"metadata"`
			Type string `yaml:"type"`
		} `yaml:"template"`
		EncryptedData map[string]string `yaml:"encryptedData"`
	} `yaml:"spec"`
}
