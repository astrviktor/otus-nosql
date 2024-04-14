### 1. Подготовка данных
```
cd homework_02
git clone https://github.com/neelabalan/mongodb-sample-dataset.git
```

### 2. Запуск mongodb через docker-compose
```
make compose-up
```


### 2. Загрузка данных
```
docker exec -it mongodb bash 
cd /mongodb-sample-dataset 
sh script.sh localhost 27017 root example
exit
```

### 3. Запуск mongosh
```
make mongosh
```

### 4. Запросы
```
show dbs
use sample_airbnb

// информация по данным
show collections
db.stats()
db.listingsAndReviews.countDocuments()
db.listingsAndReviews.findOne()
db.listingsAndReviews.distinct('address.country')

// обновление Brazil -> Бразилия
db.listingsAndReviews.updateMany({ "address.country": "Brazil" }, { $set: { "address.country": "Бразилия" } }) 
db.listingsAndReviews.distinct('address.country')

exit
```

### 5. Остановка
```
make compose-down
rm -rf mongodb-sample-dataset
```