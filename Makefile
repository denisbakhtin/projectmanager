#use fresh to watch backend.
#this is for front-end
watch:
	@echo "Running webpack watch"
	@webpack --watch --mode=development

#all-in-one ansible command for deployment
deploy: build
	ansible-playbook deploy.yml -K

build: clean vet
	@echo "Building assets"
	@webpack
	@echo "Building application"
	CGO_ENABLED=0 go build -o projectmanager-go
	CGO_ENABLED=0 go build -o periodic-go cmd/periodical/main.go

vet:
	@go vet ./...

clean:
	@echo "Cleaning binary"
	@rm -f ./projectmanager-go
	@rm -f ./periodic-go