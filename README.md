## darkmatter
Simple TUI application to run linux utilities with style.

> [!NOTE]
> - This is meant to run on Arch Linux.

### Installation
Clone the repo:

```bash
git clone --depth=1 https://github.com/g5ostXa/darkmatter.git
```

### Usage
Initialize TUI:

```bash
./init.sh
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
