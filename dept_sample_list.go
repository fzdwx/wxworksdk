package wxworksdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DeptSampleListResp struct {
	Errcode    int           `json:"errcode"`
	Errmsg     string        `json:"errmsg"`
	Department []*Department `json:"department"`
}
type Department struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	Parentid         int      `json:"parentid"`
	Order            int      `json:"order"`
}

func (d *Department) String() string {
	return fmt.Sprintf("{id: %d, name: %s, leader: %s, parentid: %d, order: %d}", d.Id, d.Name, strings.Join(d.DepartmentLeader, " "), d.Parentid, d.Order)
}

func (c *client) DeptList() ([]*Department, error) {
	resp, err := http.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s", c.accessToken))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var deptSampleListResp DeptSampleListResp
	if err := json.NewDecoder(resp.Body).Decode(&deptSampleListResp); err != nil {
		return nil, err
	}
	if deptSampleListResp.Errcode != 0 {
		return nil, fmt.Errorf("error: %s", deptSampleListResp.Errmsg)
	}

	return deptSampleListResp.Department, nil

}
