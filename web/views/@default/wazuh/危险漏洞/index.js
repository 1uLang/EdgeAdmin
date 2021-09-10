Tea.context(function () {
    
    this.ipSelectIndex = 1
    this.dangerSelectIndex=1

    this.getDangerLevel=function(status){
        switch (status) {
            case 1:
                return "低危"
            case 2:
                return "中危"
            case 3:
                return "高危"
            case 4:
                return "危急"
            default:
                return "未知"
        }
    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.ipData = [
        {id:1,ip:"192.168.0.1",},
        {id:2,ip:"192.168.0.2",},
        {id:3,ip:"192.168.0.3",},
    ]

    this.dangerData = [
        {id:1,name:"低危",},
        {id:2,name:"中危",},
        {id:3,name:"高危",},
        {id:4,name:"危急",},
        {id:5,name:"全部",},
    ]

    this.tableData = [
        {
            id:1,
            hostIP:"192.168.0.1",
            hostName:"WIN-OC98D604V3R1",
            badName:"CVE-2018-1000035 on Ubuntu 18.04 LTS (bionic) - low.",
            dangerLevel:1,
            CVEStyle:"CVE-2019-17540",
            effectPackage:"libmagickcore-6.q16-3",
            findTime:"2021-02-02T15:25:30.123",
        },
        {
            id:2,
            hostIP:"192.168.0.2",
            hostName:"WIN-OC98D604V3R2",
            badName:"CVE-2018-1000035 on Ubuntu 18.04 LTS (bionic) - low.",
            dangerLevel:2,
            CVEStyle:"CVE-2019-17540",
            effectPackage:"libmagickcore-6.q16-3",
            findTime:"2021-02-02T15:25:30.123",
        },
        {
            id:3,
            hostIP:"192.168.0.3",
            hostName:"WIN-OC98D604V3R3",
            badName:"CVE-2018-1000035 on Ubuntu 18.04 LTS (bionic) - low.",
            dangerLevel:3,
            CVEStyle:"CVE-2019-17540",
            effectPackage:"libmagickcore-6.q16-3",
            findTime:"2021-02-02T15:25:30.123",
        },
        {
            id:4,
            hostIP:"192.168.0.4",
            hostName:"WIN-OC98D604V3R4",
            badName:"CVE-2018-1000035 on Ubuntu 18.04 LTS (bionic) - low.",
            dangerLevel:4,
            CVEStyle:"CVE-2019-17540",
            effectPackage:"libmagickcore-6.q16-3",
            findTime:"2021-02-02T15:25:30.123",
        }

    ]
})