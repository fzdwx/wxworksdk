package wxworksdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserSampleListResp struct {
	Errcode  int         `json:"errcode"`
	Errmsg   string      `json:"errmsg"`
	Userlist []*UserList `json:"userlist"`
}

type UserList struct {
	Userid         string   `json:"userid"`
	Name           string   `json:"name"`
	Department     []int    `json:"department"`
	Order          []int    `json:"order"`
	Position       string   `json:"position"`
	Mobile         string   `json:"mobile"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	BizMail        string   `json:"biz_mail"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	DirectLeader   []string `json:"direct_leader"`
	Avatar         string   `json:"avatar"`
	ThumbAvatar    string   `json:"thumb_avatar"`
	Telephone      string   `json:"telephone"`
	Alias          string   `json:"alias"`
	Status         int      `json:"status"`
	Address        string   `json:"address"`
	EnglishName    string   `json:"english_name"`
	OpenUserid     string   `json:"open_userid"`
	MainDepartment int      `json:"main_department"`
	QrCode         string   `json:"qr_code"`
}

func (c *client) UserList(deptID int) ([]*UserList, error) {
	resp, err := http.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/list?access_token=%s&department_id=%d", c.accessToken, deptID))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var userSampleListResp UserSampleListResp
	if err := json.NewDecoder(resp.Body).Decode(&userSampleListResp); err != nil {
		return nil, err
	}

	if userSampleListResp.Errcode != 0 {
		return nil, fmt.Errorf("error: %s", userSampleListResp.Errmsg)
	}

	return userSampleListResp.Userlist, nil
}
