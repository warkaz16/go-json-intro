# go-json-intro
# Задача – JSON в Go: маршалинг и анмаршалинг

## Постановка задачи

Сегодня ты познакомишься с базовой работой с JSON в Go: **преобразование структур в JSON** (маршалинг) и обратное **преобразование JSON в структуры** (анмаршалинг).

### Материалы для изучения

1. Официальная документация `encoding/json` (разделы про `Marshal`/`Unmarshal`). ([Go Packages][1])
2. JSON and Go. ([Go][2])
3. Go by Example: JSON. ([Go by Example][3])
4. DigitalOcean: How To Use JSON in Go. ([DigitalOcean][4])
5. Обработка JSON в Go. ([PurpleSchool][5])

## Шпаргалка

- Пакет: `encoding/json`

  - `json.Marshal(v)` - возвращает JSON в формате `[]byte`
  - `json.MarshalIndent(v, "", "  ")` - возвращает красивый (отформатированный) JSON
  - `json.Unmarshal(data, &v)` — делает из JSON структуру Go **важно:** второй аргумент — указатель.
  - Теги полей: `` `json:"field_name,omitempty"` ``

- Чтение/запись файлов:

  - `os.ReadFile("file.json")`
  - `os.WriteFile("file.json", data, 0644)`

## Подзадача 0 — Каким будет JSON?

Создай файл `coffee.json` и накидай в нём пример JSON данных, которые будут соответствовать следующей структуре:

```go
type CoffeeShop struct {
    Name    string   `json:"name"`
    City    string   `json:"city"`
    Items   []string `json:"items"`
    OpenNow bool     `json:"open_now"`
}
```

## Подзадача 1 — Какая нужна структура?

В файле `main.go` спроектируй структуру, которая будет соответствовать следующему JSON:

```json
{
  "school": "Интукод",
  "location": "Чеченская Республика",
  "groups": [
    {
      "title": "Go Начало",
      "students": 12,
      "remote": false
    },
    {
      "title": "Go Продвинутый",
      "students": 9,
      "remote": true
    }
  ]
}
```

## Подзадача 2 — Первый маршалинг

Создай новый проект (`go mod init`), файл `main.go`.
Опиши структуру расписания рабочего дня в Интукод:

```go
type WorkDay struct {
    Title       string   `json:"title"`
    City        string   `json:"city"`
    Address     string   `json:"address"`
    IsOnline    bool     `json:"is_online"`
    Seats       int      `json:"seats"`
    Mentors     []string `json:"mentors"`
}
```

Заполни значение:

- `Title`: "Первый день в Intocode"
- `City`: "Грозный"
- `Address`: "ул. Дадин Айбики, 14"
- `IsOnline`: `false`
- `Seats`: 20
- `Mentors`: `[]string{"Кудузов Ахмад", "Назиров Расул"}`

Требуется:

1. Сериализуй структуру в JSON с помощью `json.MarshalIndent`.
2. Выведи в консоль получившийся JSON и изучи его.

## Подзадача 3 — Анмаршалинг: JSON в структуру

Создай файл `student.json` со следующим содержимым:

```json
{
  "first_name": "Иса",
  "last_name": "Исаев",
  "age": 17,
  "city": "Грозный",
  "wants_go": true
}
```

Опиши соответствующую структуру:

```go
type Student struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age"`
    City      string `json:"city"`
    WantsGo   bool   `json:"wants_go"`
}
```

Требуется:

1. Получи содержимое `student.json` с помощью `os.ReadFile`.
2. Выполни `json.Unmarshal` в переменную типа `Student`.
3. Если `WantsGo == true` и `Age >= 16`, выведи:
   `Иса Исаев готов к Буткемпу Intocode в Грозном!`
   Иначе — выведи причину – `слишком маленький возраст`.

## Подзадача 4 — Частая ошибка №1

Есть такой код для анмаршалинга, но в нём баг. Найди где он:

```go
type Lesson struct {
    Topic   string `json:"topic"`
    Minutes int    `json:"minutes"`
}

data := []byte(`{"topic":"Введение в JSON","minutes":45}`)
var l Lesson
err := json.Unmarshal(data, l)
fmt.Println("Ошибка:", err)
fmt.Printf("Результат: %+v\n", l)
```

Требуется:

1. Запусти — посмотри, что произойдёт.
2. Исправь баг, чтобы данные попали в структуру.
3. Объясни одной-двумя строками, почему изначально было неверно.

## Подзадача 5 — Частая ошибка №2: несоответствие типов в JSON

Есть JSON:

```json
{ "title": "Разбор задач", "minutes": "60" }
```

И структура:

```go
type Meetup struct {
    Title   string `json:"title"`
    Minutes int    `json:"minutes"`
}
```

Требуется:

1. Попробуй определить почему анмаршалинг приведет к ошибке.
2. Подумай, как изменить **только** структуру так, чтобы `Unmarshal` прошёл (не меняя JSON).
3. Также подумай, почему так сработало, и что лучше: менять тип поля или формат JSON.

[1]: https://pkg.go.dev/encoding/json 'encoding/json - Go Packages'
[2]: https://go.dev/blog/json 'JSON and Go - The Go Programming Language'
[3]: https://gobyexample.com/json 'JSON - Go by Example'
[4]: https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go 'How To Use JSON in Go - DigitalOcean'
[5]: https://purpleschool.ru/knowledge-base/article/encoding-json 'Обработка JSON в Golang - PurpleSchool'