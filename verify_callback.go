package wxworksdk

type WeComRecvMsg struct {
	ToUserName string `xml:"ToUserName" json:"ToUserName"`
	Encrypt    string `xml:"Encrypt" json:"Encrypt"`
	AgentID    string `xml:"AgentID" json:"AgentID"`
}

func (c *client) VerifyCallback(echostr string, msgSignature, nonce, timestamp string) ([]byte, error) {
	url, err := c.cfg.crypt.VerifyURL(msgSignature, timestamp, nonce, echostr)
	if err != nil {
		return nil, err
	}

	return url, nil
}
