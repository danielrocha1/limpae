package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)


// UploadFileToSupabase faz upload de um arquivo para o Supabase utilizando a API REST
func UploadFileToSupabase(file multipart.File, fileName string) (string, error) {
	// Bucket onde o arquivo será armazenado
	bucketName := "uploads"

	// Obter URL e chave de API do Supabase
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("As variáveis de ambiente SUPABASE_URL e SUPABASE_KEY não estão configuradas")
	}

	// Criar o request multipart
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Adicionando o arquivo no corpo da requisição
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", fmt.Errorf("erro ao criar formulário de arquivo: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("erro ao copiar o arquivo para a requisição: %v", err)
	}

	// Fechar o writer para finalizar o corpo
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("erro ao fechar o escritor multipart: %v", err)
	}

	// Criar a URL do endpoint de upload do Supabase
	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucketName, fileName)

	// Criar a requisição HTTP
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", fmt.Errorf("erro ao criar a requisição HTTP: %v", err)
	}

	// Definir os cabeçalhos da requisição
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	fileURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucketName, fileName)

	
	// Enviar a requisição HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar a requisição para o Supabase: %v", err)
	}
	defer resp.Body.Close()

	// Verificar a resposta do Supabase
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro ao fazer upload para Supabase: status %d", resp.StatusCode)
	}

	// Gerar a URL pública do arquivo

	return fileURL, nil
}
