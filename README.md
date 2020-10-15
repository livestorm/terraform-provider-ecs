# terraform-provider-ecs

This custom provider is used like this:

```
data "ecs_image_tag" "webapp_image_tag" {
  cluster = "livestorm-devops-ecs-cluster"
  service = "livestorm-devops-webapp"
}
```

If the cluster & service exist, then the current image used in the taskdefinition is outputed as the `image_tag` attribute, if not, then `dev` is outputed.
It ca be used like this: `data.ecs_image_tag.webapp_image_tag.image_tag`

## How to build it ?

After installing golang on your system, run `go build -o terraform-provider-ecs`.

## How to install it ?

Copy the binary `terraform-provider-ecs` in the `terraform.d/plugins/linux_amd64` directory of your terraform module.
NOTE: If you are on macOS the architecture is `darwin_amd64` instead of `linux_amd64`.

## Warning

This is a hack used in order to be able to deploy changes to taskdefinitions (for staging environment) without having to rollback the current images tag to dev.
It's not meant to be used with taskdefinitions that use many containers with different tags as it will always only read the first container image tag.
