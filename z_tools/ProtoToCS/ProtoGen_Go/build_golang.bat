@echo off
 
::协议文件路径, 最后不要跟“\”符号
set SOURCE_FOLDER=.\proto

::GO编译器路径
set CS_COMPILER_PATH=.\protoc.exe
::C#文件生成路径, 最后不要跟“\”符号
set CS_TARGET_PATH=.\exportFolder2

::删除之前创建的文件
del %CS_TARGET_PATH%\*.* /f /s /q

::遍历所有文件
for /f "delims=" %%i in ('dir /b "%SOURCE_FOLDER%\*.proto"') do (
	::生成 C# 代码 --csharp_out=%SOURCE_FOLDER%
	echo %CS_COMPILER_PATH% -proto_path:%SOURCE_FOLDER%\%%i -go_out:%CS_TARGET_PATH%\%%~ni.cs
	protoc.exe --proto_path=%SOURCE_FOLDER% --go_out=%CS_TARGET_PATH% %%i
	
)

echo Generate Success!!!!!
 
pause