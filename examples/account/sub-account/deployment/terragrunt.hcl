terraform {
  extra_arguments "init_args" {
    commands = [
      "init"
    ]

    arguments = [
      "-upgrade=true",
    ]
  }
}

inputs = {
  prefix = "RJPearson94"
}