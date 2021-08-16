Vue.component("nfw-dstbeginport-selector", {
    mounted: function () {


    },
    methods: {
        selectHandle() {
            for (var i = 0; i < this.dstbeginportList.length; i++) {
                if (this.dstbeginportList[i].value == this.dstbeginport ) {
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
    props: ["v-protocol", "v-dstbeginport", "v-dstbeginport-list","v-dstbeginport-input","v-dstbeginport-value"],
    data: function () {
        let protocol = this.vProtocol
        if (protocol == null) {
            protocol = "any"
        }
        let dstbeginport = this.vDstbeginport
        if (dstbeginport == null) {
            dstbeginport = "any"
        }
        let dstbeginportinput = this.vDstbeginportInput

        let dstbeginportList = this.vDstbeginportList


        let hide = this.checkInputHide(dstbeginport,dstbeginportList)

        let isdisabled = this.checkSelectDisabled(protocol)
        let dstbeginportValue = dstbeginport
        if(hide){
            dstbeginportValue = dstbeginportinput
        }
        return {
            protocol: protocol,
            dstbeginport: dstbeginport,
            dstbeginportinput: dstbeginportinput,
            dstbeginportList: dstbeginportList,
            hide: hide,
            isdisabled: isdisabled,
            dstbeginportValue: dstbeginportValue,
        }
    },
    watch: {
        vProtocol(newVal, oldVal){
            this.isdisabled = this.checkSelectDisabled(newVal)
            if(this.isdisabled){
                this.dstbeginport = "any"
                this.dstbeginportinput = ""
                this.dstbeginportValue = ""
                this.hide = false
            }
        },
        dstbeginport(newVal, oldVal) {
            // console.log("dstmask-new:", newVal);
            // console.log("dstmask-old:", oldVale);
            if (newVal !== oldVal) {
                this.$emit("update:vDstbeginport", newVal)
                this.$emit("update:vDstbeginportValue", newVal)

            }
        },
        dstbeginportinput(newVal, oldVal) {
            console.log("dstbeginportinput:", newVal);
            console.log("dstbeginportinput:", oldVal);
            if (newVal !== oldVal) {
                this.$emit("update:vDstbeginportInput", newVal)
                this.$emit("update:vDstbeginportValue", newVal)

            }
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="dstbeginport"  @change="selectHandle"  v-model="dstbeginport" :disabled="isdisabled" style="min-width: 180px;">
		<option v-for="idc in dstbeginportList" :value="idc.value">
		    {{idc.name}}
		</option>
	</select>
	<div style="display: flex;justify-content: start;flex-direction: row;margin-top: 5px;">
        <input v-show='hide' type="text" name="" maxlength="50"  v-model="dstbeginportinput" style="width: 180px;"/>
    </div>
</div>`
})