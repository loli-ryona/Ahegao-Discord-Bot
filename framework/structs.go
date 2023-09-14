package framework

/*Structure for ./cfgs/config.json to be loaded into.
 * AuthID = bot token.
 * Prefixes = prefix needed to execute commands.
 * Owner = bot owner.
 * IgnoreBots = ignore messages from other bots.
 * RespondToPings = allows bot to respond to pins if setup.
 * CheckPermissions = allows bot to check permissions of user executing command.
 * OpenWeatherAPI = API key for OpenWeather if you intended on using the weather command.
 */
type Config struct {
	AuthID           string   `json:"AuthID"`
	Prefixes         []string `json:"prefixes"`
	Owner            []string `json:"owner"`
	IgnoreBots       bool     `json:"ignorebots"`
	RespondToPings   bool     `json:"respondtopings"`
	CheckPermissions bool     `json:"checkpermissions"`
	OpenWeatherAPI   string   `json:"openweatherapi"`
}

/*Structure for ./cfgs/servers.json to be loaded into.
 * Name = name of the server.
 * Addr = server domain.
 */
type Servers struct {
	Name []string `json:"name"`
	Addr []string `json:"addr"`
}

/*Structure for ./cfgs/statusbar.json to be loaded into.
 * Name = name that is displayed on the status bar.
 * Addr = server domain.
 * Id = id of the channel to update.
 */
type StatusBarServers struct {
	Name []string `json:"name"`
	Addr []string `json:"addr"`
	Id   []string `json:"id"`
}

/*Structure for ./cfgs/lenny.json to be loaded into.
 * Expression = the name of the expression.
 * Face = the lenny associated with the expression.
 */
type LennyExpressions struct {
	Expression []string `json:"expression"`
	Face       []string `json:"face"`
}

/*Structure for ./cfgs/srcdspaths.json to be loaded into.
 * MapDir = directory to download source maps to.
 * FastDL = domain of the server hosting the maps.

Note: this is used for the download command which currently works but is not enabled by default due to it been very WIP
*/
type SrcdsPaths struct {
	MapDir string `json:"mapdir"`
	FastDL string `json:"fastdl"`
}
