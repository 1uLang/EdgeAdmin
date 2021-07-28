Tea.context(function () {

    this.severity = ''

    this.$delay(function () {
		let curSelectNode = localStorage.getItem("ddosSelectNodeId");
		if(curSelectNode){
			this.nodeId = curSelectNode
		}

        this.reloadHourChart()
        this.reloadDayChart()
        this.reloadMonthChart()
	
    })

    this.reloadHourChart = function () {
		let chartBox = document.getElementById("hour-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: ['输入流量', '输出流量']
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 10, 
                right: 60,
                bottom: 20,
                containLabel: true
            },
			xAxis: {
                name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                // boundaryGap: false,
				data: [20,22,24,2,4,6,8,10,12,14,16,18,20,22,24,2,4,6,8]
			},
			yAxis: {
                name: 'GB',
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
					name: "输入流量",
					type: "line",
					data: [1,2,5,10,2,4,6,8,12,14,16,8,5,7,20,2,4,6,8],
					itemStyle: {
						color: "#a5d9fd"
					},
					lineStyle: {
						color: "#a5d9fd"
					},
					areaStyle: {
						color: "#a5d9fd"
					}
				},
                {
					name: "输出流量",
					type: "line",
					data: [5,2,5,10,2,4,6,8,12,14,16,8,5,7,20,2,4,6,8],
					itemStyle: {
						color: "#a8ebcc"
					},
					lineStyle: {
						color: "#a8ebcc"
					},
					areaStyle: {
						color: "#a8ebcc"
					}
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

    this.reloadDayChart = function () {
		let chartBox = document.getElementById("day-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: ['输入流量', '输出流量']
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 10, 
                right: 60,
                bottom: 20,
                containLabel: true
            },
			xAxis: {
                name: 'Day',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                // boundaryGap: false,
				data: [23,25,27,29,31,1,3,5]
			},
			yAxis: {
                name: 'GB',
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
					name: "输入流量",
					type: "line",
					data: [1,2,5,10,2,4,6,8],
					itemStyle: {
						color: "#a5d9fd"
					},
					lineStyle: {
						color: "#a5d9fd"
					},
					areaStyle: {
						color: "#a5d9fd"
					}
				},
                {
					name: "输出流量",
					type: "line",
					data: [5,1,3,5,1,3,2,5],
					itemStyle: {
						color: "#a8ebcc"
					},
					lineStyle: {
						color: "#a8ebcc"
					},
					areaStyle: {
						color: "#a8ebcc"
					}
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

    this.reloadMonthChart = function () {
		let chartBox = document.getElementById("month-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: ['输入流量', '输出流量']
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 10, 
                right: 60,
                bottom: 20,
                containLabel: true
            },
			xAxis: {
                name: 'Month',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                // boundaryGap: false,
				data: [7,9,11,1,3,5]
			},
			yAxis: {
                name: 'GB',
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
					name: "输入流量",
					type: "line",
					data: [1,2,5,10,2,4],
					itemStyle: {
						color: "#a5d9fd"
					},
					lineStyle: {
						color: "#a5d9fd"
					},
					areaStyle: {
						color: "#a5d9fd"
					}
				},
                {
					name: "输出流量",
					type: "line",
					data: [5,1,4,6,1,3],
					itemStyle: {
						color: "#a8ebcc"
					},
					lineStyle: {
						color: "#a8ebcc"
					},
					areaStyle: {
						color: "#a8ebcc"
					}
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

	this.getNodeId = function () {
        let node = this.nodeId
        return node
    }

    this.showHost = function () {
		let node = this.getNodeId()
        localStorage.setItem("ddosSelectNodeId", node);
		window.location.href = '/ddos/host?nodeId=' + node
    }

	this.tableData = [
		{id:1,direction:"外网",flow:"235.04 / 215.79 Mbps",msg:"54712 / 39151 pps",anyKey:"0.01 / 0.01 Mbps"},
		{id:2,direction:"内网",flow:"295.61 / 299.79 Mbps",msg:"158644 / 164913 pps ",anyKey:"0.05 / 0.05 Mbps"},
		{id:3,direction:"内网",flow:"295.61 / 299.79 Mbps",msg:"158644 / 164913 pps ",anyKey:"0.05 / 0.05 Mbps"},
	]
})