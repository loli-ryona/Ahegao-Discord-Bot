package cmds

import (
	js "encoding/json"
	"fmt"
	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
	"os"
	str "strings"
	"time"
)

func PlayersCommand(ctx ap.Context, _ []string) error {
	// thetime
	ts := time.Now()

	//Declaring output var
	op := ""

	//Reply with embed
	embed := &dG.MessageEmbed{
		Title:       "Current players online.",
		Description: "Please wait while we query the servers for players",
		Footer: &dG.MessageEmbedFooter{
			Text:    "Calculating time to query servers.",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	msg, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}

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
	// TODO: Turn each server query into a Go Routine
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
	ed := &dG.MessageEmbed{
		Title:       "Current players online.",
		Description: op,
		Footer: &dG.MessageEmbedFooter{
			Text:    fmt.Sprintf("Took %.2fs to query servers!", time.Since(ts).Seconds()),
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	//Return embed
	_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
	return nil
}
