# Ahegao Discord Bot
**Ahegao** is the bot I made for my discord server. I have just started to learn Golang so this will be an ongoing project as I learn more about the language.

This requires [discordgo](https://github.com/bwmarrin/discordgo) to work. <br>
This bot used to be based off **[Anpan](https://github.com/MikeModder/anpan/)**. It now has Anpan integrated into it under the `./handler` folder. This is due to needing to change the way some of Anpan's functions work.

## Usage
1. Check the `example` directory for an example of the .json config files. <br>
2. Fill in the example configs with the your bot token, user id, etc.
3. Move the configs to the `./cfgs` folder. 
4. Install the dependencies by running `go mod tidy`
5. Run the bot by using `cd` to the `./` and execute `go run .`.

## Commands
By default prefixes are `_` and `.`<br>
`_help` - Displays help and info about commands. <br>
`_about` - Displays informations about the bot. <br>
`_players` - Displays players on servers provided in `servers.json`. <br>
`_currentmap` - Displays current map on servers provided in `servers.json`. <br>
`_serverinfo <url.com/ip/ip:port>` - Displays information about provided server. Only works on source games. <br>
`_urban <word>` - Searches Urban Dictioniary for provided word. <br>
`_lenny <expression>` - Returns lenny face matching expression. <br>
`_time` - Displays the time. <br>
`_weather <city>,<state>,<country code> [-c/-f/-k]` - Displays the weather in given area <br>
`_dalle <prompt> [--size=256/512/1024]` - Generates an image based on the given prompt, can optionally specify a size.

## TODO
~~Command for weather<br>~~
Command for BOM Radar<br>
Put status bar updates in a go routine for optimization <br>
