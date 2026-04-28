# Основной флоу

```bash
# Build
docker build -t gitlab.golang-school.ru:5050/potok-2/<USERNAME>/my-app:v0.1.0 .

# Login
docker login -u <GITLAB-LOGIN> -p <ACCESS-TOKEN> gitlab.golang-school.ru:5050

# Push
docker push gitlab.golang-school.ru:5050/potok-2/<USERNAME>/my-app:v0.1.0

# Pull
docker pull gitlab.golang-school.ru:5050/potok-2/<USERNAME>/my-app:v0.1.0
```

# Шпаргалка по командам

```bash
# Запущенные контейнеры
docker ps

# Спулленые образы
docker images

# Запустить контейнер
docker run -d -p 8080:80 nginx

# Остановить контейнер
docker stop <CONTAINER_ID>

# Смотреть логи контейнера
docker logs -f <CONTAINER_ID>

# Удалить контейнер
docker rm <CONTAINER_ID>

# Удалить все контейнеры
docker rm $(docker ps -a -q) --force

# Удалить все образы
docker rmi $(docker images -q) --force
```