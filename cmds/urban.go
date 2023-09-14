package cmds

import (
	ap "ahegao/handler"
	"fmt"
	ud "github.com/JoshuaDoes/urbandictionary"
	dG "github.com/bwmarrin/discordgo"
	"net/url"
	"regexp"
	"strconv"
	str "strings"
)

func UrbanCommand(ctx ap.Context, args []string) error {
	// if query fails return error
	query, err := ud.Query(str.Join(args, " "))
	if err != nil {
		errEmbed := &dG.MessageEmbed{
			Title:       "Urban Dictionary Error",
			Description: "There was an error finding a result for that term.",
		}
		ctx.ReplyEmbed(errEmbed)
		return nil
	}

	// search urbandick
	lExp := regexp.MustCompile(`\[([^]]*)]`)
	lExpFunc := func(s string) string {
		ss := lExp.FindStringSubmatch(s)
		if len(ss) == 0 {
			return s
		}
		hl := "https://www.urbandictionary.com/define.php?term=" + url.QueryEscape(ss[1])
		return fmt.Sprintf("%s(%s)", s, hl)
	}

	// return result into embed
	result := query.Results[0]
	embed := &dG.MessageEmbed{
		Title:       "Urban Dictionary - " + result.Word,
		Description: lExp.ReplaceAllStringFunc(result.Definition, lExpFunc),
		Footer: &dG.MessageEmbedFooter{
			Text:    "Results from Urban Dictionairy.",
			IconURL: "https://i.pinimg.com/originals/f2/aa/37/f2aa3712516cfd0cf6f215301d87a7c2.jpg",
		},
	}

	// Example field
	embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
		Name:  "Example",
		Value: lExp.ReplaceAllStringFunc(result.Example, lExpFunc),
	})

	// Author of the word :^)
	embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
		Name:  "Author",
		Value: result.Author,
	})

	// Ratings on the word
	embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
		Name:  "Stats",
		Value: "\U0001f44d " + strconv.Itoa(result.ThumbsUp) + " \U0001f44e " + strconv.Itoa(result.ThumbsDown),
	})

	ctx.ReplyEmbed(embed)
	return nil
}
