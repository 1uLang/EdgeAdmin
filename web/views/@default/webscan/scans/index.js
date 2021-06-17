var bScanning = false //是否在扫描中
function onStopScan(){
    if(checkValues.length > 0 && bScanning){
        console.log("touch onStopScan");
    }
}

function onCreateReport(){
    if(checkValues.length > 0){
        console.log("touch onCreateReport");
    }
}

function onDelete(){
    if(checkValues.length > 0 && !bScanning){
        console.log("touch onDelete");
    }
}

var checkValues = [] //选中的ID
function clickCheckbox () {
    var checkDomArr = document.querySelectorAll('.multi-table tbody input[type=checkbox]:checked')
    checkValues = []
    for (var i = 0, len = checkDomArr.length; i < len; i++) {
    checkValues.push(checkDomArr[i].value)
    }
    var allCheckDomArr = document.querySelectorAll('.multi-table tbody input[type=checkbox]')
    var allCheckbox = document.getElementById('js-all-checkbox')
    for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
    if (!allCheckDomArr[i].checked) {
        if (allCheckbox.checked) allCheckbox.checked = false
        break
    } else if (i === len - 1) {
        document.getElementById('js-all-checkbox').checked = true
        break
    }
    }
    updateBtnStatus()
}
function checkAll (current) {
    var allCheckDomArr = document.querySelectorAll('.multi-table tbody input[type=checkbox]')
    if (!current.checked) { // 点击的时候, 状态已经修改, 所以没选中的时候状态时true
    checkValues = []
    for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
        var checkStatus = allCheckDomArr[i].checked
        if (checkStatus) allCheckDomArr[i].checked = false
    }
    } else {
    checkValues = []
    for (var i = 0, len = allCheckDomArr.length; i < len; i++) {
        var checkStatus = allCheckDomArr[i].checked
        if (!checkStatus) allCheckDomArr[i].checked = true
        checkValues.push(allCheckDomArr[i].value)
    }
    }
    updateBtnStatus()
}
function updateBtnStatus(){
    
    const stopBtn = document.getElementById('stop-btn')
    const createBtn = document.getElementById('create-btn')
    const delBtn = document.getElementById('del-btn')
    if(checkValues.length > 0){
        stopBtn.style.backgroundColor = "#14539A"
        stopBtn.style.cursor = 'pointer'

        createBtn.style.backgroundColor = "#14539A"
        createBtn.style.cursor = 'pointer'

        delBtn.style.backgroundColor = "#D9001B"
        delBtn.style.cursor = 'pointer'
    }else{
        stopBtn.style.backgroundColor = "#AAAAAA"
        stopBtn.style.cursor = null

        createBtn.style.backgroundColor = "#AAAAAA"
        createBtn.style.cursor = null

        delBtn.style.backgroundColor = "#AAAAAA"
        delBtn.style.cursor = null
    }
}
