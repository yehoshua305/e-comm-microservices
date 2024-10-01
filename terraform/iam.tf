resource "aws_iam_openid_connect_provider" "irsa" {
  url             = var.ISSUER
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = [var.CA_THUMBPRINT]
}