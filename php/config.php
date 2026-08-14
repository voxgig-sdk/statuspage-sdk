<?php
declare(strict_types=1);

// Statuspage SDK configuration

class StatuspageConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "Statuspage",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://api.statuspage.io/v1",
                "auth" => [
                    "prefix" => "OAuth",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "component" => [],
                    "component_group_uptime" => [],
                    "group_component" => [],
                    "incident" => [],
                    "incident_postmortem" => [],
                    "incident_subscriber" => [],
                    "incident_template" => [],
                    "incident_update" => [],
                    "metric" => [],
                    "metrics_provider" => [],
                    "page" => [],
                    "page_access_group" => [],
                    "page_access_user" => [],
                    "permission" => [],
                    "postmortem" => [],
                    "status_embed_config" => [],
                    "subscriber" => [],
                    "user" => [],
                ],
            ],
            "entity" => [
        'component' => [
          'fields' => [
            [
              'name' => 'automation_email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'component',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'group',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'group_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'only_show_if_degraded',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'position',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'showcase',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'start_date',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'component',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_groups',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_group',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_users',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_user',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_access_group_id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{page_access_group_id}',
                    'components',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_access_group_id',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_access_user_id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{page_access_user_id}',
                    'components',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_access_user_id',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'end',
                        'orig' => 'end',
                        'type' => 'Any',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => 'Any',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components/{component_id}/uptime',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'uptime',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'uptime',
                    'exist' => [
                      'end',
                      'id',
                      'page_id',
                      'start',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.related_events`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_groups',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_group',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/components/{component_id}/page_access_users',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'page_access_user',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/components/{component_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'components',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'component_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'component' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'page_access_group',
              ],
              [
                'page',
                'page_access_user',
              ],
            ],
          ],
        ],
        'component_group_uptime' => [
          'fields' => [
            [
              'name' => 'component_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incidents',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'component_group_uptime',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'end',
                        'orig' => 'end',
                        'type' => 'Any',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'start',
                        'orig' => 'start',
                        'type' => 'Any',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/component-groups/{id}/uptime',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                    'uptime',
                  ],
                  'select' => [
                    'exist' => [
                      'end',
                      'id',
                      'page_id',
                      'start',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.related_events`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'group_component' => [
          'fields' => [
            [
              'name' => 'component_group',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'components',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'position',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'group_component',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/component-groups',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/component-groups',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/component-groups/{id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'component-groups',
                    '{id}',
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident' => [
          'fields' => [
            [
              'name' => 'auto_transition_deliver_notifications_at_end',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'auto_transition_deliver_notifications_at_start',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'auto_transition_to_maintenance_state',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'auto_transition_to_operational_state',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'components',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'impact',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'impact_override',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident',
              'op' => [
                'patch' => [
                  'type' => '`$OBJECT`',
                ],
                'update' => [
                  'type' => '`$OBJECT`',
                ],
              ],
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'incident_updates',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'metadata',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'monitoring_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'postmortem_body',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'postmortem_body_last_updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'postmortem_ignored',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem_notified_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem_notified_twitter',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem_published_at',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'reminder_intervals',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'resolved_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'scheduled_auto_completed',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'scheduled_auto_in_progress',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'scheduled_for',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'scheduled_remind_prior',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'scheduled_reminded_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'scheduled_until',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'shortlink',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'incident',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incidents',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                  ],
                  'select' => [
                    'exist' => [
                      'limit',
                      'page',
                      'page_id',
                      'q',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/active_maintenance',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'active_maintenance',
                  ],
                  'select' => [
                    '$action' => 'active_maintenance',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/scheduled',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'scheduled',
                  ],
                  'select' => [
                    '$action' => 'scheduled',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/unresolved',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'unresolved',
                  ],
                  'select' => [
                    '$action' => 'unresolved',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/upcoming',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    'upcoming',
                  ],
                  'select' => [
                    '$action' => 'upcoming',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident_postmortem' => [
          'fields' => [],
          'name' => 'incident_postmortem',
          'op' => [
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{id}',
                    'postmortem',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident_subscriber' => [
          'fields' => [],
          'name' => 'incident_subscriber',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'subscriber_id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{subscriber_id}',
                    'resend_confirmation',
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                      'subscriber_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
                'incident',
                'subscriber',
              ],
            ],
          ],
        ],
        'incident_template' => [
          'fields' => [
            [
              'name' => 'body',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'components',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'group_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'should_send_notifications',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'should_tweet',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'template',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'title',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'update_status',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'incident_template',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incident_templates',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incident_templates',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 1,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 100,
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incident_templates',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incident_templates',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'incident_update' => [
          'fields' => [
            [
              'name' => 'affected_components',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'body',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'custom_tweet',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'deliver_notifications',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'display_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident_update',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'status',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'tweet_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'wants_twitter_update',
              'type' => '`$BOOLEAN`',
            ],
          ],
          'name' => 'incident_update',
          'op' => [
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_update_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'incident_updates',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_update_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident_update' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'incident_update_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'incident_updates',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'incident_update_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'incident_update' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
                'incident',
              ],
            ],
          ],
        ],
        'metric' => [
          'fields' => [
            [
              'name' => 'backfill_percentage',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'backfilled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'data',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'decimal_places',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'display',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last_fetched_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'metric_identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metrics_provider_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'most_recent_data_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'reference_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'suffix',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'tooltip_description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'y_axis_hidden',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'y_axis_max',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'y_axis_min',
              'type' => '`$NUMBER`',
            ],
          ],
          'name' => 'metric',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}/data',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                    'data',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'metrics_provider_id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{metrics_provider_id}',
                    'metrics',
                  ],
                  'select' => [
                    'exist' => [
                      'metrics_provider_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics/data',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    'data',
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_access_user_id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{page_access_user_id}',
                    'metrics',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_access_user_id',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'metrics_provider_id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{metrics_provider_id}',
                    'metrics',
                  ],
                  'select' => [
                    'exist' => [
                      'metrics_provider_id',
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                  ],
                  'select' => [
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}/data',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                    'data',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'data',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/metrics/{metric_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metric_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metric' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'metrics_provider',
              ],
              [
                'page',
                'page_access_user',
              ],
            ],
          ],
        ],
        'metrics_provider' => [
          'fields' => [
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'disabled',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last_revalidated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric_base_uri',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metrics_provider',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'type',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'metrics_provider',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/metrics_providers',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'metrics_provider_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/metrics_providers/{metrics_provider_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'metrics_providers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'metrics_provider_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'metrics_provider' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'page' => [
          'fields' => [
            [
              'name' => 'activity_score',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'allow_email_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_incident_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_page_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_rss_atom_feeds',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_sms_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'allow_webhook_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'branding',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'city',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'country',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_blues',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_body_background_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_border_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_font_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_graph_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_greens',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_light_font_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_link_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_no_data',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_oranges',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_reds',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'css_yellows',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'domain',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'favicon_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'headline',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'hero_cover',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'hidden_from_search',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'ip_restrictions',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'notifications_email_footer',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'notifications_from_email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_description',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'state',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'subdomain',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'support_url',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'time_zone',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'transactional_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_logo',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'twitter_username',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'url',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'viewers_must_be_team_members',
              'type' => '`$BOOLEAN`',
            ],
          ],
          'name' => 'page',
          'op' => [
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages',
                  'parts' => [
                    'pages',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}',
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}',
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}',
                  'parts' => [
                    'pages',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'page_access_group' => [
          'fields' => [
            [
              'name' => 'component_ids',
              'op' => [
                'create' => [
                  'req' => true,
                  'type' => '`$ARRAY`',
                ],
              ],
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'external_identifier',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric_ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_group',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_access_user_ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'page_access_group',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_groups',
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups',
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_groups',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'component_id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                    '{component_id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'component_id',
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_group' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_group_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_groups/{page_access_group_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_groups',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_group_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'component',
              ],
            ],
          ],
        ],
        'page_access_user' => [
          'fields' => [
            [
              'name' => 'component_ids',
              'req' => true,
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'external_login',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'metric_ids',
              'req' => true,
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'page_access_group_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_group_ids',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_user',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'page_access_user',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/page_access_users',
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'page_access_user' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'email',
                        'orig' => 'email',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users',
                  'parts' => [
                    'pages',
                    '{id}',
                    'page_access_users',
                  ],
                  'rename' => [
                    'param' => [
                      'page_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'email',
                      'id',
                      'page',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'component_id',
                        'orig' => 'component_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                    '{component_id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'component_id',
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'metric_id',
                        'orig' => 'metric_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                    '{metric_id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'metric_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/components',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'components',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'component',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'page_access_user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/page_access_users/{page_access_user_id}/metrics',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'page_access_users',
                    '{id}',
                    'metrics',
                  ],
                  'rename' => [
                    'param' => [
                      'page_access_user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'metric',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'component',
              ],
              [
                'page',
                'metric',
              ],
            ],
          ],
        ],
        'permission' => [
          'fields' => [
            [
              'name' => 'pages',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'user_id',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'permission',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/organizations/{organization_id}/permissions/{user_id}',
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'permissions',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/organizations/{organization_id}/permissions/{user_id}',
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'permissions',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.data`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'organization',
              ],
            ],
          ],
        ],
        'postmortem' => [
          'fields' => [
            [
              'name' => 'body',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'body_draft',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'body_draft_updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'body_updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'custom_tweet',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'notify_subscribers',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'notify_twitter',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'postmortem',
              'op' => [
                'update' => [
                  'type' => '`$OBJECT`',
                ],
              ],
              'req' => true,
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'preview_key',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'published_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'postmortem',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'postmortem' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem/publish',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                    'publish',
                  ],
                  'select' => [
                    '$action' => 'publish',
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'postmortem' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/postmortem/revert',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'postmortem',
                    'revert',
                  ],
                  'select' => [
                    '$action' => 'revert',
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
                'incident',
              ],
            ],
          ],
        ],
        'status_embed_config' => [
          'fields' => [
            [
              'name' => 'incident_background_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'incident_text_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'maintenance_background_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'maintenance_text_color',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'position',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'status_embed_config',
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'status_embed_config',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/status_embed_config',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'patch' => [
              'input' => 'data',
              'name' => 'patch',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/status_embed_config',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'status_embed_config' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PUT',
                  'orig' => '/pages/{page_id}/status_embed_config',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'status_embed_config',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'status_embed_config' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
            ],
          ],
        ],
        'subscriber' => [
          'fields' => [
            [
              'name' => 'component_ids',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'components',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'display_phone_number',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'endpoint',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'integration_partner',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'mode',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'obfuscated_channel_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'page_access_user_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'phone_country',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'phone_number',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'purge_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'quarantined_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'skip_confirmation_notification',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'skip_unsubscription_notification',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'slack',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'sms',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'state',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'subscriber',
              'type' => '`$OBJECT`',
            ],
            [
              'name' => 'subscribers',
              'req' => true,
              'type' => '`$STRING`',
            ],
            [
              'name' => 'teams',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'type',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'webhook',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'workspace_name',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'subscriber',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                    'resend_confirmation',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'subscriber' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                  ],
                  'select' => [
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'subscriber' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/reactivate',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'reactivate',
                  ],
                  'select' => [
                    '$action' => 'reactivate',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/resend_confirmation',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'resend_confirmation',
                  ],
                  'select' => [
                    '$action' => 'resend_confirmation',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/pages/{page_id}/subscribers/unsubscribe',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'unsubscribe',
                  ],
                  'select' => [
                    '$action' => 'unsubscribe',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'limit',
                        'orig' => 'limit',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'example' => 0,
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'q',
                        'orig' => 'q',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'asc',
                        'kind' => 'query',
                        'name' => 'sort_direction',
                        'orig' => 'sort_direction',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'primary',
                        'kind' => 'query',
                        'name' => 'sort_field',
                        'orig' => 'sort_field',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'active',
                        'kind' => 'query',
                        'name' => 'state',
                        'orig' => 'state',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                  ],
                  'select' => [
                    'exist' => [
                      'limit',
                      'page',
                      'page_id',
                      'q',
                      'sort_direction',
                      'sort_field',
                      'state',
                      'type',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                  ],
                  'select' => [
                    'exist' => [
                      'incident_id',
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/unsubscribed',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'unsubscribed',
                  ],
                  'select' => [
                    '$action' => 'unsubscribed',
                    'exist' => [
                      'page',
                      'page_id',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 'active',
                        'kind' => 'query',
                        'name' => 'state',
                        'orig' => 'state',
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/count',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'count',
                  ],
                  'select' => [
                    '$action' => 'count',
                    'exist' => [
                      'page_id',
                      'state',
                      'type',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/pages/{page_id}/subscribers/histogram_by_state',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    'histogram_by_state',
                  ],
                  'select' => [
                    '$action' => 'histogram_by_state',
                    'exist' => [
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'incident_id',
                        'orig' => 'incident_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'incidents',
                    '{incident_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'incident_id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'skip_unsubscription_notification',
                        'orig' => 'skip_unsubscription_notification',
                        'type' => '`$BOOLEAN`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                      'skip_unsubscription_notification',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'update' => [
              'input' => 'data',
              'name' => 'update',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'subscriber_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'page_id',
                        'orig' => 'page_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'PATCH',
                  'orig' => '/pages/{page_id}/subscribers/{subscriber_id}',
                  'parts' => [
                    'pages',
                    '{page_id}',
                    'subscribers',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'subscriber_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'page_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'page',
              ],
              [
                'page',
                'incident',
              ],
            ],
          ],
        ],
        'user' => [
          'fields' => [
            [
              'name' => 'created_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'email',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'first_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'last_name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'organization_id',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'updated_at',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'user',
              'req' => true,
              'type' => '`$OBJECT`',
            ],
          ],
          'name' => 'user',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/organizations/{organization_id}/users',
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                  ],
                  'select' => [
                    'exist' => [
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => [
                      'user' => '`reqdata`',
                    ],
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'list' => [
              'input' => 'data',
              'name' => 'list',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'kind' => 'query',
                        'name' => 'page',
                        'orig' => 'page',
                        'type' => '`$INTEGER`',
                      ],
                      [
                        'kind' => 'query',
                        'name' => 'per_page',
                        'orig' => 'per_page',
                        'type' => '`$INTEGER`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/organizations/{organization_id}/users',
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                  ],
                  'select' => [
                    'exist' => [
                      'organization_id',
                      'page',
                      'per_page',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'remove' => [
              'input' => 'data',
              'name' => 'remove',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'user_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'kind' => 'param',
                        'name' => 'organization_id',
                        'orig' => 'organization_id',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'DELETE',
                  'orig' => '/organizations/{organization_id}/users/{user_id}',
                  'parts' => [
                    'organizations',
                    '{organization_id}',
                    'users',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'user_id' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                      'organization_id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'organization',
              ],
            ],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return StatuspageFeatures::make_feature($name);
    }
}
