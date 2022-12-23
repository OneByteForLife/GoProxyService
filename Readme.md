# GoProxyService

# Описание

Это микросервис который собирает прокси типа HTTP, для дальнейнего их использования при парсинге сторонних ресурсов. Так же можно использовать его в любой среде так как он независимый. 

## Структура проекта

```
├── app
│   └── cmd
│       └── main.go
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── run.go
│   ├── middleware
│   │   └── jwt.go
│   ├── models
│   │   └── proxy.go
│   └── routes
│       └── route.go
├── pkg
│   └── log.go
├── Readme.md
└── secret.txt
```

## API точки доступа

Ниже будут перечисленны точки доступа к API.

URL - http://localhost:80

#### **/api/v1/get?total=value**
* `GET` : Получение определенного кол-ва элементов.

#### Пример ответа:

```
{
    "api_version": "1.0",
    "content": [
        {
            "Type": "HTTP",
            "Data": {
                "IP": "203.13.32.73",
                "Port": "80"
            }
        }
    ],
    "description": "Success",
    "status_code": 200
}
```

## Использование

На данный момент он поддерживает как запуск средствами **Golang** так и используя **Docker**.

**Golang** - Перейдите в директорию **/app/cmd/** в ней вам нужно запустить файл **main.go**.

**Docker** - В корневой папке лежит файл **Dockerfile** запустите его используя команду. ```docker build .``` затем команду ```docker run -p 80:80 93e9a9ead376``` имя образа можно взять после сборки контейнера используя команду ```docker images```