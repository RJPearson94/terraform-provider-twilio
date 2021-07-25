resource "twilio_credentials_public_key" "public_key" {
  friendly_name = "Test Public Key Resource"
  public_key    = var.public_key
}
