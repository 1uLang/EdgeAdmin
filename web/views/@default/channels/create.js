Tea.context(function () {
    
    // this.name = ""
    // this.domain = ""
    // this.productName = ""
    // this.logo = "" ////images/logo.png
    // this.user = ""
    // this.mobile = ""
    // this.remake = ""
    // console.log(this.name)
    
    this.onImageChangeHandle = function(){
        var imageInput = document.getElementById("imageInput")
        if(imageInput && imageInput.value==""){
            return
        }
        var fileType = imageInput.files[0].type
        if (fileType != "image/png" && fileType != "image/jpeg" && fileType != "image/jpg") {
            teaweb.warn("请选择正确的图片文件")
            document.getElementById("imageInput").value=''
            return
        }
        this.onReadImage(imageInput.files[0])
    }

    this.onReadImage = function(filePath){
        let that = this
        var reader = new FileReader()
        reader.readAsDataURL(filePath)
        reader.onload = function(result){
            that.logo = reader.result
        }
    }
    
    this.delImage = function () {
        this.logo = ''
        //imageInput
        var imageInput = document.getElementById("imageInput")
        imageInput.value = ""
    }
    
})