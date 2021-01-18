@echo off
set command = " swagger mixin "
for  /r  %%i  in  (*.proto)  do  (echo  generating  %%i  &  kratos tool protoc %%i )
xcopy "*" "../*" /S/F/I/R/Y
git checkout . && git clean -xdf 
cd ../
for  /r  %%i  in  (*.json)  do  (  set command=%command%,%%i)
del *.bat&& echo %commmand%
