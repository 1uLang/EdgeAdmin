Vue.component("nfw-srcendport-selector", {
    mounted: function () {


    },
    methods: {
        selectHandle() {
            for (var i = 0; i < this.srcbeginportList.length; i++) {
                if (this.srcbeginportList[i].value == this.srcendport ) {
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
                            hide=false
                        }
                    }
                }
            }

            return hide
        },
        checkSelectDisabled(value){
            if(value && (value =="udp" || value =="tcp"  || value =="tcp/udp")){
                return false
            }
            return true
        }
    },
    props: ["v-protocol", "v-srcbeginport", "v-srcbeginport-list","v-srcendport","v-srcendport-input","v-srcbeginport-value"],
    data: function () {
        let protocol = this.vProtocol
        if (protocol == null) {
            protocol = "any"
        }
        let srcbeginport = this.vSrcbeginport
        if (srcbeginport == null) {
            srcbeginport = "any"
        }
        let srcendport = this.vSrcendport
        if (srcendport == null) {
            srcendport = "any"
        }
        let srcendportinput = this.vSrcendportInput

        let srcbeginportList = this.vSrcbeginportList


        let hide = this.checkInputHide(srcendport,srcbeginportList)

        let isdisabled = this.checkSelectDisabled(protocol)
        let srcendportValue = srcendport
        if(hide){
            srcendportValue = srcendportinput
        }
        return {
            protocol: protocol,
            srcbeginport: srcbeginport,
            srcendport:srcendport,
            srcendportinput: srcendportinput,
            srcbeginportList: srcbeginportList,
            hide: hide,
            isdisabled: isdisabled,
            srcendportValue: srcendportValue,
        }
    },
    watch: {
        vProtocol(newVal, oldVal){
            this.isdisabled = this.checkSelectDisabled(newVal)
            if(this.isdisabled){
                this.srcendport = "any"
                this.srcendportinput = ""
                this.srcendportValue = ""
                this.hide = false
            }
        },
        vSrcbeginport(newVal, oldVal){
            this.srcendport = newVal
            this.hide = this.checkInputHide(this.srcendport,this.srcbeginportList)
        },
        srcendport(newVal, oldVal) {
            if (newVal !== oldVal) {
                this.$emit("update:vSrcendport", newVal)
                this.$emit("update:vSrcendportValue", newVal)
            }
        },
        srcendportinput(newVal, oldVal) {
            if (newVal !== oldVal) {
                this.$emit("update:vSrcendportInput", newVal)
                this.$emit("update:vSrcendportValue", newVal)
                // if(newVal!=this.srcendport){
                //     this.srcendport = newVal
                //     this.$emit("update:vSrcendport", newVal)
                // }
            }
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="srcendport"  @change="selectHandle"  v-model="srcendport" :disabled="isdisabled" style="min-width: 180px;">
		<option v-for="idc in srcbeginportList" :value="idc.value">
		    {{idc.name}}
		</option>
	</select>
	<div style="display: flex;justify-content: start;flex-direction: row;margin-top: 5px;">
        <input v-show='hide' type="text" name="" maxlength="50"  v-model="srcendportinput" style="width: 180px;"/>
    </div>
</div>`
})