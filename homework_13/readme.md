### 1. Запуск Tarantool и Tarantool Cartridge CLI
```
docker run --network=host -ti --rm ubuntu:20.04 /bin/bash

apt-get update && apt-get install curl -y

curl -L https://tarantool.io/release/2/installer.sh | bash

apt-get -y install tarantool
8
34

apt-get install cartridge-cli

tarantool -version
Tarantool 2.11.3-0-gf933f77904
...

cartridge version
Tarantool Cartridge CLI
 Version:	2.12.12
 OS/Arch: 	linux/amd64
 Git commit:	7f7efcf
```

### 2. Создание шаблон приложения
```
cd /root
cartridge create --name myapp
```

### 3. Сборка и запуск приложение
```
cd /root/myapp/

cartridge build
   • Build application in /root/myapp
   • Running `cartridge.pre-build`
   • Running `tarantoolctl rocks make`
   • Application was successfully built

cartridge start
...

```

### 4. Топология кластера в UI и bootstrap
```
web-интерфейс на хосте http://0.0.0.0:8081/

Cluser / router Configure

Replica set name - Router

[v] app.roles.custom
[v] vshard-router

Create replica set

Cluser / s1-master Configure

[v] vshard-storage

Replica set name - Master


Bootstrap vshard
```
![Alt text](./tarantool.jpg?raw=true)

### 5. Остановка Tarantool и Tarantool Cartridge CLI
```
ctrl+C
exit
```
