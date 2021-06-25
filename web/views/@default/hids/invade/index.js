Tea.context(function () {

    // this.dayFrom = ""
    // this.dayTo = ""
    // this.keyword = ""
	//
    // this.$delay(function () {
    //     teaweb.datepicker("day-from-picker")
    //     teaweb.datepicker("day-to-picker")
    //     this.reloadLineTableChart()
    //     this.reloadBarTableChart()
    // })

	this.onOpenVirusPage = function(url){

		//向服务器请求获取需要显示的数据
		window.location = "/hids/invade/"+url
	}
    // this.reloadLineTableChart = function () {
	// 	let chartBox = document.getElementById("line-chart-box")
	// 	let chart = echarts.init(chartBox)
	// 	let option = {
    //         legend: {
    //             // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
    //             orient: 'horizontal',
    //             // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
    //             x: 'center',
    //             // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
    //             y: 'bottom',
    //             data: [this.pageData[0].itemData.name, this.pageData[1].itemData.name,this.pageData[2].itemData.name,this.pageData[3].itemData.name]
    //         },
    //         //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
    //         // top: '16%',   // 等价于 y: '16%'
    //         grid: {
    //             top: 30,   // 等价于 y: '16%'
    //             left: 40,
    //             right: 60,
    //             bottom: 30,
    //             containLabel: true
    //         },
	// 		xAxis: {
    //             // name: 'Hour',
    //             // boundaryGap值为false的时候，折线第一个点在y轴上
    //             boundaryGap: false,
	// 			data: this.tableData1.lineValue
	// 		},
	// 		yAxis: {
    //             // name: 'GB',
    //             min:0, // 设置y轴刻度的最小值
    //             // max:8,  // 设置y轴刻度的最大值
    //             splitNumber:5,  // 设置y轴刻度间隔个数
    //             // axisLine: {
    //             //     lineStyle: {
    //             //         // 设置y轴颜色
    //             //         color: '#fff'
    //             //     }
    //             // },
    //         },
	// 		tooltip: {
	// 			trigger: "axis",
	// 		},
	// 		series: [
	// 			{
	// 				name: this.tableData1.itemData[0].name,
	// 				type: "line",
	// 				data: this.tableData1.itemData[0].lineTableData,
	// 				itemStyle: {
	// 					color: "#a5d9fd"
	// 				},
	// 				lineStyle: {
	// 					color: "#a5d9fd"
	// 				}
	// 				// areaStyle: {
	// 				// 	color: "#a5d9fd"
	// 				// }
	// 			},
    //             {
	// 				name: this.tableData1.itemData[1].name,
	// 				type: "line",
	// 				data: this.tableData1.itemData[1].lineTableData,
	// 				itemStyle: {
	// 					color: "#a8ebcc"
	// 				},
	// 				lineStyle: {
	// 					color: "#a8ebcc"
	// 				}
	// 				// areaStyle: {
	// 				// 	color: "#a8ebcc"
	// 				// }
	// 			},
    //             {
	// 				name: this.tableData1.itemData[2].name,
	// 				type: "line",
	// 				data: this.tableData1.itemData[2].lineTableData,
	// 				itemStyle: {
	// 					color: "#dfdfdf"
	// 				},
	// 				lineStyle: {
	// 					color: "#dfdfdf"
	// 				}
	// 				// areaStyle: {
	// 				// 	color: "#a8ebcc"
	// 				// }
	// 			},
    //             {
	// 				name: this.tableData1.itemData[3].name,
	// 				type: "line",
	// 				data: this.tableData1.itemData[3].lineTableData,
	// 				itemStyle: {
	// 					color: "#000"
	// 				},
	// 				lineStyle: {
	// 					color: "#000"
	// 				}
	// 				// areaStyle: {
	// 				// 	color: "#a8ebcc"
	// 				// }
	// 			},
    //
    //
	// 		],
	// 		animation: false
	// 	}
	// 	chart.setOption(option)
	// 	chart.resize()
	// }
	//
    // this.reloadBarTableChart = function () {
	// 	let chartBox = document.getElementById("bar-chart-box")
	// 	let chart = echarts.init(chartBox)
	// 	let option = {
    //         //  图表距边框的距离,可选值：'百分比'¦ {number}（单位px）
    //         // top: '16%',   // 等价于 y: '16%'
    //         grid: {
    //             top: 30,   // 等价于 y: '16%'
    //             left: 15,
    //             right: 60,
    //             bottom: 30,
    //             containLabel: true
    //         },
	// 		xAxis: {
    //             // name: 'Hour',
    //             // boundaryGap值为false的时候，折线第一个点在y轴上
	// 			data: this.tableData2.lineValue,
    //             axisLabel: {
    //                 rotate: -30, // 旋转角度
    //                 interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
    //             },
	// 		},
	// 		yAxis: {
    //             // name: 'GB',
    //             min:0, // 设置y轴刻度的最小值
    //             // max:8,  // 设置y轴刻度的最大值
    //             splitNumber:5,  // 设置y轴刻度间隔个数
    //             // axisLine: {
    //             //     lineStyle: {
    //             //         // 设置y轴颜色
    //             //         color: '#fff'
    //             //     }
    //             // },
    //         },
	// 		tooltip: {
	// 			trigger: "axis",
	// 		},
	// 		series: [
	// 			{
	// 				type: "bar",
	// 				data: this.tableData2.itemvalue,
	// 				barWidth:"20px",
    //                 color: "#a8ebcc"
	// 			},
    //
	// 		],
	// 		animation: false
	// 	}
	// 	chart.setOption(option)
	// 	chart.resize()
	// }

});
  