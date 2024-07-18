### 1. Запуск neo4j
```
make compose-up

http://127.0.0.1:7474/browser/
```

### 2. Создание нод для туроператоров
```
create (:Operator {title:'TUI'})
create (:Operator {title:'Tez Tour'})
create (:Operator {title:'Coral Travel'})
create (:Operator {title:'Pegas Touristik'})
create (:Operator {title:'Anex Tour'})

match (n) return n
```

### 3. Создание нод для направлений, создание связей между странами и местами
```
Турция, Алания (CLUB HOTEL ANJELIQ)
Турция, Стамбул (DARU SULTAN)
Таиланд, Паттайя (WAY HOTEL)
Таиланд, Пхукет (SILK HILL HOTEL)
Куба, Гавана (PLAZA)
Куба, Варадеро (MELIA MARINA VARADERO)
Вьетнам, Нячанг (DAPHOVINA HOTEL)
Кипр, Пафос (SUNNY HILL HOTELS APTS)
ОАЭ, Дубай (NOVOTEL WORLD TRADE CENTER)
Египет, Сафага (IMPERIAL SHAMS ABU SOMA)

create (:Country {title:'Турция'})
create (:Country {title:'Таиланд'})
create (:Country {title:'Куба'})
create (:Country {title:'Вьетнам'})
create (:Country {title:'Кипр'})
create (:Country {title:'ОАЭ'})
create (:Country {title:'Египет'})

create (:Place {hotel:'CLUB HOTEL ANJELIQ'})
create (:Place {hotel:'DARU SULTAN'})
create (:Place {hotel:'WAY HOTEL'})
create (:Place {hotel:'SILK HILL HOTEL'})
create (:Place {hotel:'PLAZA'})
create (:Place {hotel:'MELIA MARINA VARADERO'})
create (:Place {hotel:'DAPHOVINA HOTEL'})
create (:Place {hotel:'SUNNY HILL HOTELS APTS'})
create (:Place {hotel:'NOVOTEL WORLD TRADE CENTER'})
create (:Place {hotel:'IMPERIAL SHAMS ABU SOMA'})

match (country:Country {title:'Турция'})
match (place:Place {hotel:'CLUB HOTEL ANJELIQ'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Турция'})
match (place:Place {hotel:'DARU SULTAN'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Таиланд'})
match (place:Place {hotel:'WAY HOTEL'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Таиланд'})
match (place:Place {hotel:'SILK HILL HOTEL'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Куба'})
match (place:Place {hotel:'PLAZA'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Куба'})
match (place:Place {hotel:'MELIA MARINA VARADERO'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Вьетнам'})
match (place:Place {hotel:'DAPHOVINA HOTEL'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Кипр'})
match (place:Place {hotel:'SUNNY HILL HOTELS APTS'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'ОАЭ'})
match (place:Place {hotel:'NOVOTEL WORLD TRADE CENTER'})
create (country) -[:LOCATED]-> (place)

match (country:Country {title:'Египет'})
match (place:Place {hotel:'IMPERIAL SHAMS ABU SOMA'})
create (country) -[:LOCATED]-> (place)

match (n)-[]-(m) return n,m
```

### 4. Создание нод для ближайших к туристическим локациям городов
```
create (:City {title:'Алания'})
create (:City {title:'Стамбул'})
create (:City {title:'Паттайя'})
create (:City {title:'Пхукет'})
create (:City {title:'Гавана'})
create (:City {title:'Варадеро'})
create (:City {title:'Нячанг'})
create (:City {title:'Пафос'})
create (:City {title:'Дубай'})
create (:City {title:'Сафага'})

match (place:Place {hotel:'CLUB HOTEL ANJELIQ'})
match (city:City {title:'Алания'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'DARU SULTAN'})
match (city:City {title:'Стамбул'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'WAY HOTEL'})
match (city:City {title:'Паттайя'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'SILK HILL HOTEL'})
match (city:City {title:'Пхукет'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'PLAZA'})
match (city:City {title:'Гавана'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'MELIA MARINA VARADERO'})
match (city:City {title:'Варадеро'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'DAPHOVINA HOTEL'})
match (city:City {title:'Нячанг'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'SUNNY HILL HOTELS APTS'})
match (city:City {title:'Пафос'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'NOVOTEL WORLD TRADE CENTER'})
match (city:City {title:'Дубай'})
merge (place) -[:LOCATED]-> (city)

match (place:Place {hotel:'IMPERIAL SHAMS ABU SOMA'})
match (city:City {title:'Сафага'})
merge (place) -[:LOCATED]-> (city)

match (n)-[]-(m) return n,m
```

### 5. Создание связей между городами, охарактеризовав каждый маршрут видом транспорта
```
match (alanya:City {title:'Алания'})
match (stanbul:City {title:'Стамбул'})
match (pattaya:City {title:'Паттайя'})
match (phuket:City {title:'Пхукет'})
match (havana:City {title:'Гавана'})
match (varadero:City {title:'Варадеро'})
match (nhatrang:City {title:'Нячанг'})
match (pathos:City {title:'Пафос'})
match (dubai:City {title:'Дубай'})
match (safaga:City {title:'Сафага'})

merge (alanya) -[:ROUTE {transport:'самолет',type:'воздушный'}]-> (stanbul)
...

match (alanya:City {title:'Алания'})
match (stanbul:City {title:'Стамбул'})
merge (alanya) -[:ROUTE {transport:'поезд',type:'наземный'}]-> (stanbul)

match (pattaya:City {title:'Паттайя'})
match (phuket:City {title:'Пхукет'})
merge (pattaya) -[:ROUTE {transport:'поезд',type:'наземный'}]-> (phuket)

match (havana:City {title:'Гавана'})
match (varadero:City {title:'Варадеро'})
merge (havana) -[:ROUTE {transport:'поезд',type:'наземный'}]-> (varadero)

match (n)-[:ROUTE]-(m) return n,m
```

### 6. Запрос, который выводит направления (со всеми промежуточными точками), доступные только наземным транспортом
```
match (p1:Place) -[:LOCATED]-> (c1:City) -[:ROUTE {type:'наземный'}]-> (c2:City) <-[:LOCATED]- (p2:Place) return p1, c1, c2, p2
```

### 7. План запроса (EXPLAIN и PROFILE)
```
explain match (p1:Place) -[:LOCATED]-> (c1:City) -[:ROUTE {type:'наземный'}]-> (c2:City) <-[:LOCATED]- (p2:Place) return p1, c1, c2, p2

profile match (p1:Place) -[:LOCATED]-> (c1:City) -[:ROUTE {type:'наземный'}]-> (c2:City) <-[:LOCATED]- (p2:Place) return p1, c1, c2, p2
Cypher version: , planner: COST, runtime: PIPELINED. 130 total db hits in 54 ms.
```

### 8. Добавление индексов, повторная проверка планов запросов
```
create index on for (r:ROUTE) on (r.type)

profile match (p1:Place) -[:LOCATED]-> (c1:City) -[:ROUTE {type:'наземный'}]-> (c2:City) <-[:LOCATED]- (p2:Place) return p1, c1, c2, p2
Cypher version: , planner: COST, runtime: PIPELINED. 130 total db hits in 2 ms.

Уменьшение с 54 ms до 2 ms
```

### 9. Остановка neo4j
```
make compose-down
```
