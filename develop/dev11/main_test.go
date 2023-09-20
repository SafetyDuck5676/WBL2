package main

import (
	"net/http"
	"testing"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Реализация обработчика CreateEvent
}

func GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	// Реализация обработчика GetEventsForDay
}

func TestCreateEvent(t *testing.T) {
	// Ваш тест для CreateEvent
}

func TestGetEventsForDay(t *testing.T) {
	// Ваш тест для GetEventsForDay
}

func TestMain(t *testing.T) {
	t.Run("TestCreateEvent", TestCreateEvent)
	t.Run("TestGetEventsForDay", TestGetEventsForDay)
}
