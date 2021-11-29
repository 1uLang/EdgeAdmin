Tea.context(function () {
    this.success = NotifyReloadSuccess("保存成功")

    this.moveIndex = ""

    let that = this

    this.onRemoveTableItem = function (code, table) {
        if (code && table && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index] == code) {
                    table.splice(index, 1)
                }
            }
        }
        return table
    }

    this.onGetTableItemInfo = function (code, table) {
        if (table && code && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index].code == code) {
                    return table[index]
                }
            }
        }
        return null
    }

    this.onCheckHadValue = function (code, table) {
        if (table && code && table.length > 0) {
            for (var index = 0; index < table.length; index++) {
                if (table[index] == code) {
                    return true
                }
            }
        }
        return false
    }

    this.getShowImageName = function (code) {

        let bAllSelect = that.onCheckHadValue(code, that.selectList)
        if (bAllSelect) {
            return "/images/select_select.png"
        } else {
            let bSelect = false
            let tempItem = that.onGetTableItemInfo(code, that.features)
            if (tempItem) {
                for (var index = 0; index < tempItem.children.length; index++) {
                    if (that.onCheckHadValue(tempItem.children[index].code, that.selectList)) {
                        bSelect = true
                        break
                    }
                }
            }
            if (bSelect) {
                return "/images/select_half_select.png"
            } else {
                return "/images/select_box.png"
            }
        }
    }


    this.getItemShowImageName = function (id) {
        if (that.onCheckHadValue) {
            if (that.onCheckHadValue(id, that.selectList)) {
                return "/images/select_select.png"
            } else {
                return "/images/select_box.png"
            }
        }
        return "/images/select_box.png"
    }

    this.onShowChildItem = function (code) {
        var tempItem = this.onGetTableItemInfo(code, this.features)
        if (tempItem) {
            tempItem.bShowChild = !tempItem.bShowChild
        }
    }

    this.onMouseEnter = function (code) {
        this.moveIndex = code
    }

    this.onMouseLeave = function () {
        this.moveIndex = ""
    }

    this.onSelectValue = function (code) {
        if (this.onCheckHadValue(code, this.selectList)) {
            this.onRemoveTableItem(code, this.selectList)
            let tempItem = this.onGetTableItemInfo(code, this.features)
            if (tempItem) {
                for (var index = 0; index < tempItem.children.length; index++) {
                    this.onRemoveTableItem(tempItem.children[index].code, this.selectList)
                }
            }
        } else {
            this.selectList.push(code)
            let tempItem = this.onGetTableItemInfo(code, this.features)
            if (tempItem && code !== 'lb-tcp') {
                for (var index = 0; index < tempItem.children.length; index++) {
                    if (!this.onCheckHadValue(tempItem.children[index].code, this.selectList)) {
                        this.selectList.push(tempItem.children[index].code)
                    }
                }
            }
        }
    }

    this.onSelectChildValue = function (code, parentId) {
        if (this.onCheckHadValue(code, this.selectList)) {
            this.onRemoveTableItem(code, this.selectList)
        } else {
            this.selectList.push(code)
        }

        let bFindAll = true
        let tempItem = this.onGetTableItemInfo(parentId, this.features)
        if (tempItem) {
            for (var index = 0; index < tempItem.children.length; index++) {
                if (!this.onCheckHadValue(tempItem.children[index].code, this.selectList)) {
                    console.log(tempItem.children[index].code)
                    bFindAll = false
                    break
                }
            }
        }
        if (bFindAll) {
            this.selectList.push(parentId)
        } else {
            this.onRemoveTableItem(parentId, this.selectList)
        }
    }

    this.onSave = function () {
        this.$post(".features").params({
            UserId: this.userId,
            Codes: this.selectList,
        }).success(resp => {
            if (resp.code === 200) {
                teaweb.success("修改成功")
            }
        })
    }
})