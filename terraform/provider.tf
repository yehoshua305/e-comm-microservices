#####################################
# Provider settings for AWS Provider
#####################################
provider "aws" {
  region = "eu-west-1"
  default_tags {
    tags = {
      Application = "e-comm-microservices"
      GitRepo     = "e-comm-microservices"
    }
  }
}
