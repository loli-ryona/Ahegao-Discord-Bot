package cmds

import (
	"fmt"
	"time"

	ap "ahegao/handler"
	dG "github.com/bwmarrin/discordgo"
)

func TheTimeCommand(ctx ap.Context, _ []string) error {
	//thetime
	ts := time.Now()

	//output var
	op := fmt.Sprintln("It is: ", ts.Format(time.Stamp))

	//create embed
	embed := &dG.MessageEmbed{
		Title:       "The Time.",
		Description: op,
		Footer: &dG.MessageEmbedFooter{
			Text:    fmt.Sprintf("Took %.2fs to get the time!", time.Since(ts).Seconds()),
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	//return embed
	_, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}
	return nil
}
