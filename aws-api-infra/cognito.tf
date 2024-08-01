resource "aws_cognito_user_pool" "user_pool" {
  name = "user-pool"
}

resource "aws_cognito_user_pool_client" "user_pool_client" {
  name         = "user-pool-client"
  user_pool_id = aws_cognito_user_pool.user_pool.id
}

resource "aws_cognito_identity_pool" "identity_pool" {
  identity_pool_name               = "identity-pool"
  allow_unauthenticated_identities = false

  cognito_identity_providers {
    client_id   = aws_cognito_user_pool_client.user_pool_client.id
    provider_name = aws_cognito_user_pool.user_pool.endpoint
  }
}
