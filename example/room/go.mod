module example.com/room

go 1.23.5

replace github.com/iyear/biligo-live => ../..

require (
	github.com/Drelf2018/go-bilibili-api v0.0.0-20250115173359-0c0d3ea89d01
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/iyear/biligo-live v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/Drelf2018/req v0.0.0-20250101101837-5b5beb05767c // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
