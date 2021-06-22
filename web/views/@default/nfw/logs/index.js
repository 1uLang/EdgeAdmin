Tea.context(function () {
    this.severity=1
    this.keyword = ""

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };

    this.onDelete = function (param) {  

    }


    this.hostData = [
        {id:1,hostAddress:"成都-ddos-192.168.1.1",},
        {id:2,hostAddress:"成都-ddos-192.168.1.2",},
        {id:3,hostAddress:"成都-ddos-192.168.1.3",},
        {id:4,hostAddress:"成都-ddos-192.168.1.4",},
    ]

    this.tableData = [
        {id:1,value1:"2021-02-06T15:12:26.000",value2:"suricata[96302]",value3:"[100221] <Notice> -- Stats for 'vtnet0': pkts: 1695, drop: 0 (0.00%), invalid chksum: 0"},
        {id:2,value1:"2021-02-06T15:12:26.000",value2:"suricata[96302]",value3:"[100221] <Notice> -- Stats for 'vtnet0': pkts: 1695, drop: 0 (0.00%), invalid chksum: 0"},
        {id:3,value1:"2021-02-06T15:12:26.000",value2:"suricata[96302]",value3:"[100221] <Notice> -- Stats for 'vtnet0': pkts: 1695, drop: 0 (0.00%), invalid chksum: 0"},
    ]
})