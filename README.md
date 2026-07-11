## darkmatter
Simple TUI application to run linux utilities with style.

> [!NOTE]
> - This is meant to run on Arch Linux.

### Installation
I recommend you download the latest release's binary and it's `sha256sum.txt` from \
the [releases page](https://github.com/g5ostXa/darkmatter/releases), so you can verify your download.

#### Using curl
Get the binary and the sha256sum:
```bash
curl -L -O "https://github.com/g5ostXa/darkmatter/releases/download/v0.1.3/darkmatter-v0.1.3-linux-amd64"
curl -L -O "https://github.com/g5ostXa/darkmatter/releases/download/v0.1.3/sha256sum.txt"
```
Verify your download:
```bash
sha256sum -c sha256sum.txt
```

Here's what the output should look like:
```
darkmatter-v0.1.3-linux-amd64: OK
```

#### Using git
Get the full source (latest git)
```bash
git clone --depth=1 https://github.com/g5ostXa/darkmatter.git
```

You'll need to manually build the binary:
```bash
go build -o "darkmatter" ./cmd/darkmatter
```

### Usage
Initialize TUI:
(Assuming the binrary was made executable)

```bash
./darkmatter-v0.1.3-linux-amd64
```

### gHosTTP
Deploy a simple HTTP server locally:

- `static` directory is where static files should go.
- In your browser, visit `http://localhost:8080/` to access the local server.
- To close the server and go back to the main menu, press `CTRL + C`.


### getarch
Quickly download latest archiso and it's signature:

> [!NOTE]
> - You can choose the mirror of your choice for better downloads.
> - See `getarch/main.go`:

```go
const mirror = "https://mirror.quantum5.ca/archlinux/iso/latest/"
```
- To cancel and go back to main menu, press `CTRL + C`.


### glyphs
Unicode symbols on the command line, built using [`huh?`](https://github.com/charmbracelet/huh):

- Add more glyphs easily to `glyphs/glyphs.json`.
- To leave the glyphs menu and go back to the main menu, press `CTRL + C`.

```json
[
  {
    "name": "Arch Linux",
    "icon": "󰣇"
  },
  {
    "name": "Secure",
    "icon": "󰦝 "
  },
  {
    "name": "Digital Key",
    "icon": "󰷖 "
  },
  {
    "name": "Lightning",
    "icon": "󰉁"
  }
]
```
