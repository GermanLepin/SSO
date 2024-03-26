# generate code from proto files
generate_proto_files:
	@echo "Starting to generate proto files..."
	protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@echo "Generating proto files was finished successfully"

