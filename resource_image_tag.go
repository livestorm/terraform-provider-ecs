package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/ecs"
        "strings"
        "fmt"
)

func dataImageTag() *schema.Resource {
        return &schema.Resource{
                // Create: dataImageTagCreate,
                Read:   dataImageTagRead,
                // Update: dataImageTagUpdate,
                // Delete: dataImageTagDelete,

                Schema: map[string]*schema.Schema{
                        "cluster": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "service": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "image_tag": {
                        				Type:     schema.TypeString,
                        				Computed: true,
                  			},
                },
        }
}

// func dataImageTagCreate(d *schema.ResourceData, m interface{}) error {
//         return dataImageTagRead(d, m)
// }

func dataImageTagRead(d *schema.ResourceData, m interface{}) error {
  cluster_name := d.Get("cluster").(string)
  service_name := d.Get("service").(string)
  d.SetId(service_name)

  sess, err := session.NewSession(&aws.Config{
      Region: aws.String("eu-west-1")},
  )
  if (err != nil) {
    fmt.Print("AWS Session Error")
    return nil
  }

  // Create ECS service client
  client := ecs.New(sess)

  services := []*string{&service_name}

  req_service, service := client.DescribeServicesRequest(&ecs.DescribeServicesInput{
      Cluster: &cluster_name,
      Services: services,
  })
  req_service.Send()
  if (len(service.Services) == 0) {
    d.Set("image_tag", "dev")
    return nil
  }

  taskdefinition_name := *service.Services[0].TaskDefinition
  // fmt.Println("> TaskDefinition is", taskdefinition_name)

  req_taskdefinition, taskdefinition := client.DescribeTaskDefinitionRequest(&ecs.DescribeTaskDefinitionInput{
      TaskDefinition: &taskdefinition_name,
  })
  req_taskdefinition.Send()
  image_url := *taskdefinition.TaskDefinition.ContainerDefinitions[0].Image
  d.Set("image_tag", strings.Split(image_url, ":")[1])
  return nil
}

// func dataImageTagUpdate(d *schema.ResourceData, m interface{}) error {
//         return dataImageTagRead(d, m)
// }
//
// func dataImageTagDelete(d *schema.ResourceData, m interface{}) error {
//         return nil
// }
