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
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
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
	newFileName := time.Now().Format("2006-01-02_15-04-05") + filepath.Ext(header.Filename)

	if err := os.WriteFile(newFileName, []byte(convertedContent), 0644); err != nil {
		log.Printf("Ошибка сохранения файла: %v", err)
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}

	log.Printf("Файл успешно создан: %s", newFileName)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := fmt.Fprintf(w, "Исходное содержимое:\n\n%s\n\nРезультат конвертации:\n\n%s\n\nФайл сохранен как: %s", string(fileContent), convertedContent, newFileName); err != nil {
		log.Printf("Ошибка отправки ответа: %v", err)
	}
}
