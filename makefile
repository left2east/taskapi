 
default:
	go build -o myapp
#编译成linux平台可执行文件
linux:export GOOS=linux
linux:export GOARCH=amd64
linux:
	go build -o myapp main.go

# 编译
build:
	go build -o myapp .

# 清理
clean:
	rm -f myapp
