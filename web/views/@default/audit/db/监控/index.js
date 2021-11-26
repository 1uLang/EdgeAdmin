Tea.context(function () {

    const STATE_IDLE = 0;
    const STATE_CONNECTING = 1;
    const STATE_WAITING = 2;
    const STATE_CONNECTED = 3;
    const STATE_DISCONNECTING = 4;
    const STATE_DISCONNECTED = 5;
    
    this.wsServer = ""

    this.showTips = ""
    this.loading = true

    this.client = null
    this.qs = null

    this.bClose = false

    this.$delay(function () {
        this.qs = Qs
        window.addEventListener('beforeunload', function () {
            if(this.client){
                this.client.disconnect();
            }
        })

    })

    this.onCreateSocket = function(connectionId){
        let tunnel = new Guacamole.WebSocketTunnel(this.wsServer + '/tunnel');
        tunnel.onstatechange = this.onClientStateChange;
        this.client = new Guacamole.Client(tunnel);

         // 处理客户端的状态变化事件
         this.client.onstatechange = this.onClientStateChange;
        let display = document.getElementById("display");

        // Add client to display div
        let element = this.client.getDisplay().getElement();
        display.appendChild(element);

        let token = ""

        let params = {
            'connectionId': connectionId,
            'X-Auth-Token': token
        };

        let paramStr = this.qs.stringify(params);

        // Connect
        this.client.connect(paramStr);

    }

    this.onTunnelStateChange = function(state){
        console.log('onTunnelStateChange', state);
        if (state === Guacamole.Tunnel.State.CLOSED) {
            this.loading = false
            this.showTips = '远程连接已关闭'
            this.bClose = true
        }
    };

    this.onClientStateChange = function(state){
        switch (state) {
            case STATE_IDLE:
                this.loading = true
                this.showTips = '正在初始化中...'
                break;
            case STATE_CONNECTING:
                this.loading = true
                this.showTips = '正在努力连接中...'
                break;
            case STATE_WAITING:
                this.loading = true
                this.showTips = '正在等待服务器响应...'
                break;
            case STATE_CONNECTED:
                this.loading = false
                this.showTips = ''
                break;
            case STATE_DISCONNECTING:
                break;
            case STATE_DISCONNECTED:
                break;
            default:
                break;
        }
    };
})