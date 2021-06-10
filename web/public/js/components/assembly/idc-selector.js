Vue.component("idc-selector", {
	mounted: function () {
		let that = this

		Tea.action("/assembly/idc/options")
			.post()
			.success(function (resp) {
				that.idcs = resp.data.idcs
			})
	},
	props: ["v-cluster-id"],
	data: function () {
		return {
			idcs: [],
			idcId: 0
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="idcId" v-model="idcId">
		<option value="0">[选择数据中心]</option>
		<option v-for="idc in idcs" :value="idc.id">{{idc.name}}</option>
	</select>
</div>`
})