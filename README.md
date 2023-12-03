
#### Оглавление:
____
0. [Сервис отправки сообщений](#сервис-отправки-сообщений).
1. [ЗАВИСИМОСТИ](#зависимости).
2. [ЗАПУСК/СБОРКА](#запусксборка).
2.1. [Конфигурация](#конфигурация).
2.1.1 [Флаги](#1-флаги).
2.1.2 [Облачные переменные](#2-облачные-переменные).
2.2. [Запуск сервера](#запуск-сервера).
3. [Для разработчиков](#для-разработчиков).
3.1. [Сервер](#сервер).
____

# Сервис отправки сообщений.

Сервис позволяет собирать телефоны клиентов, задание на рассылку, отчеты по сообщениям в БД (используется PostgreSQL:latest). Сбор данных осуществляется по протоколу HTTP.
Сервис анализирует в какое время и до какого времени необходимо отправить сообщение всем клиентам, что удовлетворяют фильтрам в задании на рассылку. Непосредственная отправка осуществляется внешним сервисом. Данный сервис собирает данные об успешных или неуспешных отправках и добавляет их в БД.

Структура [клиентов](https://github.com/CyrilSbrodov/mailingAPI/blob/master/internal/storage/models/models.go)
```GO
type Client struct {
	ID             int    `json:"id"`
	PhoneNumber    string `json:"phone_number"`
	MobileOperator string `json:"mobile_operator"`
	Tag            string `json:"tag"`
	TimeZone       string `json:"time_zone"`
}
```
Структура [рассылок](https://github.com/CyrilSbrodov/mailingAPI/blob/master/internal/storage/models/models.go)
```GO
type Mailing struct {
	ID        int       `json:"id"`
	Message   string    `json:"messages"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
	Filter    Filter    `json:"filter"`
}

type Filter struct {
	MobileOperator string `json:"mobile_operator"`
	Tag            string `json:"tag"`
}
```

Структура сервиса следующая:
1) Сервер - обработка полученных данных, отправка на внешний сервис, обогащение статусами и вставка их в БД PostgreSQL.
2) БД - прием получаемых данных.
____
# ЗАВИСИМОСТИ.

Используется язык go версии 1.21. Используемые библиотеки:
- github.com/caarlos0/env/v6 v6.10.1
- github.com/go-chi/chi/v5 v5.0.10
- github.com/jackc/pgx/v5 v5.5.0
- github.com/lib/pq v1.10.9
- github.com/rs/zerolog v1.31.0
- POSTGRESQL latest
____

# ЗАПУСК/СБОРКА

## Конфигурация

Предусмотрены различные конфигурации:
1) флаги
2) облачные переменные 

Флаги при запуске сервера:
```
-a //адрес сервера
-c //интервал проверки новых рассылок
-e //адрес внешнего сервиса
-d //адрес бд
-token //токен для аунтификации во внешнем сервисе
```

## 1) Флаги
Параметры запуска передаются в формате: -a localhost:8080, 
где "-а" параметр адреса сервера, "localhost:8080" адрес сервера.

Это позволяет запускать утилиту следующим образом:
```
cd cmd
go run main.go -a localhost:8080
```

## 2) Облачные переменные
Перед запуском утилиты необходимо присвоить переменным значения в формате: 
```
ADDRESS='localhost:8080'
```
где "ADDRESS" - облачная переменная, 'localhost:8080' присвоение значения данной переменной.

```
## Запуск сервера

Необходимо запустить сервер из пакета [cmd](https://github.com/CyrilSbrodov/mailingAPI/blob/master/cmd/main.go)
```
cd cmd
go run main.go
```

# Для разработчиков
Структура приложения позволяет нативно вносить корректировки:

[Структура сервера](https://github.com/CyrilSbrodov/mailingAPI/blob/master/internal/app/app.go):
```GO
type ServerApp struct {
	cfg    *config.ServerConfig
	logger *loggers.Logger
	router *chi.Mux
}
```

Немного о сервере:

# Сервер
Сервер получает данные по следующим эндпоинтам:
1) [http](https://github.com/CyrilSbrodov/mailingAPI/blob/master/internal/handlers/handler.go):
```GO
	r.Group(func(r chi.Router) {
		r.Post("/api/client", h.AddClient())
		r.Post("/api/client/update", h.UpdateClient())
		r.Post("/api/client/delete", h.DeleteClient())
		r.Post("/api/mailing", h.AddMailing())
		r.Post("/api/mailing/update", h.UpdateMailing())
		r.Post("/api/mailing/delete", h.DeleteMailing())
		r.Get("/api/mailing", h.GetAllStatistic())
		r.Post("/api/mailing/get", h.GetDetailStatistic())
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"), 
		))
	})

```
