package ui

import (
	"regexp"
	"sync"

	linkshortener "github.com/Bios-Marcel/shortnotforlong"
	"github.com/cainy-a/discordgo"
	"github.com/cainy-a/gord/config"
	tcell "github.com/gdamore/tcell/v2"
)

var (
	successiveCustomEmojiRegex = regexp.MustCompile("<a?:.+?:\\d+(><)a?:.+?:\\d+>")
	customEmojiRegex           = regexp.MustCompile("(?sm)(.?)<(a?):(.+?):(\\d+)>(.?)")
	codeBlockRegex             = regexp.MustCompile("(?sm)(^|.)?(\x60\x60\x60(.*?)?\n(.+?)\x60\x60\x60)($|.)")
	colorRegex                 = regexp.MustCompile("\\[#.{6}\\]")
	channelMentionRegex        = regexp.MustCompile(`<#\d*>`)
	urlRegex                   = regexp.MustCompile(`<?(https?://)(.+?)(/.+?)?($|\s|\||>)`)
	spoilerRegex               = regexp.MustCompile(`(?s)\|\|(.+?)\|\|`)
	roleMentionRegex           = regexp.MustCompile(`<@&\d*>`)
)

// MessageView is used to render a single message in a channel, intended to be used inside of a ChatView
type MessageView struct {
	message              *discordgo.Message
	isSelected           bool
	showSpoilerContent   bool
	shortenLinks         bool
	shortenWithExtension bool
	shortener            *linkshortener.Shortener
	ownUserID            string
	format               string
	
	onAction func(event *tcell.EventKey) *tcell.EventKey

	*sync.Mutex
}

// NewMessageView creates a new MessageView ready for use
func NewMessageView(message *discordgo.Message, ownUserID string) {
	messageView := MessageView {
		message: message,
		ownUserID: ownUserID,
		isSelected: false,
		//Magic date which defines the format in which all dates will be formatted.
		//While it isn't obvious which one is month and which one is day, this is
		//is still "correctly" inferred as "year-month-day".
		format: "2006-01-02",
		shortenLinks: config.Current.ShortenLinks,
		shortenWithExtension: config.Current.ShortenWithExtension,
		Mutex: &sync.Mutex{},
	}

	if messageView.shortenLinks {
		messageView.shortener = linkshortener.NewShortener(config.Current.ShortenerPort)
		go func ()  {
			shortenerErr := messageView.shortener.Start()
			if shortenerErr != nil {
				//Disable shortening in case of start failure.
				messageView.shortenLinks = false
			}
		}()
	}
}