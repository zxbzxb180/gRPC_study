gen:
	protoc --proto_path=proto --go-grpc_opt=require_unimplemented_servers=false --go_out=pb --go_opt=paths=source_relative  --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go