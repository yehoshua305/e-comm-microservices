# oidc provider for IRSA
resource "aws_iam_openid_connect_provider" "irsa" {
  url             = var.issuer
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = [var.ca_thumbprint]
}

# IAM Role for IRSA
resource "aws_iam_role" "irsa" {
  name               = "K8SIRSARole"
  assume_role_policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Federated": ${aws_iam_openid_connect_provider.irsa.arn}
            },
            "Action": "sts:AssumeRoleWithWebIdentity"
        }
    ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "irsa" {
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
  role       = aws_iam_role.irsa.name
}