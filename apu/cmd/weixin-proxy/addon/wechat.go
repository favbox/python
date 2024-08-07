package addon

import (
	"bytes"
	"log"

	"apu/pkg/store/mysql"
	"apu/pkg/store/mysql/model"
	"apu/pkg/store/mysql/query"
	"apu/pkg/util/cookiex"
	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

// WechatAddon 微信代理插件。
// 用于拦截微信APP[PC|Mobile]包含 appmsg_token 的 cookie。
type WechatAddon struct {
	proxy.BaseAddon
}

func (a *WechatAddon) Response(f *proxy.Flow) {
	// 该插件只拦截文章详情页
	if f.Request.URL.Hostname() != "mp.weixin.qq.com" ||
		f.Request.URL.Path != "/s" {
		return
	}

	f.Response.ReplaceToDecodedBody()

	// 注入定时刷新脚本 TODO 未生效
	//go injectRefreshJavascript(f)

	// 保存请求头（包含了 可请求阅读量的 appmsg_token 令牌）
	cookie := f.Request.Header.Get("cookie")
	cookieMap := cookiex.StrToMap(cookie)
	if _, exists := cookieMap["appmsg_token"]; exists {
		mysql.Init()
		wxuin := cookieMap["wxuin"]
		err := query.WeixinRequest.
			Save(&model.WeixinRequest{
				Type:   "wechat",
				UserID: wxuin,
				Cookie: cookie,
				Status: "valid",
			})
		if err != nil {
			log.Println(err)
		}
		log.Println("已捕获新的【微信】可用会话", wxuin)
	}
}

func injectRefreshJavascript(f *proxy.Flow) {
	const reloadWindowJavascript = `
	<script type="text/javascript">
		setInterval(() => {
			window.location.href=window.location.href;
		}, 3000);
	</script>
</body>
`

	body := f.Response.Body
	newBody := bytes.Replace(body, []byte("</body>"), []byte(reloadWindowJavascript), 1)
	f.Response.Body = newBody

	log.Println("刷新窗口的脚本已注入")
}
