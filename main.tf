terraform {
  required_providers {
    tf_bad_provider = {
      source  = "mdeous-datadog/tf-bad-provider"
      version = "1.0.0"
    }
  }
}

provider "tf_bad_provider" {
  address = "HOST:PORT"
  command = ""
}

resource "dummy_resource" "p" {}
