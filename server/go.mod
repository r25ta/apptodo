module apptodo.com/server

go 1.22.2

require (
	apptodo.com/constant v0.0.0-00010101000000-000000000000
	apptodo.com/model v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.52.4
)

require (
	github.com/gofiber/template v1.8.3 // indirect
	github.com/gofiber/utils v1.1.0 // indirect
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gofiber/template/html/v2 v2.1.1
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/lib/pq v1.10.9
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace apptodo.com/model => ../model

replace apptodo.com/constant => ../constant
