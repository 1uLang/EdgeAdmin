Tea.context(function () {

    this.os = ""
    this.version = ""
    this.architecture = ""
    this.command = ""

    this.OSList = ["Red Hat / CentOS", "Debian / Ubuntu", "Windows", "MacOS"]
    this.versionList = ['CentOS5', 'CentOS6 or higher']
    this.architectureList = ["i386", "x86_64", "armhf", "aarch64"]
    this.showVersion = function () {
        console.log("===========", this.os)
        if (this.os === 'Red Hat / CentOS')
            return true
        return false
    }
    this.showArchitecture = function () {
        console.log("===========", this.version)
        if (this.version === 'Red Hat / CentOS' || this.os === 'Debian / Ubuntu')
            return true
        return false
    }
    this.changeOS = function () {
        this.version = ''
        this.architecture = ''
        this.command = ''
    }
    this.createCommand = function () {

        if (this.os === 'Windows') {
            this.command = "windows command ...."
        }
        if (this.os === 'MacOS') {
            this.command = "mac os command ..."
        }
        if (this.architecture !== '') {
            if (this.version !== '') {
                this.command = "rea heat Centos command ..."
            } else {
                this.command = "debian ubuntu command ..."
            }
        }

        return true
    }
})