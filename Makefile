cognito:
	docker run --rm --name cognito -p 9229:9229 -v ${PWD}/data:/app/.cognito jagregory/cognito-local:latest
