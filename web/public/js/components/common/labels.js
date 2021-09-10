// 启用状态标签
Vue.component("label-on", {
	props: ["v-is-on"],
	template: '<div><span v-if="vIsOn" class="ui label tiny green basic">已启用</span><span v-if="!vIsOn" class="ui label tiny red basic">已停用</span></div>'
})

// 文字代码标签
Vue.component("code-label", {
	template: `<span class="ui label basic tiny" style="padding: 3px;margin-left:2px;margin-right:2px"><slot></slot></span>`
})

// tiny标签
Vue.component("tiny-label", {
	template: `<span class="ui label tiny" style="margin-bottom: 0.5em"><slot></slot></span>`
})

Vue.component("tiny-basic-label", {
	template: `<span class="ui label tiny basic" style="margin-bottom: 0.5em"><slot></slot></span>`
})

// 更小的标签
Vue.component("micro-basic-label", {
	template: `<span class="ui label tiny basic" style="margin-bottom: 0.5em; font-size: 0.7em; padding: 4px"><slot></slot></span>`
})

Vue.component("label-on-conn", {
	props: ["v-is-conn"],
	template: '<div><span v-if="vIsConn" class="ui label tiny green basic">运行中</span><span v-if="!vIsConn" class="ui label tiny  basic" style="color: #bebebe">不可用</span></div>'
})
Vue.component("label-on-active", {
	props: ["v-is-active"],
	template: '<div><span v-if="vIsActive" class="ui label tiny green basic">已连接</span><span v-if="!vIsActive" class="ui label tiny  basic" style="color: #bebebe">未连接</span></div>'
})