package scraper

import (
	"github.com/gtuk/discordwebhook"
	"github.com/jakubje/property-aggregator/models"
	"github.com/jakubje/property-aggregator/util"
	"github.com/rs/zerolog/log"
)

func sendWebhook(config util.Config, discordData *models.NewPropertyNotification) {
	var username = "Aggregator"
	var url = config.DiscordWebhook
	var image_url = discordData.Image

	image := discordwebhook.Image{
		Url: &image_url,
	}
	price := "Price"

	fields := []discordwebhook.Field{
		{
			Name:   &price,
			Value:  &discordData.Price,
			Inline: nil,
		},
	}

	embed := discordwebhook.Embed{
		Title:  &discordData.Title,
		Url:    &discordData.Url,
		Fields: &fields,
		Image:  &image,
		Footer: &discordwebhook.Footer{
			Text:    &discordData.Website,
			IconUrl: nil,
		},
	}

	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &[]discordwebhook.Embed{embed},
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
