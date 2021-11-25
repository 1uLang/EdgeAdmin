Vue.component("traffic-limit-config-box", {
	props: ["v-traffic-limit"],
	data: function () {
		let config = this.vTrafficLimit
		if (config == null) {
			config = {
				isOn: false,
				dailySize: {
					count: -1,
					unit: "gb"
				},
				monthlySize: {
					count: -1,
					unit: "gb"
				},
				totalSize: {
					count: -1,
					unit: "gb"
				},
				noticePageBody: ""
			}
		}
		if (config.dailySize == null) {
			config.dailySize = {
				count: -1,
				unit: "gb"
			}
		}
		if (config.monthlySize == null) {
			config.monthlySize = {
				count: -1,
				unit: "gb"
			}
		}
		if (config.totalSize == null) {
			config.totalSize = {
				count: -1,
				unit: "gb"
			}
		}
		return {
			config: config
		}
	},
	methods: {
		showBodyTemplate: function () {
			this.config.noticePageBody = `<!DOCTYPE html>
<html>
<head>
<title>Traffic Limit Exceeded Warning</title>
<body>

The site traffic has exceeded the limit. Please contact with the site administrator.

</body>
</html>`
		}
	},
	template: `<div>
	<input type="hidden" name="trafficLimitJSON" :value="JSON.stringify(config)"/>
	<table class="ui table selectable definition">
		<tbody>
			<tr>
				<td class="title">�Ƿ�����</td>
				<td>
					<checkbox v-model="config.isOn"></checkbox>
					<p class="comment">ע�⣺��������ͳ����ÿ5����ͳ��һ�Σ����Գ����������ƺ󣬶��û�������Ҳ�������ӳ١�</p>
				</td>
			</tr>
		</tbody>
		<tbody v-show="config.isOn">
			<tr>
				<td>����������</td>
				<td>
					<size-capacity-box :v-value="config.dailySize"></size-capacity-box>
				</td>
			</tr>
			<tr>
				<td>����������</td>
				<td>
					<size-capacity-box :v-value="config.monthlySize"></size-capacity-box>
				</td>
			</tr>
			<!--<tr>
				<td>��������</td>
				<td>
					<size-capacity-box :v-value="config.totalSize"></size-capacity-box>
					<p class="comment"></p>
				</td>
			</tr>-->
			<tr>
				<td>��ҳ��ʾ����</td>
				<td>
					<textarea v-model="config.noticePageBody"></textarea>
					<p class="comment"><a href="" @click.prevent="showBodyTemplate">[ʹ��ģ��]</a>�����ﵽ��������ʱ��ҳ��ʾ��HTML���ݣ�����д����ʾĬ�ϵ���ʾ���ݡ�</p>
				</td>
			</tr>
		</tbody>
	</table>
	<div class="margin"></div>
</div>`
})