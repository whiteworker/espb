for  /r  %%i  in  (*.proto)  do  (echo  generating  %%i  &  kratos tool protoc protoc --proto_path=D:\Users\zlbai\go/src --proto_path=D:\code\Test\ginadmin\pkg\third_party --proto_path=D:\code\Test\ginadmin\internal\app\foroutapi\pb --go_out=plugins=grpc:. %%i )
xcopy "*" "../*" /S/F/I/R/Y
git checkout . && git clean -xdf 
cd ../
