resource "aws_iam_openid_connect_provider" "irsa" {
  url             = var.issuer
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = [var.ca_thumbprint]
}