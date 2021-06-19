Tea.context(function () {
  this.address = ''
  this.severity = ''
  this.detailInfo = null
  this.bShowDetail = false
  this.onOpenDetail = function (vul) {
      if(vul){
        detailInfo = this.$get(".details")
        .params({
          VulId: vul.vuln_id
        })
        document.getElementById("rightDiv").style.display = "flex";
      }else{
        document.getElementById("rightDiv").style.display = "none";
      }
  };

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.indexOf("."));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };
    this.onChangeSeverityFormat = function (severity) {
        var resultSeverity = severity;

        switch (severity) {
            case 3:
                return '高危'
            case 2:
                return '中危'
            case 1:
                return '低危'
            default:
                return '信息'
        }
        return resultSeverity;
    };

    this.testData = [
        {id:1,
         title1:"Elasticsearch 服务可访问",content1:"URL:           http://182.150.0.104/",
         title2:"漏洞描述",content2:"Elasticsearch ...<br /><br />Acunetix ...<br /><br />发现者 Elasticsearch 服务可访问",
         title3:"此漏洞的影响",content3:"攻击者可以在 elasticsearch 服务上查询、更新和创建新索引。",
         title4:"如何修复此漏洞",content4:"禁止外部访问 elasticsearch 服务。",
        },
        {id:2,
            title1:"Elasticsearch 服务不可访问",content1:"URL:           http://www.baidu.com/",
            title2:"漏洞描述",content2:"Elasticsearch ...<br /><br />Acunetix ...<br /><br />发现者 Elasticsearch 服务可访问",
            title3:"此漏洞的影响",content3:"攻击者可以在 elasticsearch 服务上查询、更新和创建新索引。",
            title4:"如何修复此漏洞",content4:"禁止外部访问 elasticsearch 服务。",
        },
    ]

    this.getDataiInfo = function (id) { 
        this.detailInfo = null
        for(var i = 0; i < this.testData.length; i++) {
            if(this.testData[i].id == id){
                this.detailInfo = this.testData[i]
            }
        }
        if(this.detailInfo){
            this.bShowDetail = true
            // document.getElementById("rightDiv").style.display = "flex";
        }else{
            this.bShowDetail = false
            // document.getElementById("rightDiv").style.display = "none";
        }
    }

});
