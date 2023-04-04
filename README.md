# 🌐 nostrfetch

Share your [Nostr](https://github.com/nostr-protocol/nostr) profile in your screenshots on [r/unixporn](https://www.reddit.com/r/unixporn/)

![nostrfetch_showcase](https://user-images.githubusercontent.com/69465962/229657789-5b5bca34-293e-4e44-95b2-94d7971011b3.png)
![nostrfetch_showcase](https://user-images.githubusercontent.com/69465962/229674237-abcff547-62d3-4110-93ac-ec767d9c5968.png)

## 📝 Notes
- Profile preview is only supported on [Kitty](https://github.com/kovidgoyal/kitty) terminal emulator.
- The image preview is displayed on top of the text rather than to the side; this is a feature, not a bug.

## 🚀 Installation

You can download pre-built binaries for various operating systems from the [releases](https://github.com/MAHcodes/nostrfetch/releases) page.  

Alternatively, you can build the application from source by following the instructions in the next section.

## 🔨 Build from Source

To build nostrfetch from source, follow these steps:

1. Clone the repository: `git clone https://github.com/MAHcodes/nostrfetch.git`
2. Navigate to the project directory: `cd nostrfetch`
3. Install dependencies: `go get ./...`
4. Build the application: `go build -o nostrfetch main.go`

This will build the `nostrfetch` binary in the current directory.

You can also use `go install` to install the binary to your `$GOPATH/bin` directory:
```
go install .
```

This will install the `nostrfetch` binary to `$GOPATH/bin`.

Note that building from source requires Go to be installed on your machine. You can download the latest version of Go from the [official website](https://golang.org/dl/).

## 📖 Usage

1. Create a configuration file at `$XDG_CONFIG_HOME/nostrfetch/config.yaml`. If `$XDG_CONFIG_HOME` is not set, you can use `~/.config` instead.  
Here's an example configuration file:
```yaml
# The URL of the Nostr relay to use
relayURL: wss://relay.damus.io

# The npub public key for the profile you want to fetch
npub: npub1qd3hhtge6vhwapp85q8eg03gea7ftuf9um4r8x4lh4xfy2trgvksf6dkva
```
2. Run `nostrfetch`

## 👥 Contributing

Pull requests are welcome! Please follow these guidelines:

1. Fork the repo
2. Create a new branch for your feature (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new pull request

## 📝 License

This project is licensed under the terms of the [GNU General Public License v3.0](LICENSE).

## 🧑 Authors

- [@MAHcodes](https://github.com/MAHcodes/nostrfetch)

## 🙏 Acknowledgments

- [fiatjaf](https://github.com/fiatjaf)
- [xdg-go](https://github.com/casimir/xdg-go)
- [dolmen-go/kittyimg](https://github.com/dolmen-go/kittyimg)
- [i582/cfmt](https://github.com/i582/cfmt)
