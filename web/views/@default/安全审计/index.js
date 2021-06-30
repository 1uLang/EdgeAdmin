Tea.context(function () {

    this.$delay(function () {
        this.reloadLineTableChartOne()
        this.reloadLineTableChartTwo()
        this.reloadLineTableChartThree()
    })

    this.reloadLineTableChartOne = function () {
		let chartBox = document.getElementById("lineOne-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: [this.tableData1.sqlTitle,this.tableData1.hostTitle,this.tableData1.exeTitle]
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 40,
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                boundaryGap: false,
				data: this.tableData1.lineValue
			},
			yAxis: {
                // name: 'GB',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:7,  // 设置y轴刻度间隔个数
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
					name: this.tableData1.sqlTitle,
					type: "line",
					data: this.tableData1.sqlValue,
					itemStyle: {
						color: "#0085fa"
					},
					lineStyle: {
						color: "#0085fa"
					}
				},
                {
					name: this.tableData1.hostTitle,
					type: "line",
					data: this.tableData1.hostValue,
					itemStyle: {
						color: "#00ba57"
					},
					lineStyle: {
						color: "#00ba57"
					}
				},
                {
					name: this.tableData1.exeTitle,
					type: "line",
					data: this.tableData1.exeValue,
					itemStyle: {
						color: "#ffc532"
					},
					lineStyle: {
						color: "#ffc532"
					}
				},
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}
    this.reloadLineTableChartTwo= function () {
		let chartBox = document.getElementById("lineTwo-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: [this.tableData2.title1,this.tableData2.title2,this.tableData2.title3,this.tableData2.title4,this.this.tableData2.title5]
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 40,
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                boundaryGap: false,
				data: this.tableData2.lineValue
			},
			yAxis: {
                // name: 'GB',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:7,  // 设置y轴刻度间隔个数
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
					name: this.tableData2.title1,
					type: "line",
					data: this.tableData2.value1,
					itemStyle: {
						color: "#0085fa"
					},
					lineStyle: {
						color: "#0085fa"
					}
				},
                {
					name: this.tableData2.title2,
					type: "line",
					data: this.tableData2.value2,
					itemStyle: {
						color: "#00ba57"
					},
					lineStyle: {
						color: "#00ba57"
					}
				},
                {
					name: this.tableData2.title3,
					type: "line",
					data: this.tableData2.value3,
					itemStyle: {
						color: "#ffc532"
					},
					lineStyle: {
						color: "#ffc532"
					}
				},
                {
					name: this.tableData2.title4,
					type: "line",
					data: this.tableData2.value4,
					itemStyle: {
						color: "#1b2d35"
					},
					lineStyle: {
						color: "#1b2d35"
					}
				},
                {
					name: this.tableData2.title5,
					type: "line",
					data: this.tableData2.value5,
					itemStyle: {
						color: "#7c3cd6"
					},
					lineStyle: {
						color: "#7c3cd6"
					}
				},
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

    this.reloadLineTableChartThree= function () {
		let chartBox = document.getElementById("lineThree-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: [this.tableData3.title1,this.tableData3.title2,this.tableData3.title3,this.tableData3.title4,this.this.tableData3.title5]
            },
            //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
            // top: '16%',   // 等价于 y: '16%'
            grid: {
                top: 30,   // 等价于 y: '16%'
                left: 40,
                right: 60,
                bottom: 30,
                containLabel: true
            },
			xAxis: {
                // name: 'Hour',
                // boundaryGap值为false的时候，折线第一个点在y轴上
                boundaryGap: false,
				data: this.tableData3.lineValue
			},
			yAxis: {
                // name: 'GB',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:7,  // 设置y轴刻度间隔个数
            },
			tooltip: {
				trigger: "axis",
			},
			series: [
				{
					name: this.tableData3.title1,
					type: "line",
					data: this.tableData3.value1,
					itemStyle: {
						color: "#0085fa"
					},
					lineStyle: {
						color: "#0085fa"
					}
				},
                {
					name: this.tableData3.title2,
					type: "line",
					data: this.tableData3.value2,
					itemStyle: {
						color: "#00ba57"
					},
					lineStyle: {
						color: "#00ba57"
					}
				},
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}
    this.tableData1={
        lineValue:["6-11","6-12","6-13","6-14","6-15","6-16","6-17"],
        sqlTitle:"数据库日志",
        hostTitle:"主机日志",
        exeTitle:"应用日志",
        sqlValue:[10000,8562,4568,6922,15800,16879,25465],
        hostValue:[56241,28562,44568,66922,55800,86879,75465],
        exeValue:[96241,88562,74568,66922,55800,66879,75465]
    }
    this.tableData2={
        lineValue:["6-11","6-12","6-13","6-14","6-15","6-16","6-17"],
        title1:"DML",
        title2:"DDL",
        title3:"TCL",
        title4:"DCL",
        title5:"其他",
        value1:[10000,8562,4568,6922,15800,16879,25465],
        value2:[56241,28562,44568,66922,55800,86879,75465],
        value3:[96241,88562,74568,66922,55800,66879,75465],
        value4:[12563,18648,25695,26125,65140,45621,12593],
        value5:[56971,99543,98563,36922,35412,95315,36521]
    }
    this.tableData3={
        lineValue:["6-11","6-12","6-13","6-14","6-15","6-16","6-17"],
        title1:"DML",
        title2:"DDL",
        value1:[10000,15658,19635,21301,8543,15684,36521],
        value2:[10501,28562,24568,36521,55800,15486,25487],
    }

    this.pageData = {
        titleInfo:[
            {id:1,name:"数据源",count:7,unit:"个"},
            {id:2,name:"日志",count:7,unit:"条"},
            {id:3,name:"告警",count:7,unit:"条"},
            {id:4,name:"存储容量",count:7,subCount:10,unit:"G"}
        ]
    }
})