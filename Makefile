# go mod tidy
.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

# go mod vendor
.PHONY: vendor
vendor:
	go mod vendor
# generate grpc
.PHONY: grpcgen
grpcgen:
	protoc --go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
		proto/greet/v1/greet.proto
