######################
# Container Registry #
######################

# api
data "template_file" "ecomm" {
  template = file("./policies/ecr-policy.json")
  vars = {
    account_id        = local.account_id
  }
}

resource "aws_ecr_repository" "ecomm" {
  for_each = toset(["userservice"])
  name                 = "ecomm/${each.value}"
  # image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository_policy" "ecomm" {
  for_each = toset([ "userservice" ])
  repository = aws_ecr_repository.ecomm[each.value].name
  policy     = data.template_file.ecomm.rendered
}