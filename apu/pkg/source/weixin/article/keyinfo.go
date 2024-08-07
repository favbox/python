package article

import (
	"errors"
	"fmt"
	"html"
	"net/url"
	"strings"

	"apu/pkg/schema"
	"apu/pkg/source"
)

// KeyInfo 定义文章的关键信息。
type KeyInfo struct {
	Biz string
	Mid string
	Idx string
	Sn  string

	UID uint64
	Url string
}

// GetKeyInfo 获取指定文章网址的 KeyInfo。网址必须为经典格式。
func GetKeyInfo(canonicalURL string) (*KeyInfo, error) {
	canonicalURL = html.UnescapeString(canonicalURL)
	if strings.HasPrefix(canonicalURL, "http://") {
		canonicalURL = strings.Replace(canonicalURL, "http://", "https://", 1)
	}

	parsedURL, err := url.Parse(canonicalURL)
	if err != nil {
		return nil, err
	}

	if parsedURL.Hostname() != "mp.weixin.qq.com" {
		return nil, fmt.Errorf("不是常规的公众号域名")
	}
	if parsedURL.Path != "/s" {
		return nil, errors.New("不是常规的公众号文章网址路径，应为 /s")
	}

	query := parsedURL.Query()
	keyInfo := &KeyInfo{
		Biz: query.Get("__biz"),
		Mid: query.Get("mid"),
		Idx: query.Get("idx"),
		Sn:  query.Get("sn"),
		Url: canonicalURL,
	}
	if keyInfo.Biz == "" || keyInfo.Mid == "" || keyInfo.Idx == "" {
		return nil, errors.New("图文三元组不可为空")
	}

	// 生成唯一ID
	keyInfo.UID = source.UniqueID(fmt.Sprintf("%d|%s:%s:%s", schema.SourceWeixin, keyInfo.Biz, keyInfo.Mid, keyInfo.Idx))

	return keyInfo, nil
}
