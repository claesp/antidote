module libticket

go 1.19

replace github.com/claesp/antidote/libticket/drivers => ./drivers

require (
	github.com/claesp/antidote/libticket/drivers v0.0.0-00010101000000-000000000000 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/sys v0.0.0-20200923182605-d9f96fdee20d // indirect
)
