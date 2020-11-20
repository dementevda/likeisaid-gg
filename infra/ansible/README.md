# Плэйбуки для tasker-gg

## Требования
1) Python3
2) `git clone https://github.com/kubernetes-sigs/kubespray`
3) Зависимости из requirements.txt
4) Читаем документацию kubespray по установке

## Описание
1) Конкретно у себя решил поставить flannel

## Полезные команды
### Пингануть все хосты
```bash
ansible all -m ping -i inventory.ini
```
### Пингануть все хосты через плэйбук
```bash
ansible-playbook -i inventory.ini test_playbook/ping.yml --ask-become-pass
```
### Пингануть все хосты через плэйбук
```bash
ansible-playbook -i inventory/mycluster/inventory.ini -b -v cluster.yml

```
