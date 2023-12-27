module github.com/WormholeMapper/backend

go 1.21.5

replace github.com/WormholeMapper/esiclient => ./pkg/esiclient

require (
	github.com/WormholeMapper/esiclient v0.0.0-00010101000000-000000000000
	github.com/WormholeMapper/whmcfg v0.0.0-00010101000000-000000000000
)

require (
	github.com/antihax/goesi v0.0.0-20231202031403-498c84423fe8 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20230905024940-24af94b03874 // indirect
	github.com/golang-jwt/jwt/v4 v4.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211028175245-ba495a64dcb5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/WormholeMapper/whmdb => ./pkg/whmdb

replace github.com/WormholeMapper/whmcfg => ./pkg/whmcfg
