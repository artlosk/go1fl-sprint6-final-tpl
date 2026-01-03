package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		log.Printf("Ошибка чтения index.html: %v", err)
		http.Error(w, "Не удалось загрузить страницу", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlContent)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10485760) // 10 МБ
	if err != nil {
		log.Printf("Ошибка парсинга формы: %v", err)
		http.Error(w, "Ошибка обработки формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Ошибка получения файла: %v", err)
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	convertedContent, err := service.ConvertString(string(fileContent))
	if err != nil {
		log.Printf("Ошибка конвертации: %v", err)
		http.Error(w, fmt.Sprintf("Ошибка конвертации: %v", err), http.StatusInternalServerError)
		return
	}
	timestamp := time.Now().UTC().String()
	originalExt := filepath.Ext(header.Filename)
	newFileName := timestamp + originalExt

	outputFile, err := os.Create(newFileName)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(convertedContent)
	if err != nil {
		log.Printf("Ошибка записи в файл: %v", err)
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	log.Printf("Файл успешно создан: %s", newFileName)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Результат конвертации:\n\n%s\n\nФайл сохранен как: %s", convertedContent, newFileName)
}
