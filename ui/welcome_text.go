package ui

var WelcomeText = `

Welcome to version %s of Gord. Below you can see the most
important changes of the last 5 versions officially released.

[::b]THIS VERSION (2021-07-14)
	- Replies are cancelled with esc
	- Reply mentions are toggled with ctrl-r
	- Mysterious crash reported by Alyxia fixed
[::b]2021-06-17
	!! EMERGENCY RELEASE !!
		- The severity of the bug fixed here (segfaults) was enough that I think it justifies a new release.
		  Usually I would never dream about releasing at this frequency.
	- [#25] Fix random segfaulting (oops)
	- Make Gord a bit smarter about privileged intents, they're now optional for bots.
[::b]2021-06-16
	- Finally fixed the message receive in other channel empty content / crashes for good :)
[::b]2021-04-07
	- Bugfixes
[::b]2021-04-06
	- Replies
`
