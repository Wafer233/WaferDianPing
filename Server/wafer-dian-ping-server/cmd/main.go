package main

import "github.com/Wafer233/WaferDianPing/wafer-dian-ping-server/initialize"

func main() {

	r, err := initialize.Init()
	if err != nil {
		panic(err)
	}
	_ = r.Run(":8081")
}
