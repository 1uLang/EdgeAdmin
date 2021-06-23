Tea.context(function () {

    this.dayFrom = ""
    this.dayTo = ""
    this.keyword = ""

    this.$delay(function () {
        teaweb.datepicker("day-from-picker")
        teaweb.datepicker("day-to-picker")
        this.reloadLineTableChart()
        this.reloadBarTableChart()
    })
    
    this.reloadLineTableChart = function () {
		let chartBox = document.getElementById("line-chart-box")
		let chart = echarts.init(chartBox)
		let option = {
            legend: {
                // orient 设置布局方式，默认水平布局，可选值：'horizontal'（水平） ¦ 'vertical'（垂直）
                orient: 'horizontal',
                // x 设置水平安放位置，默认全图居中，可选值：'center' ¦ 'left' ¦ 'right' ¦ {number}（x坐标，单位px）
                x: 'center',
                // y 设置垂直安放位置，默认全图顶端，可选值：'top' ¦ 'bottom' ¦ 'center' ¦ {number}（y坐标，单位px）
                y: 'bottom',
                data: [this.pageData[0].itemData.name, this.pageData[1].itemData.name,this.pageData[2].itemData.name,this.pageData[3].itemData.name]
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
                splitNumber:5,  // 设置y轴刻度间隔个数
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
					name: this.tableData1.itemData[0].name,
					type: "line",
					data: this.tableData1.itemData[0].lineTableData,
					itemStyle: {
						color: "#a5d9fd"
					},
					lineStyle: {
						color: "#a5d9fd"
					}
					// areaStyle: {
					// 	color: "#a5d9fd"
					// }
				},
                {
					name: this.tableData1.itemData[1].name,
					type: "line",
					data: this.tableData1.itemData[1].lineTableData,
					itemStyle: {
						color: "#a8ebcc"
					},
					lineStyle: {
						color: "#a8ebcc"
					}
					// areaStyle: {
					// 	color: "#a8ebcc"
					// }
				},
                {
					name: this.tableData1.itemData[2].name,
					type: "line",
					data: this.tableData1.itemData[2].lineTableData,
					itemStyle: {
						color: "#dfdfdf"
					},
					lineStyle: {
						color: "#dfdfdf"
					}
					// areaStyle: {
					// 	color: "#a8ebcc"
					// }
				},
                {
					name: this.tableData1.itemData[3].name,
					type: "line",
					data: this.tableData1.itemData[3].lineTableData,
					itemStyle: {
						color: "#000"
					},
					lineStyle: {
						color: "#000"
					}
					// areaStyle: {
					// 	color: "#a8ebcc"
					// }
				},
                
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

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
				data: this.tableData2.lineValue,
                axisLabel: {
                    rotate: -30, // 旋转角度
                    interval: 0  //设置X轴数据间隔几个显示一个，为0表示都显示
                },
			},
			yAxis: {
                // name: 'GB',
                min:0, // 设置y轴刻度的最小值
                // max:8,  // 设置y轴刻度的最大值
                splitNumber:5,  // 设置y轴刻度间隔个数
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
					data: this.tableData2.itemvalue,
					barWidth:"20px",
                    color: "#a8ebcc"
				},
                
			],
			animation: false
		}
		chart.setOption(option)
		chart.resize()
	}

    this.onOpenVirusPage = function(index){
        //向服务器请求获取需要显示的数据
        window.location = "/hids/invade/virus"
    }


    this.tableData1 = {
        lineValue:["2020-06-01","2020-06-02","2020-06-03","2020-06-04","2020-06-05","2020-06-06","2020-06-07"],
        itemData:[
            {
                name:"病毒木马",
                lineTableData:[10,12,11,12,15,8,5]
            },
            {
                name:"网页后门",
                lineTableData:[1,2,3,4,5,6,7]
            },
            {
                name:"反弹shell",
                lineTableData:[7,6,5,4,3,2,1]
            },
            {
                name:"日志异常删除",
                lineTableData:[5,15,25,35,20,30,10]
            },
        ]
    }
    this.tableData2 = {
        lineValue:["病毒木马","网页后门","反弹shell","异常账号","日志异常删除","异常登录","异常进程","系统命令篡改","文件篡改","注册表篡改","入侵扫描","暴力破解","进程异常操作","非授权进程","溢出攻击","进程行为监控"],
        itemvalue:[20,12,15,16,18,20,30,40,50,12,1,19,16,30,15,12]
    }


    this.pageData =[
        {
            id:1,
            itemData:{
                name:"病毒木马",
                count:2,
            }
        },
        {
            id:2,
            itemData:{
                name:"网页后门",
                count:2,
            }
        },
        {
            id:3,
            itemData:{
                name:"反弹shell",
                count:2,
                
            }
        },
        {
            id:4,
            itemData:{
                name:"异常账号",
                count:2,
            }
        },
        {
            id:5,
            itemData:{
                name:"日志异常删除",
                count:2,
                
            }
        },
        {
            id:6,
            itemData:{
                name:"异常登录",
                count:2,
            }
        },
        {
            id:7,itemData:{
                name:"异常进程",
                count:2,
            }
        },
        {
            id:8,
            itemData:{
                name:"系统命令篡改",
                count:2,
            }
        },
        
    ]
});
  