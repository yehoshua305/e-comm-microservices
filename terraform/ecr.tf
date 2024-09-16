######################
# Container Registry #
######################

# api
data "template_file" "ecomm" {
  template = file("./policies/ecr-cross-account-policy.json")
  vars = {
    account_id        = local.account_id
  }
}

resource "aws_ecr_repository" "ecomm" {
  name                 = "api/ecomm"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository_policy" "ecomm" {
  repository = aws_ecr_repository.ecomm_.name
  policy     = data.template_file.ecomm.rendered
}