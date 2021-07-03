Tea.context(function () {
    this.seriousPer = "0%"
    this.heightPer = "0%"
    this.middlePer = "0%"
    this.lowPer = "0%"

    this.$delay(function () {
        this.reloadBarTableChart()
        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })

    this.onOpenDetail = function(type){
        switch(type){
            case 1:
                window.location = "/hids/risk/systemRisk"
                break
            case 2:
                window.location = "/hids/risk/weak"
                break
            case 3:
                window.location = "/hids/risk/dangerAccount"
                break
            case 4:
                window.location = "/hids/risk/configDefect"
                break
            default:
                break
        }
       
    }

    this.data = {
        seriousCount:10,
        heightCount:20,
        middleCount:10,
        lowCount:1,
        errHostCount:10,
        errWinCount:5,
        errLinuxCount:12,
        value1:{//弱口令
            allCount:206,
            errCount:15
        },
        value2:{//风险账号
            allCount:206,
            errCount:15
        },
        value3:{//缺陷配置
            allCount:206,
            errCount:15
        },
    }

    this.reloadBarTableChart = function () {
        let chartBox = document.getElementById("bar-chart-box")
        let chart = echarts.init(chartBox)
        let option = {
            title:{
                text: '当前漏洞风险分布情况',
                x:'left',
                y: 'top',
                textStyle: { 
                    fontSize: 16,
                    color: '#333',
                    fontWeight:"normal"
                },
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 60,   // 等价于 y: '16%'
                left: 15,
                right: 30,
                containLabel: true
            },
            xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                data: this.tableData1.itemName,
                axisLabel: {
                    rotate: 0, // 旋转角度
                    interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
                },
            },
            yAxis: {
                // name: 'GB',
                min: 0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber: 5,  // 设置y轴刻度间隔个数
                // axisLine: {
                //     lineStyle: {
                //         // 设置y轴颜色
                //         color: '#fff'
                //     }
                // },
            },
            tooltip: {
                trigger: "item",
            },
            series: [
                {
                    type: "bar",
                    data: this.tableData1.itemValue,
                    barWidth: "70px",
                    color: "#2698fb"
                },

            ],
            animation: false
        }
        chart.setOption(option)
        chart.resize()
    }

    this.tableData1 = {
        itemName:["系统漏洞", "弱口令", "风险账号", "配置缺陷"],
        itemValue:[
            {value:20,name:'系统漏洞'}, 
            {value:15,name:'弱口令'}, 
            {value:10,name:'风险账号'},
            {value:18,name:'配置缺陷'},
        ]
    }

});
  