package component

import (
	"DiscordTemplate/internal/shared"

	"github.com/bwmarrin/discordgo"
)

type Component interface {
	CustomID() string
	Component() discordgo.MessageComponent
	Execute(args *shared.CmdArgs) (*discordgo.InteractionResponse, error)
}
