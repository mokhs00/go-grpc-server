.PHONY: run-mysql
## run-mysql : run mysql docker
run-mysql:
	docker run --platform linux/amd64 \
		--name go-rpc-server-mysql \
		-e MYSQL_ROOT_PASSWORD=1234 \
		-e MYSQL_DATABASE=go_grpc_server \
		-p 3306:3306 \
		-d mysql