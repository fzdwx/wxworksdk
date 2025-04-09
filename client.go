package wxworksdk

import (
	"log/slog"
	"time"
)

type Client interface {
	VerifyCallback(echostr string, msgSignature, nonce, timestamp string) ([]byte, error)
	RefreshCfg(cfg *Config) error
	Decode(content []byte) (string, error)
	// UserListIDs 获取用户列表
	// 获取用户id列表
	UserListIDs(limit int) ([]*UserIDWithDept, error)
	// DeptList 获取部门列表
	DeptList() ([]*Department, error)
	// UserList 根据部门id获取用户列表
	UserList(deptID int) ([]*UserList, error)
}

type client struct {
	cfg         *Config
	accessToken string
	expireTime  time.Time
	log         *slog.Logger
}

type Config struct {
	CallbackToken             string
	CallbackEncodingAESKeyRaw string
	CorpID                    string
	AppSecret                 string
	AppID                     int64
	crypt                     *WXBizMsgCrypt
}

func (c *Config) Check() error {
	if c.CallbackToken == "" && c.CallbackEncodingAESKeyRaw == "" {
		return nil
	}
	crypt := NewWXBizMsgCrypt(c.CallbackToken, c.CallbackEncodingAESKeyRaw, "", XmlType)
	c.crypt = crypt
	return nil
}

func NewClient(
	cfg *Config,
	log *slog.Logger,
) Client {
	if cfg != nil {
		if err := cfg.Check(); err != nil {
			panic(err)
		}
	}

	c := &client{
		cfg: cfg,
		log: log,
	}

	if err := c.doRefreshToken(); err != nil {
		panic(err)
	}

	//go c.refreshToken()
	return c
}
