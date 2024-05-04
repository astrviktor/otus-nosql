### 1. Запуск mongodb через docker-compose
```
make compose-up
```

### 2. Инициализация
```
make mongo-configsvr-1

// Инциализируем конфиг реплику
rs.initiate({
  "_id" : "config-replica-set", 
  members : [
    {"_id" : 0, host : "mongo-configsvr-1:40001"},
    {"_id" : 1, host : "mongo-configsvr-2:40002"},
    {"_id" : 2, host : "mongo-configsvr-3:40003"}
  ]
});

rs.status();

exit

// Инициализируем шардинги

make mongo-shard-1-rs-1

rs.initiate({
  "_id" : "shard-replica-set-1", 
  members : [
    {"_id" : 0, host : "mongo-shard-1-rs-1:40011"},
    {"_id" : 1, host : "mongo-shard-1-rs-2:40012"},
    {"_id" : 2, host : "mongo-shard-1-rs-3:40013"}
  ]
});

rs.status();

exit

make mongo-shard-2-rs-1

rs.initiate({
  "_id" : "shard-replica-set-2", 
  members : [
    {"_id" : 0, host : "mongo-shard-2-rs-1:40021"},
    {"_id" : 1, host : "mongo-shard-2-rs-2:40022"},
    {"_id" : 2, host : "mongo-shard-2-rs-3:40023"}
  ]
});

rs.status();

exit

// Добавляем шардинги

make mongos-shard

sh.addShard("shard-replica-set-1/mongo-shard-1-rs-1:40011,mongo-shard-1-rs-2:40012,mongo-shard-1-rs-3:40013");
sh.addShard("shard-replica-set-2/mongo-shard-2-rs-1:40021,mongo-shard-2-rs-2:40022,mongo-shard-2-rs-3:40023");

sh.status();

exit
```

### 3. Проверка шардирования
```
make mongos-shard

use bank;
sh.enableSharding("bank");

for (var i=0; i<10000; i++) { db.tickets.insertOne({name: "Max ammout of cost tickets", amount: Math.random()*100}) }
db.tickets.countDocuments() // 100

sh.status()
db.tickets.createIndex({amount: 1})

use admin
db.runCommand({shardCollection: "bank.tickets", key: {amount: 1}})

sh.balancerCollectionStatus("bank.tickets")
sh.splitFind( "bank.tickets", { "amount": "50" } )

use config
db.settings.updateOne(
   { _id: "chunksize" },
   { $set: { _id: "chunksize", value: 1 } },
   { upsert: true }
);

use bank
sh.status()
db.tickets.getShardDistribution();

exit
```

### 4. Проверка шардирования на реальных данных
```
docker exec -it mongos-shard bash

mongosh
sh.enableSharding("test");
db.movies.createIndex({year: 1});

use config;
db.settings.updateOne(
   { _id: "chunksize" },
   { $set: { _id: "chunksize", value: 1 } },
   { upsert: true }
);

use admin
db.runCommand({shardCollection: "test.movies", key: {year: 1}})

exit

mongoimport --host localhost --port 27017 --db "test" --collection "movies" --file /data/movies.json

mongosh
db.movies.getShardDistribution();

// вывод db.movies.getShardDistribution()

    Shard shard-replica-set-1 at shard-replica-set-1/mongo-shard-1-rs-1:40011,mongo-shard-1-rs-2:40012,mongo-shard-1-rs-3:40013
    {
      data: '35.87MiB',
      docs: 23539,
      chunks: 1,
      'estimated data per chunk': '35.87MiB',
      'estimated docs per chunk': 23539
    }
    ---
    Shard shard-replica-set-2 at shard-replica-set-2/mongo-shard-2-rs-1:40021,mongo-shard-2-rs-2:40022,mongo-shard-2-rs-3:40023
    {
      data: '16.73MiB',
      docs: 11044,
      chunks: 20,
      'estimated data per chunk': '857KiB',
      'estimated docs per chunk': 552
    }
    ---
    Totals
    {
      data: '52.61MiB',
      docs: 34583,
      chunks: 21,
      'Shard shard-replica-set-1': [
        '68.18 % data',
        '68.06 % docs in cluster',
        '1KiB avg obj size on shard'
      ],
      'Shard shard-replica-set-2': [
        '31.81 % data',
        '31.93 % docs in cluster',
        '1KiB avg obj size on shard'
      ]
    }


exit
```

### 5. Отказ инстансов
```
// Отказ SECONDARY
docker-compose stop mongo-shard-1-rs-2

docker exec -it mongo-shard-1-rs-1 mongosh "mongodb://localhost:40011"
rs.status()
// SECONDARY (not reachable/healthy), PRIMARY не поменялся
exit

docker-compose start mongo-shard-1-rs-2

// Отказ PRIMARY
docker-compose stop mongo-shard-1-rs-1

docker exec -it mongo-shard-1-rs-2 mongosh "mongodb://localhost:40012"
rs.status()
// PRIMARY стал mongo-shard-1-rs-2
exit

docker-compose start mongo-shard-1-rs-1
```

### 6. Аутентификация и многоролевой доступ
```
use admin;

db.getUsers();
db.getRoles();

db.system.roles.find();
db.system.users.find();

// Создаем root админа
db.createUser({
    user: "root",
    pwd: "strictRootPassword",
    roles: [
        { role: "root", db: "admin" }
    ]
})

// Создаем роль супер-админа
db.createRole({
    role: "superRoot",
    privileges:[
        { resource: {anyResource:true}, actions: ["anyAction"]}
    ],
    roles:[]
})

// Создаем пользователя 
db.createUser({
    user: "readWriteUser",
    pwd: "strictReadWritePassword",
    roles: [
        { role: "readWrite", db: "test" }
    ]
})

// Включаем авторизацию и рестарт mongod
nano /etc/mongod.conf
security:
  authorization: enabled

sudo service mongod restart

// Через Docker добавить к контейнерам
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: example

// Подключаемся
mongosh
mongosh -u root -p --authenticationDatabase admin
use test
show dbs

exit
```

### 7. Остановка
```
make compose-down
```