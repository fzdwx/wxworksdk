package wxworksdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserIDWithDept struct {
	Userid     string `json:"userid"`
	Department int    `json:"department"`
}

type UserListIDsResp struct {
	Errcode    int              `json:"errcode"`
	Errmsg     string           `json:"errmsg"`
	NextCursor string           `json:"next_cursor"`
	DeptUser   []UserIDWithDept `json:"dept_user"`
}

type userListIDsReq struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

func (c *client) UserListIDs() (*UserListIDsResp, error) {

	req := userListIDsReq{
		Cursor: "",
		Limit:  1000,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/user/list_id?access_token="+c.accessToken,
		"application/json",
		strings.NewReader(string(reqBody)),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error: %s", resp.Status)
	}

	var userListIDsResp UserListIDsResp
	jsonDecoder := json.NewDecoder(resp.Body)
	if err := jsonDecoder.Decode(&userListIDsResp); err != nil {
		return nil, err
	}
	if userListIDsResp.Errcode != 0 {
		return nil, fmt.Errorf("get user list ids error: %s", userListIDsResp.Errmsg)
	}
	return &userListIDsResp, nil
}
