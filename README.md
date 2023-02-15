# Test-task
Проект содержит два сервиса: первый имеет только один эндпоинт который при обращении к нему генерирует salt, второй сервис после получения этого salt хэширует пароль и заносит в бд(mongodb)

### Требования
Для работы БД, необходим [MongoDB]

### Запуск сервера
В директории app2 поправить конфиг файл под свои нужды,
Запустить два терминала и монго:
app1: 
```sh
$ go run cmd/main.go
```
app2: 
```sh
$ go run cmd/main.go --config=config/config.toml
```
### Docker
Докер находится на стадии разработки...