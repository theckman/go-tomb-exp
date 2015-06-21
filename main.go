// Copyright (c) 2015 Tim Heckman
// Released under the MIT license

// Package main is just the runner for the tombexp experiemental pkg.
package main

import (
	"fmt"

	"github.com/theckman/go-tomb-exp/tombexp"
)

func Runner() error {
	cryptKeeper := &tombexp.TalesFromTheCrypt{Death: 3}

	cryptKeeper.T.Go(cryptKeeper.Sleeper)

	return cryptKeeper.T.Wait()
}

func main() {
	err := Runner()
	if err != nil {
		fmt.Println("error:", err)
	}
}
