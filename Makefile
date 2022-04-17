.PHONY: run-mysql
# run-mysql : run mysql docker
run-mysql:
	docker run --platform linux/amd64 \
		--name go-rpc-server-mysql \
		-e MYSQL_ROOT_PASSWORD=1234 \
		-e MYSQL_DATABASE=go_grpc_server \
		-p 3306:3306 \
		-d mysql

.PHONY: generate-db-models
# generate-db-models : generate code for db models
generate-db-models:
	@go install github.com/volatiletech/sqlboiler/v4@v4.7.1
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.7.1
	sqlboiler --templates ./server/db/templates \
		--templates $$(go env GOPATH)/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.7.1/templates/main \
		--wipe --no-tests --no-auto-timestamps -p mysql -o ./server/db/mysql mysql
