package config

import (
	"context"
	"flag"
	"log"
	"net/http"
)

var cpuProf = flag.String("cpu", "", "write cpu profile to `file`")
var memProf = flag.String("mem", "", "write memory profile to `file`")

func ProcessFlags(ctx *context.Context) error {
	flag.Parse()
	if *cpuProf != "" {

		// pprof webserver
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
		//go profile(60*time.Second)
	}
	return nil
}
