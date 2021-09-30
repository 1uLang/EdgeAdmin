// 单个渠道选择
Vue.component("channels-selector", {
    mounted: function () {
        let that = this

        Tea.action("/channels")
            .params({
                pageNum: 1,
                pageSize:999,
                status: "1",
            })
            .post()
            .success(function (resp) {
                that.channels = resp.data.tableData
            })
    },
    props: ["v-channel-id"],
    data: function () {
        let channelId = this.vChannelId
        if (channelId == null) {
            channelId = 0
        }
        return {
            channels: [],
            channelId: channelId
        }
    },
    template: `<div>
	<select class="ui dropdown auto-width" name="channelId" v-model="channelId">
		<option value="0">[选择渠道]</option>
		<option v-for="chan in channels" :value="chan.id">{{chan.name}}</option>
	</select>
</div>`
})