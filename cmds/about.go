package cmds

import (
	ap "ahegao/handler"
	dG "github.com/bwmarrin/discordgo"
)

func AboutCommand(ctx ap.Context, _ []string) error {
	embed := &dG.MessageEmbed{
		Title:       "About Ahegao Discord Bot:",
		Description: "Hello, I am the official [Ahegao Discord Bot](https://ahegao.neocities.org)! Written by [Lost](https://steamcommunity.com/id/T-r-i-s-t-a-n/)!",
		Fields: []*dG.MessageEmbedField{
			{
				Name:  "Repo",
				Value: "[github.com/loli-ryona/ahegao-discord-bot](https://github.com/loli-ryona/ahegao-discord-bot)",
			},
			{
				Name:  "Version",
				Value: "Currently running on version 1.4.1",
			},
			{
				Name:  "Language",
				Value: "[Golang](https://golang.org)",
			},
			{
				Name:  "Website",
				Value: "[ahegao.neocities.org](https://ahegao.neocities.org)",
			},
		},
		Footer: &dG.MessageEmbedFooter{
			Text:    "Currently running version a1.3!",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	ctx.ReplyEmbed(embed)
	return nil
}
