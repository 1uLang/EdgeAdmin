Tea.context(function () {
    this.createPolicy = function () {
        teaweb.popup("/ddos/host/createPopup", {
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload()
                })
            }
        })
    }

})