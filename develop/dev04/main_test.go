package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	expected := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
	}

	result := FindAnagramSets(&dictionary)

	if !reflect.DeepEqual(*result, expected) {
		t.Errorf("Ошибка в тесте: ожидаемый результат %v, полученный результат %v", expected, *result)
	}
}
