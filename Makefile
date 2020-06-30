gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.

clean:
	rm proto/*.pb.go

run:
	go run .