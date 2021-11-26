Vue.component("assembly-http-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-argeement"],
	data: function () {
		let argeement = this.vArgeement
		if (argeement == null) {
			argeement = 0
		}
		argeements = [
			{"id":0,"name":"http"},
			{"id":1,"name":"https"},
		]
		return {
			assemblys: argeements,
			argeement: argeement,
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="argeement" v-model="argeement">
<!--		<option value="-1">[网络协议]</option>-->
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})