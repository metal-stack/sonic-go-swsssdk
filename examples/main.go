package main

import (
	"context"
	"time"

	"github.com/metal-stack/sonic-go-swsssdk"
)

func main() {

	client, err := sonic.New(nil)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Interface("Ethernet0").Vrf(ctx, 90)
	if err != nil {
		panic(err)
	}
}
