package yapi

import (
	"github.com/jinzhu/copier"
	"net/http"
)

type CatMenuData []CatData

type CatData struct {
	ID   int    `json:"_id" structs:"_id"`
	UID  int    `json:"uid" structs:"uid"`
	Name string `json:"name" structs:"name"`
	Desc string `json:"desc" structs:"desc"`
}

type CatMenu struct {
	ErrCode int         `json:"errcode" structs:"errcode"`
	ErrMsg  string      `json:"errmsg" structs:"errmsg"`
	Data    CatMenuData `json:"data" structs:"data"`
}

// CatMenuService .
type CatMenuService struct {
	client *Client
}

type CatMenuParam struct {
	Token     string `url:"token"`
	ProjectID int    `url:"project_id"`
}

type ModifyMenuParam struct {
	ProjectID int    `json:"project_id" url:"project_id"`
	CatData
}

type ModifyMenumReq struct {
	Token     string `json:"token"`
	ModifyMenuParam
}

type ModifyMenuResp struct {
	CommonResp
	Data interface{} `json:"data" structs:"data"`
}

func (s *CatMenuService) Get(projectId int) (*CatMenu, *http.Response, error) {
	apiEndpoint := "api/interface/getCatMenu"
	catMenuParam := new(CatMenuParam)
	catMenuParam.ProjectID = projectId
	catMenuParam.Token = s.client.Authentication.token
	url, err := addOptions(apiEndpoint, catMenuParam)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(CatMenu)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, NewServerError(resp, err)
	}
	return result, resp, err
}

func (s *CatMenuService) AddOrUpdate(param *ModifyMenuParam) (*ModifyMenuResp, *http.Response, error) {
	apiEndpoint := "api/interface/add_cat"
	modifyMenuReq := new(ModifyMenumReq)
	modifyMenuReq.Token = s.client.Authentication.token

	copier.Copy(&modifyMenuReq, param)

	req, err := s.client.NewRequest("POST", apiEndpoint, modifyMenuReq)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	result := new(ModifyMenuResp)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, NewServerError(resp, err)
	}
	return result, resp, err
}
