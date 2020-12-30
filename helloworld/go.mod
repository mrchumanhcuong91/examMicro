module helloworld

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/micro/v3 v3.0.5-0.20201219085254-c8ea24387d19
	github.com/micro/services v0.16.0 // indirect
	github.com/soheilhy/cmux v0.1.4 // indirect
	google.golang.org/protobuf v1.25.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
