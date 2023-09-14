package cmds

import (
	fwk "ahegao/framework"
	js "encoding/json"
	"fmt"
	"os"
	str "strings"

	ap "ahegao/handler"
	dG "github.com/bwmarrin/discordgo"
)

var (
	lenny fwk.LennyExpressions
)

func lennyFaces(args string) (string, bool) {
	notFound := false
	face := ""
	for i := 0; i < len(lenny.Expression); i++ {
		if args == lenny.Expression[i] {
			face = lenny.Face[i]
		}
	}

	if face == "" {
		fmt.Println("Expression not found")
		notFound = true
	}
	return face, notFound
}

func listLennyFaces() (string, string) {
	faces := ""
	expressions := ""
	var face []string
	var expression []string
	for i := 0; i < len(lenny.Face); i++ {
		face = append(face, lenny.Face[i])
		expression = append(expression, lenny.Expression[i])
	}

	faces += fmt.Sprintf("%s", str.Join(face, ", "))
	expressions += fmt.Sprintf("%s", str.Join(expression, ", "))
	return faces, expressions
}

func LennyCommand(ctx ap.Context, args []string) error {
	faces, err := os.Open("cfgs/lenny.json")
	if err != nil {
		fmt.Println("Error loading lenny faces. Error: ", err)
		os.Exit(1)
	}

	if err = js.NewDecoder(faces).Decode(&lenny); err != nil {
		fmt.Println("Error decoding lenny faces. Error: ", err)
		os.Exit(1)
	}

	if len(args) >= 1 {
		if face, notFound := lennyFaces(args[0]); notFound == true {
			fmt.Println("Provided expression not found")
			ed := &dG.MessageEmbed{
				Title:       "( ͡° ͜ʖ ͡°)",
				Description: "Provided expression not found!",
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintln("Ahegao Discord Bot"),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}

			ctx.ReplyEmbed(ed)
		} else {
			msg := ctx.Message.ID
			channel := ctx.Channel.ID
			ctx.Session.ChannelMessageDelete(channel, msg)
			ctx.Reply(face)

		}
	} else {
		fmt.Println("No expression provided")
		ed := &dG.MessageEmbed{
			Title:       "( ͡° ͜ʖ ͡°)",
			Description: "Please provide an expression!",
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintln("Ahegao Discord Bot"),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}

		faces, expressions := listLennyFaces()
		ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
			Name:  "Current expressions",
			Value: expressions,
		})
		ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
			Name:  "Current faces",
			Value: faces,
		})

		ctx.ReplyEmbed(ed)
	}

	return nil
}
