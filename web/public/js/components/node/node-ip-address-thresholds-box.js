// �ڵ�IP��ֵ
Vue.component("node-ip-address-thresholds-box", {
	props: ["v-thresholds"],
	data: function () {
		let thresholds = this.vThresholds
		if (thresholds == null) {
			thresholds = []
		} else {
			thresholds.forEach(function (v) {
				if (v.items == null) {
					v.items = []
				}
				if (v.actions == null) {
					v.actions = []
				}
			})
		}

		return {
			editingIndex: -1,
			thresholds: thresholds,
			addingThreshold: {
				items: [],
				actions: []
			},
			isAdding: false,
			isAddingItem: false,
			isAddingAction: false,

			itemCode: "nodeAvgRequests",
			itemReportGroups: [],
			itemOperator: "lte",
			itemValue: "",
			itemDuration: "5",
			allItems: window.IP_ADDR_THRESHOLD_ITEMS,
			allOperators: [
				{
					"name": "С�ڵ���",
					"code": "lte"
				},
				{
					"name": "����",
					"code": "gt"
				},
				{
					"name": "������",
					"code": "neq"
				},
				{
					"name": "С��",
					"code": "lt"
				},
				{
					"name": "���ڵ���",
					"code": "gte"
				}
			],
			allActions: window.IP_ADDR_THRESHOLD_ACTIONS,

			actionCode: "up",
			actionBackupIPs: "",
			actionWebHookURL: ""
		}
	},
	methods: {
		add: function () {
			this.isAdding = !this.isAdding
		},
		cancel: function () {
			this.isAdding = false
			this.editingIndex = -1
			this.addingThreshold = {
				items: [],
				actions: []
			}
		},
		confirm: function () {
			if (this.addingThreshold.items.length == 0) {
				teaweb.warn("��Ҫ�������һ����ֵ")
				return
			}
			if (this.addingThreshold.actions.length == 0) {
				teaweb.warn("��Ҫ�������һ������")
				return
			}

			if (this.editingIndex >= 0) {
				this.thresholds[this.editingIndex].items = this.addingThreshold.items
				this.thresholds[this.editingIndex].actions = this.addingThreshold.actions
			} else {
				this.thresholds.push({
					items: this.addingThreshold.items,
					actions: this.addingThreshold.actions
				})
			}

			// ��ԭ
			this.cancel()
		},
		remove: function (index) {
			this.cancel()
			this.thresholds.$remove(index)
		},
		update: function (index) {
			this.editingIndex = index
			this.addingThreshold = {
				items: this.thresholds[index].items.$copy(),
				actions: this.thresholds[index].actions.$copy()
			}
			this.isAdding = true
		},

		addItem: function () {
			this.isAddingItem = !this.isAddingItem
			let that = this
			setTimeout(function () {
				that.$refs.itemValue.focus()
			}, 100)
		},
		cancelItem: function () {
			this.isAddingItem = false

			this.itemCode = "nodeAvgRequests"
			this.itmeOperator = "lte"
			this.itemValue = ""
			this.itemDuration = "5"
			this.itemReportGroups = []
		},
		confirmItem: function () {
			// ������ֵ�������
			if (["nodeHealthCheck"].$contains(this.itemCode)) {
				if (this.itemValue.toString().length == 0) {
					teaweb.warn("��ѡ������")
					return
				}

				let value = parseInt(this.itemValue)
				if (isNaN(value)) {
					value = 0
				} else if (value < 0) {
					value = 0
				} else if (value > 1) {
					value = 1
				}

				// ���
				this.addingThreshold.items.push({
					item: this.itemCode,
					operator: this.itemOperator,
					value: value,
					duration: 0,
					durationUnit: "minute",
					options: {}
				})
				this.cancelItem()
				return
			}

			if (this.itemDuration.length == 0) {
				let that = this
				teaweb.warn("������ͳ������", function () {
					that.$refs.itemDuration.focus()
				})
				return
			}
			let itemDuration = parseInt(this.itemDuration)
			if (isNaN(itemDuration) || itemDuration <= 0) {
				teaweb.warn("��������ȷ��ͳ������", function () {
					that.$refs.itemDuration.focus()
				})
				return
			}

			if (this.itemValue.length == 0) {
				let that = this
				teaweb.warn("������Ա�ֵ", function () {
					that.$refs.itemValue.focus()
				})
				return
			}
			let itemValue = parseFloat(this.itemValue)
			if (isNaN(itemValue)) {
				teaweb.warn("��������ȷ�ĶԱ�ֵ", function () {
					that.$refs.itemValue.focus()
				})
				return
			}


			let options = {}

			switch (this.itemCode) {
				case "connectivity": // ��ͨ��У��
					if (itemValue > 100) {
						let that = this
						teaweb.warn("��ͨ�ԶԱ�ֵ���ܳ���100", function () {
							that.$refs.itemValue.focus()
						})
						return
					}

					options["groups"] = this.itemReportGroups
					break
			}

			// ���
			this.addingThreshold.items.push({
				item: this.itemCode,
				operator: this.itemOperator,
				value: itemValue,
				duration: itemDuration,
				durationUnit: "minute",
				options: options
			})

			// ��ԭ
			this.cancelItem()
		},
		removeItem: function (index) {
			this.cancelItem()
			this.addingThreshold.items.$remove(index)
		},
		changeReportGroups: function (groups) {
			this.itemReportGroups = groups
		},
		itemName: function (item) {
			let result = ""
			this.allItems.forEach(function (v) {
				if (v.code == item) {
					result = v.name
				}
			})
			return result
		},
		itemUnitName: function (itemCode) {
			let result = ""
			this.allItems.forEach(function (v) {
				if (v.code == itemCode) {
					result = v.unit
				}
			})
			return result
		},
		itemDurationUnitName: function (unit) {
			switch (unit) {
				case "minute":
					return "����"
				case "second":
					return "��"
				case "hour":
					return "Сʱ"
				case "day":
					return "��"
			}
			return unit
		},
		itemOperatorName: function (operator) {
			let result = ""
			this.allOperators.forEach(function (v) {
				if (v.code == operator) {
					result = v.name
				}
			})
			return result
		},

		addAction: function () {
			this.isAddingAction = !this.isAddingAction
		},
		cancelAction: function () {
			this.isAddingAction = false
			this.actionCode = "up"
			this.actionBackupIPs = ""
			this.actionWebHookURL = ""
		},
		confirmAction: function () {
			this.doConfirmAction(false)
		},
		doConfirmAction: function (validated, options) {
			// �Ƿ��Ѵ���
			let exists = false
			let that = this
			this.addingThreshold.actions.forEach(function (v) {
				if (v.action == that.actionCode) {
					exists = true
				}
			})
			if (exists) {
				teaweb.warn("�˶����Ѿ���ӹ��ˣ������ظ����")
				return
			}

			if (options == null) {
				options = {}
			}

			switch (this.actionCode) {
				case "switch":
					if (!validated) {
						Tea.action("/ui/validateIPs")
							.params({
								"ips": this.actionBackupIPs
							})
							.success(function (resp) {
								if (resp.data.ips.length == 0) {
									teaweb.warn("�����뱸��IP", function () {
										that.$refs.actionBackupIPs.focus()
									})
									return
								}
								options["ips"] = resp.data.ips
								that.doConfirmAction(true, options)
							})
							.fail(function (resp) {
								teaweb.warn("�����IP '" + resp.data.failIP + "' ��ʽ����ȷ����������ύ", function () {
									that.$refs.actionBackupIPs.focus()
								})
							})
							.post()
						return
					}
					break
				case "webHook":
					if (this.actionWebHookURL.length == 0) {
						teaweb.warn("������WebHook URL", function () {
							that.$refs.webHookURL.focus()
						})
						return
					}
					if (!this.actionWebHookURL.match(/^(http|https):\/\//i)) {
						teaweb.warn("URL��ͷ������http://����https://", function () {
							that.$refs.webHookURL.focus()
						})
						return
					}
					options["url"] = this.actionWebHookURL
			}

			this.addingThreshold.actions.push({
				action: this.actionCode,
				options: options
			})

			// ��ԭ
			this.cancelAction()
		},
		removeAction: function (index) {
			this.cancelAction()
			this.addingThreshold.actions.$remove(index)
		},
		actionName: function (actionCode) {
			let result = ""
			this.allActions.forEach(function (v) {
				if (v.code == actionCode) {
					result = v.name
				}
			})
			return result
		}
	},
	template: `<div>
	<input type="hidden" name="thresholdsJSON" :value="JSON.stringify(thresholds)"/>
		
	<!-- �������� -->
	<div v-if="thresholds.length > 0">
		<div class="ui label basic small" v-for="(threshold, index) in thresholds">
			<span v-for="(item, itemIndex) in threshold.items">
				<span v-if="item.item != 'nodeHealthCheck'">
					[{{item.duration}}{{itemDurationUnitName(item.durationUnit)}}]
				</span> 
				{{itemName(item.item)}}
				
				<span v-if="item.item == 'nodeHealthCheck'">
					<!-- ������� -->
					<span v-if="item.value == 1">�ɹ�</span>
					<span v-if="item.value == 0">ʧ��</span>
				</span>
				<span v-else>
					<!-- ��ͨ�� -->
					<span v-if="item.item == 'connectivity' && item.options != null && item.options.groups != null && item.options.groups.length > 0">[<span v-for="(group, groupIndex) in item.options.groups">{{group.name}} <span v-if="groupIndex != item.options.groups.length - 1">&nbsp; </span></span>]</span>
				
					<span  class="grey">[{{itemOperatorName(item.operator)}}]</span> &nbsp;{{item.value}}{{itemUnitName(item.item)}} 
			 	</span>
			 	&nbsp;<span v-if="itemIndex != threshold.items.length - 1" style="font-style: italic">AND &nbsp;</span>
			</span>
			-&gt;
			<span v-for="(action, actionIndex) in threshold.actions">{{actionName(action.action)}}
			<span v-if="action.action == 'switch'">��{{action.options.ips.join(", ")}}</span>
			<span v-if="action.action == 'webHook'" class="small grey">({{action.options.url}})</span>
			 &nbsp;<span v-if="actionIndex != threshold.actions.length - 1" style="font-style: italic">AND &nbsp;</span></span>
			&nbsp;
			<a href="" title="�޸�" @click.prevent="update(index)"><i class="icon pencil small"></i></a> 
			<a href="" title="ɾ��" @click.prevent="remove(index)"><i class="icon small remove"></i></a>
		</div>
	</div>
	
	<!-- ����ֵ -->
	<div v-if="isAdding" style="margin-top: 0.5em">
		<table class="ui table celled">
			<thead>
				<tr>
					<td style="width: 50%; background: #f9fafb; border-bottom: 1px solid rgba(34,36,38,.1)">��ֵ</td>
					<th>����</th>
				</tr>
			</thead>
			<tr>
				<td style="background: white">
					<!-- �Ѿ���ӵ���Ŀ -->
					<div>
						<div v-for="(item, index) in addingThreshold.items" class="ui label basic small" style="margin-bottom: 0.5em;">
							<span v-if="item.item != 'nodeHealthCheck'">
								[{{item.duration}}{{itemDurationUnitName(item.durationUnit)}}]
							</span> 
							{{itemName(item.item)}}
							
							<span v-if="item.item == 'nodeHealthCheck'">
								<!-- ������� -->
								<span v-if="item.value == 1">�ɹ�</span>
								<span v-if="item.value == 0">ʧ��</span>
							</span>
							<span v-else>
								<!-- ��ͨ�� -->
								<span v-if="item.item == 'connectivity' && item.options != null && item.options.groups != null && item.options.groups.length > 0">[<span v-for="(group, groupIndex) in item.options.groups">{{group.name}} <span v-if="groupIndex != item.options.groups.length - 1">&nbsp; </span></span>]</span>
								 <span class="grey">[{{itemOperatorName(item.operator)}}]</span> {{item.value}}{{itemUnitName(item.item)}}
							 </span> 
							 &nbsp;
							<a href="" title="ɾ��" @click.prevent="removeItem(index)"><i class="icon remove small"></i></a>
						</div>
					</div>
					
					<!-- ������ӵ���Ŀ -->
					<div v-if="isAddingItem" style="margin-top: 0.8em">
						<table class="ui table">
							<tr>
								<td style="width: 6em">ͳ����Ŀ</td>
								<td>
									<select class="ui dropdown auto-width" v-model="itemCode">
									<option v-for="item in allItems" :value="item.code">{{item.name}}</option>
									</select>
									<p class="comment" style="font-weight: normal" v-for="item in allItems" v-if="item.code == itemCode">{{item.description}}</p>
								</td>
							</tr>
							<tr v-show="itemCode != 'nodeHealthCheck'">
								<td>ͳ������</td>
								<td>
									<div class="ui input right labeled">
										<input type="text" v-model="itemDuration" style="width: 4em" maxlength="4" ref="itemDuration" @keyup.enter="confirmItem()" @keypress.enter.prevent="1"/>
										<span class="ui label">����</span>
									</div>
								</td>
							</tr>
							<tr v-show="itemCode != 'nodeHealthCheck'">
								<td>������</td>
								<td>
									<select class="ui dropdown auto-width" v-model="itemOperator">
										<option v-for="operator in allOperators" :value="operator.code">{{operator.name}}</option>
									</select>
								</td>
							</tr>
							<tr v-show="itemCode != 'nodeHealthCheck'">
								<td>�Ա�ֵ</td>
								<td>
									<div class="ui input right labeled">
										<input type="text" maxlength="20" style="width: 5em" v-model="itemValue" ref="itemValue" @keyup.enter="confirmItem()" @keypress.enter.prevent="1"/>
										<span class="ui label" v-for="item in allItems" v-if="item.code == itemCode">{{item.unit}}</span>
									</div>
								</td>
							</tr>
							<tr v-show="itemCode == 'nodeHealthCheck'">
								<td>�����</td>
								<td>
									<select class="ui dropdown auto-width" v-model="itemValue">
										<option value="1">�ɹ�</option>
										<option value="0">ʧ��</option>
									</select>
									<p class="comment" style="font-weight: normal">ֻ��״̬�����ı��ʱ��Żᴥ����</p>
								</td>
							</tr>
							
							<!-- ��ͨ�� -->
							<tr v-if="itemCode == 'connectivity'">
								<td>�ն˷���</td>
								<td style="font-weight: normal">
									<div style="zoom: 0.8"><report-node-groups-selector @change="changeReportGroups"></report-node-groups-selector></div>
								</td>
							</tr>
						</table>
						<div style="margin-top: 0.8em">
							<button class="ui button tiny" type="button" @click.prevent="confirmItem">ȷ��</button>							 &nbsp;
							<a href="" title="ȡ��" @click.prevent="cancelItem"><i class="icon remove small"></i></a>
						</div>
					</div>
					<div style="margin-top: 0.8em" v-if="!isAddingItem">
						<button class="ui button tiny" type="button" @click.prevent="addItem">+</button>
					</div>
				</td>
				<td style="background: white">
					<!-- �Ѿ���ӵĶ��� -->
					<div>
						<div v-for="(action, index) in addingThreshold.actions" class="ui label basic small" style="margin-bottom: 0.5em">
							{{actionName(action.action)}} &nbsp;
							<span v-if="action.action == 'switch'">��{{action.options.ips.join(", ")}}</span>
							<span v-if="action.action == 'webHook'" class="small grey">({{action.options.url}})</span>
							<a href="" title="ɾ��" @click.prevent="removeAction(index)"><i class="icon remove small"></i></a>
						</div>
					</div>
					
					<!-- ������ӵĶ��� -->
					<div v-if="isAddingAction" style="margin-top: 0.8em">
						<table class="ui table">
							<tr>
								<td style="width: 6em">��������</td>
								<td>
									<select class="ui dropdown auto-width" v-model="actionCode">
										<option v-for="action in allActions" :value="action.code">{{action.name}}</option>
									</select>
									<p class="comment" v-for="action in allActions" v-if="action.code == actionCode">{{action.description}}</p>
								</td>
							</tr>
							
							<!-- �л� -->
							<tr v-if="actionCode == 'switch'">
								<td>����IP *</td>
								<td>
									<textarea rows="2" v-model="actionBackupIPs" ref="actionBackupIPs"></textarea>
									<p class="comment">ÿ��һ������IP��</p>
								</td>
							</tr>
							
							<!-- WebHook -->
							<tr v-if="actionCode == 'webHook'">
								<td>URL *</td>
								<td>
									<input type="text" maxlength="1000" placeholder="https://..." v-model="actionWebHookURL" ref="webHookURL" @keyup.enter="confirmAction()" @keypress.enter.prevent="1"/>
									<p class="comment">������URL������<code-label>https://example.com/webhook/api</code-label>��ϵͳ���ڴ�����ֵ��ʱ��ͨ��GET���ô�URL��</p>
								</td>
							</tr>
						</table>
						<div style="margin-top: 0.8em">
							<button class="ui button tiny" type="button" @click.prevent="confirmAction">ȷ��</button>	 &nbsp;
							<a href="" title="ȡ��" @click.prevent="cancelAction"><i class="icon remove small"></i></a>
						</div>
					</div>
					
					<div style="margin-top: 0.8em" v-if="!isAddingAction">
						<button class="ui button tiny" type="button" @click.prevent="addAction">+</button>
					</div>	
				</td>
			</tr>
		</table>
		
		<!-- �����ֵ -->
		<div>
			<button class="ui button tiny" :class="{disabled: (isAddingItem || isAddingAction)}" type="button" @click.prevent="confirm">ȷ��</button> &nbsp;
			<a href="" title="ȡ��" @click.prevent="cancel"><i class="icon remove small"></i></a>
		</div>
	</div>
	
	<div v-if="!isAdding" style="margin-top: 0.5em">
		<button class="ui button tiny" type="button" @click.prevent="add">+</button>
	</div>
</div>`
})