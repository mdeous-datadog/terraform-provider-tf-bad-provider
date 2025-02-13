terraform {
  required_providers {
    tf_bad_provider = {
      source  = "mdeous-datadog/tf-bad-provider"
      version = "1.0.8"
    }
  }
}

provider "tf_bad_provider" {
  # address = "toolbox.p.ddtdg.com:4400"
  # command = "uname -a"
}

resource "dummy_resource" "p" {}
