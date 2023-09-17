package main

import (
	"ahegao/cmds"
	"ahegao/framework"
	"ahegao/handler"
	"encoding/json"
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var (
	cfgPath = "cfgs/config.json"
	sbsPath = "cfgs/statusbar.json"

	cfg framework.Config
	sbs framework.StatusBarServers
)

// init = Loads the required config files into the interface.
func init() {
	paths := []string{cfgPath, sbsPath}
	structs := []any{&cfg, &sbs}

	for i := 0; i < len(paths); i++ {
		j, err := os.Open(paths[i])
		if err != nil {
			fmt.Println("Error opening file. Error: ", err)
			os.Exit(1)
		}
		if err = json.NewDecoder(j).Decode(structs[i]); err != nil {
			fmt.Println("Error decoding file. Error: ", err)
			os.Exit(1)
		}
	}
}

// The heart of the bot, puts everything in its place and allows it to run. A basic layout of what this does is as follows.
/*
 1. Creates the config interface.
 2. Creates the Discord session.
 3. Creates the command handler.
 4. Adds the Status Bar process and command handler to the session.
 5. Registers commands with the command handler.
 6. Opens the Discord session.
 7. Waits for interrupt to close session and shutdown bot.
*/
func main() {
	// Create new config interface
	cfgs := framework.New(&framework.Configs{
		MainConfig:      cfg,
		StatusBarConfig: sbs,
	})

	// Create discord session
	fmt.Println("Creating Discord session")
	session, err := dG.New("Bot " + cfg.AuthID)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Add handlers to the session
	h := framework.CreateCommandHandler(session, cfgs) // Create command handler
	session.AddHandler(StartStatusBarProcess)          // Start status bar service
	session.AddHandler(h.MessageHandler)               // Start command handler
	cmds.RegisterCommands(h)                           // Register commands to the command handler

	// Open discord session
	fmt.Println("Completed loading, opening session")
	err = session.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Close on os interrupt
	handler.WaitForInterrupt()
	_ = session.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
