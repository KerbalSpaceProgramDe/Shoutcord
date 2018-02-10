<p align="center">
<a href="https://www.kerbalspaceprogram.de"><img src="https://www.kerbalspaceprogram.de/wcf/images/styleLogo-5682f67ae9016cb4825289123e9a68f199b9bd7f.png" alt="Kerbal.de"></a>
</p>

# Shoutcord
A bot that bridges Discord and the WBB Shoutbox, using the [kerbal.de API](https://github.com/KerbalSpaceProgramDe/API)

### Setting up Shoutcord
To build Shoutcord you need to install the Go language tools for your plattform first. Visit [golang.org](https://golang.org) for further instructions.
Depending on your linux distribution (if you use linux) you might be able to automatically install Go through your package manager.

Shoutcord uses the Glide package manager for managing it's dependencies. You need to install this one as well: [glide.sh](https://glide.sh/)

After installing Go and glide, you need to run the following commands to download the shoutcord source code to your GOPATH. If you haven't changed the default setting, your
GOPATH should be a folder named `go` in your home directory. After the download has finished, you need to cd to your GOPATH, and install the required dependencies.

```bash
$ go get github.com/KerbalSpaceProgramDe/Shoutcord
$ cd $GOPATH/src/github.com/KerbalSpaceProgramDe/Shoutcord
$ glide install
```

Finally, build the shoutcord application by running `go build` in the source code directory. The final executable will be named `Shoutcord` or `Shoutcord.exe`, and when you run it, it will look for a
file named `config.yml`. Check `config.example.yml` for the required values.

### Configuration File
```yaml
# The access token for the bot account on Discord
token: "" 

# The ID of the channel that should be synced with the Shoutbox.
channel: ""

# The address of the kerbal.de API
endpoint: ""
```

### License
Licensed under the [MIT license](https://opensource.org/licenses/MIT).