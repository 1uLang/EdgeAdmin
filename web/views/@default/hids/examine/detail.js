Tea.context(function () {
    

    this.$delay(function () {
        this.reloadBarTableChart()
        this.reloadCircularTableChart()
    })

    this.onGoBack = function () {
        window.location = "/hids/examine"
    }

    this.getHealthName = function (num) {
        if(num<60){
            return "不健康"
        }else if(num<90){
            return "亚健康"
        }else {
            return "健康"
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

    this.reloadBarTableChart = function () {
        let chartBox = document.getElementById("bar-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
				data: this.tableData.titleValue,
                axisLabel: {
                    // rotate: -30, // 旋转角度
                    interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
                },
			},
			yAxis: {
                name: '体检概况',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:4,  // 设置y轴刻度间隔个数
                // axisLine: {
                //     lineStyle: {
                //         // 设置y轴颜色
                //         color: '#fff'
                //     }
                // },
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
					type: "bar",
					data: this.tableData.itemValue,
					barWidth:"120px",
                    color: "#2698fb"
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
    }
    this.reloadCircularTableChart = function () {
        let chartBox = document.getElementById("circular-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
				data: this.tableData.titleValue,
                axisLabel: {
                    // rotate: -30, // 旋转角度
                    interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
                },
			},
			yAxis: {
                name: '体检概况',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:4,  // 设置y轴刻度间隔个数
                // axisLine: {
                //     lineStyle: {
                //         // 设置y轴颜色
                //         color: '#fff'
                //     }
                // },
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
					type: "bar",
					data: this.tableData.itemValue,
					barWidth:"120px",
                    color: "#2698fb"
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
    }

    this.pageData = {
        ipAddr:"45.195.61.132 （192.168.1.47;172.17.0.1;172.18.0.1）",
        checkNum:70,
        checkTime:"2021-06-05T12:15:25.000",
    }

    this.tableData = {
        titleValue:["系统漏洞","配置缺陷"],
        itemValue:[20,2]
    }
})