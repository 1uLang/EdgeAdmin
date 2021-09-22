Tea.context(function () {
    
    this.channelName = ""
    this.channelHost = ""
    this.productName = ""
    this.productLogo = "/images/logo.png"
    this.linkPerson = ""
    this.phoneNum = ""
    this.markStr = ""
    
    
    this.onImageChangeHandle = function(){
        var imageInput = document.getElementById("imageInput")
        if(imageInput && imageInput.value==""){
            return
        }
        this.onReadImage(imageInput.files[0])
    }

    this.onReadImage = function(filePath){
        let that = this
        var reader = new FileReader()
        reader.readAsDataURL(filePath)
        reader.onload = function(result){
            that.productLogo = reader.result
        }
    }
    
})