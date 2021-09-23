package cmds

import (
	fwk "Ahegao_Discord_Bot/framework"
	js "encoding/json"
	"fmt"
	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
	"os"
	str "strings"
)

var (
	srv fwk.Servers
)

func PlayersCommand(ctx ap.Context, _ []string) error {
	//Declaring output var
	op := ""

	//Load servers.json
	servers, err := os.Open("servers.json")
	if err != nil {
		fmt.Println("Error loading servers. Error: ", err)
		os.Exit(1)
	}

	if err = js.NewDecoder(servers).Decode(&srv); err != nil {
		fmt.Println("Error decoding servers. Error: ", err)
		os.Exit(1)
	}

	//Loop through each server and get players
	for i := 0; i < len(srv.Name); i++ {
		var realPlayers []string
		if client, err := a2s.NewClient(srv.Addr[i]); err != nil {
			fmt.Println("Error creating new A2S client. Error: ", err)
		} else {
			defer client.Close()
			if players, err := client.QueryPlayer(); err != nil {
				fmt.Println("Error querying players. Error: ", err)
			} else {
				for _, player := range players.Players {
					if str.Index(player.Name, "!replay") == -1 &&
						str.Index(player.Name, "WR") == -1 &&
						str.Index(player.Name, "Main") == -1 &&
						str.Index(player.Name, "Bonus") == -1 &&
						str.Index(player.Name, "GOTV") == -1 {
						realPlayers = append(realPlayers, player.Name)
					}
				}
				if len(realPlayers) > 0 {
					op += fmt.Sprintf("%s**%s**\n", srv.Name[i], str.Join(realPlayers, ", "))
				}
			}
		}
	}

	//Create embed
	embed := &dG.MessageEmbed{
		Title:       "Current players online.",
		Description: op,
		Footer: &dG.MessageEmbedFooter{
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	//Return embed
	ctx.ReplyEmbed(embed)
	return nil
}
