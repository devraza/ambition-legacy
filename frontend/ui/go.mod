module github.com/devraza/ambition/frontend/ui

go 1.20

require (
	github.com/devraza/ambition/frontend/assets/fonts v0.0.0-00010101000000-000000000000
	github.com/devraza/ambition/frontend/player v0.0.0-00010101000000-000000000000
	github.com/ebitenui/ebitenui v0.5.4
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	golang.org/x/image v0.9.0
)

require (
	github.com/devraza/ambition/frontend v0.0.0-20230804135652-eb4d84965aa5 // indirect
	github.com/ebitengine/purego v0.3.2 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20221017161538-93cebf72946b // indirect
	github.com/hajimehoshi/ebiten/v2 v2.5.5 // indirect
	github.com/jezek/xgb v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/exp/shiny v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/mobile v0.0.0-20230427221453-e8d11dd0ba41 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)

replace github.com/devraza/ambition/frontend/assets/fonts => ../assets/fonts

replace github.com/devraza/ambition/frontend/player => ../player
