protoc --proto_path ../../../ -I=./monitorproto/ --go_out=plugins=grpc:./monitorproto monitorproto/monitor.proto
