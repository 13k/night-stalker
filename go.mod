module github.com/13k/night-stalker

go 1.13

require (
	cirello.io/oversight v1.0.3
	github.com/13k/geyser v0.2.0
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/go-units v0.4.0
	github.com/faceit/go-steam v0.0.0-20190206021251-2be7df6980e1
	github.com/fatih/color v1.9.0
	github.com/galaco/KeyValues v1.4.1
	github.com/go-logfmt/logfmt v0.5.0
	github.com/go-redis/redis/v7 v7.2.0
	github.com/go-resty/resty/v2 v2.2.0
	github.com/go-stack/stack v1.8.0
	github.com/golang-migrate/migrate/v4 v4.9.1
	github.com/golang/protobuf v1.3.4
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.1
	github.com/jinzhu/gorm v1.9.12
	github.com/labstack/echo/v4 v4.1.15
	github.com/labstack/gommon v0.3.0
	github.com/lib/pq v1.3.0
	github.com/markbates/pkger v0.14.1
	github.com/mattn/go-isatty v0.0.12
	github.com/mitchellh/protoc-gen-go-json v0.0.0-20200113165135-fd297ce346f1
	github.com/olebedev/emitter v0.0.0-20190110104742-e8d1457e6aee
	github.com/panjf2000/ants/v2 v2.3.1
	github.com/paralin/go-dota2 v0.0.0-20191126225751-cae5acd7b08d
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
	gopkg.in/inconshreveable/log15.v2 v2.0.0-20200109203555-b30bc20e4fd1
)

replace github.com/paralin/go-dota2 v0.0.0-20191126225751-cae5acd7b08d => github.com/13k/go-dota2 v0.0.0-20200307085842-ca26575af454
