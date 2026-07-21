<?php
declare(strict_types=1);

// Statuspage SDK utility: prepare_body

class StatuspagePrepareBody
{
    public static function call(StatuspageContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
