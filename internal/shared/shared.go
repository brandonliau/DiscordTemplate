package shared

import (
	"github.com/bwmarrin/discordgo"
)

const (
	Blue  = 0x5865f2
	Green = 0x2dcc70
	Red   = 0xe74d3b
)

type CmdArgs struct {
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
	UserID      string
}

func Response(t discordgo.InteractionResponseType, d *discordgo.InteractionResponseData) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: t,
		Data: d,
	}
}

func ContentResponse(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: c,
		},
	}
}

func EphemeralContentResponse(c string) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: c,
		},
	}
}

func EmbedResponse(e *discordgo.MessageEmbed) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{e},
		},
	}
}

func EphemeralEmbedResponse(e *discordgo.MessageEmbed) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:  discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{e},
		},
	}
}

func AddComponent(rsp *discordgo.InteractionResponse, c ...discordgo.MessageComponent) {
	rsp.Data.Components = []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: c,
		},
	}
}
