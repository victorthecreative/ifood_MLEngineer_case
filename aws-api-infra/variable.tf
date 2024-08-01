variable "aws_region" {
  description = "The AWS region to deploy the infrastructure"
  default     = "us-east-1"
}

variable "domain_name" {
  description = "The domain name for the API"
}

variable "api_name" {
  description = "The name of the API"
  default     = "my-api"
}

variable "lambda_s3_bucket" {
  description = "S3 bucket to store Lambda function code"
}

variable "lambda_s3_key" {
  description = "S3 key for the Lambda function code zip file"
}
