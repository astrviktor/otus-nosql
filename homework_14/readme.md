### 1. Создание ClickHouse в Yandex Cloud
```
1. Зарегистрироваться в Yandex Cloud (при привязке карты дают 4000 руб для тестов)
2. Создать Managed Service for ClickHouse

Подключение к ClickHouse можно настроить из интернета или из ВМ в Yandex Cloud
Проще из ВМ в Yandex Cloud

3. Локально установить и настроить yandex cli
4. Создать ВМ в Yandex Cloud

yc compute instance create \
--name ubuntu \
--cores=2 \
--core-fraction=5 \
--memory=512MB \
--preemptible \
--create-boot-disk image-id=fd87tirk5i8vitv9uuo1,size=20GB \
--network-interface subnet-name=default-ru-central1-a,nat-ip-version=ipv4,ipv4-address=10.128.0.10 \
--zone=ru-central1-a \
--metadata serial-port-enable=1 \
--metadata ssh-keys="ubuntu:ssh-rsa ..."
 

5. Узнать координаты ClickHouse внутри сети

yc managed-clickhouse host list --cluster-name=clickhouse-test
+-------------------------------------------+----------------------+------------+------------+--------+---------------+-----------+
|                   NAME                    |      CLUSTER ID      |    TYPE    | SHARD NAME | HEALTH |    ZONE ID    | PUBLIC IP |
+-------------------------------------------+----------------------+------------+------------+--------+---------------+-----------+
| rc1d-es1icnjou95a74hl.mdb.yandexcloud.net |         ***          | CLICKHOUSE | shard1     | ALIVE  | ru-central1-d | false     |
+-------------------------------------------+----------------------+------------+------------+--------+---------------+-----------+

6. Установить на ВМ clickhouse-client

7. Проверить доступ

clickhouse-client --host rc1d-es1icnjou95a74hl.mdb.yandexcloud.net \
                  --user admin \
                  --database default \
                  --port 9000 \
                  --ask-password
 
8. Конфигурация ClickHouse

yc managed-clickhouse cluster get clickhouse-test
id: c9q8qasgt537vmo83g21
folder_id: b1g9ns51mcm65pgf1glv
created_at: "2024-09-01T06:33:30.334572Z"
name: clickhouse-test
environment: PRODUCTION
monitoring:
  - name: Console
    description: Console charts
    link: https://console.cloud.yandex.ru/folders/b1g9ns51mcm65pgf1glv/managed-clickhouse/cluster/c9q8qasgt537vmo83g21/monitoring
config:
  version: "24.3"
  clickhouse:
    config:
      user_config:
        merge_tree:
          number_of_free_entries_in_pool_to_lower_max_size_of_merge: "8"
        kafka: {}
        rabbitmq: {}
        query_cache: {}
    resources:
      resource_preset_id: s3-c2-m8
      disk_size: "10737418240"
      disk_type_id: network-ssd
  zookeeper:
    resources: {}
  backup_window_start:
    hours: 22
  access:
    web_sql: true
    yandex_query: true
  cloud_storage: {}
  sql_database_management: true
  sql_user_management: true
  embedded_keeper: false
  backup_retain_period_days: "5"
network_id: enp6fbmcf5po929pgftb
health: ALIVE
status: RUNNING
maintenance_window:
  anytime: {}
 
```

### 2. Загрузка данных

```
Скопировать на ВМ clickhouse.sql, запустить создание таблиц:

cat clickhouse.sql | clickhouse-client

Загрузка и извлечение табличных данных:

curl https://datasets.clickhouse.com/hits/tsv/hits_v1.tsv.xz | unxz --threads=`nproc` > /home/ubuntu/hits_v1.tsv
curl https://datasets.clickhouse.com/visits/tsv/visits_v1.tsv.xz | unxz --threads=`nproc` > /home/ubuntu/visits_v1.tsv

Импорт данных:

clickhouse-client --query "INSERT INTO tutorial.hits_v1 FORMAT TSV" --max_insert_block_size=100000 < /home/ubuntu/hits_v1.tsv
clickhouse-client --query "INSERT INTO tutorial.visits_v1 FORMAT TSV" --max_insert_block_size=100000 < /home/ubuntu/visits_v1.tsv

Проверка импорта:

clickhouse :) SELECT COUNT(*) FROM tutorial.hits_v1;

SELECT COUNT(*)
FROM tutorial.hits_v1

Query id: c6469cd7-2429-4b7e-8886-cd5e969a2028

┌─count()─┐
│ 8873898 │
└─────────┘

clickhouse :) SELECT COUNT(*) FROM tutorial.visits_v1;

SELECT COUNT(*)
FROM tutorial.visits_v1

Query id: 03669662-726d-477a-89f6-c6b8b33d4684

┌─count()─┐
│ 1680609 │
└─────────┘

Данные успешно загружены
```

### 3. Выполнение запросов и скорость выполнения локально в docker (из ДЗ 5)
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


### 4. Выполнение запросов и скорость выполнения в облаке
```
SELECT
    StartURL AS URL,
    AVG(Duration) AS AvgDuration
FROM tutorial.visits_v1
WHERE StartDate BETWEEN '2014-03-23' AND '2014-03-30'
GROUP BY URL
ORDER BY AvgDuration DESC
LIMIT 10;

10 rows in set. Elapsed: 0.202 sec. Processed 1.47 million rows, 114.37 MB (7.31 million rows/s., 566.96 MB/s.)
Peak memory usage: 35.63 MiB.

SELECT
    sum(Sign) AS visits,
    sumIf(Sign, has(Goals.ID, 1105530)) AS goal_visits,
    (100. * goal_visits) / visits AS goal_percent
FROM tutorial.visits_v1
WHERE (CounterID = 912887) AND (toYYYYMM(StartDate) = 201403);

1 row in set. Elapsed: 0.012 sec. Processed 46.57 thousand rows, 1.25 MB (4.04 million rows/s., 108.32 MB/s.)
Peak memory usage: 83.33 KiB.

```

### 5. Выводы

В облаке получены результаты ниже, чем локально

Может быть связано с конфигурацией в облаке, и с тем, что локально используется M2 SSD