### 1. Запуск Kafka
```
make compose-up
```

### 2. Отправка сообщений через kafka-producer
```
- Создать топик test

docker exec -ti kafka /usr/bin/kafka-topics --list --bootstrap-server localhost:9092
docker exec -ti kafka /usr/bin/kafka-topics --create --topic test --bootstrap-server localhost:9092
docker exec -ti kafka /usr/bin/kafka-topics --describe --topic test --bootstrap-server localhost:9092
docker exec -ti kafka /usr/bin/kafka-topics --list --bootstrap-server localhost:9092

- Записать несколько сообщений в топик

docker exec -ti kafka /usr/bin/kafka-console-producer --topic test --bootstrap-server localhost:9092
message 1
message 2
message 3
```

### 3. Получение сообщений через kafka-consumer
```
- Прочитать сообщения из топика

docker exec -ti kafka /usr/bin/kafka-console-consumer --topic test --from-beginning --bootstrap-server localhost:9092
message 1
message 2
message 3
```

![Alt text](./kafka-log.jpg?raw=true)

### 4. Отправка и получение сообщений программно (golang)
```
Приложения с Kafka на языке golang

Producer:
- создает два топика: topic1 и topic2
- открывает транзакцию
- отправляет по 5 сообщений в каждый топик
- подтверждает транзакцию
- открывает другую транзакцию
- отправляет по 2 сообщения в каждый топик
- отменяет транзакцию

Consumer:
- читает сообщения из топиков topic1 и topic2 так, чтобы сообщения из подтверждённой транзакции были выведены, а из неподтверждённой - нет 

Проверка:
docker logs producer
docker logs consumer
```

### 5. Остановка Kafka
```
make compose-down
```
