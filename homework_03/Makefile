compose-up:
	docker-compose -f docker-compose.yml up --remove-orphans -d

compose-down:
	docker-compose -f docker-compose.yml down


mongo-configsvr-1:
	docker exec -it mongo-configsvr-1 mongosh "mongodb://localhost:40001"

mongo-shard-1-rs-1:
	docker exec -it mongo-shard-1-rs-1 mongosh "mongodb://localhost:40011"

mongo-shard-2-rs-1:
	docker exec -it mongo-shard-2-rs-1 mongosh "mongodb://localhost:40021"


mongos-shard:
	docker exec -it mongos-shard mongosh "mongodb://localhost:27017"


clondata:
	git clone https://github.com/neelabalan/mongodb-sample-dataset.git

