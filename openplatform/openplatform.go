package openplatform

import (
	"github.com/silenceper/wechat/v2/openplatform/config"
	"github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram"
	"github.com/silenceper/wechat/v2/openplatform/officialaccount"
)

//OpenPlatform 微信开放平台相关api
type OpenPlatform struct {
	*context.Context
}

//NewOpenPlatform new openplatform
func NewOpenPlatform(cfg *config.Config) *OpenPlatform {
	if cfg.Cache == nil {
		panic("cache 未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	return &OpenPlatform{ctx}
}

//GetOfficialAccount 公众号代处理
func (openPlatform *OpenPlatform) GetOfficialAccount(appID string) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(openPlatform.Context, appID)
}

//GetMiniProgram 小程序代理
func (openPlatform *OpenPlatform) GetMiniProgram(opCtx *context.Context, appID string) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(opCtx, appID)
}
