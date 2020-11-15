module github.com/ROGGER1808/go-gin-example

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/spec v0.19.12 // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.6.9
	github.com/ugorji/go v1.2.0 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201113233024-12cec1faf1ba // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201114224030-61ea331ec02b // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	github.com/ROGGER1808/go-gin-example/conf => ./conf
	github.com/ROGGER1808/go-gin-example/middleware => ./middleware
	github.com/ROGGER1808/go-gin-example/models => ./models
	github.com/ROGGER1808/go-gin-example/pkg/setting => ./pkg/setting
	github.com/ROGGER1808/go-gin-example/runtime => ./runtime
)
