Tea.context(function () {
	this.dayFrom = ""
	this.dayTo = ""

	this.titleName = []

	this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")

		for(var index = 0;index<this.tableData.length;index++){
			this.titleName.push(this.tableData[index].name)
		}

		this.reloadPieTableChart()
		this.reloadLineTableChart()
    })

	this.onTimeChange = function () {
        let startTime = document.getElementById("day-from-picker").value
        let endTime = document.getElementById("day-to-picker").value
        if(this.dayFrom != startTime || this.dayTo != endTime) {
            this.dayFrom = startTime
            this.dayTo = endTime
        }
    }

	this.onDownloadReport = function(){

	}

	this.onOpenDetail = function(url){
		window.location = url
	}

	this.reloadPieTableChart = function () {
        let chartBox = document.getElementById("pie-chart-box")
        let chart = echarts.init(chartBox)
        let option = {
            // title:{
            //     text: '当前入侵威胁分布情况',
            //     x:'left',
            //     y: 'top',
            //     textStyle: { 
            //         fontSize: 16,
            //         color: '#333',
            //         fontWeight:"normal"
            //     },
            // },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 15,
                bottom: 30,
            },
            legend: {
                orient : 'horizontal',
                x : 'center',
                y : 'bottom',
                data:this.titleName,
            },
			tooltip : {
                trigger: 'item',
                formatter: "{a} <br/>{b} : {c} ({d}%)"
            },
            color:['#2698fb','#26c4c3','#26c46f','#ffce51','#f95c74','#8f59dd','#474fc5','#3d4c7d'],
			series: [
                {
                    name:"占比情况",
                    type: 'pie',
                    radius: "55%",
                    center: ['50%', '50%'],
                    data: this.tableData,
                    itemStyle:{ 
                        normal:{ 
                            label:{ 
                                show: true, 
                                formatter: '{b} : {c} ({d}%)' 
                            }, 
                            labelLine :{show:true} 
                        } 
                    } 
                }
            ],
			animation: false
		}
        chart.setOption(option)
        chart.resize()
    }

	this.reloadLineTableChart = function () {
		let chartBox = document.getElementById("line-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
			 title:{
                text: '当日安全事件趋势情况',
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
                top: 30,   // 等价于 y: '16%'
                left: 15, 
                right: 15,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                // boundaryGap: false,
				data: this.lineTitle
			},
			yAxis: {
                // name: 'GB',
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
                    name: this.attCountData.ddosAtt.name,
                    type: "line",
                    data: this.attCountData.ddosAtt.value,
                    itemStyle: {
						color: "#2698fb"
					},
                    lineStyle: {
                        color: "#2698fb"
                    }
                },
				{
                    name: this.attCountData.webAtt.name,
                    type: "line",
                    data: this.attCountData.webAtt.value,
                    itemStyle: {
						color: "#26c4c3"
					},
                    lineStyle: {
                        color: "#26c4c3"
                    }
                },
				{
                    name: this.attCountData.netAtt.name,
                    type: "line",
                    data: this.attCountData.netAtt.value,
                    itemStyle: {
                        color: "#26c46f"
                    },
                    lineStyle: {
                        color: "#26c46f"
                    }
                },
				{
                    name: this.attCountData.hostAtt.name,
                    type: "line",
                    data: this.attCountData.hostAtt.value,
                    itemStyle: {
                        color: "#ffce51"
                    },
                    lineStyle: {
                        color: "#ffce51"
                    }
                },
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

	
	this.lineTitle = ["2点","4点","6点","8点","10点","12点","14点","16点"]
	this.attCountData = {
		ddosAtt:{
			name:"DDos攻击",
			value:[20,25,36,11,25,30,12,40],
		},
		webAtt:{
			name:"WEB攻击",
			value:[30,45,46,41,35,40,32,30],
		},
		netAtt:{
			name:"网络入侵",
			value:[10,15,26,31,45,10,22,20],
		},
		hostAtt:{
			name:"主机入侵",
			value:[40,35,16,21,15,20,42,10],
		}
	}
		
	
	
	this.tableData = [
		{id:1,name:"DDos攻击",value:1222,url:""},
		{id:2,name:"WEB攻击",value:122222,url:""},
		{id:3,name:"网络入侵",value:12225,url:""},
		{id:4,name:"主机入侵",value:125555,url:""},
	]
})
