### 1. Запуск
```
make compose-up
```

### 2. Проверка
```
Сервис делает 10 запросов в redis на чтение и на запись, и отдает среднее время запросов

curl -X 'POST' 'http://127.0.0.1:8002/test/redis/string'
{"name":"redis string test, write / read, sec","size":20480,"write_duration":0.0002148306,"read_duration":0.0001277012}

curl -X 'POST' 'http://127.0.0.1:8002/test/redis/hset'
{"name":"redis hset test, write / read, sec","size":20480,"write_duration":0.0002046565,"read_duration":0.0001842852}

curl -X 'POST' 'http://127.0.0.1:8002/test/redis/zset'
{"name":"redis zset test, write / read, sec","size":20480,"write_duration":0.0002796277,"read_duration":0.0001878852}

curl -X 'POST' 'http://127.0.0.1:8002/test/redis/list'
{"name":"redis list test, write / read, sec","size":20480,"write_duration":0.0002038115,"read_duration":0.0001429195}

В процессе нескольких запусков тесты показали, что результаты имеют довольно большой разброс
В среднем, все запросы работают очень быстро и имеют сравнимые результаты 

```

### 3. Остановка
```
make compose-down
```
