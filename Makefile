#use fresh to watch backend.
#this is for front-end
watch:
	@echo "Running webpack watch"
	@webpack --watch --mode=development

#all-in-one ansible command for deployment
deploy:
	ansible-playbook deploy.yml -K
