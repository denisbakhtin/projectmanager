client_watch:
	@echo "Running webpack watch"
	@webpack --watch --mode=development

#Basic makefile

default: build

build: clean vet
	@echo "Building application"
	@go build -o projectmanager-go

doc:
	@godoc -http=:6060 -index

lint:
	@golint ./...

debug_server: 
	@fresh
debug_assets:
	@webpack --watch --mode=development

#run 'make -j2 debug' to launch both servers in parallel
debug: clean debug_server debug_assets 

run: build
	./projectmanager-go

test:
	@go test ./...

vet:
	@go vet ./...

clean:
	@echo "Cleaning binary"
	@rm -f ./projectmanager-go

stop: 
	@echo "Stopping projectmanager service"
	@sudo systemctl stop projectmanager

start:
	@echo "Starting projectmanager service"
	@sudo systemctl start projectmanager

pull:
	@echo "Pulling origin"
	@git pull origin master

pull_restart: stop pull clean build start

deploy:
	ansible-playbook deploy.yml -K
