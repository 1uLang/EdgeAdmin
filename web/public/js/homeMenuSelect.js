var curSelectTopMenu = 1
var curSelectLeftMenu = 1
var curSelectTopDropId = -1

function onSetTopMenuId(menuId){
    console.log("onSetTopMenuId")
    console.log(menuId)
    if(menuId){
        curSelectTopMenu = menuId
    }
}

function onSetLeftMenuId(leftMenuId) { 
    if(leftMenuId){
        curSelectLeftMenu = leftMenuId
    }
}

function onSetSelectTopDropId(dropId) { 
    if(dropId){
        curSelectTopDropId = dropId
    }
}

function getSelectTopMenu() { 
   return curSelectTopMenu
}

function getSelectTopDropId() { 
    return curSelectTopDropId
}
function getSelectLeftMenu() { 
   return curSelectLeftMenu
}