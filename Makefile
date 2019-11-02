all:
	protoc game.proto --js_out=import_style:js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js --go_out=plugins=grpc:go
