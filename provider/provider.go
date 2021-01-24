package provider

import (
	"context"

	"github.com/Gopikrishna19/terraform-provider-todo-test/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
				Default:  "http://localhost",
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8080,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"todo": resourceTodo(),
		},
		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	port := d.Get("port").(int)

	return client.NewClient(host, port), nil
}
