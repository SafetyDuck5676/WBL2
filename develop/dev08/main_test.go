package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name     string
		cmd      *exec.Cmd
		expected string
	}{
		{
			name:     "echo",
			cmd:      exec.Command("echo", "Hello, World!"),
			expected: "Hello, World!\n",
		},
		{
			name:     "ls",
			cmd:      exec.Command("ls"),
			expected: "file1.txt\nfile2.txt\n",
		},
		// Добавьте ещё тестовые случаи по необходимости
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := runCmd(test.cmd, os.Stdout)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// Подготовительные действия перед запуском тестов, если необходимо

	// Запуск тестов
	exitCode := m.Run()

	// Завершающие действия после выполнения тестов, если необходимо

	// Завершить программу с кодом выхода тестов
	os.Exit(exitCode)
}
