package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

const (
	//MinimumCharactersOnID ...
	MinimumCharactersOnID int = 16
)

var (
	//RegexUserPatternID ...
	RegexUserPatternID *regexp.Regexp = regexp.MustCompile(fmt.Sprintf(`^(<@!(\d{%d,})>)$`, MinimumCharactersOnID))
)

func userMessageHandler(s *discordgo.Session, m *discordgo.Message) {
	duckMatch, _ := regexp.MatchString(".*[Qq][Uu][Aa][Cc][Kk]*.", m.Content)
	if duckMatch {
		handleQuack(s, m)

	}
	pointsData := extractPlusMinusEventData(m.Content)
	if pointsData != nil {
		item := pointsData[0]
		operation := pointsData[1]
		user, _ := s.User(item)
		if operation == "++" || operation == "--" {
			handlePlusMinus(item, operation, s, m, user)
		}
		return
	}

	parameters := strings.Split(m.Content, " ")
	if RegexUserPatternID.MatchString(parameters[1]) {
		println("Someone Mentioned Us!")
		s.ChannelMessageSend(m.ChannelID, "Quack!")
	}

}

func handleQuack(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "Quack!")
	return
}

func handlePlusMinus(item string, operation string, s *discordgo.Session, m *discordgo.Message, user *discordgo.User) {
	println("Updating Score for" + item)
	if user == nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%[1]s has %[2]d points", item, updateScore(item, operation, m.GuildID)))
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%[1]s> has %[2]d points", item, updateScore(item, operation, m.GuildID)))
	}

}
