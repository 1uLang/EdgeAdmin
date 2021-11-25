Tea.context(function () {
	this.$delay(function () {
		teaweb.datepicker("day-from-picker")
		teaweb.datepicker("day-to-picker")
	})

	this.updateRead = function (logId) {
		this.$post(".readLogs")
			.params({
				logIds: [logId]
			})
			.success(function () {
				teaweb.reload()
			})
	}

	this.updatePageRead = function () {
		let logIds = this.logs.map(function (v) {
			return v.id
		})
		teaweb.confirm("ȷ��Ҫ���ñ�ҳ��־Ϊ�Ѷ���", function () {
			this.$post(".readLogs")
				.params({
					logIds: logIds
				})
				.success(function () {
					teaweb.reload()
				})
		})
	}

	this.updateAllRead = function () {
		teaweb.confirm("ȷ��Ҫ����������־Ϊ�Ѷ���", function () {
			this.$post(".readAllLogs")
				.params({})
				.success(function () {
					teaweb.reload()
				})
		})
	}
})