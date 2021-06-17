module shippy

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/broker/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/nats-io/gnatsd v1.4.1 // indirect
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
)
