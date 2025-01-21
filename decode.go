package wxworksdk

func (c *client) Decode(content []byte) (string, error) {
	msg, cryptError := c.cfg.crypt.DecryptMsg(content)
	if cryptError != nil {
		return "", cryptError
	}
	return string(msg), nil
}
