package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type School struct {
	school   string
	location string
	groups   []Groups
}

type Groups struct {
	title    string
	students int
	remote   bool
}

type WorkDay struct {
	Title    string   `json:"title"`
	City     string   `json:"city"`
	Address  string   `json:"address"`
	IsOnline bool     `json:"is_online"`
	Seats    int      `json:"seats"`
	Mentors  []string `json:"mentors"`
}

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
	WantsGo   bool   `json:"wants_go"`
}

func main() {

	// Подзадача 2 — Первый маршалинг

	day := WorkDay{
		Title:    "Певый день в Интокод",
		City:     "грозный",
		Address:  "Дадин Айбики 14",
		IsOnline: false,
		Seats:    20,
		Mentors:  []string{"Кудузов Ахмад", "Назиров Расул"},
	}

	js, err := json.MarshalIndent(day, "", "  ")
	if err != nil {
		fmt.Println("ошибка:", err)
	}
	fmt.Println(string(js))

	// Подзадача 3 — Анмаршалинг: JSON в структуру

	student := Student{}

	v, err := os.ReadFile("student.json")
	if err != nil {
		fmt.Println("ошибка:", err)
	}
	json.Unmarshal(v, &student)

	if student.WantsGo && student.Age >= 16 {
		fmt.Println("Иса Исаев готов к Буткемпу Intocode в Грозном!")
	} else {
		fmt.Println("слишком маленький возраст")
	}

	// Подзадача 4 — Частая ошибка №1

	type Lesson struct {
		Topic   string `json:"topic"`
		Minutes int    `json:"minutes"`
	}

	data := []byte(`{"topic":"Введение в JSON","minutes":45}`)
	var l Lesson
	err = json.Unmarshal(data, &l) // без указателя не работает
	fmt.Println("Ошибка:", err)
	fmt.Printf("Результат: %+v\n", l)

	// Подзадача 5 — Частая ошибка №2: несоответствие типов в JSON

	// подумал и решил
}
