# Discord Fortune Bot

Just a small [Discord](https://discord.gg) chat bot that responds to `/fortune`
commands using [Unix fortunes](https://en.wikipedia.org/wiki/Fortune_(Unix)).

    Has anyone realized that the purpose of the fortune cookie program is to
    defuse project tensions?  When did you ever see a cheerful cookie, a
    non-cynical, or even an informative cookie?
            Perhaps inadvertently, we have a channel for our aggressions.  This
    still begs the question [sic] of whether the cookie releases the pressure or only
    serves to blunt the warning signs.

To use it yourself, visit this [link](https://discordapp.com/api/oauth2/authorize?client_id=316630741570027521&scope=bot&permissions=0).

To run your own instance:

    $ discord-fortune-bot -t <APP_TOKEN>

This depends on the `fortune` and `fortune -o` commands being available and in
your PATH.
