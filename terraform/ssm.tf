# e-comm microservices variables
resource "aws_ssm_parameter" "symmetric_key" {
  name  = "SymmetricKey"
  type  = "String"
  value = ""
  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "access_token_duration" {
  name  = "AccessTokenDuration"
  type  = "String"
  value = ""
  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "refresh_token_duration" {
  name  = "RefreshTokenDuration"
  type  = "String"
  value = ""
  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "server_address" {
  name  = "ServerAddress"
  type  = "String"
  value = ""
  lifecycle {
    ignore_changes = [value]
  }
}
