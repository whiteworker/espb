
for  /r  %%i  in  (*.proto)  do  (echo  generating  %%i  &  kratos tool protoc %%i )
xcopy "*" "../*" /S/F/I/R/Y
git checkout . && git clean -xdf 
cd ../
set command=swagger mixin 
for  /r  %%i  in  (*.json)  do call set command=%%command%% %%i
set command=%command% -o swagger.json
%command%
del *.bat
cd pb