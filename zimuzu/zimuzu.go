package zimuzu

import (
	"net/http"
	"net/http/cookiejar"
)

var (
	domain = "www.zimuzu.tv"
	ua     = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36`
)

type Ctx struct {
	username string
	password string
	client   *http.Client
}

func New(username, password string) *Ctx {
	ctx := &Ctx{
		username: username,
		password: password,
	}

	cookie, _ := cookiejar.New(nil)

	ctx.client = &http.Client{
		Jar: cookie,
	}

	return ctx
}
