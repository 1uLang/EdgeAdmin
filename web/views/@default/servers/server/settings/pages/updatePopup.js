Tea.context(function () {
	this.success = NotifyPopup

	this.bodyType = this.pageConfig.bodyType

	this.newStatus = ""
	if (this.pageConfig.newStatus > 0) {
		this.newStatus = this.pageConfig.newStatus
	}

	this.addHTMLTemplate = function () {
		this.$refs.htmlBody.value = `<!DOCTYPE html>
<html>
<head>
\t<title>ҳ�����</title>
\t<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
</head>
<body>

<h3>���ݱ���</h3>
<p>����</p>

<footer>Powered by GoEdge.</footer>

</body>
</html>`
	}
})