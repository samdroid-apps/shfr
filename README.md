# How To Setup

First, open the `forums.json` file.  In there, change the `url` to the public url of your forum instance (eg: `socialhelp.example.com`).  Then change the `forums` mapping to fit your needs.  It is in the format of bundle id to forum name.

> **Pro Tip**: SHFR will automatically reload in response to a change in the `forums.json` file.

Then run the software.  If you are on an x64 Linux box, just run:

    ./shfr

Otherwise you will need to [install go](http://golang.org/doc/install) and run before you can run the previous command:

    go build

If go build has error messages relating to dependencies, try running before trying again:

    go get dependency

# HTTP Api

    GET /goto/ID

Redirect to the forum for ID.  If ID is not in this forum, it will be logged and sent to the error page.

    GET /records

Get a json map in the format of bundle ID to number of times requested and not found.

    PATCH /records

Save the current json map to the disk (useful before server restarts)

    GET /info

Get some useful info about this server, eg:

    {
      "forums": {
        "org.laptop.AbiWordActivity": "write", 
        "org.laptop.Chat": "chat", 
        "org.laptop.Oficina": "paint", 
        "org.laptop.Pippy": "pippy", 
        "org.laptop.Terminal": "terminal", 
        "org.laptop.WebActivity": "browse", 
        "org.laptop.physics": "physics", 
        "org.sugarlabs.MusicKeyboard": "music-keyboard", 
        "org.tuxpaint": "tuxpaint", 
        "vu.lux.olpc.Speak": "speak"
      }, 
      "url": "http://54.187.40.150"
    }
