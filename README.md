# Emergency Notification System

## TODO
- [ ] Загрузка пользователей с контактами через .cvs или .xls файл
- [ ] Создание шаблонов нотификаций и конфигурация нужных получателей
- [ ] Отправка нотификаций на любой девайс получателя автоматически при отравке нотификаций всей группе получателей
- [ ] Подключить Redis

## СТЭК
- Golang - v1.22
- Docker
- Kafka
- Twilio
- Postgres
- Redis

## Запуск приложения 

- `go-task run`


## taskfile
### ENV
```
    CONFIG_PATH: /path/to/app/config/config.yaml
```