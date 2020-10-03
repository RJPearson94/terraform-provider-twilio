resource "twilio_account_address" "address" {
  account_sid   = var.account_sid
  customer_name = var.customer_name
  street        = var.street
  city          = var.city
  region        = var.region
  postal_code   = var.postal_code
  iso_country   = var.iso_country
}
