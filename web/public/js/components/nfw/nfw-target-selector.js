Vue.component("nfw-target-selector", {
    mounted: function () {
        let that = this

        // Tea.action("/nfw/nat/detail/options")
        //     .params({
        //         NodeId: that.nodeId,
        //         id: that.id,
        //     })
        //     .get()
        //     .success(function (resp) {
        //         that.dsts = resp.data.dst
        //         for(var i = 0; i < that.dsts.length; i++){
        //             if(that.dsts[i].selected == true && that.dsts[i].data_other == true){
        //                 that.hide = false
        //             }
        //
        //             if(that.dsts[i].selected == true){
        //                 that.dst = that.dsts[i].value
        //             }
        //         }
        //         that.dstmask = "32"
        //         let dstmask = resp.data.dstmask
        //         for(var i = 0; i < dstmask.length; i++){
        //             if(dstmask[i].selected == true ){
        //                 that.dstmask = dstmask[i].value
        //             }
        //         }
        //     })


    },
    methods: {
        selectSrc(event) {
            let value = event.target.value
            let vmap = ["bogons", "bogonsv6", "virusprot", "sshlockout", "any", "(self)", "lan", "lanip",
                "lo0", "wan", "wanip"]
            if (vmap.indexOf(value) == -1) {
                this.hide = false
            } else {
                this.hide = true
            }
            this.dstinput = value
        },
        checkHide(value){
            let vmap = ["bogons", "bogonsv6", "virusprot", "sshlockout", "any", "(self)", "lan", "lanip",
                "lo0", "wan", "wanip"]
            if (vmap.indexOf(value) == -1) {
                return false
            } else {
                return true
            }
        }
    },
    props: ["v-dst", "v-node-id", "v-id", "v-masks", "v-dstmask", "v-dstinput", "v-dsts"],
    data: function () {
        let dst = this.vDst
        let nodeId = this.vNodeId
        if (nodeId == null) {
            nodeId = 0
        }
        let id = this.vId
        if (id == null) {
            id = ""
        }
        let masks = this.vMasks
        let dstmask = this.vDstmask
        let dsts = this.vDsts
        let dstinput = this.vDstinput

        let hide = true
        if(!this.checkHide(dstinput)){
            hide = false
        }
        return {
            dsts: dsts,
            dst: dst,
            nodeId: nodeId,
            id: id,
            masks: masks,
            dstmask: dstmask,
            hide: hide,
            dstinput: dstinput,
        }
    }, watch: {
        dstmask(newVal, oldVale) {
            // console.log("dstmask-new:", newVal);
            // console.log("dstmask-old:", oldVale);
            if (newVal !== oldVale) {
                this.$emit("update:vDstmask", newVal)
            }
        },
        dstinput(newVal, oldVale) {
            // console.log("dstinput-new:", newVal);
            // console.log("dstinput-old:", oldVale);
            if (newVal !== oldVale) {
                this.$emit("update:vDstinput", newVal)
            }
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="dst"  @change="selectSrc($event)"  v-model="dst" style="min-width: 180px;">
		<option v-for="(idc,k) in dsts" :value="idc.value" :v-data="idc.data_other" :key="k" >
		    {{idc.name}}
		</option>
	</select>
	<div style="display: flex;justify-content: start;flex-direction: row;margin-top: 5px;">
        <input type="text" name="dst" ref="dst" maxlength="50"  v-model="dstinput" :class="{'hide':hide==true}" style="width: 180px;"/>
        <select class="ui dropdown auto-width" ref="dstmask" name="dstmask" v-model="dstmask" :class="{'hide':hide==true}" style="min-width: 80px;margin-left: 10px;">
            <option v-for="item in masks" :value="item.id" >{{item.value}}</option>
        </select>
    </div>
</div>`
})