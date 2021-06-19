Vue.component("user-selector", {
	mounted: function () {
		let that = this

		Tea.action("/assembly/idc/options")
			.post()
			.success(function (resp) {
				that.idcs = resp.data.idcs
			})
	},
	props: ["v-idc-id"],
	data: function () {
		let idcId = this.vIdcId
		if (idcId == null) {
			idcId = -1
		}
		return {
			idcs: [],
			idcId: idcId
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="idcId" v-model="idcId">
		<option value="-1">[选择数据中心]</option>
		<option v-for="idc in idcs" :value="idc.id">{{idc.name}}</option>
	</select>
</div>`
})