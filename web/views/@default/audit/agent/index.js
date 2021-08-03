Tea.context(function () {
    this.uploadFileSuccess = NotifyReloadSuccess("保存成功")
    this.pageState = 1

    this.fileDesc = ""

    this.bShowDialog = false

    this.onChangeState=function (id) {
        if(this.pageState!=id){
            this.pageState = id
        }
    }
    
    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
          var tempTime = time.substring(0, time.indexOf("."));
          resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
      };

    this.onDownFile = function (fileName,fileType) {
        // teaweb.confirm("确定下载该文件？",function() {
        //     this.$get("/databackup/download").params({
        //         name: fileName,
        //     }).success((res)=>{
        //         this.onDownloadFlie(res,fileType)
        //     })
        // })
    }

    this.onEdit = function (id) {
        teaweb.popup(Tea.url(".create"), {
			height: "300px",
            width:"460px",
			callback: function () {
				
			}
		})
    }

    this.onDelete = function (name) {
        // teaweb.confirm("确定要删除该文件？",function() {
        //     this.$post("/databackup/delete").params({
        //         Opt: "delete",
        //         name: name,
        //     }).refresh()
        // })
    }


    this.onDownloadFlie = function(res,fileType){
        var bstr = atob(res.data.body)
        let n = bstr.length
        let u8arr =new Uint8Array(n)
        while (n--) {
            u8arr[n] = bstr.charCodeAt(n);
        }

        const blob = new Blob([u8arr], { type:fileType});
        const reader = new FileReader();
        reader.readAsDataURL(blob);
        reader.onload = (e) => {
          const a = document.createElement('a');
          a.download = res.data.fileName;
          a.href = e.target.result;
          document.body.appendChild(a);
          a.click();
          document.body.removeChild(a);
        }
    }

    this.onuploadFile = function () {
        this.onShowLoading()
        setTimeout(() => {
            this.onHideLoading()
        }, 2000);
        
        // let that = this
        // var uploadFile = document.getElementById("uploadFile");
        // if(uploadFile.value==""){
        //     teaweb.warn("请选择上传文件")
        //     return
        // }
        // console.log(uploadFile.files[0].type)
        // var fm = document.getElementById('formData');
        // var fd = new FormData(fm);

        // this.$post("/databackup").params(fd)
        // .success(()=>{
        //     that.uploadFileSuccess()
        //     return true
        // })
        
        
    }

    this.onShowLoading =  function() {
        this.bShowDialog = true
    }
     
     
    this.onHideLoading = function () {
        this.bShowDialog = false
    }

    this.list = [
        {id:1,name:"数据库agent.exe",content_type:"CloudShield用户手册V1.0",used_bytes:"100K",last_modified:"2021-06-30 21:52:20"}
    ]
})