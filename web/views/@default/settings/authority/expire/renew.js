Tea.context(function () {

    this.onRenewal = function (){
        var secret = document.getElementById('secret').value

        var tempFormData = new FormData();
        tempFormData.append("secret",secret)
        reqApi("post","/renewal",tempFormData,null,
            (res)=>{
                teaweb.success("修改成功",function(){
                    var loginPart =  document.getElementById("loginPart")
                    var renewalPart =  document.getElementById("renewalPart")
                    var resetPwdPart =  document.getElementById("resetPart")
                    document.getElementById('password').value = ""
                    resetPwdPart.style.display = "none"
                    renewalPart.style.display = "none"
                    loginPart.style.display = "block"
                    onGetRefreshToken()
                    onGetToken()
                })
            },
            (res)=>{
                teaweb.warn(res.message,function(){
                    onGetRefreshToken()
                    onGetToken()
                })
            }
        )
    }
})