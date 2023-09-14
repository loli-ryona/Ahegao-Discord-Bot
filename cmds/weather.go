package cmds

import (
	js "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	fwk "ahegao/framework"

	ap "ahegao/handler"
	dG "github.com/bwmarrin/discordgo"
)

var (
	cfg fwk.Config
)

type ApiResp struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone   int    `json:"timezone"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Cod        int    `json:"cod"`
	CodMessage string `json:"message"`
}

func WeatherCommand(ctx ap.Context, args []string) error {
	//vars
	ts := time.Now()
	measurement := "celsius"

	embed := &dG.MessageEmbed{
		Title:       fmt.Sprintf("Weather: Searching..."),
		Description: "Please wait while we query the API for info",
		Footer: &dG.MessageEmbedFooter{
			Text:    "Calculating time to query servers.",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	msg, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}

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

	if len(args) >= 1 {
		if len(args) >= 2 {
			if args[1] == "-c" {
				measurement = "celsius"
			}
			if args[1] == "-f" {
				measurement = "fahrenheit"
			}
			if args[1] == "-k" {
				measurement = "kelvin"
			}
		}
		q := args[0]
		qurl := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", q, cfg.OpenWeatherAPI)

		weatherResp := &ApiResp{}

		u, err := url.Parse(qurl)
		if err != nil {
			fmt.Println("Error parsing OpenWeatherAPI. Error: ", err)
			return err
		}

		fmt.Println(u)

		r, err := http.Get(u.String())
		if err != nil {
			fmt.Println("Error getting API response. Error: ", err)
			return err
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading response body. Error: ", err)
			return err
		}

		if err := js.Unmarshal(body, &weatherResp); err != nil {
			fmt.Println("Error unmarshalling json. Error: ", err)
			ed := &dG.MessageEmbed{
				Title:       fmt.Sprintln("Weather: ", q),
				Description: fmt.Sprintln("The provided query did not work."),
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Took %.2fs to query API!", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}

			_, err := ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
			if err != nil {
				return err
			}
			return err
		}

		_ = fmt.Sprintf("Json Results: %+v\n", weatherResp)

		embed := &dG.MessageEmbed{
			Title:       fmt.Sprintf("Weather: %v", q),
			Description: fmt.Sprintf("The weather is currently **%v** with **%v**", weatherResp.Weather[0].Main, weatherResp.Weather[0].Description),
			Thumbnail: &dG.MessageEmbedThumbnail{
				URL: fmt.Sprintf("http://openweathermap.org/img/wn/%v@2x.png", weatherResp.Weather[0].Icon),
			},
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintf("Took %.2fs to query API!", time.Since(ts).Seconds()),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}

		embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
			Name:  "Temperature",
			Value: fmt.Sprintf("It is currently: **%.2f**\nIt currently feels like: **%.2f**\nToday is a minimum of **%.2f** and a maximum of **%.2f**.", tempConvert(weatherResp.Main.Temp, measurement), tempConvert(weatherResp.Main.FeelsLike, measurement), tempConvert(weatherResp.Main.TempMin, measurement), tempConvert(weatherResp.Main.TempMax, measurement)),
		})

		embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
			Name:  "Stats",
			Value: fmt.Sprintf("The humidity is: **%v%%**\nThe pressure is: **%vhPa**\nThe visibility is: **%v meters**\nThe wind is blowing **%vm/s** towards **%v°**", weatherResp.Main.Humidity, weatherResp.Main.Pressure, weatherResp.Visibility, weatherResp.Wind.Speed, weatherResp.Wind.Deg),
		})

		ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, embed)

	} else {
		fmt.Println("No expression provided")
		ed := &dG.MessageEmbed{
			Title:       "Weather",
			Description: fmt.Sprintln("Please provide atleast a city, state/county and country are optional. Please note that not every location will work, if a location doesnt work try choosing a bigger city/town near the one that didnt work\n *.weather <city>,<state>,<country code>*"),
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintf("Took %.2fs to return!", time.Since(ts).Seconds()),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}

		ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
			Name:  "City",
			Value: "If only a city is provided, it should aim for the most popular city out of all the cities with that name. You can also provide a country code to be more specific.\n *.weather melbourne*\n *.weather melbourne,au*",
		})

		ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
			Name:  "State/County",
			Value: "When provided with a state/county/region/prefecture code, a country must and a city must also be provided.\n *.weather melbourne,vic,au*",
		})

		ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
			Name:  "Country",
			Value: "When specifying a city, you can specify a country by its code to make the search more specific.\n *.weather constantine,gb*",
		})

		ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
			Name:  "Measurements",
			Value: "You can specifiy the temperature measurement to be used by adding either -c, -k, or -f to the end.\n *.weather noojee,vic,au -f*",
		})

		_, err := ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
		if err != nil {
			return err
		}
	}

	return nil
}

func tempConvert(k float64, measurement string) float64 {
	value := k
	if measurement == "celsius" {
		value = k - 273.15
	} else if measurement == "kelvin" {
		value = k
	} else if measurement == "fahrenheit" {
		value = (k-273.15)*1.8 + 32
	}

	return value
}
