Vue.component("assembly-selector", {
	mounted: function () {
		let that = this

		Tea.action("/assembly/options")
			.post()
			.success(function (resp) {
				that.assemblys = resp.data.assemblys
			})
	},
	props: ["v-cluster-id"],
	data: function () {
		return {
			assemblys: [],
			assemblyId: 0
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="assemblyId" v-model="assemblyId">
		<option value="0">[选择节点类型]</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})