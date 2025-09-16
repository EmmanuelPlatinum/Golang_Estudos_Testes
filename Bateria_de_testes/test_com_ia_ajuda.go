package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Define os nomes dos arquivos de origem e destino
	arquivoOrigem := "arquivoA.txt"
	arquivoDestino := "arquivoB.txt"

	// Lê todo o conteúdo do arquivo de origem
	conteudo, err := os.ReadFile(arquivoOrigem)
	if err != nil {
		log.Fatalf("Falha ao ler o arquivo de origem: %s", err)
	}

	// Escreve o conteúdo lido no arquivo de destino
	// O terceiro argumento (0644) define as permissões do arquivo
	err = os.WriteFile(arquivoDestino, conteudo, 0644)
	if err != nil {
		log.Fatalf("Falha ao escrever no arquivo de destino: %s", err)
	}

	fmt.Println("Conteúdo copiado com sucesso de", arquivoOrigem, "para", arquivoDestino)
}