module github.com/TeaOSLab/EdgeAdmin

go 1.16

replace github.com/1uLang/zhiannet-api => ../zhiannet-api

replace github.com/TeaOSLab/EdgeCommon => ../EdgeCommon

require (
	github.com/1uLang/zhiannet-api v0.0.0-00010101000000-000000000000
	github.com/TeaOSLab/EdgeCommon v0.0.0-00010101000000-000000000000
	github.com/cespare/xxhash v1.1.0
	github.com/dlclark/regexp2 v1.4.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/iwind/TeaGo v0.0.0-20210809112119-a57ed0e84e34
	github.com/iwind/gosock v0.0.0-20210722083328-12b2d66abec3
	github.com/miekg/dns v1.1.35
	github.com/sirupsen/logrus v1.8.1
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/xlzd/gotp v0.0.0-20181030022105-c8557ba2c119
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
