module github.com/devraza/ambition/frontend/player

go 1.20

require github.com/devraza/ambition/frontend/server v0.0.0-00010101000000-000000000000

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.7.1 // indirect
	github.com/charmbracelet/log v0.2.3 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

replace github.com/devraza/ambition/frontend/server => ../server
