server-local:
	go run cmd/sorting_server/main.go
.PHONY: server-local

server-container-start:
	docker build -t suhail12/sorting_server:latest .
	docker run -p 8000:8000 -d --name sorting_server suhail12/sorting_server:latest
.PHONY: server-container-start

server-container-stop:
	docker container stop sorting_server
	docker rm -f sorting_server
.PHONY: server-container-stop
