Vue.component("user-selector", {
	mounted: function () {
		let that = this

		Tea.action("/users/option")
			.get()
			.success(function (resp) {
				that.users = resp.data.users
			})
	},
	props: [""],
	data: function () {

		return {
			users: [],
			user:"0",
		}
	},
	template: `<div>
	<select class="ui dropdown auto-width" name="user_id" v-model="user">
		<option value="0">[选择所属用户]</option>
		<option v-for="idc in users" :value="idc.id">{{idc.username}}</option>
	</select>
</div>`
})