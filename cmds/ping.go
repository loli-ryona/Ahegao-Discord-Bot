package cmds

import (
	hnd "ahegao/handler"
	"fmt"
	"time"
)

//Debug ping command
func PingCommand(ctx hnd.Context, _ []string) error {
	// We need to know what time it is now.
	ts := time.Now()

	msg, err := ctx.Reply("Pong!")
	if err != nil {
		return err
	}

	// Now we can compare it to the current time to see how much time went away during the process of sending a message.
	_, err = ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, msg.ID, fmt.Sprintf("Pong! Ping took **%dms**!", time.Since(ts).Milliseconds()))
	return err
}
