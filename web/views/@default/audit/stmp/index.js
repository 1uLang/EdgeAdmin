Tea.context(function () {
    this.success = NotifyReloadSuccess("保存成功")


    this.check = function () {
        this.isRunning = true
        console.log(this.host)
        console.log(this.port)
        console.log(this.email)
        console.log(this.password)
        this.$post("/audit/stmp/check").params({
            host: this.host,
            port: this.port,
            email: this.email,
            password: this.password,
            to: this.to,
        }).success(function () {
            teaweb.success("发送成功", function () {
                teaweb.reload()
            })

        })
        this.isRunning = false
    }
})