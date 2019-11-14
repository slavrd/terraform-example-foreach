package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformOutput(t *testing.T) {
	input := make(map[string]int)
	input["pet1"] = 2
	input["pet2"] = 4
	input["pet3"] = 7

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"pets": input,
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	defer terraform.Destroy(t, terraformOptions)

	_, err := terraform.InitAndApplyE(t, terraformOptions)

	if err != nil {
		t.Fatalf("terraform init/apply failed: %v", err)
	}

	petsNames, err := terraform.OutputListE(t, terraformOptions, "pets")

	if err != nil {
		t.Fatalf("error reading terraform output 'pets': %v", err)
	}

	if len(petsNames) != len(input) {
		t.Fatalf("terraform output pets has unexpected length, want %d, got %d", len(input), len(petsNames))
	}

	for _, pet := range petsNames {
		splitName := strings.Split(pet, "-")
		if input[splitName[0]] != len(splitName)-1 {
			t.Fatalf("unexpected name, want prefix: %s and length: %d, got: %s", splitName[0], len(splitName)-1, pet)
		}
	}
}
