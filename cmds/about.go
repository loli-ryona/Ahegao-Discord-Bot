package cmds

import (
	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
)

func AboutCommand(ctx ap.Context, _ []string) error {
	embed := &dG.MessageEmbed{
		Title:       "About Ahegao Discord Bot:",
		Description: "Hello, I am the official [Ahegao Discord](https://discord.com/invite/FFb59tM)! Written in [Golang](https://golang.org) by [Lost](https://steamcommunity.com/id/T-r-i-s-t-a-n/)!",
		Footer: &dG.MessageEmbedFooter{
			Text:    "Currently running version a1.2!",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	ctx.ReplyEmbed(embed)
	return nil
}
