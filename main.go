package main

import (
	"Ahegao_Discord_Bot/cmds" //imagine a world where i could "./cmds" instead
	"Ahegao_Discord_Bot/framework"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	_ "strings"
	"syscall"
	"time"

	ap "github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

//Channels and bhop server vars
type statusbarServer struct {
	name string
	addr string
	id   string
}

var statusbarServers = []statusbarServer{
	{"🍺 Pub: ", "144.48.37.114:27015", "883671130286739476"},
	{"🤍 WL: ", "144.48.37.118:27015", "883671190521147412"},
	{"🦘 Kanga: ", "146.185.214.33:27015", "883671212784488498"},
}

//Regularly updates servers in status bar/channel list
func ready(s *discordgo.Session, event *discordgo.Ready) {
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		for ; true; <-ticker.C {
			for _, statusbarServer := range statusbarServers {
				process(s, statusbarServer.name, statusbarServer.addr, statusbarServer.id)
			}
		}
	}()
}

//Actual program shit now
var (
	cfg framework.Config
)

func onError(ctx ap.Context, cmd *ap.Command, context []string, err error) {
	if err == ap.ErrCommandNotFound {
		return
	}

	fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", cmd.Name, err.Error())
}

func cmdPrerun(ctx ap.Context, cmd *ap.Command, content []string) bool {
	fmt.Printf("Command \"%s\" is being run by \"%s#%s\" (ID: %s).\n", cmd.Name, ctx.User.Username, ctx.User.Discriminator, ctx.User.ID)
	return true
}

func init() {
	//Load config
	bV, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error loading config. Error: ", err)
	}

	err = json.NewDecoder(bV).Decode(&cfg)
	if err != nil {
		fmt.Println("Error decoding config. Error: ", err)
		os.Exit(1)
	}
}
func main() {
	discord, err := discordgo.New("Bot " + cfg.AuthID)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	//Anpan command handler
	handler := ap.New(cfg.Prefixes, []string{"Lost#5712", "251959155382812673"}, discord.StateEnabled, true, true, true, cmdPrerun, onError, nil)
	discord.AddHandler(handler.MessageHandler)

	//Command registers
	handler.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, ap.CommandTypeEverywhere, pingCommand)
	handler.AddCommand("players", "Lists players on Bhop Servers", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.PlayersCommand)
	handler.AddCommand("about", "Shows bot information", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.AboutCommand)
	handler.AddCommand("urban", "Search a word on urban dictionairy", []string{"ud"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.UrbanCommand)

	//Help command
	handler.SetHelpCommand("help", []string{}, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, cmds.HelpCommand)

	//Needed for status bar updates to work
	discord.AddHandler(ready)

	//Open session
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}
	ap.WaitForInterrupt()

	//Close on os interrupt
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}

//Debug ping command
func pingCommand(ctx ap.Context, _ []string) error {
	// We need to know what time it is now.
	timestamp := time.Now()

	msg, err := ctx.Reply("Pong!")
	if err != nil {
		return err
	}

	// Now we can compare it to the current time to see how much time went away during the process of sending a message.
	_, err = ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, msg.ID, fmt.Sprintf("Pong! Ping took **%dms**!", time.Since(timestamp).Milliseconds()))
	return err
}
