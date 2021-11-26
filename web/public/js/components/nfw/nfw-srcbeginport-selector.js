Vue.component("nfw-srcbeginport-selector", {
    mounted: function () {


    },
    methods: {
        selectHandle() {
            for (var i = 0; i < this.srcbeginportList.length; i++) {
                if (this.srcbeginportList[i].value == this.srcbeginport ) {
                    if(this.srcbeginportList[i].data_other){
                        this.hide = true
                    }else{
                        this.hide = false
                    }
                    break
                }
            }
        },
        checkInputHide(value,table){
            let hide = false
            if(table&&table.length > 0){
                for (var i = 0; i < table.length; i++) {
                    if (table[i].value == value ) {
                        if(table[i].data_other){
                            hide= true
                        }else{
                            hide = false
                        }

                    }
                }
            }

            return hide
        },
        checkSelectDisabled(value){
            if(value && (value =="udp" || value =="tcp" || value =="tcp/udp")){
                return false
            }
            return true
        }
    },
    props: ["v-protocol", "v-srcbeginport", "v-srcbeginport-list","v-srcbeginport-input","v-srcbeginport-value"],
    data: function () {
        let protocol = this.vProtocol
        if (protocol == null) {
            protocol = "any"
        }
        let srcbeginport = this.vSrcbeginport
        if (srcbeginport == null) {
            srcbeginport = "any"
        }
        let srcbeginportinput = this.vSrcbeginportInput

        let srcbeginportList = this.vSrcbeginportList


        let hide = this.checkInputHide(srcbeginport,srcbeginportList)

        let isdisabled = this.checkSelectDisabled(protocol)
        let srcbeginportValue = srcbeginport
        if(hide){
            srcbeginportValue = srcbeginportinput
        }
        return {
            protocol: protocol,
            srcbeginport: srcbeginport,
            srcbeginportinput: srcbeginportinput,
            srcbeginportList: srcbeginportList,
            hide: hide,
            isdisabled: isdisabled,
            srcbeginportValue: srcbeginportValue,
        }
    },
    watch: {
        vProtocol(newVal, oldVal){
            this.isdisabled = this.checkSelectDisabled(newVal)
            if(this.isdisabled){
                this.srcbeginport = "any"
                this.srcbeginportinput = ""
                this.srcbeginportValue = ""
                this.hide = false
            }
        },
        srcbeginport(newVal, oldVal) {
            // console.log("dstmask-new:", newVal);
            // console.log("dstmask-old:", oldVale);
            if (newVal !== oldVal) {
                this.$emit("update:vSrcbeginport", newVal)
                this.$emit("update:vSrcbeginportValue", newVal)
            }
        },
        srcbeginportinput(newVal, oldVal) {
            console.log("srcbeginportinput:", newVal);
            console.log("srcbeginportinput:", oldVal);
            if (newVal !== oldVal) {
                this.$emit("update:vSrcbeginportInput", newVal)
                this.$emit("update:vSrcbeginportValue", newVal)

            }
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="srcbeginport"  @change="selectHandle()"  v-model="srcbeginport" :disabled="isdisabled" style="min-width: 180px;">
		<option v-for="idc in srcbeginportList" :value="idc.value">
		    {{idc.name}}
		</option>
	</select>
	<div style="display: flex;justify-content: start;flex-direction: row;margin-top: 5px;">
        <input v-show='hide' type="text" name="" maxlength="50"  v-model="srcbeginportinput" style="width: 180px;"/>
    </div>
</div>`
})