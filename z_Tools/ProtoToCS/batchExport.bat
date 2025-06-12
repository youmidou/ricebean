@echo off
 
::协议文件路径, 最后不要跟“\”符号
set SOURCE_FOLDER=.\proto

::C#编译器路径
set CS_COMPILER_PATH=.\ProtoGen\protogen.exe
::C#文件生成路径, 最后不要跟“\”符号
set CS_TARGET_PATH=.\exportFolder

::删除之前创建的文件
del %CS_TARGET_PATH%\*.* /f /s /q

::读取版本
echo %CS_COMPILER_PATH% --version
%CS_COMPILER_PATH% --version


::遍历所有文件
for /f "delims=" %%i in ('dir /b "%SOURCE_FOLDER%\*.proto"') do (
	::生成 C# 代码 --csharp_out=%SOURCE_FOLDER%
	echo %CS_COMPILER_PATH% -I:%SOURCE_FOLDER%\%%i -o:%CS_TARGET_PATH%\%%~ni.cs
	::%CS_COMPILER_PATH% --csharp_out:%CS_TARGET_PATH% --proto_path:%SOURCE_FOLDER% -I:%SOURCE_FOLDER%\%%i -o:%CS_TARGET_PATH%\%%~ni.cs
	%CS_COMPILER_PATH% --proto_path=%SOURCE_FOLDER% --csharp_out=%CS_TARGET_PATH% %%i
)

echo Generate Success!!!!!
 
pause