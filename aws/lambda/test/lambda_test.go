package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestLambda(t *testing.T) {

	functionName := fmt.Sprintf("function-%s", uuid.New().String())
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",

		Vars: map[string]interface{}{
			"function_name": functionName,
			"aws_region": awsRegion,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	response := aws.InvokeFunction(t, awsRegion, functionName, AddRequest{A: 2, B: 3})

	assert.Equal(t, "5", string(response))
}

type AddRequest struct {
	A int
	B int
}