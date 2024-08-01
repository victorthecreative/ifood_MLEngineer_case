resource "aws_s3_bucket" "my_bucket" {
  bucket = "my-bucket"

  versioning {
    enabled = true
  }

  tags = {
    Name = "my-bucket"
  }
}
