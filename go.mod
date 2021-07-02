module github.com/TeaOSLab/EdgeAdmin

go 1.15

replace github.com/TeaOSLab/EdgeCommon => ../EdgeCommon

require (
	github.com/1uLang/zhiannet-api v0.0.0-20210702022407-3eff3ee726f5
	github.com/PuerkitoBio/goquery v1.7.0 // indirect
	github.com/TeaOSLab/EdgeCommon v0.0.0-00010101000000-000000000000
	github.com/cespare/xxhash v1.1.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/iwind/TeaGo v0.0.0-20210628135026-38575a4ab060
	github.com/miekg/dns v1.1.35
	github.com/sirupsen/logrus v1.8.1
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/tealeg/xlsx/v3 v3.2.3
	github.com/xlzd/gotp v0.0.0-20181030022105-c8557ba2c119
	golang.org/x/sys v0.0.0-20210330210617-4fbd30eecc44
	google.golang.org/grpc v1.32.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)
