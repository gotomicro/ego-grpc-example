APP_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SCRIPT_PATH:=$(APP_PATH)/scripts

SERVER_APP_NAME:=ego-grpc-server
SERVER_COMPILE_OUT:=$(APP_PATH)/server/bin/$(SERVER_APP_NAME)

server_local:export EGO_DEBUG=true
server_local:
	@cd $(APP_PATH) && go run server/main.go --config=config/local.toml

server_dev:export EGO_DEBUG=true
server_dev:server_build server_devrun

server_devrun:export EGO_DEBUG=true
server_devrun:
	@cd $(APP_PATH) && ./server/bin/ego-grpc-server --config='etcd://127.0.0.1:2379/?configKey=/ego/config/ego-grpc-server&configType=toml'

server_build:
	@echo ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>making build ego server gRPC<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
	@chmod +x $(SCRIPT_PATH)/build/*.sh
	@cd $(APP_PATH)/server && $(SCRIPT_PATH)/build/gobuild.sh $(SERVER_APP_NAME) $(SERVER_COMPILE_OUT)
