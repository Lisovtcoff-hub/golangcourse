### Установка k3d

`k3d` - это инструмент для запуска k3s (легковесной версии Kubernetes от Rancher) в Docker контейнерах.

Документация: https://k3d.io/stable

```shell
# Для Linux
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

### Использование локального Registry для k3d

Чтобы локальный registry работал в /etc/hosts надо добавить
```shell
127.0.1.1 k3d-registry
```

Подробнее см. https://k3d.io/stable/usage/registries/#using-a-local-registry

### Установка kubectl

`kubectl` - это командная утилита для взаимодействия с Kubernetes кластерами.

https://kubernetes.io/docs/tasks/tools/

```shell
# Для Linux (нужен VPN)
curl -LO "https://dl.k8s.io/release/v1.33.1/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
kubectl version --client
```

### Удобный UI для k8s

https://k8slens.dev/