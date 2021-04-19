module ego-grpc-example

go 1.16

require (
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/gotomicro/ego v0.5.3
	github.com/gotomicro/ego-component/eetcd v0.1.5-0.20210413100321-4c4b6ad1bac9
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/gotomicro/ego => ../../gotomicro/ego
	github.com/gotomicro/ego-component/eetcd => ../../gotomicro/ego-component/eetcd
)
