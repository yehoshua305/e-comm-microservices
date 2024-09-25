# Public Issuer

resource "random_string" "random" {
  length           = 8
  special          = false
  override_special = "/@£$"
}

module "issuer" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket = lower(random_string.random.result)
  acl    = "public-read"

  versioning = {
    enabled = true
  }
}