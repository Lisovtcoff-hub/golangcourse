Логинимся под своим логином и созданным паролем на предыдущей шаге (access token).
```bash
  docker login -u <GITLAB-LOGIN> -p <ACCESS-TOKEN> gitlab.golang-school.ru:5050
```

Полезные команды:
```bash
docker ps # Запущенные контейнеры
docker images # Спулленые образы

docker rm $(docker ps -aq) --force # Удалить все контейнеры
docker rmi $(docker images -q) --force # Удалить все образы
```

Пуллим к себе образ с docker hub
```bash
docker pull alpine:latest
```

Смотрим что спуллилось
```bash
docker images

# Output
# alpine       latest    9234e8fb04c4   7 weeks ago   8.31MB
```

В Gitlab образы привязываются к проекту. Поэтому надо создать новый проект в вашей папке.
Предлагаю назвать его `my-app`.

После того как создали репозиторий, следуйте по инструкции дальше.

Ставим новый тег:
<USERNAME> - ваше имя пользователя
```bash
docker tag alpine:latest gitlab.golang-school.ru:5050/potok-2/<USERNAME>/my-app:v0.1.0
```

Смотрим на изменения и видим новый образ в списке:
```bash
docker images

# Output
# harbor.goscl.ru/mnepryakhin/alpine       v0.1.0         4048db5d3672   8 days ago      7.84MB
```

Пушим образ:
<USERNAME> - ваше имя пользователя
```bash
docker push gitlab.golang-school.ru:5050/potok-2/<USERNAME>/my-app:v0.1.0
```

Смотрим на запушеный образ на странице:
<USERNAME> - ваше имя пользователя
```bash
https://gitlab.golang-school.ru/potok-2/<USERNAME>/my-app/container_registry/
```

# Пуллим образ с Gitlab Registry к себе на машину:
```bash
docker pull gitlab.golang-school.ru:5050/potok-2/<USERNAME>/my-app:v0.1.0
```