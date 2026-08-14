package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Statuspage",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.statuspage.io/v1",
			"auth": map[string]any{
				"prefix": "OAuth",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"component": map[string]any{},
				"component_group_uptime": map[string]any{},
				"group_component": map[string]any{},
				"incident": map[string]any{},
				"incident_postmortem": map[string]any{},
				"incident_subscriber": map[string]any{},
				"incident_template": map[string]any{},
				"incident_update": map[string]any{},
				"metric": map[string]any{},
				"metrics_provider": map[string]any{},
				"page": map[string]any{},
				"page_access_group": map[string]any{},
				"page_access_user": map[string]any{},
				"permission": map[string]any{},
				"postmortem": map[string]any{},
				"status_embed_config": map[string]any{},
				"subscriber": map[string]any{},
				"user": map[string]any{},
			},
		},
		"entity": map[string]any{
			"component": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "automation_email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "component",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "group",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "group_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "only_show_if_degraded",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "showcase",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "start_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_group",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_users",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_user",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_access_group_id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{page_access_group_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_group_id",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{page_access_user_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_user_id",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "end",
											"orig": "end",
											"type": "Any",
										},
										map[string]any{
											"kind": "query",
											"name": "start",
											"orig": "start",
											"type": "Any",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components/{component_id}/uptime",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"uptime",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "uptime",
									"exist": []any{
										"end",
										"id",
										"page_id",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.related_events`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_group",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_users",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "page_access_user",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"component_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"page_access_group",
						},
						[]any{
							"page",
							"page_access_user",
						},
					},
				},
			},
			"component_group_uptime": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incidents",
						"type": "`$OBJECT`",
					},
				},
				"name": "component_group_uptime",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "end",
											"orig": "end",
											"type": "Any",
										},
										map[string]any{
											"kind": "query",
											"name": "start",
											"orig": "start",
											"type": "Any",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups/{id}/uptime",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
									"uptime",
								},
								"select": map[string]any{
									"exist": []any{
										"end",
										"id",
										"page_id",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.related_events`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"group_component": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_group",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "components",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "group_component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/component-groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "auto_transition_deliver_notifications_at_end",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "auto_transition_deliver_notifications_at_start",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "auto_transition_to_maintenance_state",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "auto_transition_to_operational_state",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "components",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "impact",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "impact_override",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident",
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$OBJECT`",
							},
							"update": map[string]any{
								"type": "`$OBJECT`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "incident_updates",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "metadata",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "monitoring_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postmortem_body",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postmortem_body_last_updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postmortem_ignored",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem_notified_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem_notified_twitter",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem_published_at",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "reminder_intervals",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "resolved_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_auto_completed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "scheduled_auto_in_progress",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "scheduled_for",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_remind_prior",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "scheduled_reminded_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled_until",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shortlink",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "incident",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"page",
										"page_id",
										"q",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/active_maintenance",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"active_maintenance",
								},
								"select": map[string]any{
									"$action": "active_maintenance",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/scheduled",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"scheduled",
								},
								"select": map[string]any{
									"$action": "scheduled",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/unresolved",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"unresolved",
								},
								"select": map[string]any{
									"$action": "unresolved",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/upcoming",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"upcoming",
								},
								"select": map[string]any{
									"$action": "upcoming",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident_postmortem": map[string]any{
				"fields": []any{},
				"name": "incident_postmortem",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{id}",
									"postmortem",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident_subscriber": map[string]any{
				"fields": []any{},
				"name": "incident_subscriber",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "subscriber_id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{subscriber_id}",
									"resend_confirmation",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
										"subscriber_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
							"incident",
							"subscriber",
						},
					},
				},
			},
			"incident_template": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "body",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "components",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "group_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "should_send_notifications",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "should_tweet",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "update_status",
						"type": "`$STRING`",
					},
				},
				"name": "incident_template",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incident_templates",
								"parts": []any{
									"pages",
									"{page_id}",
									"incident_templates",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incident_templates",
								"parts": []any{
									"pages",
									"{page_id}",
									"incident_templates",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"incident_update": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "affected_components",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "body",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "custom_tweet",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "deliver_notifications",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "display_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident_update",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tweet_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "wants_twitter_update",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "incident_update",
				"op": map[string]any{
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_update_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"incident_updates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_update_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident_update": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "incident_update_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"incident_updates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"incident_update_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"incident_update": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
							"incident",
						},
					},
				},
			},
			"metric": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "backfill_percentage",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "backfilled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "data",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "decimal_places",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "display",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_fetched_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "metric_identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metrics_provider_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "most_recent_data_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "suffix",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tooltip_description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "y_axis_hidden",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "y_axis_max",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "y_axis_min",
						"type": "`$NUMBER`",
					},
				},
				"name": "metric",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics/{metric_id}/data",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
									"data",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{metrics_provider_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"metrics_provider_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics/data",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"data",
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{page_access_user_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_user_id",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{metrics_provider_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"metrics_provider_id",
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics/{metric_id}/data",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
									"data",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metric_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"metrics_provider",
						},
						[]any{
							"page",
							"page_access_user",
						},
					},
				},
			},
			"metrics_provider": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "disabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_revalidated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric_base_uri",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metrics_provider",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "metrics_provider",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics_providers",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"metrics_provider_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "activity_score",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "allow_email_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_incident_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_page_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_rss_atom_feeds",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_sms_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "allow_webhook_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "branding",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_blues",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_body_background_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_border_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_font_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_graph_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_greens",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_light_font_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_link_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_no_data",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_oranges",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_reds",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "css_yellows",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "domain",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "favicon_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "headline",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hero_cover",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hidden_from_search",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ip_restrictions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notifications_email_footer",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notifications_from_email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "state",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subdomain",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "support_url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "time_zone",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactional_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_username",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "viewers_must_be_team_members",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/pages",
								"parts": []any{
									"pages",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}",
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}",
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}",
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"page_access_group": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "external_identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_group",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_access_user_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "page_access_group",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_groups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "component_id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
									"{component_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"component_id",
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_group_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"component",
						},
					},
				},
			},
			"page_access_user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "external_login",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metric_ids",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "page_access_group_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_group_ids",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_user",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "page_access_user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_user": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "email",
											"orig": "email",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users",
								"parts": []any{
									"pages",
									"{id}",
									"page_access_users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email",
										"id",
										"page",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "component_id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
									"{component_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"component_id",
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "metric_id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
									"{metric_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"metric_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"components",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{id}",
									"metrics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_access_user_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"component",
						},
						[]any{
							"page",
							"metric",
						},
					},
				},
			},
			"permission": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "pages",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "user_id",
						"type": "`$STRING`",
					},
				},
				"name": "permission",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/organizations/{organization_id}/permissions/{user_id}",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"permissions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/organizations/{organization_id}/permissions/{user_id}",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"permissions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"organization",
						},
					},
				},
			},
			"postmortem": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "body",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "body_draft",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "body_draft_updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "body_updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "custom_tweet",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notify_subscribers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "notify_twitter",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "postmortem",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$OBJECT`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "preview_key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "published_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
				},
				"name": "postmortem",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"postmortem": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/publish",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
									"publish",
								},
								"select": map[string]any{
									"$action": "publish",
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"postmortem": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/revert",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
									"revert",
								},
								"select": map[string]any{
									"$action": "revert",
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
							"incident",
						},
					},
				},
			},
			"status_embed_config": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "incident_background_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incident_text_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maintenance_background_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maintenance_text_color",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status_embed_config",
						"type": "`$OBJECT`",
					},
				},
				"name": "status_embed_config",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/status_embed_config",
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/status_embed_config",
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"status_embed_config": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/status_embed_config",
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"status_embed_config": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
					},
				},
			},
			"subscriber": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "components",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "display_phone_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "endpoint",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "integration_partner",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "mode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "obfuscated_channel_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_user_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_country",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "purge_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quarantined_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "skip_confirmation_notification",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "skip_unsubscription_notification",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "slack",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sms",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "state",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subscriber",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "subscribers",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "teams",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "webhook",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "workspace_name",
						"type": "`$STRING`",
					},
				},
				"name": "subscriber",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
									"resend_confirmation",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"subscriber": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"subscriber": "`reqdata`",
									},
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/reactivate",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"reactivate",
								},
								"select": map[string]any{
									"$action": "reactivate",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/resend_confirmation",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"resend_confirmation",
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/unsubscribe",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"unsubscribe",
								},
								"select": map[string]any{
									"$action": "unsubscribe",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 0,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "asc",
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "primary",
											"kind": "query",
											"name": "sort_field",
											"orig": "sort_field",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "active",
											"kind": "query",
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"page",
										"page_id",
										"q",
										"sort_direction",
										"sort_field",
										"state",
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/unsubscribed",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"unsubscribed",
								},
								"select": map[string]any{
									"$action": "unsubscribed",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": "active",
											"kind": "query",
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/count",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"count",
								},
								"select": map[string]any{
									"$action": "count",
									"exist": []any{
										"page_id",
										"state",
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/histogram_by_state",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"histogram_by_state",
								},
								"select": map[string]any{
									"$action": "histogram_by_state",
									"exist": []any{
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "skip_unsubscription_notification",
											"orig": "skip_unsubscription_notification",
											"type": "`$BOOLEAN`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
										"skip_unsubscription_notification",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"page",
						},
						[]any{
							"page",
							"incident",
						},
					},
				},
			},
			"user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "first_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "last_name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "organization_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "user",
						"req": true,
						"type": "`$OBJECT`",
					},
				},
				"name": "user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/organizations/{organization_id}/users",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
								},
								"select": map[string]any{
									"exist": []any{
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"user": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/organizations/{organization_id}/users",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
								},
								"select": map[string]any{
									"exist": []any{
										"organization_id",
										"page",
										"per_page",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/organizations/{organization_id}/users/{user_id}",
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"user_id": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"organization",
						},
					},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
