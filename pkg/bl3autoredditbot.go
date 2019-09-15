package bl3autoredditbot

import (
	"fmt"
	"github.com/matt1484/bl3-auto-reddit-bot/internal/redditbot"
	"github.com/rs/zerolog/log"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

//func setup() (database.DB, error) {
//
//}

func Run(connectionString str) {
	log.Info().Msg("Starting Reddit bot")

	if bot, err := reddit.NewBotFromAgentFile("/Users/brad/Development/BradLugo/bl3_auto_reddit_bot/codeBot.agent", 0); err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to create bot handle: %s", err))
	} else {
		log.Info().Msg("Running bot...")

		cfg := graw.Config{Messages: true, Mentions: true}
		handler := &redditbot.CodeBot{Bot: bot}

		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			log.Fatal().Msg(fmt.Sprintf("Failed to start graw run: %s", err))
		} else {
			log.Fatal().Msg(fmt.Sprintf("graw run failed: %s", wait()))
		}
	}

	log.Fatal().Msg("Reddit bot is stopping")
}


