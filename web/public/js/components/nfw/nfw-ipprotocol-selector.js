Vue.component("nfw-ipprotocol-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/idc/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.idcs = resp.data.idcs
		// 	})
	},
	props: ["v-ipprotocol","v-ipprotocols"],
	data: function () {
		let ipprotocol = this.vIpprotocol
		if (ipprotocol == null) {
			ipprotocol = "inet"
		}
		let ipprotocols = this.vIpprotocols
		return {
			ipprotocol: ipprotocol,
			ipprotocols: ipprotocols
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="ipprotocol" v-model="ipprotocol">
<!--		<option value="-1">[选择数据中心]</option>-->
		<option v-for="idc in ipprotocols" :value="idc.value">{{idc.name}}</option>
	</select>
</div>`
})