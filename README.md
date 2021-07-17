# Установка
```kubectl apply -f infra/.```

# Tecтирование

`curl arch.homework/health`
`curl arch.homework/health/`
`curl arch.homework/otusapp/igor/account/email`
`curl arch.homework/otusapp/igor/`

# Описание
Сервер на Go представляет два эндпоинта
* `/health` возвращает `{"status": "OK"}`
* `/greeting/<name>` возвращает строку `Hello <name>!`

Ingress перенаправляет запрос на путь `^/health/?$` в эндпоинт `/health` и запрос на путь `^/otusapp/<name>/*` в эндпоинт `/greeting/<name>`