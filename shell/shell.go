package shell

import (
	"fmt"
	"os"

	"github.com/abiosoft/ishell"
	"github.com/rbonichon/rmapi/api"
	"github.com/rbonichon/rmapi/model"
)

type ShellCtxt struct {
	node           *model.Node
	api            api.ApiCtx
	path           string // remote path
	local_path     string
	useHiddenFiles bool
	UserInfo       api.UserInfo
}

func (ctx *ShellCtxt) prompt() string {
	return fmt.Sprintf("[%s::%s]>", ctx.local_path, ctx.path)
}

func setCustomCompleter(shell *ishell.Shell) {
	cmdCompleter := make(cmdToCompleter)
	for _, cmd := range shell.Cmds() {
		cmdCompleter[cmd.Name] = cmd.Completer
	}

	completer := shellPathCompleter{cmdCompleter}
	shell.CustomCompleter(completer)
}

func useHiddenFiles() bool {
	val, ok := os.LookupEnv("RMAPI_USE_HIDDEN_FILES")

	if !ok {
		return false
	}

	return val != "0"
}

func RunShell(apiCtx api.ApiCtx, userInfo *api.UserInfo, args []string) error {
	shell := ishell.New()
	cwd, _ := os.Getwd()
	ctx := &ShellCtxt{
		node:           apiCtx.Filetree().Root(),
		api:            apiCtx,
		path:           apiCtx.Filetree().Root().Name(),
		local_path:     cwd,
		useHiddenFiles: useHiddenFiles(),
		UserInfo:       *userInfo,
	}

	shell.SetPrompt(ctx.prompt())

	shell.AddCmd(lsCmd(ctx))
	shell.AddCmd(llsCmd(ctx))
	shell.AddCmd(pwdCmd(ctx))
	shell.AddCmd(cdCmd(ctx))
	shell.AddCmd(lcdCmd(ctx))
	shell.AddCmd(getCmd(ctx))
	shell.AddCmd(mgetCmd(ctx))
	shell.AddCmd(mkdirCmd(ctx))
	shell.AddCmd(rmCmd(ctx))
	shell.AddCmd(mvCmd(ctx))
	shell.AddCmd(putCmd(ctx))
	shell.AddCmd(mputCmd(ctx))
	shell.AddCmd(versionCmd(ctx))
	shell.AddCmd(statCmd(ctx))
	shell.AddCmd(getACmd(ctx))
	shell.AddCmd(findCmd(ctx))
	shell.AddCmd(nukeCmd(ctx))
	shell.AddCmd(accountCmd(ctx))
	shell.AddCmd(refreshCmd(ctx))

	setCustomCompleter(shell)

	if len(args) > 0 {
		return shell.Process(args...)
	} else {
		shell.Printf("ReMarkable Cloud API Shell, User: %s, SyncVersion: %s\n", userInfo.User, userInfo.SyncVersion)
		shell.Run()

		return nil
	}
}
