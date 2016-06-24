protoc --proto_path ../../../ -I=./proto/ --go_out=plugins=grpc:./proto proto/monitor.proto
