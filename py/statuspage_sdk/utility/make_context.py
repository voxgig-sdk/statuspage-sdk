# Statuspage SDK utility: make_context

from statuspage_sdk.core.context import StatuspageContext


def make_context_util(ctxmap, basectx):
    return StatuspageContext(ctxmap, basectx)
