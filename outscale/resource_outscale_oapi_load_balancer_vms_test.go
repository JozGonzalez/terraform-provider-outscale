package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	oscgo "github.com/marinsalinas/osc-sdk-go"
)

func TestAccOutscaleOAPILBUAttachment_basic(t *testing.T) {
	var conf oscgo.LoadBalancer

	testCheckInstanceAttached := func(count int) resource.TestCheckFunc {
		return func(*terraform.State) error {
			if len(*conf.BackendVmIds) != count {
				return fmt.Errorf("backend_vm_id count does not match")
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "outscale_load_balancer.bar",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckOutscaleOAPILBUDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleOAPILBUAttachmentConfig1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPILBUExists("outscale_load_balancer.bar", &conf),
					testCheckInstanceAttached(1),
				),
			},
		},
	})
}

// add one attachment
const testAccOutscaleOAPILBUAttachmentConfig1 = `
resource "outscale_load_balancer" "bar" {
	load_balancer_name = "load-test"
	
	availability_zones = ["eu-west-2a"]
    listeners {
    backend_port = 8000
    backend_protocol = "HTTP"
    load_balancer_port = 80
    load_balancer_protocol = "HTTP"
  }
}

resource "outscale_vm" "foo1" {
  image_id = "ami-8a6a0120"
	type = "t2.micro"
}

resource "outscale_load_balancer_vms" "foo1" {
  load_balancer_name      = "${outscale_load_balancer.bar.id}"
  backend_vm_id = ["${outscale_vm.foo1.id}"]
}
`
