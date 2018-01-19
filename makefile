build:
	docker build -t gorestapi:latest .
run:
	docker run -d --rm --name gorestapi -p 8080:8080 gorestapi:latest
clean:
	docker stop gorestapi || true
	docker rm gorestapi || true
	docker rmi gorestapi:latest
