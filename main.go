package main

import (
	"Ahegao_Discord_Bot/cmds" //imagine a world where i could "./cmds" instead
	fwk "Ahegao_Discord_Bot/framework"
	js "encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
)

var (
	cfg        fwk.Config
	srv        fwk.Servers
	statBarSrv fwk.StatusBarServers
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
	config, err := os.Open("cfgs/config.json")
	if err != nil {
		fmt.Println("Error loading config. Error: ", err)
		os.Exit(1)
	}

	if err = js.NewDecoder(config).Decode(&cfg); err != nil {
		fmt.Println("Error decoding config. Error: ", err)
		os.Exit(1)
	}

	//Load status bar servers
	sbs, err := os.Open("cfgs/statusbar.json")
	if err != nil {
		fmt.Println("Error loading status bar servers. Error: ", err)
		os.Exit(1)
	}

	if err = js.NewDecoder(sbs).Decode(&statBarSrv); err != nil {
		fmt.Println("Error decoding status bar servers. Error: ", err)
		os.Exit(1)
	}
}

//Regularly updates servers in status bar/channel list
func ready(s *dG.Session, event *dG.Ready) {
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		for ; true; <-ticker.C {
			for i := 0; i < len(statBarSrv.Name); i++ {
				process(s, statBarSrv.Name[i], statBarSrv.Addr[i], statBarSrv.Id[i])
			}
		}
	}()
}

func main() {
	d, err := dG.New("Bot " + cfg.AuthID)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	//Anpan command handler
	h := ap.New(cfg.Prefixes, cfg.Owner, d.StateEnabled, cfg.IgnoreBots, cfg.RespondToPings, cfg.CheckPermissions, cmdPrerun, onError, nil)
	d.AddHandler(h.MessageHandler)

	//Command registers
	h.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, pingCommand)
	h.AddCommand("players", "Lists players on Bhop Servers", []string{"on"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.PlayersCommand)
	h.AddCommand("about", "Shows bot information", []string{"bot"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.AboutCommand)
	h.AddCommand("urban", "Search a word on urban dictionary", []string{"ud"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.UrbanCommand)
	h.AddCommand("currentmap", "List current maps on Bhop Servers", []string{"cm"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.CurrentMapCommand)
	h.AddCommand("serverinfo", "Shows information of a server. _serverinfo 127.0.0.1:27015", []string{"si"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.ServerInfoCommand)
	h.AddCommand("lenny", "Replies with the specified lenny.", []string{}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, ap.CommandTypeEverywhere, cmds.LennyCommand)

	//Help command
	h.SetHelpCommand("help", []string{}, dG.PermissionSendMessages, dG.PermissionSendMessages, cmds.HelpCommand)

	//Needed for status bar updates to work
	d.AddHandler(ready)

	//Open session
	err = d.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}
	ap.WaitForInterrupt()
	d.Close()

	//Close on os interrupt
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

//Debug ping command
func pingCommand(ctx ap.Context, _ []string) error {
	// We need to know what time it is now.
	ts := time.Now()

	msg, err := ctx.Reply("Pong!")
	if err != nil {
		return err
	}

	// Now we can compare it to the current time to see how much time went away during the process of sending a message.
	_, err = ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, msg.ID, fmt.Sprintf("Pong! Ping took **%dms**!", time.Since(ts).Milliseconds()))
	return err
}
