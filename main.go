package main

import (
	"ahegao/cmds"
	fwk "ahegao/framework"
	hnd "ahegao/handler"
	js "encoding/json"
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cfg        fwk.Config
	srv        fwk.Servers
	statBarSrv fwk.StatusBarServers
)

func init() {
	//Load config.json
	fmt.Println("Loading Config (1/2)")
	config, err := os.Open("cfgs/config.json")
	if err != nil {
		fmt.Println("Error loading config. Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Loading Config (2/2)")
	if err = js.NewDecoder(config).Decode(&cfg); err != nil {
		fmt.Println("Error decoding config. Error: ", err)
		os.Exit(1)
	}
	fmt.Println(cfg.Prefixes)

	//Load statusbar.json
	fmt.Println("Loading Status Bar Config (1/2)")
	sbs, err := os.Open("cfgs/statusbar.json")
	if err != nil {
		fmt.Println("Error loading status bar servers. Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Loading Status Bar Config (2/2)")
	if err = js.NewDecoder(sbs).Decode(&statBarSrv); err != nil {
		fmt.Println("Error decoding status bar servers. Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Configs are loaded!")
}

func main() {
	// Create discord session
	fmt.Println("Creating Discord session")
	session, err := dG.New("Bot " + cfg.AuthID)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	session.AddHandler(startStatusBarTicker) // Start status bar service

	// Create command handler
	fmt.Println("Creating command handler")
	h := hnd.New(cfg.Prefixes, cfg.Owner, session.StateEnabled, cfg.IgnoreBots, cfg.RespondToPings, cfg.CheckPermissions, fwk.CmdPrerun, fwk.OnError, nil)
	h.SetDebugFunc(fwk.HandlerDebug) // Comment this out if you dont want debug logs
	session.AddHandler(h.MessageHandler)

	// Register commands
	fmt.Println("Registering commands")
	h.AddCommand("ping", "Check the bot's ping. Used for debugging", []string{"pong"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.PingCommand)
	h.AddCommand("players", "Lists players on Bhop Servers", []string{"on"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.PlayersCommand)
	h.AddCommand("about", "Shows bot information", []string{"bot"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.AboutCommand)
	h.AddCommand("urban", "Search a word on urban dictionary", []string{"ud"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.UrbanCommand)
	h.AddCommand("currentmap", "List current maps on Bhop Servers", []string{"cm"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.CurrentMapCommand)
	h.AddCommand("serverinfo", "Shows information of a server. _serverinfo 127.0.0.1:27015", []string{"si"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.ServerInfoCommand)
	h.AddCommand("lenny", "Replies with the specified lenny.", []string{}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.LennyCommand)
	h.AddCommand("thetime", "Replies with the time.", []string{"time", "tt"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.TheTimeCommand)
	h.AddCommand("weather", "Replies with weather for provided location. Execute command without any arguments for more info.", []string{"temp", "bom"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, cmds.WeatherCommand)

	fmt.Println("Building Help command")
	h.SetHelpCommand("help", []string{}, dG.PermissionSendMessages, dG.PermissionSendMessages, cmds.HelpCommand)

	// Open discord session
	fmt.Println("Completed loading, opening session")
	err = session.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}
	hnd.WaitForInterrupt()
	session.Close()

	//Close on os interrupt
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

/*Starts the status bar service.
 * session = *discordgo.Session
 * event = *discordgo.Ready
 */
func startStatusBarTicker(session *dG.Session, event *dG.Ready) {
	fmt.Println("Starting Status Bar Ticker")
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		for ; true; <-ticker.C {
			for i := 0; i < len(statBarSrv.Name); i++ {
				process(session, statBarSrv.Name[i], statBarSrv.Addr[i], statBarSrv.Id[i])
			}
		}
	}()
}
