compose-up:
	docker-compose -f build/docker-compose.yml up --remove-orphans -d

compose-down:
	docker-compose -f build/docker-compose.yml down

docker-build-service:
	docker build --no-cache \
		--platform linux/amd64 \
		-t "astrviktor/service:1.0.0" \
		-f ./build/service/Dockerfile .

docker-push:
	docker push astrviktor/service:1.0.0
