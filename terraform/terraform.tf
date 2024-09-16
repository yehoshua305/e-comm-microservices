terraform {

  backend "s3" {
    bucket         = "maverick305-tfstate-bucket"
    region         = "eu-west-1"
    key            = "ecomm-microservices-statefile"
    dynamodb_table = "github-actions-tf-state-lock"
  }
}
