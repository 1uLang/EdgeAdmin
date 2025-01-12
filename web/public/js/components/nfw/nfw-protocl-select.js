Vue.component("nfw-protocol-selector", {
    mounted: function () {
        let that = this

        // Tea.action("/assembly/idc/options")
        // 	.post()
        // 	.success(function (resp) {
        // 		that.idcs = resp.data.idcs
        // 	})
    }, watch: {
        protocol(newVal, oldVale) {
            // console.log("dstmask-new:", newVal);
            // console.log("dstmask-old:", oldVale);
            if (newVal !== oldVale) {
                this.$emit("update:vProtocol", newVal)
                console.log(1111);
            }
        },
    },
    props: ["v-protocol","v-protocols"],
    data: function () {
        let protocol = this.vProtocol
        if (protocol == null) {
            protocol = "inet"
        }
        let protocols = this.vProtocols
        return {
            protocol: protocol,
            protocols: protocols
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="protocol" v-model="protocol" style="min-width: 180px;">
<!--		<option value="-1">[选择数据中心]</option>-->
		<option v-for="idc in protocols" :value="idc.value">{{idc.name}}</option>
	</select>
</div>`
})