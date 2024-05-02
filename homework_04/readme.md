### 1. Запуск Couchbase кластера
```
make compose-up
```

### 2. Инициализация
```
http://127.0.0.1:8091

Setup New Cluster

Cluster Name: test
Create Admin Username: admin
Create Password: example

Next -> Finish with defaults

// Подключение серверов

http://127.0.0.1:8094

Join Cluster -> couchbase3, admin, example

http://127.0.0.1:8095

Join Cluster -> couchbase3, admin, example

// Перебалансировка

Servers -> Rebalance

```

### 3. Наполнение тестовыми данными
```
http://127.0.0.1:8095

Перейти в Buckets -> sampleBuckets -> "travel-sample" -> Load Sample Data

После загрузки проверить данные в Buckets
```

### 4. Выполнение запросов
```
Data Tools -> Query

SELECT count(route.airlineid)
FROM `travel-sample` route;

[
  {
    "$1": 24024
  }
]
```

### 5. Проверка отказоустойчивости
```
Можно остановить контейнер

make accident

Проверить статусы и базу

http://127.0.0.1:8094

Servers - один сервер недоступен

172.21.0.3 сделать Failover и Rebalance

Data Tools -> Query

SELECT count(route.airlineid)
FROM `travel-sample` route;

[
  {
    "$1": 24024
  }
]

Данные не потерялись
```

### 5. Остановка
```
make compose-down
```