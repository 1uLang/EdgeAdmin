package helpers

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/TeaOSLab/EdgeAdmin/internal/configloaders"
	teaconst "github.com/TeaOSLab/EdgeAdmin/internal/const"
	"github.com/TeaOSLab/EdgeAdmin/internal/utils/numberutils"
	"github.com/iwind/TeaGo/actions"
	"net"
	"net/http"
)

type UserShouldAuth struct {
	action *actions.ActionObject
}

func (this *UserShouldAuth) BeforeAction(actionPtr actions.ActionWrapper, paramName string) (goNext bool) {
	if teaconst.IsRecoverMode {
		actionPtr.Object().RedirectURL("/recover")
		return false
	}

	this.action = actionPtr.Object()

	// 安全相关
	action := this.action
	securityConfig, _ := configloaders.LoadSecurityConfig()
	if securityConfig == nil {
		action.AddHeader("X-Frame-Options", "SAMEORIGIN")
	} else if len(securityConfig.Frame) > 0 {
		action.AddHeader("X-Frame-Options", securityConfig.Frame)
	}
	action.AddHeader("Content-Security-Policy", "default-src 'self' data:; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'")

	// 检查IP
	if !checkIP(securityConfig, action.RequestRemoteIP()) {
		action.ResponseWriter.WriteHeader(http.StatusForbidden)
		return false
	}
	remoteAddr, _, _ := net.SplitHostPort(action.Request.RemoteAddr)
	if len(remoteAddr) > 0 && remoteAddr != action.RequestRemoteIP() && !checkIP(securityConfig, remoteAddr) {
		action.ResponseWriter.WriteHeader(http.StatusForbidden)
		return false
	}

	return true
}

// StoreAdmin 存储用户名到SESSION
func (this *UserShouldAuth) StoreAdmin(adminId int64, remember bool) {
	// 修改sid的时间
	if remember {
		cookie := &http.Cookie{
			Name:     teaconst.CookieSID,
			Value:    this.action.Session().Sid,
			Path:     "/",
			MaxAge:   14 * 86400,
			HttpOnly: true,
		}
		if this.action.Request.TLS != nil {
			cookie.SameSite = http.SameSiteStrictMode
			cookie.Secure = true
		}
		this.action.AddCookie(cookie)
	} else {
		cookie := &http.Cookie{
			Name:     teaconst.CookieSID,
			Value:    this.action.Session().Sid,
			Path:     "/",
			MaxAge:   0,
			HttpOnly: true,
		}
		if this.action.Request.TLS != nil {
			cookie.SameSite = http.SameSiteStrictMode
			cookie.Secure = true
		}
		this.action.AddCookie(cookie)
	}
	this.action.Session().Write("adminId", numberutils.FormatInt64(adminId))
}

func (this *UserShouldAuth) IsUser() bool {
	return this.action.Session().GetInt("adminId") > 0
}

func (this *UserShouldAuth) AdminId() int {
	return this.action.Session().GetInt("adminId")
}

func (this *UserShouldAuth) Logout() {
	cache.DelKey(fmt.Sprintf("login_success_adminid_%v", this.AdminId()))
	this.action.Session().Delete()
}

func (this *UserShouldAuth) SetUpdatePwdToken(value int64) {
	this.action.Session().Write("update_pwd_admin_id", numberutils.FormatInt64(value))
}

func (this *UserShouldAuth) GetUpdatePwdToken() int64 {
	return this.action.Session().GetInt64("update_pwd_admin_id")
}
