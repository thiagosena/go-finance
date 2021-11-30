build-image:
	docker build -t thiagosena/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

lint:
	golint ./...
	go fmt -n ./...

u-test:
	go test ./...