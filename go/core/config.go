package core

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
				"prefix": "Bearer",
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
						"active": true,
						"name": "automation_email",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "component",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "group",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "group_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "major_outage",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "only_show_if_degraded",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "partial_outage",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "position",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "range_end",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "range_start",
						"req": false,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "related_event",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "showcase",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "start_date",
						"req": false,
						"type": "`$STRING`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": false,
						"type": "`$STRING`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "uptime_percentage",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "warning",
						"req": false,
						"type": "`$STRING`",
						"index$": 21,
					},
				},
				"name": "component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 2,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_access_group_id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "end",
											"orig": "end",
											"reqd": false,
											"type": "Any",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "skip_related_event",
											"orig": "skip_related_event",
											"reqd": false,
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "start",
											"orig": "start",
											"reqd": false,
											"type": "Any",
										},
									},
								},
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
										"skip_related_event",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "major_outage",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "partial_outage",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "range_end",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "range_start",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "related_event",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "uptime_percentage",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "warning",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
				},
				"name": "component_group_uptime",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "end",
											"orig": "end",
											"reqd": false,
											"type": "Any",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "skip_related_event",
											"orig": "skip_related_event",
											"reqd": false,
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "start",
											"orig": "start",
											"reqd": false,
											"type": "Any",
										},
									},
								},
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
										"skip_related_event",
										"start",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
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
						"active": true,
						"name": "component",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "component_group",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "description",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "position",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
				},
				"name": "group_component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "auto_transition_deliver_notifications_at_end",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "auto_transition_deliver_notifications_at_start",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "auto_transition_to_maintenance_state",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "auto_transition_to_operational_state",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "component",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "impact",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "impact_override",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "incident",
						"op": map[string]any{
							"patch": map[string]any{
								"req": false,
								"type": "`$OBJECT`",
							},
							"update": map[string]any{
								"req": false,
								"type": "`$OBJECT`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "incident_impact",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "incident_update",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "metadata",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "monitoring_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "postmortem_body",
						"req": false,
						"type": "`$STRING`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "postmortem_body_last_updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "postmortem_ignored",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "postmortem_notified_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "postmortem_notified_twitter",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "postmortem_published_at",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "reminder_interval",
						"req": false,
						"type": "`$STRING`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "resolved_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "scheduled_auto_completed",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 24,
					},
					map[string]any{
						"active": true,
						"name": "scheduled_auto_in_progress",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 25,
					},
					map[string]any{
						"active": true,
						"name": "scheduled_for",
						"req": false,
						"type": "`$STRING`",
						"index$": 26,
					},
					map[string]any{
						"active": true,
						"name": "scheduled_remind_prior",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 27,
					},
					map[string]any{
						"active": true,
						"name": "scheduled_reminded_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 28,
					},
					map[string]any{
						"active": true,
						"name": "scheduled_until",
						"req": false,
						"type": "`$STRING`",
						"index$": 29,
					},
					map[string]any{
						"active": true,
						"name": "shortlink",
						"req": false,
						"type": "`$STRING`",
						"index$": 30,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": false,
						"type": "`$STRING`",
						"index$": 31,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 32,
					},
				},
				"name": "incident",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 3,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 4,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "remove",
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
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "subscriber_id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "create",
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
						"active": true,
						"name": "body",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "component",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "group_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "should_send_notification",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "should_tweet",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "template",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "title",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "update_status",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
				},
				"name": "incident_template",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 100,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
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
						"active": true,
						"name": "affected_component",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "body",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "custom_tweet",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "deliver_notification",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "display_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "incident_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "incident_update",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "status",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "tweet_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "twitter_updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "wants_twitter_update",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 13,
					},
				},
				"name": "incident_update",
				"op": map[string]any{
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_update_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "incident_update_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "backfill_percentage",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "backfilled",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "data",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "decimal_place",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "display",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "last_fetched_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "metric",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "metric_identifier",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "metrics_provider_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "most_recent_data_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "reference_name",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "suffix",
						"req": false,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "tooltip_description",
						"req": false,
						"type": "`$STRING`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "y_axis_hidden",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "y_axis_max",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "y_axis_min",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 19,
					},
				},
				"name": "metric",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "disabled",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "last_revalidated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "metric_base_uri",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "metrics_provider",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "type",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
				},
				"name": "metrics_provider",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "metrics_provider_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "activity_score",
						"req": false,
						"type": "`$NUMBER`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "allow_email_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "allow_incident_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "allow_page_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "allow_rss_atom_feed",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "allow_sms_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "allow_webhook_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "branding",
						"req": false,
						"type": "`$STRING`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "city",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "country",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "css_blue",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "css_body_background_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "css_border_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "css_font_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "css_graph_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "css_green",
						"req": false,
						"type": "`$STRING`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "css_light_font_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "css_link_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "css_no_data",
						"req": false,
						"type": "`$STRING`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "css_orange",
						"req": false,
						"type": "`$STRING`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "css_red",
						"req": false,
						"type": "`$STRING`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "css_yellow",
						"req": false,
						"type": "`$STRING`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "domain",
						"req": false,
						"type": "`$STRING`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "email_logo",
						"req": false,
						"type": "`$STRING`",
						"index$": 24,
					},
					map[string]any{
						"active": true,
						"name": "favicon_logo",
						"req": false,
						"type": "`$STRING`",
						"index$": 25,
					},
					map[string]any{
						"active": true,
						"name": "headline",
						"req": false,
						"type": "`$STRING`",
						"index$": 26,
					},
					map[string]any{
						"active": true,
						"name": "hero_cover",
						"req": false,
						"type": "`$STRING`",
						"index$": 27,
					},
					map[string]any{
						"active": true,
						"name": "hidden_from_search",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 28,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 29,
					},
					map[string]any{
						"active": true,
						"name": "ip_restriction",
						"req": false,
						"type": "`$STRING`",
						"index$": 30,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 31,
					},
					map[string]any{
						"active": true,
						"name": "notifications_email_footer",
						"req": false,
						"type": "`$STRING`",
						"index$": 32,
					},
					map[string]any{
						"active": true,
						"name": "notifications_from_email",
						"req": false,
						"type": "`$STRING`",
						"index$": 33,
					},
					map[string]any{
						"active": true,
						"name": "page",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 34,
					},
					map[string]any{
						"active": true,
						"name": "page_description",
						"req": false,
						"type": "`$STRING`",
						"index$": 35,
					},
					map[string]any{
						"active": true,
						"name": "state",
						"req": false,
						"type": "`$STRING`",
						"index$": 36,
					},
					map[string]any{
						"active": true,
						"name": "subdomain",
						"req": false,
						"type": "`$STRING`",
						"index$": 37,
					},
					map[string]any{
						"active": true,
						"name": "support_url",
						"req": false,
						"type": "`$STRING`",
						"index$": 38,
					},
					map[string]any{
						"active": true,
						"name": "time_zone",
						"req": false,
						"type": "`$STRING`",
						"index$": 39,
					},
					map[string]any{
						"active": true,
						"name": "transactional_logo",
						"req": false,
						"type": "`$STRING`",
						"index$": 40,
					},
					map[string]any{
						"active": true,
						"name": "twitter_logo",
						"req": false,
						"type": "`$STRING`",
						"index$": 41,
					},
					map[string]any{
						"active": true,
						"name": "twitter_username",
						"req": false,
						"type": "`$STRING`",
						"index$": 42,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 43,
					},
					map[string]any{
						"active": true,
						"name": "url",
						"req": false,
						"type": "`$STRING`",
						"index$": 44,
					},
					map[string]any{
						"active": true,
						"name": "viewers_must_be_team_member",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 45,
					},
				},
				"name": "page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"page_access_group": map[string]any{
				"fields": []any{
					map[string]any{
						"active": true,
						"name": "component_id",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"req": false,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "external_identifier",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "metric_id",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "name",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "page_access_group",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "page_access_user_id",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
				},
				"name": "page_access_group",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 1,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "component_id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_group_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "component_id",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "external_login",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "metric_id",
						"req": true,
						"type": "`$ARRAY`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "page_access_group_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "page_access_user",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
				},
				"name": "page_access_user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 2,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "email",
											"orig": "email",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "patch",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "component_id",
											"orig": "component_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "metric_id",
											"orig": "metric_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 3,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 4,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "page_access_user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "data",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "page",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 1,
					},
				},
				"name": "permission",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "body",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "body_draft",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "body_draft_updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "body_updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "custom_tweet",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "notify_subscriber",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "notify_twitter",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "postmortem",
						"op": map[string]any{
							"update": map[string]any{
								"req": false,
								"type": "`$OBJECT`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "preview_key",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "published_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
				},
				"name": "postmortem",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "incident_background_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "incident_text_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "maintenance_background_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "maintenance_text_color",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "page_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "position",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "status_embed_config",
						"req": false,
						"type": "`$OBJECT`",
						"index$": 6,
					},
				},
				"name": "status_embed_config",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "load",
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "patch",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "component",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "component_id",
						"req": false,
						"type": "`$ARRAY`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "display_phone_number",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "endpoint",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "integration_partner",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 7,
					},
					map[string]any{
						"active": true,
						"name": "mode",
						"req": false,
						"type": "`$STRING`",
						"index$": 8,
					},
					map[string]any{
						"active": true,
						"name": "obfuscated_channel_name",
						"req": false,
						"type": "`$STRING`",
						"index$": 9,
					},
					map[string]any{
						"active": true,
						"name": "page_access_user_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 10,
					},
					map[string]any{
						"active": true,
						"name": "phone_country",
						"req": false,
						"type": "`$STRING`",
						"index$": 11,
					},
					map[string]any{
						"active": true,
						"name": "phone_number",
						"req": false,
						"type": "`$STRING`",
						"index$": 12,
					},
					map[string]any{
						"active": true,
						"name": "purge_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 13,
					},
					map[string]any{
						"active": true,
						"name": "quarantined_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 14,
					},
					map[string]any{
						"active": true,
						"name": "skip_confirmation_notification",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 15,
					},
					map[string]any{
						"active": true,
						"name": "skip_unsubscription_notification",
						"req": false,
						"type": "`$BOOLEAN`",
						"index$": 16,
					},
					map[string]any{
						"active": true,
						"name": "slack",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 17,
					},
					map[string]any{
						"active": true,
						"name": "sms",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 18,
					},
					map[string]any{
						"active": true,
						"name": "state",
						"req": false,
						"type": "`$STRING`",
						"index$": 19,
					},
					map[string]any{
						"active": true,
						"name": "subscriber",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"req": false,
						"type": "`$OBJECT`",
						"index$": 20,
					},
					map[string]any{
						"active": true,
						"name": "team",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 21,
					},
					map[string]any{
						"active": true,
						"name": "type",
						"req": false,
						"type": "`$STRING`",
						"index$": 22,
					},
					map[string]any{
						"active": true,
						"name": "webhook",
						"req": false,
						"type": "`$INTEGER`",
						"index$": 23,
					},
					map[string]any{
						"active": true,
						"name": "workspace_name",
						"req": false,
						"type": "`$STRING`",
						"index$": 24,
					},
				},
				"name": "subscriber",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 3,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 4,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 5,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"example": 0,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "q",
											"orig": "q",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "asc",
											"kind": "query",
											"name": "sort_direction",
											"orig": "sort_direction",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "primary",
											"kind": "query",
											"name": "sort_field",
											"orig": "sort_field",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"example": "active",
											"kind": "query",
											"name": "state",
											"orig": "state",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "type",
											"orig": "type",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 2,
							},
						},
						"key$": "list",
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"example": "active",
											"kind": "query",
											"name": "state",
											"orig": "state",
											"reqd": false,
											"type": "`$STRING`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "type",
											"orig": "type",
											"reqd": false,
											"type": "`$STRING`",
										},
									},
								},
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
								"index$": 1,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 2,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
								"index$": 3,
							},
						},
						"key$": "load",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "incident_id",
											"orig": "incident_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 2,
										},
									},
								},
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
								"index$": 0,
							},
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "skip_unsubscription_notification",
											"orig": "skip_unsubscription_notification",
											"reqd": false,
											"type": "`$BOOLEAN`",
										},
									},
								},
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
								"index$": 1,
							},
						},
						"key$": "remove",
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "subscriber_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "page_id",
											"orig": "page_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "update",
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
						"active": true,
						"name": "created_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 0,
					},
					map[string]any{
						"active": true,
						"name": "email",
						"req": false,
						"type": "`$STRING`",
						"index$": 1,
					},
					map[string]any{
						"active": true,
						"name": "first_name",
						"req": false,
						"type": "`$STRING`",
						"index$": 2,
					},
					map[string]any{
						"active": true,
						"name": "id",
						"req": false,
						"type": "`$STRING`",
						"index$": 3,
					},
					map[string]any{
						"active": true,
						"name": "last_name",
						"req": false,
						"type": "`$STRING`",
						"index$": 4,
					},
					map[string]any{
						"active": true,
						"name": "organization_id",
						"req": false,
						"type": "`$STRING`",
						"index$": 5,
					},
					map[string]any{
						"active": true,
						"name": "updated_at",
						"req": false,
						"type": "`$STRING`",
						"index$": 6,
					},
					map[string]any{
						"active": true,
						"name": "user",
						"req": true,
						"type": "`$OBJECT`",
						"index$": 7,
					},
				},
				"name": "user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
								},
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
									"req": "`reqdata`",
									"res": "`body`",
								},
								"index$": 0,
							},
						},
						"key$": "create",
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
									},
									"query": []any{
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"active": true,
											"kind": "query",
											"name": "per_page",
											"orig": "per_page",
											"reqd": false,
											"type": "`$INTEGER`",
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "list",
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"active": true,
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 0,
										},
										map[string]any{
											"active": true,
											"kind": "param",
											"name": "organization_id",
											"orig": "organization_id",
											"reqd": true,
											"type": "`$STRING`",
											"index$": 1,
										},
									},
								},
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
								"index$": 0,
							},
						},
						"key$": "remove",
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
