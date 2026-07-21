<?php
declare(strict_types=1);

// Statuspage SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class StatuspageMakeContext
{
    public static function call(array $ctxmap, ?StatuspageContext $basectx): StatuspageContext
    {
        return new StatuspageContext($ctxmap, $basectx);
    }
}
