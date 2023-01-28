package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/mysql/mgmt/mysqlflexibleservers"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/supertrampr/modules/azure"
)

func TestTerraformAzureMySQLFlexibleDBExample(t *testing.T) {
	t.Parallel()

	uniquePostfix := strings.ToLower(random.UniqueId())
	expectedServerSkuName := "Standard_B1ms"
	expectedServerStoragemGb := "20"
	expectedDatabaseCharSet := "utf8"
	expectedDatabaseCollation := "utf8_unicode_ci"

	// website::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-mysqlflexibledb-example",
		Vars: map[string]interface{}{
			"postfix":                        uniquePostfix,
			"mysql_server_flexible_sku_name": "B_Standard_B1ms",
			"mysql_server_storage_size_gb":   expectedServerStoragemGb,
			"mysqldb_flexible_charset":       expectedDatabaseCharSet,
		},
	}

	// website::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// website::tag::3:: Run `terraform output` to get the values of output variables
	expectedResourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	expectedMYSQLServerFlexibleName := terraform.Output(t, terraformOptions, "mysql_server_flexible_name")

	expectedMYSQLDBFlexibleName := terraform.Output(t, terraformOptions, "mysql_database_flexible_name")

	// website::tag::4:: Get mySQL server details and assert them against the terraform output
	actualMYSQLFlexibleServer := azure.GetMySQLServer(t, expectedResourceGroupName, expectedMYSQLServerFlexibleName, "")

	assert.Equal(t, expectedServerSkuName, *actualMYSQLFlexibleServer.Sku.Name)
	assert.Equal(t, expectedServerStoragemGb, fmt.Sprint(*actualMYSQLFlexibleServer.ServerProperties.Storage.StorageSizeGB))

	assert.Equal(t, mysqlflexibleservers.ServerStateReady, actualMYSQLFlexibleServer.ServerProperties.State)

	// website::tag::5:: Get  mySQL server DB details and assert them against the terraform output
	actualDatabase := azure.GetMySQLFlexibleDB(t, expectedResourceGroupName, expectedMYSQLServerFlexibleName, expectedMYSQLDBFlexibleName, "")

	assert.Equal(t, expectedDatabaseCharSet, *actualDatabase.DatabaseProperties.Charset)
	assert.Equal(t, expectedDatabaseCollation, *actualDatabase.DatabaseProperties.Collation)
}
