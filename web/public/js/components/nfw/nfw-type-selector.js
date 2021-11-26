Vue.component("nfw-type-selector", {
    mounted: function () {
        let that = this

        // Tea.action("/nfw/nat/detail/options")
        //     .params({
        //         NodeId: that.nodeId,
        //         id: that.id,
        //     })
        //     .get()
        //     .success(function (resp) {
        //         that.types = resp.data.type
        //     })
    }, watch: {
        type(newVal, oldVale) {
            // console.log("dstmask-new:", newVal);
            // console.log("dstmask-old:", oldVale);
            if (newVal !== oldVale) {
                this.$emit("update:vType", newVal)
            }
        },
    },
    props: ["v-type", "v-node-id", "v-id", "v-types"],
    data: function () {
        let type = this.vType
        if (type == null) {
            type = -1
        }
        let nodeId = this.vNodeId
        if (nodeId == null) {
            nodeId = 0
        }
        let id = this.vId
        if (id == null) {
            id = ""
        }
        let types = this.vTypes
        return {
            types: types,
            type: type,
            nodeId: nodeId,
            id: id,
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="type" v-model="type" style="min-width: 180px;">
<!--		<option value="-1">[选择数据中心]</option>-->
		<option v-for="idc in types" :value="idc.value"  >{{idc.name}}</option>
<!--		<option :value="idc.value" v-else>{{idc.name}}</option>-->

	</select>
</div>`
})