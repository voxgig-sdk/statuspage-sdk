# Statuspage SDK utility: feature_add
module StatuspageUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
