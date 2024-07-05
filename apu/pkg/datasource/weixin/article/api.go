package article

import (
	"encoding/json"
	"errors"

	"apu/internal/cookieutil"
	"apu/pkg/store/mysql/query"
	"github.com/imroc/req/v3"
)

// GetArticles 利用微信读书 headers 获取公众号的文章列表。
func GetArticles(bookId string, count, offset, syncKey int) (articles []*BookArticle, nextSyncKey int, err error) {
	// 获取微信读书请求头
	//mysql.Init()
	weRequest, err := query.WeRequest.Where(
		query.WeRequest.Type.Eq("weread"),
		query.WeRequest.Status.Eq("valid"),
	).First()
	if err != nil {
		return
	}
	var headers map[string]string
	err = json.Unmarshal([]byte(weRequest.Headers), &headers)
	if err != nil {
		return
	}
	if len(headers) == 0 {
		err = errors.New("无可用请求头")
		return
	}

	request := req.R()

	// 设置查询参数
	request.SetQueryParamsAnyType(map[string]any{
		"bookId":  bookId,
		"count":   count,
		"offset":  offset,
		"synckey": syncKey,
	})

	// 设置请求头
	request.SetHeadersNonCanonical(headers)

	// 设置响应结果
	var result ArticlesResult
	request.SetSuccessResult(&result)

	// 发起请求
	resp, err := request.Get("https://i.weread.qq.com/book/articles")
	if err != nil {
		return
	}
	if resp.IsErrorState() {
		err = errors.New(resp.GetStatus())
		return
	}
	if result.Errmsg != "" {
		err = errors.New(result.Errmsg)
		return
	}

	nextSyncKey = result.SyncKey
	for _, review := range result.Reviews {
		articles = append(articles, review.Review.MpInfo)
	}

	return
}

// GetArticle 获取公开的公众号文章详情。
func GetArticle(biz, mid, idx, sn string) {
}

// GetStat 利用微信 cookie 获取文章统计信息。 https://www.cnblogs.com/jianpansangejian/p/17970546
func GetStat(biz, mid, idx, sn string) (*Stat, error) {
	// 获取微信阅读量请求 cookie
	//mysql.Init()
	weRequest, err := query.WeRequest.Where(
		query.WeRequest.Type.Eq("wechat"),
		query.WeRequest.Status.Eq("valid"),
	).First()
	if err != nil {
		return nil, err
	}

	request := req.R()

	// 设置查询参数
	cookieMap := cookieutil.StrToMap(weRequest.Cookie)
	request.SetQueryParams(map[string]string{
		"appmsg_token": cookieMap["appmsg_token"],
		"x5":           "0",
	})

	// 设置请求头
	request.SetHeaders(map[string]string{
		"User-Agent":   "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0Chrome/57.0.2987.132 MQQBrowser/6.2 Mobile",
		"Cookie":       weRequest.Cookie,
		"Origin":       "https://mp.weixin.qq.com",
		"Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
		"Host":         "mp.weixin.qq.com",
	})

	// 设置表单数据
	request.SetFormData(map[string]string{
		"is_only_read": "1",
		"is_temp_url":  "0",
		"appmsg_type":  "9",
		"__biz":        biz,
		"mid":          mid,
		"idx":          idx,
		"sn":           sn,
	})

	// 设置响应结果
	var result StatResult
	request.SetSuccessResult(&result)

	// 发起请求
	resp, err := request.Post("https://mp.weixin.qq.com/mp/getappmsgext")
	if err != nil {
		return nil, err
	}
	if resp.IsErrorState() {
		return nil, errors.New(resp.GetStatus())
	}
	if result.ArticleStat == nil {
		return nil, errors.New("会话已过期")
	}

	return result.ArticleStat, nil
}
