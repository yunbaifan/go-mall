// Code generated by goctl. DO NOT EDIT.
package types

type AdminInfo struct {
	Roles []string `json:"roles" default:"[]" description:"角色列表"`
	Icon  string   `json:"icon" default:"" description:"头像"`
	Menus []Menus  `json:"menus" default:"[]" description:"菜单列表"`
}

type Menus struct {
	Id         int    `json:"id" description:"菜单ID"`
	ParentId   int    `json:"parentId" description:"父级ID"`
	CreateTime string `json:"createTime" description:"创建时间"`
	Level      int    `json:"level" description:"层级"`
	Name       string `json:"name" description:"菜单名称"`
	Sort       int    `json:"sort" description:"排序"`
	Icon       string `json:"icon" description:"图标"`
	Hidden     bool   `json:"hidden" description:"是否隐藏"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}
