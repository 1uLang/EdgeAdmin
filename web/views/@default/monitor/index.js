Tea.context(function () {

    this.nShowState = 1
    
    this.onChangeShowState = function(state){
        if(this.nShowState!=state){
            this.nShowState = state
        }
    }

    this.onOpenPost = function () {
      teaweb.popup(Tea.url(".createPopup"), {
        callback: function () {
          
        },
      });
    };

    this.onOpenWeb = function () {
        teaweb.popup(Tea.url(".create"), {
          callback: function () {
            
          },
        });
      };

    this.onAdd = function(){
        if(this.nShowState==1){
            this.onOpenPost()
        }else{
            this.onOpenWeb()
        }
    }
    this.onDeletePost = function(id){
        teaweb.confirm("确定要删除所选端口？", function () {
           
        })
    }
    this.onDeleteWeb = function(id){
        teaweb.confirm("确定要删除所选web？", function () {
           
        })
    }

    this.getStatusImgName = function(status){
        return status == 1? "/images/icon_success.png":"/images/icon_error.png"
    }

    this.tableData = [
        {id:1,status:1,ip:"192.168.0.1",post:"4433",user:"robin"},
        {id:2,status:2,ip:"192.168.0.2",post:"8080",user:"lusir"},
        {id:3,status:1,ip:"192.168.0.3",post:"3306",user:"robin"},
    ]

    this.tableData1 = [
        {id:1,status:1,statusCode:200,url:"https://www.fapiao.com",user:"robin"},
        {id:2,status:2,statusCode:500,url:"https://www.fapiao.com",user:"lusir"},
        {id:3,status:1,statusCode:404,url:"https://www.fapiao.com",user:"robin"}
    ]

});
  