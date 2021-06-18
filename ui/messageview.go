package ui

import (
	"github.com/cainy-a/gord/discordutil"
	"github.com/cainy-a/gord/tview"
	"regexp"
	"sync"

	linkShortener "github.com/Bios-Marcel/shortnotforlong"
	"github.com/cainy-a/discordgo"
	"github.com/cainy-a/gord/config"
	"github.com/gdamore/tcell/v2"
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
	shortener            *linkShortener.Shortener
	ownUserID            string
	dateTimeFormat       string

	onAction func(event *tcell.EventKey) *tcell.EventKey

	*sync.Mutex

	state *discordgo.State

	// the item to be rendered will be set to this
	uiToRender *tview.Box

	uiFlex      *tview.Flex
	uiTimestamp *tview.TextView
	uiAuthor    *tview.TextView
	uiContent   *tview.TextView

	// if the message is a reply, uiReplyStackFlex will contain uiReply on top and uiFlex underneath
	uiReplyStackFlex *tview.Flex
	uiReply          *tview.Flex
	uiReplyAuthor    *tview.TextView
	uiReplyContent   *tview.TextView
}

// NewMessageView creates a new MessageView ready for use
func NewMessageView(message *discordgo.Message, ownUserID string, shortener *linkShortener.Shortener) {
	messageView := MessageView{
		message:    message,
		ownUserID:  ownUserID,
		isSelected: false,
		//Magic date which defines the format in which all dates will be formatted.
		//While it isn't obvious which one is month and which one is day, this is
		//is still "correctly" inferred as "year-month-day".
		dateTimeFormat:       "2006-01-02",
		shortenLinks:         config.Current.ShortenLinks,
		shortenWithExtension: config.Current.ShortenWithExtension,
		Mutex:                &sync.Mutex{},
	}

	messageView.buildRawUI()

	if messageView.shortenLinks {
		if shortener == nil {
			messageView.shortenLinks = false
		} else {
			messageView.shortener = shortener
		}
	}
}

//////////////////////////////////
// UI BUILDING AND MODIFICATION //
//////////////////////////////////

// buildRawUI builds the UI for a MessageView, and returns relevant UI items: The UI is not populated - call render()
func (messageView *MessageView) buildRawUI() {
	messageView.uiTimestamp = tview.NewTextView().SetTextColor(tcell.ColorGray)
	messageView.uiAuthor = tview.NewTextView()
	messageView.uiContent = tview.NewTextView()
	messageView.uiFlex = tview.NewFlex().
		AddItem(messageView.uiTimestamp, 1, 0, false).
		AddItem(messageView.uiAuthor, 1, 0, false).
		AddItem(messageView.uiContent, 0, 1, true)
}

// makeUIReply will set the structure of the MessageView's UI to have replies included
func (messageView *MessageView) makeUIReply() {
	messageView.uiReplyAuthor = tview.NewTextView()
	messageView.uiReplyContent = tview.NewTextView()

	messageView.uiReply = tview.NewFlex().
		AddItem(messageView.uiReplyAuthor, 1, 0, false).
		AddItem(messageView.uiReplyContent, 0, 1, false)

	messageView.uiReplyStackFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(messageView.uiReply, 1, 0, false).
		AddItem(messageView.uiFlex, 0, 1, true)

	messageView.uiToRender = messageView.uiReplyStackFlex.Box
}

// makeUISimple will set the structure of  the MessageView's UI to not have replies
func (messageView *MessageView) makeUISimple() {
	messageView.uiToRender = messageView.uiFlex.Box

	messageView.uiReplyStackFlex = nil
	messageView.uiReply = nil
	messageView.uiReplyAuthor = nil
	messageView.uiReplyContent = nil
}

///////////////////////////////
// UI CONTENT (RE-)RENDERING //
///////////////////////////////

func (messageView *MessageView) render() {
	// set the appropriate UI structure, and if necessary render the reply
	if messageView.message.ReferencedMessage != nil {
		messageView.makeUIReply()
		messageView.renderReply()
	} else {
		messageView.makeUISimple()
	}

	// render the message (populate relevant UI elements)
	messageView.renderTimestamp()
	messageView.renderAuthor()
	messageView.renderContent()
}

func (messageView *MessageView) renderTimestamp() {
	time, err := messageView.message.Timestamp.Parse()
	if err != nil {
		return
	}
	formatted := time.Format(messageView.dateTimeFormat)

	messageView.uiTimestamp = messageView.uiTimestamp.SetText(formatted)
}

func (messageView *MessageView) renderAuthorBase(message *discordgo.Message) (string, string) {
	// code taken from chatview.go:434 before the rewrite (dfce984d)
	var member *discordgo.Member
	if message.GuildID != "" {
		member, _ = messageView.state.Member(message.GuildID, message.Author.ID)
	}

	var messageAuthor string
	var userColor string
	if member != nil {
		messageAuthor = discordutil.GetMemberName(member)
		userColor = discordutil.GetMemberColor(messageView.state, member)
	}
	if messageAuthor == "" {
		messageAuthor = discordutil.GetUserName(message.Author)
		userColor = discordutil.GetUserColor(message.Author)
	}

	return messageAuthor, userColor
}

func (messageView *MessageView) renderAuthor() {
	author, rawColour := messageView.renderAuthorBase(messageView.message)

	colour := tcell.Color(tcell.Color.Hex(rawColour))
	messageView.uiAuthor = messageView.uiAuthor.SetText(author).SetTextColor(colour)
}

func (messageView *MessageView) renderContent() {
	messageView.uiContent = tview.NewTextView().SetText("TODO: make messages display properly\n" + messageView.message.Content)
}

// reply rendering

func (messageView *MessageView) renderReply() {
	messageView.renderReplyAuthor()
	messageView.renderReplyContent()
}

func (messageView *MessageView) renderReplyAuthor() {
	author, rawColour := messageView.renderAuthorBase(messageView.message)

	colour := tcell.Color(tcell.Color.Hex(rawColour))
	messageView.uiReplyAuthor = messageView.uiReplyAuthor.SetText(author).SetTextColor(colour)
}

func (messageView *MessageView) renderReplyContent() {
	messageView.uiReplyContent = messageView.uiReplyContent.SetText(messageView.message.ReferencedMessage.Content)
}

/////////////////////////////
// SETTING VIEW ATTRIBUTES //
/////////////////////////////

func (messageView *MessageView) SetMessage(message *discordgo.Message) {
	messageView.message = message
	messageView.render()
}

func (messageView *MessageView) SetSelected(selected bool) {
	messageView.isSelected = selected
}

/////////////////////////////
// DISCORD MESSAGE ACTIONS //
/////////////////////////////

func (messageView *MessageView) Delete(window *Window) {
	window.askForMessageDeletion(messageView.message.ID, messageView.isSelected)
}

func (messageView *MessageView) Edit(window *Window) {
	window.startEditingMessage(messageView.message)
}

func (messageView *MessageView) Reply(window *Window) {
	window.currentReplyMsg = messageView.message
}
