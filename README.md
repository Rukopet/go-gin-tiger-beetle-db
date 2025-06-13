# finance-api

Базовая структура REST API сервера на Go с использованием Gin для работы с TigerBeetle.

## Что создано

Базовая структура проекта с HTTP handlers без бизнес-логики, готовая для дальнейшей разработки.

## Структура проекта

```
finance-api/
├── cmd/server/main.go              # Точка входа приложения
├── internal/
│   ├── config/config.go            # Конфигурация
│   ├── handlers/                   # HTTP handlers (view layer)
│   │   ├── account.go              # Handlers для счетов
│   │   ├── transfer.go             # Handlers для переводов  
│   │   ├── stats.go                # Handlers для статистики
│   │   ├── middleware.go           # Middleware
│   │   └── router.go               # Настройка роутов
│   ├── services/                   # Интерфейсы и заглушки сервисов
│   │   ├── interfaces.go           # Интерфейсы сервисов
│   │   └── mock_services.go        # Mock реализации
│   └── models/                     # Модели данных
│       ├── account.go              # Модели счетов
│       ├── transfer.go             # Модели переводов
│       └── common.go               # Общие модели ответов
├── go.mod                          # Go модуль
├── Makefile                        # Команды для сборки и запуска
└── README.md
```

## Реализованные endpoint'ы

#### Счета
- `POST /api/v1/accounts` - создание счета
- `GET /api/v1/accounts/{id}` - получение счета
- `GET /api/v1/accounts` - список счетов (с пагинацией)
- `GET /api/v1/accounts/{id}/balance` - баланс счета
- `POST /api/v1/accounts/{id}/deposit` - пополнение
- `POST /api/v1/accounts/{id}/withdraw` - снятие
- `GET /api/v1/accounts/{id}/statement` - выписка

#### Переводы
- `POST /api/v1/transfers` - создание перевода
- `GET /api/v1/transfers/{id}` - получение перевода
- `GET /api/v1/transfers` - список переводов (с фильтрацией)

#### Статистика
- `GET /api/v1/stats/summary` - общая статистика

#### Системные
- `GET /health` - health check

## Запуск

```bash
# Загрузить зависимости
make mod-tidy

# Запустить в режиме разработки
make dev

# Собрать проект
make build

# Показать все команды
make help
```

## Тестирование API

После запуска сервер будет доступен на `http://localhost:8080`

```bash
curl http://localhost:8080/health
```

## ⚠️ Важно

Это базовая структура БЕЗ бизнес-логики! Все сервисы возвращают "not implemented".

### Что работает:
- ✅ HTTP сервер на Gin
- ✅ Роутинг всех endpoint'ов  
- ✅ Валидация входных данных
- ✅ Архитектурные слои

### Что НЕ реализовано:
- ❌ Интеграция с TigerBeetle
- ❌ Бизнес-логика операций
- ❌ Тесты
