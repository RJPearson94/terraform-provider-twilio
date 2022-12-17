terraform {
  source = "../..//modules/sub-account"
}

include {
  path = find_in_parent_folders()
}
