Tea.context(function () {
    
    this.ipSelectIndex = 1

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

    this.tableData = [
        {
            id:1,
            hostIP:"192.168.0.1",
            hostName:"WIN-OC98D604V3R1",
            shaValue:"ee68269a278bb549cef2fe2aeac21e6ce7c1d60d ",
            pathValue:"/tmp/virustotal/notavirus   ",
            findTime:"2021-02-02T15:25:30.123",
        },
        {
            id:2,
            hostIP:"192.168.0.2",
            hostName:"WIN-OC98D604V3R2",
            shaValue:"ee68269a278bb549cef2fe2aeac21e6ce7c1d60d ",
            pathValue:"/tmp/virustotal/notavirus   ",
            findTime:"2021-02-02T15:25:30.123",
        },
        {
            id:3,
            hostIP:"192.168.0.3",
            hostName:"WIN-OC98D604V3R3",
            shaValue:"ee68269a278bb549cef2fe2aeac21e6ce7c1d60d ",
            pathValue:"/tmp/virustotal/notavirus   ",
            findTime:"2021-02-02T15:25:30.123",
        },
        {
            id:4,
            hostIP:"192.168.0.4",
            hostName:"WIN-OC98D604V3R4",
            shaValue:"ee68269a278bb549cef2fe2aeac21e6ce7c1d60d ",
            pathValue:"/tmp/virustotal/notavirus   ",
            findTime:"2021-02-02T15:25:30.123",
        }

    ]
})