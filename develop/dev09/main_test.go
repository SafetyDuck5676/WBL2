package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetDirName(t *testing.T) {
	url := "https://example.com/"
	expected := "example.com"

	dirname := getDirName(url)
	if dirname != expected {
		t.Errorf("Expected directory name: %s, but got: %s", expected, dirname)
	}
}

func TestDownloadSite(t *testing.T) {
	// Создаем временный HTTP-сервер с тестовым HTML-контентом
	content := "<html><body>Test content</body></html>"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, content)
	}))
	defer server.Close()

	// Создаем временную директорию для сохранения файлов
	dirname, err := ioutil.TempDir("", "test_wget")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(dirname)

	// Скачиваем сайт
	err = downloadSite(server.URL, dirname)
	if err != nil {
		t.Fatalf("Error downloading site: %v", err)
	}

	// Проверяем, что файл был создан и содержит ожидаемый контент
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		t.Fatalf("Error reading directory: %v", err)
	}

	if len(files) != 1 {
		t.Errorf("Expected 1 file, but got %d", len(files))
	}

	filepath := dirname + "/" + files[0].Name()
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	if string(fileContent) != content {
		t.Errorf("Expected file content: %s, but got: %s", content, fileContent)
	}
}

func TestMain(m *testing.M) {
	// Запуск всех тестов
	exitCode := m.Run()

	os.Exit(exitCode)
}
