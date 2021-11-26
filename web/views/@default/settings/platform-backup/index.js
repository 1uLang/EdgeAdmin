Tea.context(function () {

    this.onCopyFile = function () {
        // teaweb.confirm("确定启用规则？", function () {
        let that = this
        that.$post(".backup")
            .params({
                // id: id,
            }).success(resp => {
            if (resp.code === 200) {
                teaweb.success("备份成功", function () {
                    teaweb.reload()
                })
            }
        })
        //.refresh()
        // })
    }

    this.onCleanCopy = function () {
        teaweb.confirm("确定清除备份？", function () {
            let that = this
            that.$post(".clean")
                .params({
                    // id: id,
                })
                .refresh()
        })
    }

    this.onDownLoadFlie = function (item) {
        let that = this
        let url = "/remote.php/dav/files/admin/平台数据/" + item.name
        let filename=item.name
        that.$get("/databackup/download").params({
            name: filename,
            fp: url
        }).timeout(120).success((res) => {
            that.onDownload(res, "application/zip")
        }).fail((res) => {
            // that.onHideLoading()
            teaweb.warn(res.message)
        })
    }

    this.onDownload = function (res, fileType) {
        try {
            let that = this
            var bstr = atob(res.data.body)
            let n = bstr.length
            let u8arr = new Uint8Array(n)
            while (n--) {
                u8arr[n] = bstr.charCodeAt(n);
            }

            const blob = new Blob([u8arr], {type: fileType});
            const reader = new FileReader();
            reader.readAsDataURL(blob);
            reader.onload = (e) => {
                const a = document.createElement('a');
                a.download = res.data.fileName;
                a.href = e.target.result;
                document.body.appendChild(a);
                a.click();
                document.body.removeChild(a);
                // that.onHideLoading()
                // that.uploadFileSuccess()
            }
        } catch (e) {
            // this.onHideLoading()
        }

    }

    this.onResetFile = function (item) {
        let that = this
        that.$post(".recover")
            .params({
                id: item.id,
            }).success(resp => {
            if (resp.code === 200) {
                teaweb.success("恢复成功", function () {
                    teaweb.reload()
                })
            }
        })
        // .refresh()
    }
    this.onDeleteFile = function (item) {
        teaweb.confirm("确定删除备份？", function () {
            let that = this
            that.$post(".delete")
                .params({
                    id: item.id,
                })
                .refresh()
        })
    }

    // this.tableData = [
    //     {id: 1, name: "20211213.zip", create_time: "2021-09-15 17:25:28", size: "500k"},
    //     {id: 2, name: "20211213.zip", create_time: "2021-09-15 17:25:28", size: "500k"},
    //     {id: 3, name: "20211213.zip", create_time: "2021-09-15 17:25:28", size: "500k"},
    // ]

    this.getTime = function (timestamp) {
        return new Date(timestamp * 1000).format("Y-m-d H:i:s")
    }

})