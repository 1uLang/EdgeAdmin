Vue.component("http-webp-config-box", {
	props: ["v-webp-config", "v-is-location", "v-is-group"],
	data: function () {
		let config = this.vWebpConfig
		if (config == null) {
			config = {
				isPrior: false,
				isOn: false,
				quality: 50,
				minLength: {count: 0, "unit": "kb"},
				maxLength: {count: 0, "unit": "kb"},
				mimeTypes: ["image/png", "image/jpeg", "image/bmp", "image/x-ico", "image/gif"],
				extensions: [".png", ".jpeg", ".jpg", ".bmp", ".ico"],
				conds: null
			}
		}

		if (config.mimeTypes == null) {
			config.mimeTypes = []
		}
		if (config.extensions == null) {
			config.extensions = []
		}

		return {
			config: config,
			moreOptionsVisible: false,
			quality: config.quality
		}
	},
	watch: {
		quality: function (v) {
			let quality = parseInt(v)
			if (isNaN(quality)) {
				quality = 90
			} else if (quality < 1) {
				quality = 1
			} else if (quality > 100) {
				quality = 100
			}
			this.config.quality = quality
		}
	},
	methods: {
		isOn: function () {
			return ((!this.vIsLocation && !this.vIsGroup) || this.config.isPrior) && this.config.isOn
		},
		changeExtensions: function (values) {
			values.forEach(function (v, k) {
				if (v.length > 0 && v[0] != ".") {
					values[k] = "." + v
				}
			})
			this.config.extensions = values
		},
		changeMimeTypes: function (values) {
			this.config.mimeTypes = values
		},
		changeAdvancedVisible: function () {
			this.moreOptionsVisible = !this.moreOptionsVisible
		},
		changeConds: function (conds) {
			this.config.conds = conds
		}
	},
	template: `<div>
	<input type="hidden" name="webpJSON" :value="JSON.stringify(config)"/>
	<table class="ui table definition selectable">
		<prior-checkbox :v-config="config" v-if="vIsLocation || vIsGroup"></prior-checkbox>
		<tbody v-show="(!vIsLocation && !vIsGroup) || config.isPrior">
			<tr>
				<td class="title">�Ƿ�����</td>
				<td>
					<div class="ui checkbox">
						<input type="checkbox" value="1" v-model="config.isOn"/>
						<label></label>
					</div>
					<p class="comment">ѡ�к��ʾ�����Զ�WebPѹ����</p>
				</td>
			</tr>
		</tbody>
		<tbody v-show="isOn()">
			<tr>
				<td>ͼƬ����</td>
				<td>
					<div class="ui input right labeled">
						<input type="text" v-model="quality" style="width: 5em" maxlength="4"/>
						<span class="ui label">%</span>
					</div>
					<p class="comment">ȡֵ��0��100֮�䣬��ֵԽ�����ɵ�ͼ��Խ������ͬʱ�ļ��ߴ�Ҳ��Խ��</p>
				</td>
			</tr>
			<tr>
				<td>֧�ֵ���չ��</td>
				<td>
					<values-box :values="config.extensions" @change="changeExtensions" placeholder="���� .html"></values-box>
					<p class="comment">������Щ��չ����URL���ᱻת��WebP�������ִ�Сд��</p>
				</td>
			</tr>
			<tr>
				<td>֧�ֵ�MimeType</td>
				<td>
					<values-box :values="config.mimeTypes" @change="changeMimeTypes" placeholder="���� text/*"></values-box>
					<p class="comment">��Ӧ��Content-Type�������ЩMimeType�����ݽ��ᱻת��WebP��</p>
				</td>
			</tr>
		</tbody>
		<more-options-tbody @change="changeAdvancedVisible" v-if="isOn()"></more-options-tbody>
		<tbody v-show="isOn() && moreOptionsVisible">
				<tr>
					<td>������С����</td>
				<td>
					<size-capacity-box :v-name="'minLength'" :v-value="config.minLength" :v-unit="'kb'"></size-capacity-box>
					<p class="comment">0��ʾ�����ƣ����ݳ��ȴ��ļ��ߴ��Content-Length�л�ȡ��</p>
				</td>
			</tr>
			<tr>
				<td>������󳤶�</td>
				<td>
					<size-capacity-box :v-name="'maxLength'" :v-value="config.maxLength" :v-unit="'mb'"></size-capacity-box>
					<p class="comment">0��ʾ�����ƣ����ݳ��ȴ��ļ��ߴ��Content-Length�л�ȡ��</p>
				</td>
			</tr>
			<tr>
				<td>ƥ������</td>
				<td>
					<http-request-conds-box :v-conds="config.conds" @change="changeConds"></http-request-conds-box>
	</td>
			</tr>
		</tbody>
	</table>			
	<div class="ui margin"></div>
</div>`
})