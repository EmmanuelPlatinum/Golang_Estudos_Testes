package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	arquivoOrigem := "arqivoA.txt"
	arquivoDestino := "arquivoB.txt"

	conteudo, err := os.ReadFile(arquivoOrigem)
	if err != nil {
		log.Fatalf("Falha ao ler o arquivo de origem: %s", err)
	}

	err = os.WriteFile(arquivoDestino, conteudo, 0644)
	if err != nil {
		log.Fatalf("Falha ao escrever no arquivo de destino: %s", err)
	}

	fmt.Println("Conteúdo copiado com sucesso de", arquivoOrigem, "para", arquivoDestino)
}
