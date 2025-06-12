@echo off
 
::协议文件路径, 最后不要跟“\”符号
set SOURCE_FOLDER=.\Protol
set GoSOURCE_FOLDER=.\ProtolServer

::GO编译器路径
set CS_COMPILER_PATH=.\ProtoGen_Go\protoc.exe
::C#文件生成路径, 最后不要跟“\”符号
set CS_TARGET_PATH=.\Protocal

::C#文件生成路径, 最后不要跟“\”符号
::set GoServerPath=.\ProtolServer

::删除之前创建的文件
del %CS_TARGET_PATH%\*.* /f /s /q

::遍历所有文件 生成go文件
for /f "delims=" %%i in ('dir /b "%SOURCE_FOLDER%\*.proto"') do (
	echo %CS_COMPILER_PATH% -proto_path:%SOURCE_FOLDER%\%%i -go_out:%CS_TARGET_PATH%\%%~ni.go
	%CS_COMPILER_PATH% --proto_path=%SOURCE_FOLDER% --go_out=%CS_TARGET_PATH% %%i --go-grpc_out=%CS_TARGET_PATH% %%i
)
::遍历所有文件 生成go server文件
for /f "delims=" %%i in ('dir /b "%GoSOURCE_FOLDER%\*.proto"') do (
	echo %CS_COMPILER_PATH% -proto_path:%GoSOURCE_FOLDER%\%%i -go_out:%CS_TARGET_PATH%\%%~ni.go
	%CS_COMPILER_PATH% --proto_path=%GoSOURCE_FOLDER% --go_out=%CS_TARGET_PATH% %%i --go-grpc_out=%CS_TARGET_PATH% %%i
)

echo Generate Success!!!!!
 
pause