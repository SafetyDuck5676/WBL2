package main

import (
	"reflect"
	"testing"
)

// Функция для проверки сортировки по числовому значению
func TestSortSliceNumeric(t *testing.T) {
	arr := []string{"10", "2", "5", "1", "20"}
	expected := []string{"1", "2", "5", "10", "20"}

	sortSliceNumeric(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, but got %v", expected, arr)
	}
}

// Функция для проверки чтения строк из файла
func TestReadLines(t *testing.T) {
	lines, err := readLines("/home/safetyduck/WBL2/develop/dev03/filename")
	expected := []string{"1", "2", "3", "4"}

	if err != nil {
		t.Errorf("Error reading lines from file: %v", err)
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("Expected %v, but got %v", expected, lines)
	}
}
