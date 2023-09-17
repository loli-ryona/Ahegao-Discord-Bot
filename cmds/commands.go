package cmds

import (
	"ahegao/framework"
	hnd "ahegao/handler"
	"encoding/json"
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"os"
)

var (
	cfgPath     = "cfgs/config.json"
	serversPath = "cfgs/servers.json"
	lennyPath   = "cfgs/lenny.json"
	srcdsPath   = "cfgs/srcdspaths.json"

	cfg   framework.Config
	srv   framework.Servers
	srcds framework.SrcdsPaths
	lenny framework.LennyExpressions
)

// init = Loads the required config files into the interface.
func init() {
	paths := []string{cfgPath, serversPath, lennyPath, srcdsPath}
	status := []any{&cfg, &srv, &lenny, &srcds}

	for i := 0; i < len(paths); i++ {
		j, err := os.Open(paths[i])
		if err != nil {
			fmt.Println("Error opening file. Error: ", err)
			os.Exit(1)
		}
		if err = json.NewDecoder(j).Decode(status[i]); err != nil {
			fmt.Println("Error decoding file. Error: ", err)
			os.Exit(1)
		}
	}
}

// RegisterCommands = Registers the commands in the handler which allows them to be used.
/*
 * handler = The command handler to register the commands to, accepted as *handler.CommandHandler
 */
func RegisterCommands(handler *hnd.CommandHandler) {
	// Register new commands
	fmt.Println("Registering commands")
	handler.AddCommand("ping", "Check the bot's ping. Used for debugging", []string{"pong"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, PingCommand)
	handler.AddCommand("players", "Lists players on Bhop Servers", []string{"on"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, PlayersCommand)
	handler.AddCommand("about", "Shows bot information", []string{"bot"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, AboutCommand)
	handler.AddCommand("urban", "Search a word on urban dictionary", []string{"ud"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, UrbanCommand)
	handler.AddCommand("currentmap", "List current maps on Bhop Servers", []string{"cm"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, CurrentMapCommand)
	handler.AddCommand("serverinfo", "Shows information of a server. _serverinfo 127.0.0.1:27015", []string{"si"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, ServerInfoCommand)
	handler.AddCommand("lenny", "Replies with the specified lenny.", []string{}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, LennyCommand)
	handler.AddCommand("thetime", "Replies with the time.", []string{"time", "tt"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, TheTimeCommand)
	handler.AddCommand("weather", "Replies with weather for provided location. Execute command without any arguments for more info.", []string{"temp", "bom"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, WeatherCommand)
	handler.AddCommand("dalle", "Generate an image based on a prompt.", []string{"sd", "imggen"}, false, false, dG.PermissionSendMessages, dG.PermissionSendMessages, hnd.CommandTypeEverywhere, DalleCommand)

	// Build help command using the registered commands
	fmt.Println("Building Help command")
	handler.SetHelpCommand("help", []string{}, dG.PermissionSendMessages, dG.PermissionSendMessages, HelpCommand)
}
