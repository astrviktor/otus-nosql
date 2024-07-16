## ETCD

### 1. Запуск etcd
```
make etcd-up
```

### 2. Проверка etcd
```
docker exec -it etcd-1 etcdctl endpoint health
127.0.0.1:2379 is healthy: successfully committed proposal: took = 2.933277ms

docker exec -it etcd-2 etcdctl endpoint health
127.0.0.1:2379 is healthy: successfully committed proposal: took = 3.471154ms

docker exec -it etcd-3 etcdctl endpoint health
127.0.0.1:2379 is healthy: successfully committed proposal: took = 3.388803ms

```

### 3. Создание и получение ключа в etcd

```
docker exec -it etcd-1 etcdctl put key1 value1
OK

docker exec -it etcd-1 etcdctl get key1
key1
value1

docker exec -it etcd-2 etcdctl get key1
key1
value1

docker exec -it etcd-3 etcdctl get key1
key1
value1

```

### 4. Проверка отказоустойчивости etcd

```
docker stop etcd-2

docker exec -it etcd-1 etcdctl endpoint health
127.0.0.1:2379 is healthy: successfully committed proposal: took = 2.309533ms

docker exec -it etcd-2 etcdctl endpoint health
Error response from daemon: Container 2ca74b67b0c37a674e1013ca1cf3d4a7a858c86cbf2bc2d22d8ff79b5a7b8eb1 is not running

docker exec -it etcd-3 etcdctl endpoint health
127.0.0.1:2379 is healthy: successfully committed proposal: took = 2.682818ms

docker exec -it etcd-1 etcdctl get key1
key1
value1

docker exec -it etcd-3 etcdctl get key1
key1
value1

```

### 5. Остановка etcd
```
make etcd-down
```

## CONSUL

### 1. Запуск consul
```
make consul-up
```

### 2. Проверка consul
```
curl http://127.0.0.1:8501/v1/status/peers
["172.26.0.3:8300","172.26.0.2:8300","172.26.0.4:8300"]

curl http://127.0.0.1:8501/v1/status/leader
"172.26.0.3:8300"

```

### 3. Проверка отказоустойчивости consul

```
docker stop consul-server-1

curl http://127.0.0.1:8502/v1/status/peers
["172.26.0.2:8300","172.26.0.4:8300"]

curl http://127.0.0.1:8502/v1/status/leader
"172.26.0.4:8300"
```

### 4. Остановка consul
```
make consul-down
```

