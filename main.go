package main

import (
	"os"
	"pipeline-common-k8s-sidecar/app"
	"time"

	"github.com/spf13/cast"
)

func main() {
	exitTime := cast.ToInt(os.Getenv("TIMEOUT"))
	if exitTime == 0 {
		exitTime = 1800
	}
	go func() {
		time.AfterFunc(time.Second*time.Duration(exitTime), func() {
			os.Exit(0)
		})
	}()
	app.RegisterRouter().Run(":9999")
}

//172-17-0-3.default.pod
