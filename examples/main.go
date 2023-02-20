package main

import (
	"context"
	"time"

	"github.com/metal-stack/sonic-go-swsssdk"
)

func main() {

	s, err := sonic.New(nil)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()



	err = s.Interface().MTU(ctx, "Ethernet0", 9000)
	if err != nil {
		panic(err)
	}
}
