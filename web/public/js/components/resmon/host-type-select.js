Vue.component("host-type-selector", {
	mounted: function () {
		let that = this

		// Tea.action("/assembly/options")
		// 	.post()
		// 	.success(function (resp) {
		// 		that.assemblys = resp.data.assemblys
		// 	})
	},
	props: ["v-assembly-type"],
	data: function () {

        let assemblyType = this.vAssemblyType
        // console.log(assemblyType)
        if (assemblyType == null) {
            assemblyType = 1
        }

        
        let assembly = [
            {id:1,name:"Linux 64位"},
            {id:2,name:"Windows 64位"},
            {id:3,name:"MacOS 64位"},
            {id:4,name:"Freebsd 64位"},
			{id:5,name:"Linux ARM 64位，ARM架构CPU专用"},
            {id:6,name:"Linux Mips 64，MIPS架构CPU专用"},
            {id:7,name:"Linux Mips 64 LE，MIPS64LE架构CPU专用"},
        ]
    
		return {
			assemblys: assembly,
			assemblyType: assemblyType,
		}
	},
	methods:{
 
    },
	watch:{
		assemblyType(newVal, oldVale) {
            if (newVal !== oldVale) {
                this.$emit("update:vAssemblyType", newVal)
            }
        },

	},
	template: `<div>
	<select name="assemblyType" v-model="assemblyType" style="width: 250px;height: 30px;padding: 0 0 0 5px;line-height: 30px;font-size: 13px;border: 1px solid #d7d7d7;">
		<option v-for="assembly in assemblys" :value="assembly.id">{{assembly.name}}</option>
	</select>
</div>`
})