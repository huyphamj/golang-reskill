gen-files:
	protoc -Icommon/proto --go_out=. \
	--go_opt=module=github.com/huyphamj/flight-management \
	--go-grpc_out=. \
	--go-grpc_opt=module=github.com/huyphamj/flight-management \
	common/proto/*.proto
	#go build -o bin/greet/server ./greet/server
	#go build -o bin/greet/client ./greet/client