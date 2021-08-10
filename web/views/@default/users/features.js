Tea.context(function () {
	this.success = NotifyReloadSuccess("保存成功")

	this.selectList = []
    this.moveIndex=-1

	let that = this

    this.onRemoveTableItem=function (id,table) { 
        if(id && table && table.length>0 ){
            for(var index=0;index<table.length;index++){
                if(table[index]==id){
                    table.splice(index,1)
                }
            }
        }
        return table
    }

    this.onGetTableItemInfo = function (id, table) {
        if (table && id && table.length > 0 && id > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].id == id) {
                    return table[index]
                }
            }
        }
        return null
    }

	this.onCheckHadValue = function (id, table) {
        if (table && id && table.length > 0 && id > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index] == id) {
                    return true
                }
            }
        }
        return false
    }

	this.getShowImageName=function (id) { 
		
		let bAllSelect = that.onCheckHadValue(id,that.selectList)
		if(bAllSelect){
			return "/images/select_select.png"
		}else{
            let bSelect = false
            let tempItem = that.onGetTableItemInfo(id,that.menuData)
            if(tempItem){
                for(var index=0;index<tempItem.children.length;index++){
                    if(that.onCheckHadValue(tempItem.children[index].id,that.selectList)){
                        bSelect = true
                        break
                    }
                }
            }
			if(bSelect){
				return "/images/select_half_select.png"
			}else{
				return "/images/select_box.png"
			}
		}
	}


	this.getItemShowImageName = function (id) { 
		if(that.onCheckHadValue){
			if(that.onCheckHadValue(id,that.selectList)){
				return "/images/select_select.png"
			}else{
				return "/images/select_box.png"
			}
		}
		return "/images/select_box.png" 
	}

    this.onShowChildItem=function (id) { 
        var tempItem = this.onGetTableItemInfo(id,this.menuData)
        if(tempItem){
            tempItem.bShowChild = !tempItem.bShowChild
        }
    }

    this.onMouseEnter = function (id) { 
        this.moveIndex = id
    }

    this.onMouseLeave = function (id) { 
        this.moveIndex = -1
    }

    this.onSelectValue = function (id) { 
        if(this.onCheckHadValue(id,this.selectList)){
            this.onRemoveTableItem(id,this.selectList)
            let tempItem = this.onGetTableItemInfo(id,this.menuData)
            if(tempItem){
                for(var index=0;index<tempItem.children.length;index++){
                    this.onRemoveTableItem(tempItem.children[index].id,this.selectList)
                }
            }
        }else{
            this.selectList.push(id)
            let tempItem = this.onGetTableItemInfo(id,this.menuData)
            if(tempItem){
                for(var index=0;index<tempItem.children.length;index++){
                    if(!this.onCheckHadValue(tempItem.children[index].id,this.selectList)){
                        this.selectList.push(tempItem.children[index].id)
                    }
                }
            }
        }
    }

    this.onSelectChildValue = function (id,parentId) { 
        if(this.onCheckHadValue(id,this.selectList)){
            this.onRemoveTableItem(id,this.selectList)
        }else{
            this.selectList.push(id)
        }
        
        let bFindAll = true
        let tempItem = this.onGetTableItemInfo(parentId,this.menuData)
        if(tempItem){
            for(var index=0;index<tempItem.children.length;index++){
                if(!this.onCheckHadValue(tempItem.children[index].id,this.selectList)){
                    console.log(tempItem.children[index].id)
                    bFindAll = false
                    break
                }
            }
        }
        if(bFindAll){
            this.selectList.push(parentId)
        }else{
            this.onRemoveTableItem(parentId,this.selectList)
        }
    }


	this.menuData = [
		{
			id:1,
			name:"这是父节点1",
			description:"这是副标题1",
            bShowChild:false,
			children:[
				{id:2,name:"父节点1-id=>2",description:"这是副标题"},
				{id:3,name:"父节点1-id=>3",description:"这是副标题"},
				{id:4,name:"父节点1-id=>4",description:"这是副标题"},
				{id:5,name:"父节点1-id=>5",description:"这是副标题"},
				{id:6,name:"父节点1-id=>6",description:"这是副标题"},
			],
		},
		{
			id:7,
			name:"这是父节点2",
			description:"这是副标题2",
            bShowChild:false,
			children:[
				{id:8, name:"父节点2-id=>8", description:"这是副标题"},
				{id:9, name:"父节点2-id=>9", description:"这是副标题"},
				{id:10,name:"父节点2-id=>10",description:"这是副标题"},
				{id:11,name:"父节点2-id=>11",description:"这是副标题"},
				{id:12,name:"父节点2-id=>12",description:"这是副标题"},
			],
		}
	]
})