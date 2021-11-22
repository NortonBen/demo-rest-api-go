package model

import "apm/pkg/util"

type Permission struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var ListPermission = []Permission{
	{
		Id:   "login",
		Name: "Login System",
	},
	{
		Id:   "own_system",
		Name: "Admin System",
	},
	{
		Id:   "read_user",
		Name: "View User",
	},
	{
		Id:   "writer_user",
		Name: "Write User",
	},
	{
		Id:   "read_project",
		Name: "View Project",
	},
}

func FilterPermission(list []string) []string {
	rs := make([]string, 0)

	for _, item := range ListPermission {
		if util.StringInSlice(item.Id, list) {
			rs = append(rs, item.Id)
		}
	}

	return rs
}
