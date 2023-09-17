package framework

import (
	hnd "ahegao/handler"
	"errors"
	"fmt"
)

// OnError = Prints an error if a command fails.
func OnError(ctx hnd.Context, cmd *hnd.Command, context []string, err error) {
	if errors.Is(err, hnd.ErrCommandNotFound) {
		return
	}
	fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", cmd.Name, err.Error())
}

// CmdPrerun = Prints who has run a command.
func CmdPrerun(ctx hnd.Context, cmd *hnd.Command, content []string) bool {
	fmt.Printf("Command \"%s\" is being run by \"%s#%s\" (ID: %s).\n", cmd.Name, ctx.User.Username, ctx.User.Discriminator, ctx.User.ID)
	return true
}

// HandlerDebug = Functions for handling debug messages.
func HandlerDebug(df string) {
	fmt.Println(df)
	return
}
