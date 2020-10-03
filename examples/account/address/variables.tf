variable "account_sid" {
  description = "The account SID to associate the address with"
  type        = string
}

variable "customer_name" {
  description = "The customer/ business name"
  type        = string
}

variable "street" {
  description = "The address street"
  type        = string
}

variable "city" {
  description = "The address city"
  type        = string
}

variable "region" {
  description = "The address region"
  type        = string
}

variable "postal_code" {
  description = "The address postal code"
  type        = string
}

variable "iso_country" {
  description = "The address ISO country"
  type        = string
}
