Tea.context(function () {

        this.onOpenRenew = function () {

            teaweb.popup(Tea.url("/settings/authority/expire/renew"),
                {
                    height: '18em',
                    // width: '50em',
                    callback: function () {
                        teaweb.success("续订成功", function () {
                            teaweb.reload()
                        })
                    }
                }
            );
        }
    }
)