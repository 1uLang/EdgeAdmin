Vue.component("nfw-srcbeginport-selector", {
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
            // let vmap = ["bogons", "bogonsv6", "virusprot", "sshlockout", "any", "(self)", "lan", "lanip",
            //     "lo0", "wan", "wanip"]
            // if (vmap.indexOf(value) == -1) {
            //     this.hide = false
            // } else {
            //     this.hide = true
            // }
            this.srcbeginportinput = value

            let srcportlist =  this.srcbeginportList
            for (var i = 0; i < srcportlist.length; i++) { //目标下拉
                if (!srcportlist[i].data_other ) {
                    if(value == srcportlist[i].value){
                        this.hide = false
                    }

                }
            }
            this.hide = true
        },
        checkHide(value){
            this.srcbeginportinput = value

            let srcportlist =  this.srcbeginportList
            for (var i = 0; i < srcportlist.length; i++) { //目标下拉
                if (!srcportlist[i].data_other ) {
                    if(value == srcportlist[i].value){
                        this.hide = true
                    }
                }
            }
            this.hide = true
        }
    },
    props: ["v-protocol", "v-srcbeginport", "v-srcbeginport-list","srcbeginportinput"],
    data: function () {
        let protocol = this.vProtocol
        if (protocol == null) {
            protocol = "any"
        }
        let srcbeginport = this.vSrcbeginport
        if (srcbeginport == null) {
            srcbeginport = "any"
        }
        let srcbeginportinput = this.vSrcbeginportinput
        if (srcbeginportinput == null) {
            srcbeginportinput = "any"
        }
        let srcbeginportList = this.vSrcbeginportList


        let hide = true
        if(!this.checkHide(srcbeginportinput)){
            hide = false
        }

        let isdisabled = true
        if(!this.checkHide(srcbeginportinput)){
            isdisabled = false
        }
        return {
            protocol: protocol,
            srcbeginport: srcbeginport,
            srcbeginportinput: srcbeginportinput,
            srcbeginportList: srcbeginportList,
            hide: hide,
            isdisabled: isdisabled,
        }
    }, watch: {
        srcbeginport(newVal, oldVale) {
            // console.log("dstmask-new:", newVal);
            // console.log("dstmask-old:", oldVale);
            if (newVal !== oldVale) {
                this.$emit("update:vSrcbeginportinput", newVal)
            }
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="srcbeginport"  @change="selectSrc($event)"  v-model="srcbeginport" :disabled="isdisabled" style="min-width: 180px;">
		<option v-for="(idc,k) in srcbeginportList" :value="idc.value" :v-data="idc.data_other" :key="k" >
		    {{idc.name}}
		</option>
	</select>
	<div style="display: flex;justify-content: start;flex-direction: row;margin-top: 5px;">
        <input type="text" name="" ref="" maxlength="50"  v-model="srcbeginportinput" :class="{'hide':hide==true}" style="width: 180px;"/>
<!--        <select class="ui dropdown auto-width" ref="dstmask" name="dstmask" v-model="dstmask" :class="{'hide':hide==true}" style="min-width: 80px;margin-left: 10px;">-->
<!--            <option v-for="item in masks" :value="item.id" >{{item.value}}</option>-->
<!--        </select>-->
    </div>
</div>`
})