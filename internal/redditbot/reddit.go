package redditbot

import (
	"fmt"
	"github.com/matt1484/bl3-auto-reddit-bot/internal/bot"
	"github.com/rs/zerolog/log"
	"github.com/turnage/graw/reddit"
)

type CodeBot struct {
	Bot reddit.Bot
}

func (c *CodeBot) Mention(p *reddit.Message) error {
	params := getParams(`^(?:u/)?bl3-code-bot (?P<codetype>\w+) (?P<code>\w+)$`, p.Body)
	if params == nil {
		log.Warn().Msg("didn't match")
		return nil
	}

	log.Debug().Msg(fmt.Sprintf("Received message from %s: %s ", p.Author, p.Body))
	log.Debug().Msg(fmt.Sprintf("params: %#v ", params))

	bot.AddCode(params["codetype"], params["code"])

	return nil
}

func (c *CodeBot) Message(p *reddit.Message) error {
	params := getParams(`^(?:u/)?bl3-code-bot (?P<codetype>\w+) (?P<code>\w+)$`, p.Body)
	if params == nil {
		params = getParams(`^(?P<codetype>\w+) (?P<code>\w+)$`, p.Body)
		if params == nil {
			log.Warn().Msg("didn't match")
			return nil
		}
	}

	log.Debug().Msg(fmt.Sprintf("Received message from %s: %s ", p.Author, p.Body))
	log.Debug().Msg(fmt.Sprintf("params: %#v ", params))

	bot.AddCode(params["codetype"], params["code"])

	return nil
}