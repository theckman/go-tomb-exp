// Copyright (c) 2015 Tim Heckman
// Released under the MIT license

// Package tombexp is an experiment playing with the tomb package.
// This isn't really meant to be used for anything useful.
package tombexp

import (
	"fmt"
	"time"

	"gopkg.in/tomb.v2"
)

type TalesFromTheCrypt struct {
	T     tomb.Tomb
	Death int
}

func (tftc TalesFromTheCrypt) Sleeper() error {
	var delay int

	// loop over a channel select
	for {
		// look for the Dying channel to send a message
		// or just delay a bit, looking for the time to die
		select {
		case <-tftc.T.Dying():
			fmt.Println("killing off goroutine")
		default:
			// increase the delay
			// and exit if it's past the death time
			delay++
			if delay >= tftc.Death {
				fmt.Println("goroutine exiting")
				return nil
			}
			// print how long we are sleeping for, and then sleep
			fmt.Printf("sleeping for %d seconds\n", delay)
			time.Sleep(time.Duration(delay) * time.Second)
		}
	}
	return nil
}
