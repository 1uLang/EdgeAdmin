Vue.component("nfw-agreement-selector", {
    mounted: function () {
        let that = this

        // Tea.action("/assembly/idc/options")
        // 	.post()
        // 	.success(function (resp) {
        // 		that.idcs = resp.data.idcs
        // 	})
    },methods: {
        selectSrc(event) {
            let value = event.target.value
            let vmap = ["bogons","bogonsv6","virusprot","sshlockout","any","(self)","lan","lanip",
                "lo0","wan","wanip"]
            if(vmap.indexOf(value) == -1){
                this.hide = false
            }else{
                this.hide = true
            }
            console.log(this.hide)
            console.log(value)
        }
    },
    props: ["v-srcs", "v-src","v-srcmask","v-masks","v-srcinput"],
    data: function () {
        let src = this.vSrc
        if (src == null) {
            src = ""
        }
        let srcs = this.vSrcs
        let srcmask = this.vSrcmask
        let masks = this.vMasks
        let srcinput = this.vSrcinput
        console.log("select == "+src);

        let hide = true
        let vmap = ["bogons","bogonsv6","virusprot","sshlockout","any","(self)","lan","lanip",
            "lo0","wan","wanip"]
        if(vmap.indexOf(src) == -1){
            hide = false
        }else{
            hide = true
        }
        return {
            srcs: srcs,
            src: src,
			masks: masks,
			srcmask: srcmask,
            srcinput: srcinput,
            hide:hide,
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="src" v-model="src" @change="selectSrc($event)">
<!--		<option value="-1">[选择数据中心]</option>-->
		<option v-for="idc in srcs" :value="idc.value">{{idc.name}}</option>
	</select>
	 <div style="display: flex;justify-content: start;flex-direction: row;" :class="{'hide':hide==true}">
        <input type="text" name="srcinput" maxlength="50" v-model="srcinput"/>
        <!-- 网络请求用这个 -->
        <!-- <nfw-source-selector :v-idc-id="sourceId"></nfw-source-selector> -->
        <!-- 本地数据用这个 -->
        <select class="ui dropdown auto-width" name="srcmask" v-model="srcmask">
            <option v-for="item in masks" :value="item.id">{{item.value}}</option>
        </select>
    </div>
</div>

`
})