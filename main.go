package main

import (
	"github.com/mkideal/cli"
	sv "example/server"
)

type cliArgs struct {
	cli.Helper
	//ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	//ListenAddr    string `cli:"*l,*listen" usage:"gateway listen host and port" dft:"$GATEWAY_LS"`
}


func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		sv.NewApiGateWayServer("root:123456@tcp(localhost:3306)/microservice?parseTime=true&collation=utf8mb4_unicode_ci").Run(":9091")
		return nil
	})
	
}

