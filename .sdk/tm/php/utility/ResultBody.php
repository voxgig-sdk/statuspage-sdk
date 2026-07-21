<?php
declare(strict_types=1);

// Statuspage SDK utility: result_body

class StatuspageResultBody
{
    public static function call(StatuspageContext $ctx): ?StatuspageResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
