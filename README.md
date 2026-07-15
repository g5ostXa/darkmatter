# darkmatter
Simple TUI application to run linux utilities with style.

> [!NOTE]
> - This is meant to run on Arch Linux.

## Installation
I recommend you download the latest release's binary and it's `sha256sum.txt` from \
the [releases page](https://github.com/g5ostXa/darkmatter/releases), so you can verify your download.

#### Get the binary and the sha256sum using curl:
```bash
curl -L -O "https://github.com/g5ostXa/darkmatter/releases/download/v0.1.5/darkmatter-v0.1.5-linux-amd64"
curl -L -O "https://github.com/g5ostXa/darkmatter/releases/download/v0.1.5/sha256sum.txt"
```
Verify your download:
```bash
sha256sum -c sha256sum.txt
```

Here's what the output should look like:
```
darkmatter-v0.1.5-linux-amd64: OK
```

#### Using git:
Get the full source (latest git)
```bash
git clone --depth=1 https://github.com/g5ostXa/darkmatter.git
```

<br>

## Usage
**Initialize TUI**

First, make sure the binary is executable:
```bash
chmod +x darkmatter-v0.1.5-linux-amd64
```

You can rename the binary to nake you life simpler:
```bash
mv darkmatter-v0.1.5-linux-amd64  "new-name"
```

I recommend putting the binary anywhere on your `$PATH`, so you can run it \
from anywhere just by typing the binary name.

Otherwise, cd into the directory where you store the binary and run it:
```bash
./darkmatter-v0.1.5-linux-amd64
```

<br>

## gHosTTP
Deploy a simple HTTP server locally:

> [!TIP]
> - The application will prompt you to specify the full path to serve your website locally.
> - As for an example: `/home/user/path/to/website-folder`
> - In your browser, visit `http://localhost:8080/` to access the local server.
> - To close the server and go back to the main menu, press `CTRL + C`.

<br>

## getarch
Quickly download latest archiso and it's signature

- You can choose the mirror of your choice for better downloads.
- See `getarch/main.go`:

```go
const mirror = "https://mirror.quantum5.ca/archlinux/iso/latest/"
```

<br>

## glyphs
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