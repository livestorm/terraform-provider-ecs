package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{},
                DataSourcesMap: map[string]*schema.Resource{
                  "ecs_image_tag": dataImageTag(),
                },
        }
}
