resource "aws_lambda_function" "my_lambda" {
  function_name = "my_lambda_function"
  s3_bucket     = var.lambda_s3_bucket
  s3_key        = var.lambda_s3_key
  handler       = "main.handler"
  runtime       = "python3.8"
  role          = aws_iam_role.lambda_exec.arn

  environment {
    variables = {
      REDIS_HOST = "your_redis_host"
    }
  }
}

resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
