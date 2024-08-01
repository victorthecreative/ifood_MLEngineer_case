output "api_gateway_url" {
  value = aws_api_gateway_rest_api.my_api.invoke_url
}

output "cloudfront_domain_name" {
  value = aws_cloudfront_distribution.my_distribution.domain_name
}
