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
\t<title>页面标题</title>
\t<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
</head>
<body>

<h3>内容标题</h3>
<p>内容</p>

<footer>Powered by GoEdge.</footer>

</body>
</html>`
	}
})