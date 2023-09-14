package framework

import (
	hnd "ahegao/handler"
	"errors"
	"fmt"
)

// Prints download progress.
func ByteString(n int64) string {
	if n < 1<<10 {
		return fmt.Sprintf("%dB", n)
	}
	if n < 1<<20 {
		return fmt.Sprintf("%dKB", n>>10)
	}
	if n < 1<<30 {
		return fmt.Sprintf("%dMB", n>>20)
	}
	if n < 1<<40 {
		return fmt.Sprintf("%dGB", n>>30)
	}
	return fmt.Sprintf("%dTB", n>>40)
}

// Prints an error if a command fails.
func OnError(ctx hnd.Context, cmd *hnd.Command, context []string, err error) {
	if errors.Is(err, hnd.ErrCommandNotFound) {
		return
	}
	fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", cmd.Name, err.Error())
}

// Prints who has run a command.
func CmdPrerun(ctx hnd.Context, cmd *hnd.Command, content []string) bool {
	fmt.Printf("Command \"%s\" is being run by \"%s#%s\" (ID: %s).\n", cmd.Name, ctx.User.Username, ctx.User.Discriminator, ctx.User.ID)
	return true
}

// Functions for handling debug messages.
func HandlerDebug(df string) {
	fmt.Println(df)
	return
}
