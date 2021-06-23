Vue.component("nfw-interface-selector", {
    mounted: function () {
        let that = this

        Tea.action("/nfw/nat/detail/options")
            .params({
                NodeId: that.nodeId,
                id: that.id,
            })
            .get()
            .success(function (resp) {
                that.interfaces = resp.data.interface
            })
    },
    props: ["v-node-id", "v-interface-id", "v-id"],
    data: function () {
        let interface = this.vInterfaceId
        if (interface == null) {
            interface = "wan"
        }
        let nodeId = this.vNodeId
        if (nodeId == null) {
            nodeId = 0
        }
        let id = this.vId
        if (id == null) {
            id = ""
        }
        return {
            interfaces: [],
            interface: interface,
            nodeId: nodeId,
            id: id,
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="interface" :value="interface" v-model="interface"> 
<!--		<option value="-1">[选择接口]</option>-->
		<option v-for="idc in interfaces" :value="idc.value" v-if="idc.selected == true" selected>{{idc.name}}</option>
		<option :value="idc.value" v-else>{{idc.name}}</option>
	</select>
</div>`
})
