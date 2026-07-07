# Statuspage SDK configuration

module StatuspageConfig
  def self.make_config
    {
      "main" => {
        "name" => "Statuspage",
      },
      "feature" => {
        "test" => {
          "options" => {
            "active" => false,
          },
        },
      },
      "options" => {
        "base" => "https://api.statuspage.io/v1",
        "auth" => {
          "prefix" => "OAuth",
        },
        "headers" => {
          "content-type" => "application/json",
        },
        "entity" => {
          "component" => {},
          "component_group_uptime" => {},
          "group_component" => {},
          "incident" => {},
          "incident_postmortem" => {},
          "incident_subscriber" => {},
          "incident_template" => {},
          "incident_update" => {},
          "metric" => {},
          "metrics_provider" => {},
          "page" => {},
          "page_access_group" => {},
          "page_access_user" => {},
          "permission" => {},
          "postmortem" => {},
          "status_embed_config" => {},
          "subscriber" => {},
          "user" => {},
        },
      },
      "entity" => {
        "component" => {
          "fields" => [
            {
              "active" => true,
              "name" => "automation_email",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "component",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "description",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "group",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "group_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "major_outage",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "only_show_if_degraded",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "partial_outage",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "position",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "range_end",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "range_start",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "related_event",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "showcase",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "start_date",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "status",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 19,
            },
            {
              "active" => true,
              "name" => "uptime_percentage",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 20,
            },
            {
              "active" => true,
              "name" => "warning",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 21,
            },
          ],
          "name" => "component",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/components/{component_id}/page_access_groups",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                    "page_access_groups",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "page_access_group",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/components/{component_id}/page_access_users",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                    "page_access_users",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "page_access_user",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_access_group_id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{page_access_group_id}",
                    "components",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_access_group_id",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_access_user_id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{page_access_user_id}",
                    "components",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_access_user_id",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "end",
                        "orig" => "end",
                        "reqd" => false,
                        "type" => "Any",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "skip_related_event",
                        "orig" => "skip_related_event",
                        "reqd" => false,
                        "type" => "`$BOOLEAN`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "start",
                        "orig" => "start",
                        "reqd" => false,
                        "type" => "Any",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/components/{component_id}/uptime",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                    "uptime",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "uptime",
                    "exist" => [
                      "end",
                      "id",
                      "page_id",
                      "skip_related_event",
                      "start",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/components/{component_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/components/{component_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/components/{component_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/components/{component_id}/page_access_groups",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                    "page_access_groups",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "page_access_group",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/components/{component_id}/page_access_users",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                    "page_access_users",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "page_access_user",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/components/{component_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "components",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "component_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
              [
                "page",
                "page_access_group",
              ],
              [
                "page",
                "page_access_user",
              ],
            ],
          },
        },
        "component_group_uptime" => {
          "fields" => [
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "major_outage",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "partial_outage",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "range_end",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "range_start",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "related_event",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "uptime_percentage",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "warning",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
          ],
          "name" => "component_group_uptime",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "end",
                        "orig" => "end",
                        "reqd" => false,
                        "type" => "Any",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "skip_related_event",
                        "orig" => "skip_related_event",
                        "reqd" => false,
                        "type" => "`$BOOLEAN`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "start",
                        "orig" => "start",
                        "reqd" => false,
                        "type" => "Any",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/component-groups/{id}/uptime",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                    "{id}",
                    "uptime",
                  ],
                  "select" => {
                    "exist" => [
                      "end",
                      "id",
                      "page_id",
                      "skip_related_event",
                      "start",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "group_component" => {
          "fields" => [
            {
              "active" => true,
              "name" => "component",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "component_group",
              "req" => true,
              "type" => "`$OBJECT`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "description",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "position",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
          ],
          "name" => "group_component",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/component-groups",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/component-groups",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/component-groups/{id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                    "{id}",
                  ],
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/component-groups/{id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                    "{id}",
                  ],
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/component-groups/{id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                    "{id}",
                  ],
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/component-groups/{id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "component-groups",
                    "{id}",
                  ],
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "incident" => {
          "fields" => [
            {
              "active" => true,
              "name" => "auto_transition_deliver_notifications_at_end",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "auto_transition_deliver_notifications_at_start",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "auto_transition_to_maintenance_state",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "auto_transition_to_operational_state",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "component",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "impact",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "impact_override",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "incident",
              "op" => {
                "patch" => {
                  "req" => false,
                  "type" => "`$OBJECT`",
                },
                "update" => {
                  "req" => false,
                  "type" => "`$OBJECT`",
                },
              },
              "req" => true,
              "type" => "`$OBJECT`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "incident_impact",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "incident_update",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "metadata",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "monitoring_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "postmortem_body",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "postmortem_body_last_updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "postmortem_ignored",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "postmortem_notified_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 19,
            },
            {
              "active" => true,
              "name" => "postmortem_notified_twitter",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 20,
            },
            {
              "active" => true,
              "name" => "postmortem_published_at",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 21,
            },
            {
              "active" => true,
              "name" => "reminder_interval",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 22,
            },
            {
              "active" => true,
              "name" => "resolved_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 23,
            },
            {
              "active" => true,
              "name" => "scheduled_auto_completed",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 24,
            },
            {
              "active" => true,
              "name" => "scheduled_auto_in_progress",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 25,
            },
            {
              "active" => true,
              "name" => "scheduled_for",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 26,
            },
            {
              "active" => true,
              "name" => "scheduled_remind_prior",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 27,
            },
            {
              "active" => true,
              "name" => "scheduled_reminded_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 28,
            },
            {
              "active" => true,
              "name" => "scheduled_until",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 29,
            },
            {
              "active" => true,
              "name" => "shortlink",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 30,
            },
            {
              "active" => true,
              "name" => "status",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 31,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 32,
            },
          ],
          "name" => "incident",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/incidents",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "limit",
                        "orig" => "limit",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "q",
                        "orig" => "q",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                  ],
                  "select" => {
                    "exist" => [
                      "limit",
                      "page",
                      "page_id",
                      "q",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "example" => 1,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "example" => 100,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/active_maintenance",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "active_maintenance",
                  ],
                  "select" => {
                    "$action" => "active_maintenance",
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "example" => 1,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "example" => 100,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/scheduled",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "scheduled",
                  ],
                  "select" => {
                    "$action" => "scheduled",
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "example" => 1,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "example" => 100,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/unresolved",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "unresolved",
                  ],
                  "select" => {
                    "$action" => "unresolved",
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 3,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "example" => 1,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "example" => 100,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/upcoming",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "upcoming",
                  ],
                  "select" => {
                    "$action" => "upcoming",
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 4,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "incident_postmortem" => {
          "fields" => [],
          "name" => "incident_postmortem",
          "op" => {
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/postmortem",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{id}",
                    "postmortem",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "remove",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "incident_subscriber" => {
          "fields" => [],
          "name" => "incident_subscriber",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "subscriber_id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "subscribers",
                    "{subscriber_id}",
                    "resend_confirmation",
                  ],
                  "select" => {
                    "exist" => [
                      "incident_id",
                      "page_id",
                      "subscriber_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
                "incident",
                "subscriber",
              ],
            ],
          },
        },
        "incident_template" => {
          "fields" => [
            {
              "active" => true,
              "name" => "body",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "component",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "group_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "should_send_notification",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "should_tweet",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "template",
              "req" => true,
              "type" => "`$OBJECT`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "title",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "update_status",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
          ],
          "name" => "incident_template",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/incident_templates",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incident_templates",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "example" => 1,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "example" => 100,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incident_templates",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incident_templates",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "incident_update" => {
          "fields" => [
            {
              "active" => true,
              "name" => "affected_component",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "body",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "custom_tweet",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "deliver_notification",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "display_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "incident_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "incident_update",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "status",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "tweet_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "twitter_updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "wants_twitter_update",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 13,
            },
          ],
          "name" => "incident_update",
          "op" => {
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_update_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "incident_updates",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_update_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "incident_update_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "incident_updates",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "incident_update_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
                "incident",
              ],
            ],
          },
        },
        "metric" => {
          "fields" => [
            {
              "active" => true,
              "name" => "backfill_percentage",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "backfilled",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "data",
              "req" => true,
              "type" => "`$OBJECT`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "decimal_place",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "display",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "last_fetched_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "metric",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "metric_identifier",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "metrics_provider_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "most_recent_data_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "reference_name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "suffix",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "tooltip_description",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "y_axis_hidden",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "y_axis_max",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "y_axis_min",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 19,
            },
          ],
          "name" => "metric",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/metrics/{metric_id}/data",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "{id}",
                    "data",
                  ],
                  "rename" => {
                    "param" => {
                      "metric_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "data",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "metrics_provider_id",
                        "orig" => "metrics_provider_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                    "{metrics_provider_id}",
                    "metrics",
                  ],
                  "select" => {
                    "exist" => [
                      "metrics_provider_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/metrics/data",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "data",
                  ],
                  "select" => {
                    "$action" => "data",
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_access_user_id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{page_access_user_id}",
                    "metrics",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_access_user_id",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "metrics_provider_id",
                        "orig" => "metrics_provider_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                    "{metrics_provider_id}",
                    "metrics",
                  ],
                  "select" => {
                    "exist" => [
                      "metrics_provider_id",
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                  ],
                  "select" => {
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/metrics/{metric_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metric_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/metrics/{metric_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metric_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/metrics/{metric_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metric_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/metrics/{metric_id}/data",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "{id}",
                    "data",
                  ],
                  "rename" => {
                    "param" => {
                      "metric_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "data",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/metrics/{metric_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metric_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
              [
                "page",
                "metrics_provider",
              ],
              [
                "page",
                "page_access_user",
              ],
            ],
          },
        },
        "metrics_provider" => {
          "fields" => [
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "disabled",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "last_revalidated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "metric_base_uri",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "metrics_provider",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "type",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
          ],
          "name" => "metrics_provider",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/metrics_providers",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/metrics_providers",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metrics_provider_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metrics_provider_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metrics_provider_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metrics_provider_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metrics_provider_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metrics_provider_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "metrics_provider_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/metrics_providers/{metrics_provider_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "metrics_providers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "metrics_provider_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "page" => {
          "fields" => [
            {
              "active" => true,
              "name" => "activity_score",
              "req" => false,
              "type" => "`$NUMBER`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "allow_email_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "allow_incident_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "allow_page_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "allow_rss_atom_feed",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "allow_sms_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "allow_webhook_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "branding",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "city",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "country",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "css_blue",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "css_body_background_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "css_border_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "css_font_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "css_graph_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "css_green",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "css_light_font_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "css_link_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "css_no_data",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 19,
            },
            {
              "active" => true,
              "name" => "css_orange",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 20,
            },
            {
              "active" => true,
              "name" => "css_red",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 21,
            },
            {
              "active" => true,
              "name" => "css_yellow",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 22,
            },
            {
              "active" => true,
              "name" => "domain",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 23,
            },
            {
              "active" => true,
              "name" => "email_logo",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 24,
            },
            {
              "active" => true,
              "name" => "favicon_logo",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 25,
            },
            {
              "active" => true,
              "name" => "headline",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 26,
            },
            {
              "active" => true,
              "name" => "hero_cover",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 27,
            },
            {
              "active" => true,
              "name" => "hidden_from_search",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 28,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 29,
            },
            {
              "active" => true,
              "name" => "ip_restriction",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 30,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 31,
            },
            {
              "active" => true,
              "name" => "notifications_email_footer",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 32,
            },
            {
              "active" => true,
              "name" => "notifications_from_email",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 33,
            },
            {
              "active" => true,
              "name" => "page",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 34,
            },
            {
              "active" => true,
              "name" => "page_description",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 35,
            },
            {
              "active" => true,
              "name" => "state",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 36,
            },
            {
              "active" => true,
              "name" => "subdomain",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 37,
            },
            {
              "active" => true,
              "name" => "support_url",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 38,
            },
            {
              "active" => true,
              "name" => "time_zone",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 39,
            },
            {
              "active" => true,
              "name" => "transactional_logo",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 40,
            },
            {
              "active" => true,
              "name" => "twitter_logo",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 41,
            },
            {
              "active" => true,
              "name" => "twitter_username",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 42,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 43,
            },
            {
              "active" => true,
              "name" => "url",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 44,
            },
            {
              "active" => true,
              "name" => "viewers_must_be_team_member",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 45,
            },
          ],
          "name" => "page",
          "op" => {
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {},
                  "method" => "GET",
                  "orig" => "/pages",
                  "parts" => [
                    "pages",
                  ],
                  "select" => {},
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}",
                  "parts" => [
                    "pages",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}",
                  "parts" => [
                    "pages",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}",
                  "parts" => [
                    "pages",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [],
          },
        },
        "page_access_group" => {
          "fields" => [
            {
              "active" => true,
              "name" => "component_id",
              "op" => {
                "create" => {
                  "req" => true,
                  "type" => "`$ARRAY`",
                },
              },
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "external_identifier",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "metric_id",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "page_access_group",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "page_access_user_id",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
          ],
          "name" => "page_access_group",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/page_access_groups",
                  "parts" => [
                    "pages",
                    "{id}",
                    "page_access_groups",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_groups",
                  "parts" => [
                    "pages",
                    "{id}",
                    "page_access_groups",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "component_id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                    "components",
                    "{component_id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "component_id",
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_group_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/page_access_groups/{page_access_group_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_groups",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_group_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
              [
                "page",
                "component",
              ],
            ],
          },
        },
        "page_access_user" => {
          "fields" => [
            {
              "active" => true,
              "name" => "component_id",
              "req" => true,
              "type" => "`$ARRAY`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "email",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "external_login",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "metric_id",
              "req" => true,
              "type" => "`$ARRAY`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "page_access_group_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "page_access_user",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
          ],
          "name" => "page_access_user",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "metrics",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "metric",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/page_access_users",
                  "parts" => [
                    "pages",
                    "{id}",
                    "page_access_users",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "email",
                        "orig" => "email",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_users",
                  "parts" => [
                    "pages",
                    "{id}",
                    "page_access_users",
                  ],
                  "rename" => {
                    "param" => {
                      "page_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "email",
                      "id",
                      "page",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "metrics",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "metric",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "patch",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "component_id",
                        "orig" => "component_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "components",
                    "{component_id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "component_id",
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "metric_id",
                        "orig" => "metric_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "metrics",
                    "{metric_id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "metric_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 3,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "metrics",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "metric",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 4,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/components",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "components",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "component",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "page_access_user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "page_access_users",
                    "{id}",
                    "metrics",
                  ],
                  "rename" => {
                    "param" => {
                      "page_access_user_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "metric",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
              [
                "page",
                "component",
              ],
              [
                "page",
                "metric",
              ],
            ],
          },
        },
        "permission" => {
          "fields" => [
            {
              "active" => true,
              "name" => "data",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "page",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 1,
            },
          ],
          "name" => "permission",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "organization_id",
                        "orig" => "organization_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/organizations/{organization_id}/permissions/{user_id}",
                  "parts" => [
                    "organizations",
                    "{organization_id}",
                    "permissions",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "organization_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "organization_id",
                        "orig" => "organization_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/organizations/{organization_id}/permissions/{user_id}",
                  "parts" => [
                    "organizations",
                    "{organization_id}",
                    "permissions",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "organization_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "organization",
              ],
            ],
          },
        },
        "postmortem" => {
          "fields" => [
            {
              "active" => true,
              "name" => "body",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "body_draft",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "body_draft_updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "body_updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "custom_tweet",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "notify_subscriber",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "notify_twitter",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "postmortem",
              "op" => {
                "update" => {
                  "req" => false,
                  "type" => "`$OBJECT`",
                },
              },
              "req" => true,
              "type" => "`$OBJECT`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "preview_key",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "published_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
          ],
          "name" => "postmortem",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/postmortem",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "postmortem",
                  ],
                  "select" => {
                    "exist" => [
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/postmortem",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "postmortem",
                  ],
                  "select" => {
                    "exist" => [
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/postmortem/publish",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "postmortem",
                    "publish",
                  ],
                  "select" => {
                    "$action" => "publish",
                    "exist" => [
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/postmortem/revert",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "postmortem",
                    "revert",
                  ],
                  "select" => {
                    "$action" => "revert",
                    "exist" => [
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
                "incident",
              ],
            ],
          },
        },
        "status_embed_config" => {
          "fields" => [
            {
              "active" => true,
              "name" => "incident_background_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "incident_text_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "maintenance_background_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "maintenance_text_color",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "page_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "position",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "status_embed_config",
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 6,
            },
          ],
          "name" => "status_embed_config",
          "op" => {
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/status_embed_config",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "status_embed_config",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "load",
            },
            "patch" => {
              "input" => "data",
              "name" => "patch",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/status_embed_config",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "status_embed_config",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "patch",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "PUT",
                  "orig" => "/pages/{page_id}/status_embed_config",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "status_embed_config",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
            ],
          },
        },
        "subscriber" => {
          "fields" => [
            {
              "active" => true,
              "name" => "component",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "component_id",
              "req" => false,
              "type" => "`$ARRAY`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "display_phone_number",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "email",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "endpoint",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "integration_partner",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 7,
            },
            {
              "active" => true,
              "name" => "mode",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 8,
            },
            {
              "active" => true,
              "name" => "obfuscated_channel_name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 9,
            },
            {
              "active" => true,
              "name" => "page_access_user_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 10,
            },
            {
              "active" => true,
              "name" => "phone_country",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 11,
            },
            {
              "active" => true,
              "name" => "phone_number",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 12,
            },
            {
              "active" => true,
              "name" => "purge_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 13,
            },
            {
              "active" => true,
              "name" => "quarantined_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 14,
            },
            {
              "active" => true,
              "name" => "skip_confirmation_notification",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 15,
            },
            {
              "active" => true,
              "name" => "skip_unsubscription_notification",
              "req" => false,
              "type" => "`$BOOLEAN`",
              "index$" => 16,
            },
            {
              "active" => true,
              "name" => "slack",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 17,
            },
            {
              "active" => true,
              "name" => "sms",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 18,
            },
            {
              "active" => true,
              "name" => "state",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 19,
            },
            {
              "active" => true,
              "name" => "subscriber",
              "op" => {
                "create" => {
                  "req" => true,
                  "type" => "`$STRING`",
                },
              },
              "req" => false,
              "type" => "`$OBJECT`",
              "index$" => 20,
            },
            {
              "active" => true,
              "name" => "team",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 21,
            },
            {
              "active" => true,
              "name" => "type",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 22,
            },
            {
              "active" => true,
              "name" => "webhook",
              "req" => false,
              "type" => "`$INTEGER`",
              "index$" => 23,
            },
            {
              "active" => true,
              "name" => "workspace_name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 24,
            },
          ],
          "name" => "subscriber",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "{id}",
                    "resend_confirmation",
                  ],
                  "rename" => {
                    "param" => {
                      "subscriber_id" => "id",
                    },
                  },
                  "select" => {
                    "$action" => "resend_confirmation",
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/subscribers",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "subscribers",
                  ],
                  "select" => {
                    "exist" => [
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/subscribers",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                  ],
                  "select" => {
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/subscribers/reactivate",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "reactivate",
                  ],
                  "select" => {
                    "$action" => "reactivate",
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 3,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/subscribers/resend_confirmation",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "resend_confirmation",
                  ],
                  "select" => {
                    "$action" => "resend_confirmation",
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 4,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/pages/{page_id}/subscribers/unsubscribe",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "unsubscribe",
                  ],
                  "select" => {
                    "$action" => "unsubscribe",
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 5,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "limit",
                        "orig" => "limit",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "example" => 0,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "q",
                        "orig" => "q",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "asc",
                        "kind" => "query",
                        "name" => "sort_direction",
                        "orig" => "sort_direction",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "primary",
                        "kind" => "query",
                        "name" => "sort_field",
                        "orig" => "sort_field",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "example" => "active",
                        "kind" => "query",
                        "name" => "state",
                        "orig" => "state",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "type",
                        "orig" => "type",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/subscribers",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                  ],
                  "select" => {
                    "exist" => [
                      "limit",
                      "page",
                      "page_id",
                      "q",
                      "sort_direction",
                      "sort_field",
                      "state",
                      "type",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/subscribers",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "subscribers",
                  ],
                  "select" => {
                    "exist" => [
                      "incident_id",
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/subscribers/unsubscribed",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "unsubscribed",
                  ],
                  "select" => {
                    "$action" => "unsubscribed",
                    "exist" => [
                      "page",
                      "page_id",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
              ],
              "key$" => "list",
            },
            "load" => {
              "input" => "data",
              "name" => "load",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "subscribers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "subscriber_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "example" => "active",
                        "kind" => "query",
                        "name" => "state",
                        "orig" => "state",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "type",
                        "orig" => "type",
                        "reqd" => false,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/subscribers/count",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "count",
                  ],
                  "select" => {
                    "$action" => "count",
                    "exist" => [
                      "page_id",
                      "state",
                      "type",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/subscribers/{subscriber_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "subscriber_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 2,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/pages/{page_id}/subscribers/histogram_by_state",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "histogram_by_state",
                  ],
                  "select" => {
                    "$action" => "histogram_by_state",
                    "exist" => [
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 3,
                },
              ],
              "key$" => "load",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "incident_id",
                        "orig" => "incident_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 2,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "incidents",
                    "{incident_id}",
                    "subscribers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "subscriber_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "incident_id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "skip_unsubscription_notification",
                        "orig" => "skip_unsubscription_notification",
                        "reqd" => false,
                        "type" => "`$BOOLEAN`",
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/pages/{page_id}/subscribers/{subscriber_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "subscriber_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                      "skip_unsubscription_notification",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 1,
                },
              ],
              "key$" => "remove",
            },
            "update" => {
              "input" => "data",
              "name" => "update",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "subscriber_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "page_id",
                        "orig" => "page_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "PATCH",
                  "orig" => "/pages/{page_id}/subscribers/{subscriber_id}",
                  "parts" => [
                    "pages",
                    "{page_id}",
                    "subscribers",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "subscriber_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "page_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "update",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "page",
              ],
              [
                "page",
                "incident",
              ],
            ],
          },
        },
        "user" => {
          "fields" => [
            {
              "active" => true,
              "name" => "created_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 0,
            },
            {
              "active" => true,
              "name" => "email",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 1,
            },
            {
              "active" => true,
              "name" => "first_name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 2,
            },
            {
              "active" => true,
              "name" => "id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 3,
            },
            {
              "active" => true,
              "name" => "last_name",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 4,
            },
            {
              "active" => true,
              "name" => "organization_id",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 5,
            },
            {
              "active" => true,
              "name" => "updated_at",
              "req" => false,
              "type" => "`$STRING`",
              "index$" => 6,
            },
            {
              "active" => true,
              "name" => "user",
              "req" => true,
              "type" => "`$OBJECT`",
              "index$" => 7,
            },
          ],
          "name" => "user",
          "op" => {
            "create" => {
              "input" => "data",
              "name" => "create",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "organization_id",
                        "orig" => "organization_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                  },
                  "method" => "POST",
                  "orig" => "/organizations/{organization_id}/users",
                  "parts" => [
                    "organizations",
                    "{organization_id}",
                    "users",
                  ],
                  "select" => {
                    "exist" => [
                      "organization_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "create",
            },
            "list" => {
              "input" => "data",
              "name" => "list",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "organization_id",
                        "orig" => "organization_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                    ],
                    "query" => [
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "page",
                        "orig" => "page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                      {
                        "active" => true,
                        "kind" => "query",
                        "name" => "per_page",
                        "orig" => "per_page",
                        "reqd" => false,
                        "type" => "`$INTEGER`",
                      },
                    ],
                  },
                  "method" => "GET",
                  "orig" => "/organizations/{organization_id}/users",
                  "parts" => [
                    "organizations",
                    "{organization_id}",
                    "users",
                  ],
                  "select" => {
                    "exist" => [
                      "organization_id",
                      "page",
                      "per_page",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "list",
            },
            "remove" => {
              "input" => "data",
              "name" => "remove",
              "points" => [
                {
                  "active" => true,
                  "args" => {
                    "params" => [
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "id",
                        "orig" => "user_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 0,
                      },
                      {
                        "active" => true,
                        "kind" => "param",
                        "name" => "organization_id",
                        "orig" => "organization_id",
                        "reqd" => true,
                        "type" => "`$STRING`",
                        "index$" => 1,
                      },
                    ],
                  },
                  "method" => "DELETE",
                  "orig" => "/organizations/{organization_id}/users/{user_id}",
                  "parts" => [
                    "organizations",
                    "{organization_id}",
                    "users",
                    "{id}",
                  ],
                  "rename" => {
                    "param" => {
                      "user_id" => "id",
                    },
                  },
                  "select" => {
                    "exist" => [
                      "id",
                      "organization_id",
                    ],
                  },
                  "transform" => {
                    "req" => "`reqdata`",
                    "res" => "`body`",
                  },
                  "index$" => 0,
                },
              ],
              "key$" => "remove",
            },
          },
          "relations" => {
            "ancestors" => [
              [
                "organization",
              ],
            ],
          },
        },
      },
    }
  end


  def self.make_feature(name)
    require_relative 'features'
    StatuspageFeatures.make_feature(name)
  end
end
