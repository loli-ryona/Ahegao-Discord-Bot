package cmds

import (
	"fmt"
	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
	"strings"
)

type server struct {
	name string
	addr string
}

var servers = []server{
	{"🍺 Pub: ", "144.48.37.114:27015"},
	{"🤍 WL: ", "144.48.37.118:27015"},
	{"🛹 Trikz: ", "144.48.37.119:27015"},
	{"🦘 Kanga: ", "146.185.214.33:27015"},
	{"🌌 Solitude: ", "51.161.131.99:27015"},
	{"🚸 IMK Easy: ", "139.99.209.158:27016"},
	{"🚸 IMK Hard: ", "139.99.209.158:27017"},
	{"🏳️‍🌈 Gay Tradies: ", "203.28.238.134:27015"},
	{"☭ Luchshe Veteranov: ", "46.174.52.164:27015"},
}

func PlayersCommand(ctx anpan.Context, _ []string) error {
	op := ""
	for _, server := range servers {
		client, err := a2s.NewClient(server.addr)
		var realPlayers []string
		if err != nil {
			//kill yourself
		} else {
			defer client.Close()
			players, err := client.QueryPlayer()
			if err != nil {
				//dont care
			} else {
				for _, player := range players.Players {
					if strings.Index(player.Name, "!replay") == -1 &&
						strings.Index(player.Name, "WR") == -1 &&
						strings.Index(player.Name, "Main") == -1 &&
						strings.Index(player.Name, "Bonus") == -1 &&
						strings.Index(player.Name, "GOTV") == -1 {
						realPlayers = append(realPlayers, player.Name)
					}
				}
				if len(realPlayers) > 0 {
					op += fmt.Sprintf("%s**%s**\n", server.name, strings.Join(realPlayers, ", "))
				}
			}
		}
	}

	embed := &discordgo.MessageEmbed{
		Title: "Current players online.",
		Description: op,
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	ctx.ReplyEmbed(embed)
	return nil
}