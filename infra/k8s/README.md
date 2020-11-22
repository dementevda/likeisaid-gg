# Плэйбуки для tasker-gg

## Требования
1) kubectl
2) helm3

## Описание

## Полезные команды
### Добавляем заразу к 3му мастеру - не шедулить туда поды кроме ингерсса
```bash
kubectl label node vs03.k8s.master node-role/ingress=

kubectl taint node vs03.k8s.master node-role/ingress="":NoSchedule

```
### Установка ингресса через хелм
```bash
helm install ingress ingress-nginx/ingress-nginx -f ingress-nginx-values.yaml
helm install --name metallb stable/metallb -f metallb-values.yaml
```
### 
```bash
```
