Tea.context(function () {

    this.onChangeShowState = function (state) {
        if (this.nShowState != state) {
            window.location.href = '/monitor?nShowState=' + state
        }
    }

    this.$delay(function () {
        if (this.errorMessage !== "" && this.errorMessage !== undefined) {
            teaweb.warn(this.errorMessage, function () {
            })
        }
    })
    this.onOpenPost = function () {
        teaweb.popup(Tea.url(".createPopup"), {
            height: "16em",
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload()
                })
            }
        })
    };

    this.onOpenWeb = function () {
        teaweb.popup(Tea.url(".create"), {
            height: "16em",
            callback: function () {
                teaweb.success("保存成功", function () {
                    teaweb.reload()
                })
            }
        })
    };

    this.onAdd = function () {
        if (this.nShowState == 1) {
            this.onOpenPost()
        } else {
            this.onOpenWeb()
        }
    }
    this.onDeletePost = function (id) {

        teaweb.confirm("确定要删除所选端口监控？", function () {
            this.$post(".delete")
                .params({
                    id: id
                })
                .refresh()
        })
    }
    this.onDeleteWeb = function (id) {
        teaweb.confirm("确定要删除所选web监控？", function () {

        })
    }

    this.getStatusImgName = function (status) {
        return status == 1 ? "/images/icon_success.png" : "/images/icon_error.png"
    }

    this.tableData = [
        {id: 1, status: 1, ip: "192.168.0.1", post: "4433", user: "robin"},
        {id: 2, status: 2, ip: "192.168.0.2", post: "8080", user: "lusir"},
        {id: 3, status: 1, ip: "192.168.0.3", post: "3306", user: "robin"},
    ]

    this.tableData1 = [
        {id: 1, status: 1, statusCode: 200, url: "https://www.fapiao.com", user: "robin"},
        {id: 2, status: 2, statusCode: 500, url: "https://www.fapiao.com", user: "lusir"},
        {id: 3, status: 1, statusCode: 404, url: "https://www.fapiao.com", user: "robin"}
    ]

});
  