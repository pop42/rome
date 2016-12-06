package main

import (
	"log"
	"runtime"

	"github.com/bluefoxcode/rome/services/Nero/lib/boot"
	"github.com/bluefoxcode/rome/services/Nero/lib/router"

	"github.com/urfave/negroni"
)

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	info := boot.LoadConfig()
	boot.RegisterServices(info)

	n := negroni.New(negroni.NewLogger())
	// n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router.Instance())

	n.Run(":" + info.Port)
}
