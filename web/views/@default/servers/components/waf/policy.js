Tea.context(function () {
	this.upgradeTemplate = function () {
		teaweb.confirm("ȷ��Ҫ������Щ�¹�����", function () {
			this.$post(".upgradeTemplate")
				.params({
					policyId: this.firewallPolicy.id
				})
				.refresh()
		})
	}
})