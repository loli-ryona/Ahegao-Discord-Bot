package cmds

import (
	"fmt"
	"github.com/JoshuaDoes/urbandictionary"
	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func UrbanCommand(ctx anpan.Context, args []string) error {
	results, err := urbandictionary.Query(strings.Join(args, " "))
	if err != nil {
		errEmbed := &discordgo.MessageEmbed{
			Title:       "Urban Dictionary Error",
			Description: "There was an error finding a result for that term.",
		}
		ctx.ReplyEmbed(errEmbed)
		return nil
	}

	linkExp := regexp.MustCompile(`\[([^\]]*)\]`)
	linkExpFunc := func(s string) string {
		ss := linkExp.FindStringSubmatch(s)
		if len(ss) == 0 {
			return s
		}
		hyperlink := "https://www.urbandictionary.com/define.php?term=" + url.QueryEscape(ss[1])
		return fmt.Sprintf("%s(%s)", s, hyperlink)
	}

	result := results.Results[0]
	resultEmbed := &discordgo.MessageEmbed{
		Title:       "Urban Dictionary - " + result.Word,
		Description: linkExp.ReplaceAllStringFunc(result.Definition, linkExpFunc),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Results from Urban Dictionairy.",
			IconURL: "https://res.cloudinary.com/hrscywv4p/image/upload/c_limit,fl_lossy,h_300,w_300,f_auto,q_auto/v1/1194347/vo5ge6mdw4creyrgaq2m.png",
		},
	}

	resultEmbed.Fields = append(resultEmbed.Fields, &discordgo.MessageEmbedField{
		Name:  "Example",
		Value: linkExp.ReplaceAllStringFunc(result.Example, linkExpFunc),
	})

	resultEmbed.Fields = append(resultEmbed.Fields, &discordgo.MessageEmbedField{
		Name:  "Author",
		Value: result.Author,
	})

	resultEmbed.Fields = append(resultEmbed.Fields, &discordgo.MessageEmbedField{
		Name:  "Stats",
		Value: "\U0001f44d " + strconv.Itoa(result.ThumbsUp) + " \U0001f44e " + strconv.Itoa(result.ThumbsDown),
	})

	ctx.ReplyEmbed(resultEmbed)
	return nil
}
