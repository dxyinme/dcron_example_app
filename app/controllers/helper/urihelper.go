package helper

import "github.com/gin-gonic/gin"

type UriHelper struct {
	mp map[string]string
}

func (uriHelper *UriHelper) Parse(c *gin.Context) (err error) {
	uriHelper.mp = make(map[string]string, 0)
	return c.ShouldBindUri(uriHelper.mp)
}

func (UriHelper *UriHelper) Get(key string) (value string, ok bool) {
	value, ok = UriHelper.mp[key]
	return
}

func GetUriHelperFromGinCtx(c *gin.Context) (uriHelper *UriHelper, err error) {
	uriHelper = &UriHelper{}
	err = uriHelper.Parse(c)
	return uriHelper, err
}
