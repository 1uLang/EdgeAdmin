Tea.context(function () {

    this.hanleOpen = function () {
        teaweb.popup(Tea.url("/hids/agents/create"), {height:'23em',width:'50em'});
    };
    this.onDelete = function (agent){

        teaweb.confirm("确定要删除所选资产吗？", function () {
            this.$post("/hids/agents/delete")
                .params({
                    agent: agent,
                }).success(function () {
                window.location.reload()
            })
        })

    }

    this.onChangeTimeFormat = function (time) {
        var resultTime = "";
        if (time) {
            var tempTime = time.substring(0, time.lastIndexOf("Z"));
            resultTime = tempTime.replace("T", " ");
        }
        return resultTime;
    };

    this.dbClickId = -1
    this.onDoubleClick = function(id){
        this.dbClickId = id
    }

    this.onCheckKeyDown = function(frm, event){
        var event = window.event ? window.event : event;
        if (event.keyCode == 13) {
            this.onSaveRemark();
        }
    }

    this.onSaveRemark = function(){
        var tempInput = document.getElementById("inputTxt")
        if(tempInput){
            console.log(tempInput.value)
        }
        
        this.dbClickId =-1
    }

});
