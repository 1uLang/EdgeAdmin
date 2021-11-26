Tea.context(function () {

    //这边数据是从服务器获取的 本地只做展示
    this.name = ""
    this.address = "" 
    this.typeSelect = 1 
    this.openState = 1

    this.id = 0
    this.$delay = function(){
        this.id = this.getUrlParam("id")
        if(this.id && this.id.length>0){
            this.name = this.getUrlParam("name")
            if(this.name.length>0){
                this.name = unescape(this.name)
            }
            this.address = this.getUrlParam("address")
            if(this.address.length>0){
                if(this.address.indexOf("\/")!=-1){
                    let tempList = this.address.split("\/")
                    this.address = ""
                    for(var index=0;index<tempList.length;index++){
                        if(index==tempList.length-1){
                            this.address = this.address + tempList[index]
                        }else{
                            this.address = this.address + tempList[index]+"."
                        }
                    }
                }
            }
            this.openState = this.getUrlParam("enabled")
            this.typeSelect = this.getUrlParam("selectType")
        }
        
    }

    this.getUrlParam = function(param){
        var reg = new RegExp("(^|&)"+param+"=([^&]*)(&|$)")
        var r = window.location.search.substr(1).match(reg)
        if(r!=null){
            return unescape(r[2]);
        }
        return ""
    }

    this.onSelectOpenState = function () {
        this.openState = this.openState == 1 ? 1:0
    }
    
    this.onSave = function(){
        if(this.id && this.id.length>0){
            this.onEdit()
        }else{
            this.onAdd()
        }
    }

    this.onAdd =function(){
        let curEnabled = false
        if(this.openState == 1){
            curEnabled = true
        }
        this.$post("/resmon").params(
            {
                name:this.name,
                host:this.address,
                aid:this.typeSelect,
                on:curEnabled,
            }
        ).success(
            function(){
                teaweb.closePopup()
                window.location.reload()
            }
        )
    }

    this.onEdit = function(){
        let curEnabled = false
        if(this.openState == 1){
            curEnabled = true
        }
        this.$post("/resmon/update").params(
            {
                id:this.id,
                name:this.name,
                host:this.address,
                aid:this.typeSelect,
                on:curEnabled,
            }
        ).success(
            function(){
                teaweb.closePopup()
                window.location.reload()
            }
        )
    }
})