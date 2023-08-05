module github.com/devraza/ambition/frontend

go 1.20

require (
	github.com/devraza/ambition/frontend/player v0.0.0-00010101000000-000000000000
	github.com/devraza/ambition/frontend/ui v0.0.0-00010101000000-000000000000
	github.com/hajimehoshi/ebiten/v2 v2.5.5
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.7.1 // indirect
	github.com/charmbracelet/log v0.2.3 // indirect
	github.com/devraza/ambition/frontend/assets/fonts v0.0.0-00010101000000-000000000000 // indirect
	github.com/devraza/ambition/frontend/server v0.0.0-00010101000000-000000000000 // indirect
	github.com/ebitengine/purego v0.3.2 // indirect
	github.com/ebitenui/ebitenui v0.5.4 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20221017161538-93cebf72946b // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/jezek/xgb v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/image v0.9.0 // indirect
	golang.org/x/mobile v0.0.0-20230427221453-e8d11dd0ba41 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

replace (
	github.com/devraza/ambition/frontend/assets/fonts => ./assets/fonts
	github.com/devraza/ambition/frontend/player => ./player
	github.com/devraza/ambition/frontend/server => ./server
	github.com/devraza/ambition/frontend/ui => ./ui
)
