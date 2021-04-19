resource "twilio_phone_number" "phone_number" {
  account_sid = data.twilio_account_details.account_details.sid

  search_criteria {
    type        = "mobile"
    iso_country = "GB"

    exclude_address_requirements {
      all = true
    }
  }
}
