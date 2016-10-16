package zimuzu

import (
	"encoding/json"
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"
)

var (
	loginPageURL = "http://" + domain + "/user/login/index"
	loginFormURL = "http://" + domain + "/User/Login/ajaxLogin"
	currUserURL  = "http://" + domain + "/user/login/getCurUserTopInfo"
)

func (c *Ctx) testLogin() (bool, error) {
	resp, err := c.client.Get(currUserURL)
	if err != nil {
		return false, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var ret struct {
		Status int
	}
	if err := json.Unmarshal(buf, &ret); err != nil {
		return false, err
	}

	return ret.Status == 1, nil
}

func (c *Ctx) Login() error {
	if ret, err := c.testLogin(); err == nil && ret {
		return nil
	}

	r, err := http.NewRequest("POST", loginFormURL, strings.NewReader(fmt.Sprintf("account=%s&password=%s&remember=1", c.username, c.password)))
	if err != nil {
		return err
	}
	r.Header.Set("Origin", "http://"+domain)
	r.Header.Set("Referer", loginPageURL)
	r.Header.Set("User-Agent", ua)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		//TODO: error
		return fmt.Errorf("login: http %d", resp.StatusCode)
	}

	var ret struct {
		Status int
		Info   string
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, &ret); err != nil {
		return err
	}
	fmt.Printf("Login: %+v\n", ret)

	if ret.Status != 1 {
		return fmt.Errorf("login: %s", ret.Info)
	}

	return nil
}
