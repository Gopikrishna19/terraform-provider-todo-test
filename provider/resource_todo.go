package provider

import (
	"github.com/Gopikrishna19/terraform-provider-todo-test/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTodo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"task": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"completed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
		Read: resourceGetTodos,
	}
}

func resourceGetTodos(data *schema.ResourceData, i interface{}) error {
	c := i.(client.Client)

	todos, err := c.Get()

	if err != nil {
		return err
	}

	_ = data.Set("todos", todos)
	return nil
}
