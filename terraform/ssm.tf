# e-comm microservices variables
resource "aws_ssm_parameter" "e_comm_microservices" {
  name  = "/ecomm/parameters"
  type  = "StringList"
  value = "string,string,string,string"
  lifecycle {
    ignore_changes = [value]
  }
}