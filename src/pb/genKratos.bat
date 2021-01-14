
for  /r  %%i  in  (*.proto)  do  (echo  正在生成  %%i  &  kratos tool protoc *.proto && xcopy "*" "../../*" /S/F/I/R/Y && del *.json &&del *.go && cd ../../&&  del *.proto &&cd pb)


