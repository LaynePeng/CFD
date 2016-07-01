@ECHO OFF

SET PWD=%cd%
IF NOT EXIST %PWD%\src\github.com\LaynePeng (
    ECHO Start to Create Link
    md %PWD%\src\github.com\LaynePeng
    mklink /J  %PWD%\src\github.com\LaynePeng\cfd %PWD%
) 

ECHO Start to Build
go build -o cfd.exe