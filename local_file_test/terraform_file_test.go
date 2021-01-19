package test

import (
	"io/ioutil"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformFile(t *testing.T) {
	fileName := "helloWorld.txt"

	options := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../local_file",
		Vars: map[string]interface{}{
			"file_name": fileName,
		},
		VarFiles:     []string{"varfile.tfvars"},
		NoColor:      true,
	})

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)

	humanOutput := terraform.Output(t, options, "human")
	magicTupleOutput := terraform.Output(t, options, "magic_tuple")
	filePathOutput := terraform.Output(t, options, "file_path")

	expectedHuman := "map[age:22 name:Kuba]"
	expectedMagicTuple := "[terraform 2021 true]"
	expectedContent := "Hello World!"

	assert.Equal(t, expectedHuman, humanOutput)
	assert.Equal(t, expectedMagicTuple, magicTupleOutput)

	file, err := ioutil.ReadFile(filePathOutput + "/" + fileName)
	assert.NoError(t, err)

	content := string(file)
	assert.Equal(t, content, expectedContent)
}