module github.com/ROGGER1808/go-gin-example

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	github.com/ROGGER1808/go-gin-example/conf => ./conf
	github.com/ROGGER1808/go-gin-example/middleware => ./middleware
	github.com/ROGGER1808/go-gin-example/models => ./models
	github.com/ROGGER1808/go-gin-example/pkg/setting => ./pkg/setting
	github.com/ROGGER1808/go-gin-example/runtime => ./runtime
)
