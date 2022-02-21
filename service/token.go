package service

import (
	"encoding/json"
	"myproject/model"
)

func Token() string {
	rs , _ := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
		RefreshToken: "{refresh_token}",
		Token: "{token}",
	})
	return string(rs)
}