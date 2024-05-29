### 1. Запуск Cassandra
```
make compose-up
```

### 2. Проверка
```
docker exec -it cassandra-1 nodetool status

Datacenter: my-datacenter-1
===========================
Status=Up/Down
|/ State=Normal/Leaving/Joining/Moving
--  Address     Load       Tokens  Owns (effective)  Host ID                               Rack 
UN  172.23.0.3  75.2 KiB   16      59.3%             51537332-684b-4545-ba30-ddffa04d8255  rack1
UN  172.23.0.4  75.27 KiB  16      76.0%             b11f53bb-17bf-4cd6-97f6-d4bc510b5d99  rack1
UN  172.23.0.2  109.4 KiB  16      64.7%             95bab0f0-7900-4324-9788-15fc2a0cf6f4  rack1


docker exec -it cassandra-1 nodetool info

ID                     : 95bab0f0-7900-4324-9788-15fc2a0cf6f4
Gossip active          : true
Native Transport active: true
Load                   : 109.4 KiB
Generation No          : 1716960135
Uptime (seconds)       : 355
Heap Memory (MB)       : 573.30 / 3916.00
Off Heap Memory (MB)   : 0.00
Data Center            : my-datacenter-1
Rack                   : rack1
Exceptions             : 0
Key Cache              : entries 2, size 168 bytes, capacity 100 MiB, 78 hits, 85 requests, 0.918 recent hit rate, 14400 save period in seconds
Row Cache              : entries 0, size 0 bytes, capacity 0 bytes, 0 hits, 0 requests, NaN recent hit rate, 0 save period in seconds
Counter Cache          : entries 0, size 0 bytes, capacity 50 MiB, 0 hits, 0 requests, NaN recent hit rate, 7200 save period in seconds
Network Cache          : size 8 MiB, overflow size: 0 bytes, capacity 128 MiB
Percent Repaired       : 100.0%
Token                  : (invoke with -T/--tokens to see all 16 tokens)


docker exec -it cassandra-1 cqlsh

Connected to my-cluster at 127.0.0.1:9042
[cqlsh 6.1.0 | Cassandra 4.1.5 | CQL spec 3.4.6 | Native protocol v5]
Use HELP for help.
cqlsh> 

```

### 3. Создание таблиц

```
docker exec -it cassandra-1 cqlsh

CREATE KEYSPACE test  
  WITH REPLICATION = { 
   'class' : 'NetworkTopologyStrategy', 
   'my-datacenter-1' : 3 
  };
  
  
DESCRIBE keyspaces;

SELECT * FROM system_schema.keyspaces;

CREATE TABLE test.customers (  
  id UUID,
  title text,
  first_name text,
  last_name text,
  language text,
  PRIMARY KEY (id, language)
)
WITH CLUSTERING ORDER BY (language ASC);

CREATE TABLE test.addresses (  
  id UUID PRIMARY KEY,
  postal_code text,
  region text,
  city text,
  street text,
  building_number text
);

DESCRIBE tables;
SELECT * FROM system_schema.tables WHERE keyspace_name = 'test';

```

### 4. Загрузка данных

```
docker cp ./data/customers.csv cassandra-1:/tmp/customers.csv
docker cp ./data/addresses.csv cassandra-1:/tmp/addresses.csv

docker exec -ti cassandra-1 /bin/bash
ls /tmp | grep ".csv"

docker exec -it cassandra-1 cqlsh

COPY test.customers (id,title,first_name,last_name,language) FROM '/tmp/customers.csv' WITH DELIMITER=',' AND HEADER=TRUE;

    Starting copy of test.customers with columns [id, title, first_name, last_name, language].
    Processed: 1154 rows; Rate:    1938 rows/s; Avg. rate:    2883 rows/s
    1154 rows imported from 1 files in 0.400 seconds (0 skipped).

SELECT count(*) FROM test.customers;

     count
    -------
    1154

    (1 rows)

COPY test.addresses (id,postal_code,region,city,street,building_number) FROM '/tmp/addresses.csv' WITH DELIMITER=',' AND HEADER=TRUE;

    Starting copy of test.addresses with columns [id, postal_code, region, city, street, building_number].
    Processed: 1154 rows; Rate:    2121 rows/s; Avg. rate:    3086 rows/s
    1154 rows imported from 1 files in 0.374 seconds (0 skipped).

SELECT count(*) FROM test.addresses;

    count
    -------
    1154

    (1 rows)

```

### 5. Запросы

```
SELECT count(*) FROM test.customers WHERE language='EN' ALLOW FILTERING;
cqlsh> SELECT count(*) FROM test.customers WHERE language='EN' ALLOW FILTERING;

 count
-------
   488

(1 rows)

Warnings :
Aggregation query used without partition key

SELECT * FROM test.customers WHERE id=02BDAF3B-D712-B7AD-B958-D3AE0D65F358;
cqlsh> SELECT * FROM test.customers WHERE id=02BDAF3B-D712-B7AD-B958-D3AE0D65F358;

 id                                   | language | first_name | last_name | title
--------------------------------------+----------+------------+-----------+-------
 02bdaf3b-d712-b7ad-b958-d3ae0d65f358 |       EN |   Rosemary |    Bailey |   Mrs

(1 rows)


SELECT * FROM test.addresses WHERE city='Berlin' ALLOW FILTERING;
cqlsh> SELECT * FROM test.addresses WHERE city='Berlin' ALLOW FILTERING;

 id                                   | building_number | city   | postal_code | region | street
--------------------------------------+-----------------+--------+-------------+--------+------------------------
 17b71247-40c8-69f1-9f51-6c6aec468581 |             48a | Berlin |       13403 |   null |            Antonienstr
 2cb9f874-2ef9-d785-121c-1cf6f961575a |            null | Berlin |       10319 |   null |           Balatonstr 2
 72bf5d03-781a-e788-1d16-943e3dc1f44e |            null | Berlin |       12527 |   null |             Dahmestr 1
 6d3ada82-1ad7-1561-b9db-f6d81f703fbd |            null | Berlin |       10119 |   null |         Brunnenstr 196
 b7b19f17-7f11-d8e8-b09f-6e76aafccbad |            null | Berlin |       10585 |   null |  Richard wagner str 44
 caa3063f-4f2e-b7b2-6593-b72b52bf2bc3 |            null | Berlin |       10115 |   null |       Invalidenstr 150
 130bc770-193a-c94a-6097-1bf5d479ea3d |            null | Berlin |       10779 |   null |           Stübbenstr 7
 615620e5-c209-4e33-8707-c77e7eb3723b |            null | Berlin |       13409 |   null |             Am stand 5
 f6107b1e-cdae-e404-7b7b-d08dbbd4b218 |            null | Berlin |       10365 |   null | Wilhelm-guddorf-str 10
 ee04f74d-620e-f045-ed0d-dac6fe18d86d |            null | Berlin |       12249 |   null |    Am gemeindepark 40a
 d742d26e-d3ad-8633-8422-551cb1b6c9e7 |            null | Berlin |       13189 |   null |             Granitzstr
 9d944b8a-c758-a09d-74d9-e535cbd1dabb |          str 44 | Berlin |       10585 |   null |         Richard wagner
 514672db-7d1b-29cc-31cf-3c37ab4e4ee0 |            null | Berlin |       13125 |   null |              Straße 67
 863d81f9-570e-2b83-c3ba-d95b42ae182f |            null | Berlin |       10249 |   null |        Pufendorfstr 6a

(14 rows)

Warnings :
Read 1154 live rows and 1700 tombstone cells for query SELECT * FROM test.addresses WHERE city = 'Berlin' LIMIT 100 ALLOW FILTERING; token 9206687968599054459 (see tombstone_warn_threshold)

```

### 6. Вторичный индекс

```
CREATE INDEX city_idx ON test.addresses (city);

SELECT * FROM test.addresses WHERE city='Berlin';
cqlsh> SELECT * FROM test.addresses WHERE city='Berlin';

 id                                   | building_number | city   | postal_code | region | street
--------------------------------------+-----------------+--------+-------------+--------+------------------------
 17b71247-40c8-69f1-9f51-6c6aec468581 |             48a | Berlin |       13403 |   null |            Antonienstr
 2cb9f874-2ef9-d785-121c-1cf6f961575a |            null | Berlin |       10319 |   null |           Balatonstr 2
 72bf5d03-781a-e788-1d16-943e3dc1f44e |            null | Berlin |       12527 |   null |             Dahmestr 1
 6d3ada82-1ad7-1561-b9db-f6d81f703fbd |            null | Berlin |       10119 |   null |         Brunnenstr 196
 b7b19f17-7f11-d8e8-b09f-6e76aafccbad |            null | Berlin |       10585 |   null |  Richard wagner str 44
 caa3063f-4f2e-b7b2-6593-b72b52bf2bc3 |            null | Berlin |       10115 |   null |       Invalidenstr 150
 130bc770-193a-c94a-6097-1bf5d479ea3d |            null | Berlin |       10779 |   null |           Stübbenstr 7
 615620e5-c209-4e33-8707-c77e7eb3723b |            null | Berlin |       13409 |   null |             Am stand 5
 f6107b1e-cdae-e404-7b7b-d08dbbd4b218 |            null | Berlin |       10365 |   null | Wilhelm-guddorf-str 10
 ee04f74d-620e-f045-ed0d-dac6fe18d86d |            null | Berlin |       12249 |   null |    Am gemeindepark 40a
 d742d26e-d3ad-8633-8422-551cb1b6c9e7 |            null | Berlin |       13189 |   null |             Granitzstr
 9d944b8a-c758-a09d-74d9-e535cbd1dabb |          str 44 | Berlin |       10585 |   null |         Richard wagner
 514672db-7d1b-29cc-31cf-3c37ab4e4ee0 |            null | Berlin |       13125 |   null |              Straße 67
 863d81f9-570e-2b83-c3ba-d95b42ae182f |            null | Berlin |       10249 |   null |        Pufendorfstr 6a

```

### 7. Остановка Cassandra
```
make compose-down
```