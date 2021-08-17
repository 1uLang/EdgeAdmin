Vue.component("nfw-dstendport-selector", {
    mounted: function () {


    },
    methods: {
        selectHandle() {
            for (var i = 0; i < this.dstbeginportList.length; i++) {
                if (this.dstbeginportList[i].value == this.dstendport ) {
                    if(this.dstbeginportList[i].data_other){
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
    props: ["v-protocol", "v-dstbeginport", "v-dstbeginport-list","v-dstendport","v-dstendport-input","v-dstendport-value"],
    data: function () {
        let protocol = this.vProtocol
        if (protocol == null) {
            protocol = "any"
        }
        let dstbeginport = this.vDstbeginport
        if (dstbeginport == null) {
            dstbeginport = "any"
        }
        let dstendport = this.vDstendport
        if (dstendport == null) {
            dstendport = "any"
        }
        let dstendportinput = this.vDstendportInput

        let dstbeginportList = this.vDstbeginportList


        let hide = this.checkInputHide(dstendport,dstbeginportList)

        let isdisabled = this.checkSelectDisabled(protocol)

        let dstendportValue = dstendport
        if(hide){
            dstendportValue = dstendportinput
        }
        return {
            protocol: protocol,
            dstbeginport: dstbeginport,
            dstendport:dstendport,
            dstendportinput: dstendportinput,
            dstbeginportList: dstbeginportList,
            hide: hide,
            isdisabled: isdisabled,
            dstendportValue: dstendportValue,
        }
    },
    watch: {
        vProtocol(newVal, oldVal){
            this.isdisabled = this.checkSelectDisabled(newVal)
            if(this.isdisabled){
                this.dstendport = "any"
                this.dstendportinput = ""
                this.dstendportValue = ""
                this.hide = false
            }
        },
        vDstbeginport(newVal, oldVal){
            this.dstendport = newVal
            this.hide = this.checkInputHide(this.dstendport,this.dstbeginportList)
        },
        dstendport(newVal, oldVal) {
            if (newVal !== oldVal) {
                this.$emit("update:vDstendport", newVal)
                this.$emit("update:vDstendportValue", newVal)

            }
        },
        dstendportinput(newVal, oldVal) {
            if (newVal !== oldVal) {
                this.$emit("update:vDstendportInput", newVal)
                this.$emit("update:vDstendportValue", newVal)

            }
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="dstendport"  @change="selectHandle"  v-model="dstendport" :disabled="isdisabled" style="min-width: 180px;">
		<option v-for="idc in dstbeginportList" :value="idc.value">
		    {{idc.name}}
		</option>
	</select>
	<div style="display: flex;justify-content: start;flex-direction: row;margin-top: 5px;">
        <input v-show='hide' type="text" name="" maxlength="50"  v-model="dstendportinput" style="width: 180px;"/>
    </div>
</div>`
})