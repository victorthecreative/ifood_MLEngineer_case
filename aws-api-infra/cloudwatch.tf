resource "aws_cloudwatch_log_group" "my_log_group" {
  name              = "/aws/lambda/my_lambda_function"
  retention_in_days = 14
}

resource "aws_cloudwatch_metric_alarm" "my_alarm" {
  alarm_name          = "lambda_error_alarm"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "Errors"
  namespace           = "AWS/Lambda"
  period              = "60"
  statistic           = "Sum"
  threshold           = "1"
  alarm_description   = "Alarm if Lambda function has errors"
  actions_enabled     = true

  alarm_actions = [
    "arn:aws:sns:us-east-1:123456789012:my-sns-topic",
  ]

  dimensions = {
    FunctionName = aws_lambda_function.my_lambda.function_name
  }
}
