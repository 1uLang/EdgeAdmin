// ��ʾ��������˵��
Vue.component("traffic-limit-view", {
	props: ["v-traffic-limit"],
	data: function () {
		return {
			config: this.vTrafficLimit
		}
	},
	template: `<div>
	<div v-if="config.isOn">
		<span v-if="config.dailySize != null && config.dailySize.count > 0">���������ƣ�{{config.dailySize.count}}{{config.dailySize.unit.toUpperCase()}}<br/></span>
		<span v-if="config.monthlySize != null && config.monthlySize.count > 0">���������ƣ�{{config.monthlySize.count}}{{config.monthlySize.unit.toUpperCase()}}<br/></span>
	</div>
	<span v-else class="disabled">û�����ơ�</span>
</div>`
})