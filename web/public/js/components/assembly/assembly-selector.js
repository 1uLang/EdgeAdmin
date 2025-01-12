Vue.component("assembly-selector", {
	mounted: function () {
		let that = this

		Tea.action("/assembly/options")
			.post()
			.success(function (resp) {
				that.assemblys = resp.data.assemblys
			})
	},
	props: ["v-assembly-id"],
	data: function () {
		let assemblyType = this.vAssemblyId
		if (assemblyType == null) {
			assemblyType = -1
		}
		return {
			assemblys: [],
			assemblyType: assemblyType,
		}
	},
	watch:{
		assemblyType:function (){
			if (Tea.Vue != null) {
				Tea.Vue.showAPIAUTHVisible = this.assemblyType
			}
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="assemblyType" v-model="assemblyType">
		<option value="-1">[选择节点类型]</option>
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})