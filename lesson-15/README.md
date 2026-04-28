## Запуск

У вас должен быть установлен Docker и Docker Compose.

```bash
# Устанавливаем утилиту migrate (для запуска миграций)
make migrate-install

# Поднимаем контейнеры с БД, Графаной и т.д.
make up

# Прогоняем миграции в БД
make migrate-up

# Запуск приложения
make run
```

Grafana: http://localhost:3001
Смотреть трейсы на странице Explore, в пустое поле ввести: {}