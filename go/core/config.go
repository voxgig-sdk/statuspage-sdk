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
			"slug": "statuspage",
			"version": "0.1.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
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
						"title": "Automation Email",
						"type": "`$STRING`",
						"short": "Requires a special feature flag to be enabled",
					},
					map[string]any{
						"name": "component",
						"title": "Component",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "More detailed description for component",
					},
					map[string]any{
						"name": "group",
						"title": "Group",
						"type": "`$BOOLEAN`",
						"short": "Is this component a group",
					},
					map[string]any{
						"name": "group_id",
						"title": "Group Id",
						"type": "`$STRING`",
						"short": "Component Group identifier",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Identifier for component",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Display name for component",
					},
					map[string]any{
						"name": "only_show_if_degraded",
						"title": "Only Show If Degraded",
						"type": "`$BOOLEAN`",
						"short": "Requires a special feature flag to be enabled",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$STRING`",
						"short": "Page identifier",
					},
					map[string]any{
						"name": "position",
						"title": "Position",
						"type": "`$INTEGER`",
						"short": "Order the component will appear on the page",
						"format": "int32",
					},
					map[string]any{
						"name": "showcase",
						"title": "Showcase",
						"type": "`$BOOLEAN`",
						"short": "Should this component be showcased",
					},
					map[string]any{
						"name": "start_date",
						"title": "Start Date",
						"type": "`$STRING`",
						"short": "The date this component started being used",
						"format": "date",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "Status of component",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "page_access_group",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_users",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "page_access_user",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "page_access_group_id",
									},
									map[string]any{
										"lit": "components",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_groups",
									"{page_access_group_id}",
									"components",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_access_group_id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_group_id",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "page_access_user_id",
									},
									map[string]any{
										"lit": "components",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{page_access_user_id}",
									"components",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_user_id",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"components",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components/{component_id}/uptime",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "uptime",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.related_events`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "end",
											"orig": "end",
											"type": "Any",
											"kind": "query",
										},
										map[string]any{
											"name": "start",
											"orig": "start",
											"type": "Any",
											"kind": "query",
										},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/components/{component_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/components/{component_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_groups",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "page_access_group",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/components/{component_id}/page_access_users",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "page_access_user",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/components/{component_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"component": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.page_access_group",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.page_access_user",
						},
					},
				},
			},
			"component_group_uptime": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_id",
						"title": "Component Id",
						"type": "`$STRING`",
						"short": "Component identifier",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "incidents",
						"title": "Incidents",
						"type": "`$OBJECT`",
						"short": "Related incidents",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "component_group_uptime",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups/{id}/uptime",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "uptime",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
									"uptime",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.related_events`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "end",
											"orig": "end",
											"type": "Any",
											"kind": "query",
										},
										map[string]any{
											"name": "start",
											"orig": "start",
											"type": "Any",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"end",
										"id",
										"page_id",
										"start",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"group_component": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_group",
						"title": "Component Group",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "components",
						"title": "Components",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "Description of the component group.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Component Group Identifier",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"title": "Position",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "group_component",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/component-groups",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/component-groups/{id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "component-groups",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"component-groups",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"incident": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "auto_transition_deliver_notifications_at_end",
						"title": "Auto Transition Deliver Notifications At End",
						"type": "`$BOOLEAN`",
						"short": "Controls whether send notification when scheduled maintenances auto transition to completed.",
					},
					map[string]any{
						"name": "auto_transition_deliver_notifications_at_start",
						"title": "Auto Transition Deliver Notifications At Start",
						"type": "`$BOOLEAN`",
						"short": "Controls whether send notification when scheduled maintenances auto transition to started.",
					},
					map[string]any{
						"name": "auto_transition_to_maintenance_state",
						"title": "Auto Transition To Maintenance State",
						"type": "`$BOOLEAN`",
						"short": "Controls whether change components status to under_maintenance once scheduled maintenance is in progress.",
					},
					map[string]any{
						"name": "auto_transition_to_operational_state",
						"title": "Auto Transition To Operational State",
						"type": "`$BOOLEAN`",
						"short": "Controls whether change components status to operational once scheduled maintenance completes.",
					},
					map[string]any{
						"name": "components",
						"title": "Components",
						"type": "`$ARRAY`",
						"short": "Incident components",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The timestamp when the incident was created at.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Incident Identifier",
					},
					map[string]any{
						"name": "impact",
						"title": "Impact",
						"type": "`$STRING`",
						"short": "The impact of the incident.",
					},
					map[string]any{
						"name": "impact_override",
						"title": "Impact Override",
						"type": "`$STRING`",
						"short": "value to override calculated impact value",
					},
					map[string]any{
						"name": "incident",
						"title": "Incident",
						"type": "`$OBJECT`",
						"req": true,
						"op": map[string]any{
							"patch": map[string]any{
								"type": "`$OBJECT`",
							},
							"update": map[string]any{
								"type": "`$OBJECT`",
							},
						},
					},
					map[string]any{
						"name": "incident_updates",
						"title": "Incident Updates",
						"type": "`$ARRAY`",
						"short": "The incident updates for incident.",
					},
					map[string]any{
						"name": "metadata",
						"title": "Metadata",
						"type": "`$OBJECT`",
						"short": "Metadata attached to the incident.",
					},
					map[string]any{
						"name": "monitoring_at",
						"title": "Monitoring At",
						"type": "`$STRING`",
						"short": "The timestamp when incident entered monitoring state.",
						"format": "date-time",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Incident Name.",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$STRING`",
						"short": "Incident Page Identifier",
					},
					map[string]any{
						"name": "postmortem_body",
						"title": "Postmortem Body",
						"type": "`$STRING`",
						"short": "Body of the Postmortem.",
					},
					map[string]any{
						"name": "postmortem_body_last_updated_at",
						"title": "Postmortem Body Last Updated At",
						"type": "`$STRING`",
						"short": "The timestamp when the incident postmortem body was last updated at.",
						"format": "date-time",
					},
					map[string]any{
						"name": "postmortem_ignored",
						"title": "Postmortem Ignored",
						"type": "`$BOOLEAN`",
						"short": "Controls whether the incident will have postmortem.",
					},
					map[string]any{
						"name": "postmortem_notified_subscribers",
						"title": "Postmortem Notified Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Indicates whether subscribers are already notificed about postmortem.",
					},
					map[string]any{
						"name": "postmortem_notified_twitter",
						"title": "Postmortem Notified Twitter",
						"type": "`$BOOLEAN`",
						"short": "Controls whether to decide if notify postmortem on twitter.",
					},
					map[string]any{
						"name": "postmortem_published_at",
						"title": "Postmortem Published At",
						"type": "`$BOOLEAN`",
						"short": "The timestamp when the postmortem was published.",
					},
					map[string]any{
						"name": "reminder_intervals",
						"title": "Reminder Intervals",
						"type": "`$STRING`",
						"short": "Custom reminder intervals for unresolved/open incidents.",
					},
					map[string]any{
						"name": "resolved_at",
						"title": "Resolved At",
						"type": "`$STRING`",
						"short": "The timestamp when incident was resolved.",
						"format": "date-time",
					},
					map[string]any{
						"name": "scheduled_auto_completed",
						"title": "Scheduled Auto Completed",
						"type": "`$BOOLEAN`",
						"short": "Controls whether the incident is scheduled to automatically change to complete.",
					},
					map[string]any{
						"name": "scheduled_auto_in_progress",
						"title": "Scheduled Auto In Progress",
						"type": "`$BOOLEAN`",
						"short": "Controls whether the incident is scheduled to automatically change to in progress.",
					},
					map[string]any{
						"name": "scheduled_for",
						"title": "Scheduled For",
						"type": "`$STRING`",
						"short": "The timestamp the incident is scheduled for.",
						"format": "date-time",
					},
					map[string]any{
						"name": "scheduled_remind_prior",
						"title": "Scheduled Remind Prior",
						"type": "`$BOOLEAN`",
						"short": "Controls whether to remind subscribers prior to scheduled incidents.",
					},
					map[string]any{
						"name": "scheduled_reminded_at",
						"title": "Scheduled Reminded At",
						"type": "`$STRING`",
						"short": "The timestamp when the scheduled incident reminder was sent at.",
						"format": "date-time",
					},
					map[string]any{
						"name": "scheduled_until",
						"title": "Scheduled Until",
						"type": "`$STRING`",
						"short": "The timestamp the incident is scheduled until.",
						"format": "date-time",
					},
					map[string]any{
						"name": "shortlink",
						"title": "Shortlink",
						"type": "`$STRING`",
						"short": "Incident Shortlink.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The incident status.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"short": "The timestamp when the incident was updated at.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "incident",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"limit",
										"page",
										"page_id",
										"q",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/active_maintenance",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"lit": "active_maintenance",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"active_maintenance",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"$action": "active_maintenance",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/scheduled",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"lit": "scheduled",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"scheduled",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"$action": "scheduled",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/unresolved",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"lit": "unresolved",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"unresolved",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"$action": "unresolved",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/upcoming",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"lit": "upcoming",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"upcoming",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"$action": "upcoming",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"incident": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"incident_postmortem": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "incident_postmortem",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "postmortem",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"incident_template": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "body",
						"title": "Body",
						"type": "`$STRING`",
						"short": "Body of the incident or maintenance update to be applied when selecting this template",
					},
					map[string]any{
						"name": "components",
						"title": "Components",
						"type": "`$ARRAY`",
						"short": "Affected components",
					},
					map[string]any{
						"name": "group_id",
						"title": "Group Id",
						"type": "`$STRING`",
						"short": "Identifier of Template Group this template belongs to",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Incident Template Identifier",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page",
					},
					map[string]any{
						"name": "should_send_notifications",
						"title": "Should Send Notifications",
						"type": "`$BOOLEAN`",
						"short": "Whether the \"deliver notifications\" checkbox should be selected when selecting this template",
					},
					map[string]any{
						"name": "should_tweet",
						"title": "Should Tweet",
						"type": "`$BOOLEAN`",
						"short": "Whether the \"tweet update\" checkbox should be selected when selecting this template",
					},
					map[string]any{
						"name": "template",
						"title": "Template",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"short": "Title to be applied to the incident or maintenance when selecting this template",
					},
					map[string]any{
						"name": "update_status",
						"title": "Update Status",
						"type": "`$STRING`",
						"short": "The status the incident or maintenance should transition to when selecting this template",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "incident_template",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incident_templates",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incident_templates",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incident_templates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incident_templates",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incident_templates",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incident_templates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 100,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"incident_update": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "affected_components",
						"title": "Affected Components",
						"type": "`$ARRAY`",
						"short": "Affected components associated with the incident update.",
					},
					map[string]any{
						"name": "body",
						"title": "Body",
						"type": "`$STRING`",
						"short": "Incident update body.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The timestamp when the incident update was created at.",
						"format": "date-time",
					},
					map[string]any{
						"name": "custom_tweet",
						"title": "Custom Tweet",
						"type": "`$STRING`",
						"short": "An optional customized tweet message for incident postmortem.",
					},
					map[string]any{
						"name": "deliver_notifications",
						"title": "Deliver Notifications",
						"type": "`$BOOLEAN`",
						"short": "Controls whether to delivery notifications.",
					},
					map[string]any{
						"name": "display_at",
						"title": "Display At",
						"type": "`$STRING`",
						"short": "Timestamp when incident update is happened.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Incident Update Identifier.",
					},
					map[string]any{
						"name": "incident_id",
						"title": "Incident Id",
						"type": "`$STRING`",
						"short": "Incident Identifier.",
					},
					map[string]any{
						"name": "incident_update",
						"title": "Incident Update",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The incident status.",
					},
					map[string]any{
						"name": "tweet_id",
						"title": "Tweet Id",
						"type": "`$STRING`",
						"short": "Tweet identifier associated to this incident update.",
					},
					map[string]any{
						"name": "twitter_updated_at",
						"title": "Twitter Updated At",
						"type": "`$STRING`",
						"short": "The timestamp when twitter updated at.",
						"format": "date-time",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"short": "The timestamp when the incident update is updated.",
						"format": "date-time",
					},
					map[string]any{
						"name": "wants_twitter_update",
						"title": "Wants Twitter Update",
						"type": "`$BOOLEAN`",
						"short": "Controls whether to create twitter update.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "incident_update",
				"op": map[string]any{
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "incident_updates",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"incident_update": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_update_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "incident_updates",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"incident_update": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "incident_update_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.incident",
						},
					},
				},
			},
			"metric": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "backfill_percentage",
						"title": "Backfill Percentage",
						"type": "`$INTEGER`",
						"format": "int32",
					},
					map[string]any{
						"name": "backfilled",
						"title": "Backfilled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "decimal_places",
						"title": "Decimal Places",
						"type": "`$INTEGER`",
						"format": "int32",
					},
					map[string]any{
						"name": "display",
						"title": "Display",
						"type": "`$BOOLEAN`",
						"short": "Should the metric be displayed",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Metric identifier",
					},
					map[string]any{
						"name": "last_fetched_at",
						"title": "Last Fetched At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "metric",
						"title": "Metric",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "metric_identifier",
						"title": "Metric Identifier",
						"type": "`$STRING`",
						"short": "Metric Display identifier used to look up the metric data from the provider",
					},
					map[string]any{
						"name": "metrics_provider_id",
						"title": "Metrics Provider Id",
						"type": "`$STRING`",
						"short": "Metric Provider identifier",
					},
					map[string]any{
						"name": "most_recent_data_at",
						"title": "Most Recent Data At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of metric",
					},
					map[string]any{
						"name": "reference_name",
						"title": "Reference Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "suffix",
						"title": "Suffix",
						"type": "`$STRING`",
						"short": "Suffix to describe the units on the graph",
					},
					map[string]any{
						"name": "tooltip_description",
						"title": "Tooltip Description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "y_axis_hidden",
						"title": "Y Axis Hidden",
						"type": "`$BOOLEAN`",
						"short": "Should the values on the y axis be hidden on render",
					},
					map[string]any{
						"name": "y_axis_max",
						"title": "Y Axis Max",
						"type": "`$NUMBER`",
						"format": "float",
					},
					map[string]any{
						"name": "y_axis_min",
						"title": "Y Axis Min",
						"type": "`$NUMBER`",
						"format": "float",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "metric",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics/{metric_id}/data",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "data",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
									map[string]any{
										"var": "metrics_provider_id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{metrics_provider_id}",
									"metrics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"metrics_provider_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics/data",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"lit": "data",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
									"data",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "page_access_user_id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"page_access_users",
									"{page_access_user_id}",
									"metrics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_access_user_id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_access_user_id",
										"page_id",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
									map[string]any{
										"var": "metrics_provider_id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
									"{metrics_provider_id}",
									"metrics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "metrics_provider_id",
											"orig": "metrics_provider_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"metrics_provider_id",
										"page",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics/{metric_id}/data",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "data",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "data",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/metrics/{metric_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"metric": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.metrics_provider",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.page_access_user",
						},
					},
				},
			},
			"metrics_provider": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "disabled",
						"title": "Disabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Identifier for Metrics Provider",
					},
					map[string]any{
						"name": "last_revalidated_at",
						"title": "Last Revalidated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "metric_base_uri",
						"title": "Metric Base Uri",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "metrics_provider",
						"title": "Metrics Provider",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$INTEGER`",
						"format": "int32",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "metrics_provider",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/metrics_providers",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"metrics_providers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metrics_provider_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metrics_provider_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metrics_provider_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "metrics_providers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"metrics_provider": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "metrics_provider_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "activity_score",
						"title": "Activity Score",
						"type": "`$NUMBER`",
						"format": "float",
					},
					map[string]any{
						"name": "allow_email_subscribers",
						"title": "Allow Email Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Can your users choose to receive notifications via email",
					},
					map[string]any{
						"name": "allow_incident_subscribers",
						"title": "Allow Incident Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Can your users subscribe to notifications for a single incident",
					},
					map[string]any{
						"name": "allow_page_subscribers",
						"title": "Allow Page Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Can your users subscribe to all notifications on the page",
					},
					map[string]any{
						"name": "allow_rss_atom_feeds",
						"title": "Allow Rss Atom Feeds",
						"type": "`$BOOLEAN`",
						"short": "Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)",
					},
					map[string]any{
						"name": "allow_sms_subscribers",
						"title": "Allow Sms Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Can your users choose to receive notifications via SMS",
					},
					map[string]any{
						"name": "allow_webhook_subscribers",
						"title": "Allow Webhook Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Can your users choose to receive notifications via Webhooks",
					},
					map[string]any{
						"name": "branding",
						"title": "Branding",
						"type": "`$STRING`",
						"short": "The main template your statuspage will use",
					},
					map[string]any{
						"name": "city",
						"title": "City",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "country",
						"title": "Country",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp the record was created",
						"format": "date-time",
					},
					map[string]any{
						"name": "css_blues",
						"title": "Css Blues",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_body_background_color",
						"title": "Css Body Background Color",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_border_color",
						"title": "Css Border Color",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_font_color",
						"title": "Css Font Color",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_graph_color",
						"title": "Css Graph Color",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_greens",
						"title": "Css Greens",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_light_font_color",
						"title": "Css Light Font Color",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_link_color",
						"title": "Css Link Color",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_no_data",
						"title": "Css No Data",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_oranges",
						"title": "Css Oranges",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_reds",
						"title": "Css Reds",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "css_yellows",
						"title": "Css Yellows",
						"type": "`$STRING`",
						"short": "CSS Color",
					},
					map[string]any{
						"name": "domain",
						"title": "Domain",
						"type": "`$STRING`",
						"short": "CNAME alias for your status page",
					},
					map[string]any{
						"name": "email_logo",
						"title": "Email Logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "favicon_logo",
						"title": "Favicon Logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "headline",
						"title": "Headline",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hero_cover",
						"title": "Hero Cover",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hidden_from_search",
						"title": "Hidden From Search",
						"type": "`$BOOLEAN`",
						"short": "Should your page hide itself from search engines",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Page identifier",
					},
					map[string]any{
						"name": "ip_restrictions",
						"title": "Ip Restrictions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of your page to be displayed",
					},
					map[string]any{
						"name": "notifications_email_footer",
						"title": "Notifications Email Footer",
						"type": "`$STRING`",
						"short": "Allows you to customize the footer appearing on your notification emails.",
					},
					map[string]any{
						"name": "notifications_from_email",
						"title": "Notifications From Email",
						"type": "`$STRING`",
						"short": "Allows you to customize the email address your page notifications come from",
					},
					map[string]any{
						"name": "page",
						"title": "Page",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_description",
						"title": "Page Description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "state",
						"title": "State",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "subdomain",
						"title": "Subdomain",
						"type": "`$STRING`",
						"short": "Subdomain at which to access your status page",
					},
					map[string]any{
						"name": "support_url",
						"title": "Support Url",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "time_zone",
						"title": "Time Zone",
						"type": "`$STRING`",
						"short": "Timezone configured for your page",
					},
					map[string]any{
						"name": "transactional_logo",
						"title": "Transactional Logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_logo",
						"title": "Twitter Logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "twitter_username",
						"title": "Twitter Username",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"short": "Timestamp the record was last updated",
						"format": "date-time",
					},
					map[string]any{
						"name": "url",
						"title": "Url",
						"type": "`$STRING`",
						"short": "Website of your page.",
					},
					map[string]any{
						"name": "viewers_must_be_team_members",
						"title": "Viewers Must Be Team Members",
						"type": "`$BOOLEAN`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
								},
								"parts": []any{
									"pages",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"pages",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"page_id": "id",
									},
								},
								"transform": map[string]any{
									"req": map[string]any{
										"page": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Component Ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "external_identifier",
						"title": "External Identifier",
						"type": "`$STRING`",
						"short": "Associates group with external group.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Page Access Group Identifier",
					},
					map[string]any{
						"name": "metric_ids",
						"title": "Metric Ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name for this Group.",
					},
					map[string]any{
						"name": "page_access_group",
						"title": "Page Access Group",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_access_user_ids",
						"title": "Page Access User Ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$STRING`",
						"short": "Page Identifier.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "page_access_group",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_groups",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "component_id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "component_id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"component_id",
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_group": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_groups",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_group_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.component",
						},
					},
				},
			},
			"page_access_user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "external_login",
						"title": "External Login",
						"type": "`$STRING`",
						"short": "IDP login user id.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Page Access User Identifier",
					},
					map[string]any{
						"name": "page_access_group_id",
						"title": "Page Access Group Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_group_ids",
						"title": "Page Access Group Ids",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "page_access_user",
						"title": "Page Access User",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "page_access_user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/page_access_users",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
								},
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
								"transform": map[string]any{
									"req": map[string]any{
										"page_access_user": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "email",
											"orig": "email",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
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
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
									map[string]any{
										"var": "component_id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "component_id",
											"orig": "component_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"component_id",
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "metrics",
									},
									map[string]any{
										"var": "metric_id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "metric_id",
											"orig": "metric_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"metric_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "components",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "component",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "page_access_users",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "page_access_user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "metric",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.component",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.metric",
						},
					},
				},
			},
			"permission": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pages",
						"title": "Pages",
						"type": "`$OBJECT`",
						"short": "Pages accessible by the user.",
					},
					map[string]any{
						"name": "user_id",
						"title": "User Id",
						"type": "`$STRING`",
						"short": "User identifier",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "permission",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/organizations/{organization_id}/permissions/{user_id}",
								"segments": []any{
									map[string]any{
										"lit": "organizations",
									},
									map[string]any{
										"var": "organization_id",
									},
									map[string]any{
										"lit": "permissions",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "organization_id",
											"orig": "organization_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/organizations/{organization_id}/permissions/{user_id}",
								"segments": []any{
									map[string]any{
										"lit": "organizations",
									},
									map[string]any{
										"var": "organization_id",
									},
									map[string]any{
										"lit": "permissions",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "organization_id",
											"orig": "organization_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"postmortem": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "body",
						"title": "Body",
						"type": "`$STRING`",
						"short": "Postmortem body",
					},
					map[string]any{
						"name": "body_draft",
						"title": "Body Draft",
						"type": "`$STRING`",
						"short": "Body draft",
					},
					map[string]any{
						"name": "body_draft_updated_at",
						"title": "Body Draft Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "body_updated_at",
						"title": "Body Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "custom_tweet",
						"title": "Custom Tweet",
						"type": "`$STRING`",
						"short": "Custom tweet for Incident Postmortem",
					},
					map[string]any{
						"name": "notify_subscribers",
						"title": "Notify Subscribers",
						"type": "`$BOOLEAN`",
						"short": "Should email subscribers be notified.",
					},
					map[string]any{
						"name": "notify_twitter",
						"title": "Notify Twitter",
						"type": "`$BOOLEAN`",
						"short": "Should Twitter followers be notified.",
					},
					map[string]any{
						"name": "postmortem",
						"title": "Postmortem",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "preview_key",
						"title": "Preview Key",
						"type": "`$STRING`",
						"short": "Preview Key",
					},
					map[string]any{
						"name": "published_at",
						"title": "Published At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
				},
				"name": "postmortem",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "postmortem",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "postmortem",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"postmortem": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/publish",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "postmortem",
									},
									map[string]any{
										"lit": "publish",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
									"publish",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"postmortem": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "publish",
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/revert",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "postmortem",
									},
									map[string]any{
										"lit": "revert",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"postmortem",
									"revert",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "revert",
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.incident",
						},
					},
				},
			},
			"status_embed_config": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "incident_background_color",
						"title": "Incident Background Color",
						"type": "`$STRING`",
						"short": "Color of status embed iframe background when displaying incident",
					},
					map[string]any{
						"name": "incident_text_color",
						"title": "Incident Text Color",
						"type": "`$STRING`",
						"short": "Color of status embed iframe text when displaying incident",
					},
					map[string]any{
						"name": "maintenance_background_color",
						"title": "Maintenance Background Color",
						"type": "`$STRING`",
						"short": "Color of status embed iframe background when displaying maintenance",
					},
					map[string]any{
						"name": "maintenance_text_color",
						"title": "Maintenance Text Color",
						"type": "`$STRING`",
						"short": "Color of status embed iframe text when displaying maintenance",
					},
					map[string]any{
						"name": "page_id",
						"title": "Page Id",
						"type": "`$STRING`",
						"short": "Page identifier",
					},
					map[string]any{
						"name": "position",
						"title": "Position",
						"type": "`$STRING`",
						"short": "Corner where status embed iframe will appear on page",
					},
					map[string]any{
						"name": "status_embed_config",
						"title": "Status Embed Config",
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
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/status_embed_config",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "status_embed_config",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"patch": map[string]any{
						"input": "data",
						"name": "patch",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/status_embed_config",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "status_embed_config",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"status_embed_config": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/pages/{page_id}/status_embed_config",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "status_embed_config",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"status_embed_config",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"status_embed_config": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
					},
				},
			},
			"subscriber": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "component_ids",
						"title": "Component Ids",
						"type": "`$ARRAY`",
						"short": "A list of component ids for which the subscriber should recieve updates for.",
					},
					map[string]any{
						"name": "components",
						"title": "Components",
						"type": "`$STRING`",
						"short": "The components for which the subscriber has elected to receive updates.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "display_phone_number",
						"title": "Display Phone Number",
						"type": "`$STRING`",
						"short": "A formatted version of the phone_number and phone_country pair, nicely formatted for display.",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"short": "The email address to use to contact the subscriber.",
					},
					map[string]any{
						"name": "endpoint",
						"title": "Endpoint",
						"type": "`$STRING`",
						"short": "The URL where a webhook subscriber elects to receive updates.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Subscriber Identifier",
					},
					map[string]any{
						"name": "mode",
						"title": "Mode",
						"type": "`$STRING`",
						"short": "The communication mode of the subscriber.",
					},
					map[string]any{
						"name": "obfuscated_channel_name",
						"title": "Obfuscated Channel Name",
						"type": "`$STRING`",
						"short": "Obfuscated slack channel name",
					},
					map[string]any{
						"name": "page_access_user_id",
						"title": "Page Access User Id",
						"type": "`$STRING`",
						"short": "The Page Access user this subscriber belongs to (only for audience-specific pages).",
					},
					map[string]any{
						"name": "phone_country",
						"title": "Phone Country",
						"type": "`$STRING`",
						"short": "The two-character country code representing the country of which the phone_number is a part.",
					},
					map[string]any{
						"name": "phone_number",
						"title": "Phone Number",
						"type": "`$STRING`",
						"short": "The phone number used to contact an SMS subscriber",
					},
					map[string]any{
						"name": "purge_at",
						"title": "Purge At",
						"type": "`$STRING`",
						"short": "The timestamp when a quarantined subscriber will be purged (unsubscribed).",
						"format": "date-time",
					},
					map[string]any{
						"name": "quarantined_at",
						"title": "Quarantined At",
						"type": "`$STRING`",
						"short": "The timestamp when the subscriber was quarantined due to an issue reaching them.",
						"format": "date-time",
					},
					map[string]any{
						"name": "skip_confirmation_notification",
						"title": "Skip Confirmation Notification",
						"type": "`$BOOLEAN`",
						"short": "If this is true, do not notify the user with changes to their subscription.",
					},
					map[string]any{
						"name": "subscriber",
						"title": "Subscriber",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "workspace_name",
						"title": "Workspace Name",
						"type": "`$STRING`",
						"short": "The workspace name of the slack subscriber.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "subscriber",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "resend_confirmation",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
									"{id}",
									"resend_confirmation",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"subscriber_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "resend_confirmation",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"subscriber": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"subscriber": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/reactivate",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"lit": "reactivate",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"reactivate",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "reactivate",
									"exist": []any{
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/resend_confirmation",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"lit": "resend_confirmation",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"resend_confirmation",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "resend_confirmation",
									"exist": []any{
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/pages/{page_id}/subscribers/unsubscribe",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"lit": "unsubscribe",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"unsubscribe",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "unsubscribe",
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 0,
										},
										map[string]any{
											"name": "q",
											"orig": "q",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "sort_direction",
											"orig": "sort_direction",
											"type": "`$STRING`",
											"kind": "query",
											"example": "asc",
										},
										map[string]any{
											"name": "sort_field",
											"orig": "sort_field",
											"type": "`$STRING`",
											"kind": "query",
											"example": "primary",
										},
										map[string]any{
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
											"kind": "query",
											"example": "active",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
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
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"incidents",
									"{incident_id}",
									"subscribers",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"incident_id",
										"page",
										"page_id",
										"per_page",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/unsubscribed",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"lit": "unsubscribed",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"unsubscribed",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "unsubscribed",
									"exist": []any{
										"page",
										"page_id",
										"per_page",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/count",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"lit": "count",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"count",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "state",
											"orig": "state",
											"type": "`$STRING`",
											"kind": "query",
											"example": "active",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "count",
									"exist": []any{
										"page_id",
										"state",
										"type",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/pages/{page_id}/subscribers/histogram_by_state",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"lit": "histogram_by_state",
									},
								},
								"parts": []any{
									"pages",
									"{page_id}",
									"subscribers",
									"histogram_by_state",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "histogram_by_state",
									"exist": []any{
										"page_id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "incidents",
									},
									map[string]any{
										"var": "incident_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "incident_id",
											"orig": "incident_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"incident_id",
										"page_id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "skip_unsubscription_notification",
											"orig": "skip_unsubscription_notification",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
										"skip_unsubscription_notification",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/pages/{page_id}/subscribers/{subscriber_id}",
								"segments": []any{
									map[string]any{
										"lit": "pages",
									},
									map[string]any{
										"var": "page_id",
									},
									map[string]any{
										"lit": "subscribers",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "subscriber_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "page_id",
											"orig": "page_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"page_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.page",
						},
						[]any{
							"$.main.kit.entity.page",
							"$.main.kit.entity.incident",
						},
					},
				},
			},
			"user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"short": "Email address for the team member",
					},
					map[string]any{
						"name": "first_name",
						"title": "First Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "User identifier",
					},
					map[string]any{
						"name": "last_name",
						"title": "Last Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "organization_id",
						"title": "Organization Id",
						"type": "`$STRING`",
						"short": "Organization identifier",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "user",
						"title": "User",
						"type": "`$OBJECT`",
						"req": true,
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "user",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/organizations/{organization_id}/users",
								"segments": []any{
									map[string]any{
										"lit": "organizations",
									},
									map[string]any{
										"var": "organization_id",
									},
									map[string]any{
										"lit": "users",
									},
								},
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"user": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "organization_id",
											"orig": "organization_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"organization_id",
									},
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/organizations/{organization_id}/users",
								"segments": []any{
									map[string]any{
										"lit": "organizations",
									},
									map[string]any{
										"var": "organization_id",
									},
									map[string]any{
										"lit": "users",
									},
								},
								"parts": []any{
									"organizations",
									"{organization_id}",
									"users",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "organization_id",
											"orig": "organization_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "per_page",
											"orig": "per_page",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"organization_id",
										"page",
										"per_page",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/organizations/{organization_id}/users/{user_id}",
								"segments": []any{
									map[string]any{
										"lit": "organizations",
									},
									map[string]any{
										"var": "organization_id",
									},
									map[string]any{
										"lit": "users",
									},
									map[string]any{
										"var": "id",
									},
								},
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "user_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "organization_id",
											"orig": "organization_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"organization_id",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
