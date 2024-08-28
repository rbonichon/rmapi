package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rbonichon/rmapi/api"
	"github.com/rbonichon/rmapi/config"
	"github.com/rbonichon/rmapi/log"
	"github.com/rbonichon/rmapi/shell"
	"github.com/rbonichon/rmapi/version"
)

const AUTH_RETRIES = 3

func parseOfflineCommands(cmd []string) bool {
	if len(cmd) == 0 {
		return false
	}

	switch cmd[0] {
	case "reset":
		configFile, err := config.ConfigPath()
		if err != nil {
			log.Error.Fatalln(err)
		}
		if err := os.Remove(configFile); err != nil {
			log.Error.Fatalln(err)
		}
		return true
	case "version":
		fmt.Println(version.Version)
		return true
	}
	return false
}

func main() {
	ni := flag.Bool("ni", false, "not interactive (prevents asking for code)")
	flag.Usage = func() {
		fmt.Println(`
  help		detailed commands, but the user needs to be logged in

Offline Commands:
  version	prints the version
  reset		removes the config file `)

		flag.PrintDefaults()
	}
	flag.Parse()
	otherFlags := flag.Args()
	if parseOfflineCommands(otherFlags) {
		return
	}

	var ctx api.ApiCtx
	var err error
	var userInfo *api.UserInfo

	for i := 0; i < AUTH_RETRIES; i++ {
		authCtx := api.AuthHttpCtx(i > 0, *ni)

		userInfo, err = api.ParseToken(authCtx.Tokens.UserToken)
		if err != nil {
			log.Trace.Println(err)
			continue
		}

		ctx, err = api.CreateApiCtx(authCtx, userInfo.SyncVersion)
		if err != nil {
			log.Trace.Println(err)
		} else {
			break
		}
	}

	if err != nil {
		log.Error.Fatal("failed to build documents tree, last error: ", err)
	}

	err = shell.RunShell(ctx, userInfo, otherFlags)

	if err != nil {
		log.Error.Println("Error: ", err)

		os.Exit(1)
	}
}
