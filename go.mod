module example

go 1.14

require (
	example/helloworld v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/micro/v3 v3.0.5-0.20201219085254-c8ea24387d19
	github.com/micro/services v0.16.0 // indirect
	github.com/mkideal/cli v0.2.3
)

replace example/helloworld => ./helloworld
