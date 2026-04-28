## Настройки доступа к Gitlab Registry

Нужно создать персональный токен доступа (access token) в настройках своего профиля на Gitlab.

Переходим по ссылке:
```web
https://gitlab.golang-school.ru/-/user_settings/personal_access_tokens
```

Нажимаем кнопку `Add new token`
Token name: Registry

В **Select scopes** нужны права:
- api
- read_api
- read_user
- read_repository
- write_repository

Нажимаем кнопку `Create personal access token`.

Сохраняем сгенерированный токен для дальнейшего использования.



