package ui

var WelcomeText = `

Welcome to version %s of Gord. Below you can see the most
important changes of the last two versions officially released.

[::b]THIS VERSION (2021-06-16)
	- Finally fixed the message receive in other channel empty content / crashes for good :)
[::b]2021-04-07
	- Bugfixes
[::b]2021-04-06
	- Replies
[::b]2020-10-24
	- Features
		- DM people via "p" in the chatview or use the dm-open command
		- Mark guilds as read
		- Mark guild channels as read
		- Write to logfile by setting "--log"
		- Mentions are now displayed in the guild list
		- You can now bulk send folders and files
	- Changes
		- Dialogs shown at the bottom of the chatview now allow tab / backtab
		- There's now a double-colon to separate author and messages
		- There's more customizable shortcuts now
	- Bugfixes
		- Guilds and channels were sometimes falsely seen as muted
		- Deleting / Leaving guilds now properly deletes them from the UI
		- Jumping to guilds / channels you were mentioned in, now works by
		  by typing their name again
		- Fixed deadlock when spamming "Switch to previous channel"
		- "Switch to previous channel" doesn't jumble the state anymore
		  when switching between different guilds and DMs
		- Muted guilds, channels and categories shouldn't be displayed as
		  unread anymore
		- @everyone works again, so you can piss of others again
		- Messages containing links won't disappear anymore after sending
		- Messages from blocked users won't trigger notifications anymore
		- No more spammed empty error messages when receiving notifications
[::b]2020-08-30
	- Features
		- Nicknames can now be disabled via the configuration
		- Files from messages can now be downloaded (key d) or opened (key o)
		- New parameter "--account" to start cordless with a certain account
	- Changes
		- The "friends" command now has "friend" as an alias
		- "logout" is now a separate command, but "account logout" still works
		- Currently active account is now highlight in "account list" output
		- Password input dialog now uses the configured shortcut for paste
		- Baremode
			- Now includes the message input
			- The command view will hide when entering baremode
	- Bugfixes
		- Fix crash due to race condition in readmarker feature
		- Embed-Edits won't be ignored anymore
		- Names with role colors now respect their role order
		- Unread message numbers now always update when loading a channel instead of when leaving it
		- UTF-8 disabling wasn't taken into account when rendering the channel tree
[::b]2020-08-11 - 2020-06-30
	- Features
		- Notifications for servers and DMs are now displayed in the containers header row 
		- Embeds can now be rendered
		- Usernames can now be rendered with their respective role color.
		  Bots however can't have colors, to avoid confusion with real users.
		  The default is set to "single", meaning it uses the default user
		  color from the specified theme. The setting "UseRandomUserColors" has
		  been removed.
	- Changes
		- The button to switch between DMs and servers is gone. Instead you can
		  click the containers, since the header row is always visible now
		- Token input now ingores surrounding spaces
		- Bot token syntax is more lenient now
	-  Bugfixes
		- Bot login works again
		- Holding down your left mouse and moving it on the chatview won't
		  cause lags anymore
		- No more false positives for unread dm-channels
[::b]20-06-26
	- Features
		- you can now define a custom status
		- shortened URLs optionally can display a file suffix (extension)
		- You can now cycle through message in edit-mode by repeatedly hitting KeyUp/Down
	- Bugfixes
		- config directory path now read from "XDF_CONFIG_HOME" instead of "XDG_CONFIG_DIR"
		- the delete message shortcut was pointing to the same value as "show spoilered message"
		- the lack of the config directory would cause a crash
		- nitro users couldn't use emojis anymore
		- several typos have been corrected
		- the "version" command printed it's help output to stdout
		- the "man" command now searches through the content of pages and suggests those
[::b]2020-01-05
	- Features
		- VT320 terminals are now supported
		- quoted messages now preserve attachment URLs
		- Ctrl-W now deletes the word to the left
		- announcement channels are now shown as well
		- Cordless now has an amazing autocompletion
		- support for TFA
		- user-set command allows supplying emojis
		- custom emojis are now rendered as links
		- login now navigable via arrow keys
		- Ctrl-B now toggles the so called "bare mode", giving all space to the chat
		- configuration path is now customizable via parameters
	- Bugfixes
		- emoji sequences with underscores now work
		- text channels sometimes didn't show up'
		- Cordless doesn't crash anymore when sending a message into an empty channel
		- attachment links are now copied as well
	- Performance improvements
		- the usertree will now load lazily
		- dummycall to validate session token has been removed
	- Changes
		- login button has been removed ... just hit enter ;)
		- tokeninput on login is now masked
		- Docs have been improved
	- JS API
		- there's now an "init" function that gets called on script load
`
