currencyservice-gen:
	protoc --proto_path=currencyservice/proto --go_out=currencyservice/pb --go_opt=paths=source_relative --go-grpc_out=currencyservice/pb --go-grpc_opt=paths=source_relative currencyservice/proto/*.proto

currencyservice-run:
	go run currencyservice/cmd/main.go

productCatalogService-gen:
	protoc --proto_path=productCatalogService/proto --go_out=productCatalogService/pb --go_opt=paths=source_relative --go-grpc_out=productCatalogService/pb --go-grpc_opt=paths=source_relative productCatalogService/proto/*.proto

productCatalogService-run:
	go run productCatalogService/cmd/main.go