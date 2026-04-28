В Windows WSL лучше использовать Ubuntu и делать всё через Windows Terminal.

Установить докер в Ubuntu 22.04 или 24.04
```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
```

Дать права пользователю на запуск докера без sudo
<YOUR-USER-NAME> - заменить на ваше имя пользователя
```bash
sudo usermod -a -G docker <YOUR-USER-NAME>
```

Сделайте ребут
```bash
sudo reboot
```