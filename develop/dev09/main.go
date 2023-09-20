package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Проверяем, что переданы аргументы командной строки
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	url := os.Args[1]

	// Создаем директорию для сохранения файлов
	dirname := getDirName(url)
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Скачиваем сайт
	err = downloadSite(url, dirname)
	if err != nil {
		fmt.Println("Error downloading site:", err)
		return
	}

	fmt.Println("Site downloaded successfully!")
}

// Функция для получения имени директории, основанной на URL сайта
func getDirName(url string) string {
	// Удаляем протокол и слэши из URL
	// Например, для "https://example.com/" остается только "example.com"
	domain := strings.TrimPrefix(url, "https://")
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimSuffix(domain, "/")

	return domain
}

// Функция для скачивания сайта
func downloadSite(url string, dirname string) error {
	// Открываем URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Генерируем имя файла основываясь на текущем времени
	filename := time.Now().Format("20060102150405") + ".html"

	// Создаем файл для записи
	filepath := dirname + "/" + filename
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем тело ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Добавьте ещё тестовые случаи по необходимости
