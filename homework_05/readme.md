### 1. Запуск Clickhouse
```
make compose-up
```

### 2. Проверка таблиц

Использовался startup-скрипт, должны создаться таблицы:

- tutorial.hits_v1
- tutorial.visits_v1

```
docker exec -ti clickhouse clickhouse-client

clickhouse :) SHOW TABLES FROM tutorial;

SHOW TABLES FROM tutorial

Query id: 3c2fc6ed-69a7-4ef4-b9e0-6ebd404f0541

┌─name──────┐
│ hits_v1   │
│ visits_v1 │
└───────────┘

2 rows in set. Elapsed: 0.002 sec. 
```

### 3. Загрузка данных

Загрузка и извлечение табличных данных

```
curl https://datasets.clickhouse.com/hits/tsv/hits_v1.tsv.xz | unxz --threads=`nproc` > ./clickhouse/hits_v1.tsv
curl https://datasets.clickhouse.com/visits/tsv/visits_v1.tsv.xz | unxz --threads=`nproc` > ./clickhouse/visits_v1.tsv
```

Импорт данных

```
docker exec -ti clickhouse /bin/bash
clickhouse-client --query "INSERT INTO tutorial.hits_v1 FORMAT TSV" --max_insert_block_size=100000 < /clickhouse/hits_v1.tsv
clickhouse-client --query "INSERT INTO tutorial.visits_v1 FORMAT TSV" --max_insert_block_size=100000 < /clickhouse/visits_v1.tsv
exit
```

Проверка импорта

```
docker exec -ti clickhouse clickhouse-client

clickhouse :) SELECT COUNT(*) FROM tutorial.hits_v1;

SELECT COUNT(*)
FROM tutorial.hits_v1

Query id: c6469cd7-2429-4b7e-8886-cd5e969a2028

┌─count()─┐
│ 8873898 │
└─────────┘

1 row in set. Elapsed: 0.009 sec. 

clickhouse :) SELECT COUNT(*) FROM tutorial.visits_v1;

SELECT COUNT(*)
FROM tutorial.visits_v1

Query id: 03669662-726d-477a-89f6-c6b8b33d4684

┌─count()─┐
│ 1680609 │
└─────────┘

1 row in set. Elapsed: 0.013 sec. 
```

### 4. Выполнение запросов и скорость выполнения
```
SELECT
    StartURL AS URL,
    AVG(Duration) AS AvgDuration
FROM tutorial.visits_v1
WHERE StartDate BETWEEN '2014-03-23' AND '2014-03-30'
GROUP BY URL
ORDER BY AvgDuration DESC
LIMIT 10;

10 rows in set. Elapsed: 0.078 sec. Processed 1.47 million rows, 114.72 MB (18.93 million rows/s., 1.47 GB/s.)
Peak memory usage: 54.12 MiB.

SELECT
    sum(Sign) AS visits,
    sumIf(Sign, has(Goals.ID, 1105530)) AS goal_visits,
    (100. * goal_visits) / visits AS goal_percent
FROM tutorial.visits_v1
WHERE (CounterID = 912887) AND (toYYYYMM(StartDate) = 201403);

1 row in set. Elapsed: 0.011 sec. Processed 46.57 thousand rows, 1.25 MB (4.37 million rows/s., 117.03 MB/s.)
Peak memory usage: 1.74 MiB.

```

### 5. Остановка Clickhouse
```
make compose-down
```