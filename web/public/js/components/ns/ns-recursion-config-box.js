// �ݹ�DNS����
Vue.component("ns-recursion-config-box", {
	props: ["v-recursion-config"],
	data: function () {
		let recursion = this.vRecursionConfig
		if (recursion == null) {
			recursion = {
				isOn: false,
				hosts: [],
				allowDomains: [],
				denyDomains: [],
				useLocalHosts: false
			}
		}
		if (recursion.hosts == null) {
			recursion.hosts = []
		}
		if (recursion.allowDomains == null) {
			recursion.allowDomains = []
		}
		if (recursion.denyDomains == null) {
			recursion.denyDomains = []
		}
		return {
			config: recursion,
			hostIsAdding: false,
			host: "",
			updatingHost: null
		}
	},
	methods: {
		changeHosts: function (hosts) {
			this.config.hosts = hosts
		},
		changeAllowDomains: function (domains) {
			this.config.allowDomains = domains
		},
		changeDenyDomains: function (domains) {
			this.config.denyDomains = domains
		},
		removeHost: function (index) {
			this.config.hosts.$remove(index)
		},
		addHost: function () {
			this.updatingHost = null
			this.host = ""
			this.hostIsAdding = !this.hostIsAdding
			if (this.hostIsAdding) {
				var that = this
				setTimeout(function () {
					let hostRef = that.$refs.hostRef
					if (hostRef != null) {
						hostRef.focus()
					}
				}, 200)
			}
		},
		updateHost: function (host) {
			this.updatingHost = host
			this.host = host.host
			this.hostIsAdding = !this.hostIsAdding

			if (this.hostIsAdding) {
				var that = this
				setTimeout(function () {
					let hostRef = that.$refs.hostRef
					if (hostRef != null) {
						hostRef.focus()
					}
				}, 200)
			}
		},
		confirmHost: function () {
			if (this.host.length == 0) {
				teaweb.warn("������DNS��ַ")
				return
			}

			// TODO У��Host
			// TODO ��������˿ں�
			// TODO ����ѡ��Э��

			this.hostIsAdding = false
			if (this.updatingHost == null) {
				this.config.hosts.push({
					host: this.host
				})
			} else {
				this.updatingHost.host = this.host
			}
		},
		cancelHost: function () {
			this.hostIsAdding = false
		}
	},
	template: `<div>
	<input type="hidden" name="recursionJSON" :value="JSON.stringify(config)"/>
	<table class="ui table definition selectable">
		<tbody>
			<tr>
				<td class="title">�Ƿ�����</td>
				<td>
					<div class="ui checkbox">
						<input type="checkbox" name="isOn" value="1" v-model="config.isOn"/>
						<label></label>
					</div>
					<p class="comment">���ú�����Ҳ���ĳ�������Ľ�����¼��������һ��DNS���ҡ�</p>
				</td>
			</tr>
		</tbody>
		<tbody v-show="config.isOn">
			<tr>
				<td>�ӽڵ㱾����ȡ<br/>�ϼ�DNS����</td>
				<td>
					<div class="ui checkbox">
						<input type="checkbox" name="useLocalHosts" value="1" v-model="config.useLocalHosts"/>
						<label></label>
					</div>
					<p class="comment">ѡ�к󣬽ڵ����ͼ��<code-label>/etc/resolv.conf</code-label>�ļ��ж�ȡDNS���á� </p>
				</td>
			</tr>
			<tr v-show="!config.useLocalHosts">
				<td>�ϼ�DNS������ַ *</td>
				<td>
					<div v-if="config.hosts.length > 0">
						<div v-for="(host, index) in config.hosts" class="ui label tiny basic">
							{{host.host}} &nbsp;
							<a href="" title="�޸�" @click.prevent="updateHost(host)"><i class="icon pencil tiny"></i></a>
							<a href="" title="ɾ��" @click.prevent="removeHost(index)"><i class="icon remove small"></i></a>
						</div>
						<div class="ui divider"></div>
					</div>
					<div v-if="hostIsAdding">
						<div class="ui fields inline">
							<div class="ui field">
								<input type="text" placeholder="DNS������ַ" v-model="host" ref="hostRef" @keyup.enter="confirmHost" @keypress.enter.prevent="1"/>
							</div>
							<div class="ui field">
								<button class="ui button tiny" type="button" @click.prevent="confirmHost">ȷ��</button> &nbsp; <a href="" title="ȡ��" @click.prevent="cancelHost"><i class="icon remove small"></i></a>
							</div>
						</div>
					</div>
					<div style="margin-top: 0.5em">
						<button type="button" class="ui button tiny" @click.prevent="addHost">+</button>
					</div>
				</td>
			</tr>
			<tr>
				<td>���������</td>
				<td><values-box name="allowDomains" :values="config.allowDomains" @change="changeAllowDomains"></values-box>
					<p class="comment">֧���Ǻ�ͨ���������<code-label>*.example.org</code-label>��</p>
				</td>
			</tr>
			<tr>
				<td>�����������</td>
				<td>
					<values-box name="denyDomains" :values="config.denyDomains" @change="changeDenyDomains"></values-box>
					<p class="comment">֧���Ǻ�ͨ���������<code-label>*.example.org</code-label>�����ȼ�������������ߡ�</p>
				</td>
			</tr>
		</tbody>
	</table>
	<div class="margin"></div>
</div>`
})