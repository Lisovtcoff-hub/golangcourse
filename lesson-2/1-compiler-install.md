Установка компилятора
```
https://go.dev/doc/install
```

Проверка доступности
```bash
go version
```
Если не доступен, то добавить в PATH
```bash
export PATH=$PATH:/usr/local/go/bin
```

Хранить проекты можно в любой удобной папке.
Если у вас Windows, то в файловой системе WSL.

Я обычно храню в
```
~/Projects
```

Зарегистрируйтесь в GitLab
```
https://gitlab.golang-school.ru/users/sign_up
```

Создайте SSH ключ для удобства
```
https://gitlab.golang-school.ru/-/user_settings/ssh_keys
```

Генерация ключей в Linux
```bash
ssh-keygen
```
- Укажите название файла ключей
- Пароль можно не указывать
- Далее скопируйте содержание `.pub` файла в поле `Key` GitLab'а

Если что вот инструкция от GitLab
```
https://docs.gitlab.com/17.6/ee/user/ssh.html
```