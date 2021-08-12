
del test.exe

go build  -o test.exe .

go get github.com/GeertJohan/go.rice/rice@v1.0.2

go run github.com/GeertJohan/go.rice/rice append --exec test.exe -i ./build

pause