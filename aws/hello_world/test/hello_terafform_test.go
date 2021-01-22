package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"net/http"
	"testing"
	"time"
)

func TestAwsHelloTerraform(t *testing.T) {
	options := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
	})

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)

	ip := terraform.Output(t, options, "public_ip")
	url := fmt.Sprintf("http://%s:8888", ip)
	http_helper.HttpGetWithRetry(t, url, nil, http.StatusOK, "Hello Terraform!", 50, 5 * time.Second)
}