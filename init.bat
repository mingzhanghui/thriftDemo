cd %GOPATH%\src\ThriftDemo
mkdir client server thrift_file

@doskey touch=copy /b NUL $*
touch .\client\client.go
touch .\server\server.go
touch .\thrift_file\example.thrift