run:
	go run main.go

build:
	go build -o main main.go

docker-up:
	docker-compose up  --build -docker

docker-down:
	docker-compose down