package config

import (
	"github.com/herhe-com/framework/contracts/auth"
	"github.com/herhe-com/framework/facades"
)

func init() {

	cfg := facades.Cfg
	cfg.Add("auth", map[string]any{
		"casbin": map[string]any{
			"table": cfg.Env("auth.casbin.table", "sys_casbin"),
		},
		"platforms": []uint16{auth.CodeOfPlatform, auth.CodeOfStore},
		"permissions": []auth.Permission{
			shop(),
			hr(),
			site(),
		},
	})
}

func site() auth.Permission {
	return auth.Permission{
		Code: "site",
		Name: "站点设置",
		Children: []auth.Permission{
			{
				Code: "role",
				Name: "角色",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "user",
				Name: "账号",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "创建",
						Common: true,
					},
					{
						Code:   "update",
						Name:   "修改",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "enable",
						Name:   "启禁",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
			{
				Code: "secret",
				Name: "密钥",
				Children: []auth.Permission{
					{
						Code:   "create",
						Name:   "生成",
						Common: true,
					},
					{
						Code:   "delete",
						Name:   "删除",
						Common: true,
					},
					{
						Code:   "paginate",
						Name:   "列表",
						Common: true,
					},
				},
			},
		},
	}
}

func hr() auth.Permission {
	return auth.Permission{
		Code: "hr",
		Name: "人员架构",
		Children: []auth.Permission{
			{
				Code: "org",
				Name: "组织架构",
				Children: []auth.Permission{
					{
						Code: "organization",
						Name: "组织门店",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfRegion},
							},
							{
								Code:   "update",
								Name:   "修改",
								Common: true,
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfRegion},
							},
							{
								Code:      "enable",
								Name:      "启禁",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfRegion},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfRegion},
							},
							{
								Code:      "enter",
								Name:      "进入",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfClique, auth.CodeOfRegion},
							},
							{
								Code:      "store",
								Name:      "门店",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "brand",
						Name: "品牌",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfClique},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfClique},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfClique},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfClique},
							},
						},
					},
				},
			},
		},
	}
}

func shop() auth.Permission {

	return auth.Permission{
		Code: "shop",
		Name: "商城中心",
		Children: []auth.Permission{
			{
				Code: "order",
				Name: "订单管理",
				Children: []auth.Permission{
					{
						Code: "ordinary",
						Name: "产品订单",
						Children: []auth.Permission{
							{
								Code:      "shipment",
								Name:      "发货",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "refund",
								Name:      "退款",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfStore},
							},
						},
					},
					{
						Code: "service",
						Name: "售后订单",
						Children: []auth.Permission{
							{
								Code:      "handle",
								Name:      "处理",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfStore},
							},
						},
					},
					{
						Code: "appraisal",
						Name: "订单评价",
						Children: []auth.Permission{
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfPlatform, auth.CodeOfStore},
							},
						},
					},
				},
			},
			{
				Code: "commodity",
				Name: "商品信息",
				Children: []auth.Permission{
					{
						Code: "product",
						Name: "产品",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "enable",
								Name:      "上下架",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "specification",
								Name:      "规格",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "category",
						Name: "栏目",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "tree",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "specification",
						Name: "规格模板",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "enable",
								Name:      "启禁",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
				},
			},
			{
				Code: "member",
				Name: "客户管理",
				Children: []auth.Permission{
					{
						Code: "user",
						Name: "用户",
						Children: []auth.Permission{
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
				},
			},
			{
				Code: "pay",
				Name: "支付管理",
				Children: []auth.Permission{
					{
						Code: "channel",
						Name: "渠道配置",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "enable",
								Name:      "启禁",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
				},
			},
			{
				Code: "basic",
				Name: "基础配置",
				Children: []auth.Permission{
					{
						Code: "banner",
						Name: "轮播",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "advertise",
						Name: "广告",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "enable",
								Name:      "启禁",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "blog",
						Name: "博客",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "top",
								Name:      "置顶",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "page",
						Name: "页面",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "system",
								Name:      "内置",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "shipping",
						Name: "物流",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "enable",
								Name:      "启禁",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "seo",
						Name: "SEO",
						Children: []auth.Permission{
							{
								Code:      "create",
								Name:      "创建",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "delete",
								Name:      "删除",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "paginate",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
					{
						Code: "setting",
						Name: "设置",
						Children: []auth.Permission{
							{
								Code:      "update",
								Name:      "修改",
								Platforms: []uint16{auth.CodeOfStore},
							},
							{
								Code:      "list",
								Name:      "列表",
								Platforms: []uint16{auth.CodeOfStore},
							},
						},
					},
				},
			},
		},
	}
}
