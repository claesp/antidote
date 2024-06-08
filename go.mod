module antidote

replace github.com/claesp/antidote/libticket => ./libticket

replace github.com/claesp/antidote/libticket/drivers => ./libticket/drivers

replace github.com/claesp/antidote/templates => ./templates

go 1.19

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/claesp/antidote/libticket v0.0.0-00010101000000-000000000000 // indirect
	github.com/claesp/antidote/libticket/drivers v0.0.0-00010101000000-000000000000 // indirect
	github.com/claesp/antidote/templates v0.0.0-00010101000000-000000000000 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.43.0 // indirect
	github.com/valyala/fasthttprouter v0.0.0-20160217050331-24073dd8f323 // indirect
	github.com/valyala/quicktemplate v1.7.0 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
)
