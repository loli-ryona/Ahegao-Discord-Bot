# Ahegao Discord Bot
**Ahegao** is the bot I made for my discord server. I have just started to learn Golang so this will be an ongoing project as I learn more about the language.

This requires [discordgo](https://github.com/bwmarrin/discordgo) and **[Anpan](https://github.com/MikeModder/anpan/)** to work. <br>
This bot is roughly based off **[Anpan](https://github.com/MikeModder/anpan/)** too.

## Usage
Check the `example` directory for an example of the .json config files. <br>
These config files go in `./` by default. <br>
To run `cd` to `./` and execute `go run .`.

## Commands
By default prefixes are `_` and `.`<br>
`_help` - Displays help and info about commands. <br>
`_about` - Displays informations about the bot. <br>
`_players` - Displays players on servers provided in `servers.json`. <br>
`_currentmap` - Displays current map on servers provided in `servers.json`. <br>
`_serverinfo <url.com/ip/ip:port>` - Displays information about provided server. Only works on source games. <br>
`_urban <word>` - Searches Urban Dictioniary for provided word.