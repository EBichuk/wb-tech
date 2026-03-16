### HTTP-сервер для работы с небольшим календарем событий

### API endpoints

- POST /create_event — создание нового события;
- POST /update_event — обновление существующего;
- POST /delete_event — удаление;
- GET /events_for_day/user_id=&date=YYYY-MM-DD — получить все события на день;
- GET /events_for_week/user_id=&date=YYYY-MM-DD — события на неделю;
- GET /events_for_month/user_id=&date=YYYY-MM-DD — события на месяц.

### Тесты
запустит юнит тесты
```bash
make unit-test
```
запустит линтер
```bash
make lint
```
### Структура
```
calendar-service/
├── cmd/
│   └── main.go                     # Точка входа приложения              
├── internal/
│   ├── api/                        # Запуск и остановка приложения
│   │   └── api.go                       
│   ├── config/                     # Конфигурация
│   │   └── config.go
│   ├── errs/                       # Кастомные ошибки
│   │   └── errs.go
│   ├── handler/
│   |   ├── event/                  # Слой handler             
|   |   |   ├── dto.go              
|   |   |   └── handler.go                        
|   |   ├── middleware.go           # Мидлвари для логирования
|   |   └── router.go               
│   ├── models/                     # Модели данных
│   │   ├── event.go
│   │   └── event_test.go
│   ├── repository/                 # Слой repository
│   |   ├── mocks/                  # Слой handler             
|   |   |   └── mock_repository.go  # Моки для тестов
│   │   └── repository.go
│   ├── server/                     # HTTP сервер
│   |   └── server.go
│   └── service/                    # Слой service
│   │   ├── service.go
│   │   └── service_test.go
├── prg/
│   └── logger/                     # Настройка логера                                       
│       └── logger.go                       
├── .golangci.yml                   # Настройки линтера                         
├── .env.example                    # Пример переменных окружения
├── go.mod                          
├── go.sum                          
└── makefile                
```
