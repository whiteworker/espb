protoc --proto_path=D:\Users\zlbai\go/src --proto_path=D:\code\Test\ginadmin\pkg\third_party --proto_path=D:\code\Test\ginadmin\internal\app\foroutapi\pb --go_out=plugins=grpc:. *.proto
xcopy "*" "../*" /S/F/I/R/Y
del *.json &&del *.go && cd ../&& del *.bat&& del *.proto &&cd pb