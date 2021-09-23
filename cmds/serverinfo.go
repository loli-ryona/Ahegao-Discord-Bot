package cmds

import (
	js "encoding/json"
	"fmt"
	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
	"os"
	"strconv"
	"time"
)

func ServerInfoCommand(ctx ap.Context, args []string) error {
	// thetime
	ts := time.Now()

	//Check server was supplied
	if len(args) >= 1 {
		//Reply with embed
		embed := &dG.MessageEmbed{
			Title:       "Server info.",
			Description: "Please wait while we query the server.",
			Footer: &dG.MessageEmbedFooter{
				Text:    "Calculating time to query server.",
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

		//Get server from command
		server := args[0]
		var sName []string
		var sMap []string
		var sPlayers []uint8
		var sMaxPlayers []uint8
		var sVac []bool
		var sOs []a2s.ServerOS
		var sPass []bool
		var sGame []string
		var sGameID []uint16

		if client, err := a2s.NewClient(server); err != nil {
			fmt.Println("Error creating new A2S client. Error: ", err)
		} else {
			defer client.Close()
			if info, err := client.QueryInfo(); err != nil {
				fmt.Println("Error querying server. Error: ", err)
			} else {
				sName := append(sName, info.Name)
				sMap := append(sMap, info.Map)
				sPlayers := append(sPlayers, info.Players)
				sMaxPlayers := append(sMaxPlayers, info.MaxPlayers)
				sVac := append(sVac, info.VAC)
				sOs := append(sOs, info.ServerOS)
				sPass := append(sPass, info.Visibility)
				sGame := append(sGame, info.Game)
				sGameID := append(sGameID, info.ID)
				srvOs := ""

				if sOs[0] == a2s.ServerOS_Mac {
					srvOs = "Mac OSX"
				} else if sOs[0] == a2s.ServerOS_Linux {
					srvOs = "Linux"
				} else if sOs[0] == a2s.ServerOS_Windows {
					srvOs = "Windows"
				} else {
					srvOs = "Unknown OS"
				}

				//Edit embed
				ed := &dG.MessageEmbed{
					Title:       sName[0],
					Description: "Currently " + strconv.Itoa(int(sPlayers[0])) + "/" + strconv.Itoa(int(sMaxPlayers[0])) + " online.",
					Footer: &dG.MessageEmbedFooter{
						Text:    fmt.Sprintf("Took %.2fs to query servers!", time.Since(ts).Seconds()),
						IconURL: ctx.Session.State.User.AvatarURL("512"),
					},
				}

				//Fields
				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "Current map:",
					Value: sMap[0],
				})

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "Security:",
					Value: "VAC: " + strconv.FormatBool(sVac[0]) + " | Password: " + strconv.FormatBool(sPass[0]),
				})

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "System Info",
					Value: "OS: " + srvOs + " | Game: " + sGame[0] + " | Game ID: " + strconv.Itoa(int(sGameID[0])),
				})
				//Return embed
				_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
				return nil
			}
		}
	}
	return nil
}
