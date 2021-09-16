package cmds

import (
	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func AboutCommand(ctx anpan.Context, _ []string) error {
	embed := &discordgo.MessageEmbed{
		Title:       "About Ahegao Discord Bot:",
		Description: "Hello, I am the official [Ahegao Discord](https://discord.com/invite/FFb59tM)! Written in [Golang](https://golang.org) by [Lost](https://steamcommunity.com/id/T-r-i-s-t-a-n/)!",
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Currently running version a0.0.1!",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	ctx.ReplyEmbed(embed)
	return nil
}
