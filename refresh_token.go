package wxworksdk

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type RefreshTokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type PhoneToUserIDResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Userid  string `json:"userid"`
}

func (c *client) RefreshCfg(cfg *Config) error {
	if err := cfg.Check(); err != nil {
		return err
	}

	c.cfg = cfg
	return c.doRefreshToken()
}

func (c *client) refreshToken() {
	for {
		if c.cfg == nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if err := c.doRefreshToken(); err != nil {
			c.log.Error("refresh token", slog.Any("err", err))
		}

		for {
			if time.Now().After(c.expireTime) {
				break
			}

			time.Sleep(5 * time.Second)
		}
	}
}

func (c *client) doRefreshToken() error {
	if c.cfg == nil {
		return nil
	}

	resp, err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + c.cfg.CorpID + "&corpsecret=" + c.cfg.AppSecret)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var tokenResp RefreshTokenResp
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return err
	}

	if tokenResp.Errcode != 0 {
		return fmt.Errorf("get token err: %s", tokenResp.Errmsg)
	}

	c.accessToken = tokenResp.AccessToken
	c.expireTime = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	return nil
}
