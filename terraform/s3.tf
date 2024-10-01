# Public Issuer

resource "random_string" "random" {
  length           = 8
  special          = false
  override_special = "/@£$"
}

module "issuer" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket                   = lower(random_string.random.result)
  acl                      = "public-read"
  block_public_acls        = false
  block_public_policy      = false
  ignore_public_acls       = false
  restrict_public_buckets  = false
  control_object_ownership = true
  object_ownership         = "BucketOwnerPreferred"
  versioning = {
    enabled = true
  }
}