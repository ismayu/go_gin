module github.com/EDDYCJY/go-gin-example

go 1.14

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => /d/go_blog/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => /d/go_blog/middleware
	github.com/EDDYCJY/go-gin-example/models => /d/go_blog/models
	github.com/EDDYCJY/go-gin-example/pkg/e => /d/go_blog/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => /d/go_blog/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => /d/go_blog/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => /d/go_blog/routers

)
