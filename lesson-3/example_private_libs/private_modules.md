### Приватный репозиторий для модулей в Gitlab

Версионируем libs через semver теги: v0.1.0, далее v0.2.0 и т.д.

Создаём токен на чтение в Gitlab

Добавляем токен в файл ~/.netrc:
```
machine gitlab.golang-school.ru login <GITLAB_LOGIN> password <TOKEN>
```

Добавляем репозиторий в go env:
```sh
go env -w GOPRIVATE="gitlab.golang-school.ru"
```

Проверяем, что всё работает:
```sh
go get gitlab.golang-school.ru/potok-2/mnepryakhin/libs
```

