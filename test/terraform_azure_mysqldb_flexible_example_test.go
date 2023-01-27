//go:build azure
// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

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
	expectedServerSkuName := "B_Standard_B1ms"
	expectedServerStoragemGb := "20"
	expectedDatabaseCharSet := "utf8"
	expectedDatabaseCollation := "utf8_unicode_ci"

	// website::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/azure/terraform-azure-mysqldb-example",
		Vars: map[string]interface{}{
			"postfix":                         uniquePostfix,
			"mysqlserver_flexible_sku_name":   expectedServerSkuName,
			"mysqlserver_flexible_storage_mb": expectedServerStoragemGb,
			"mysqldb_flexible_charset":        expectedDatabaseCharSet,
		},
	}

	// website::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// website::tag::3:: Run `terraform output` to get the values of output variables
	expectedResourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	expectedMYSQLServerFlexibleName := terraform.Output(t, terraformOptions, "mysql_server_name")

	expectedMYSQLDBFlexibleName := terraform.Output(t, terraformOptions, "mysql_database_name")

	// website::tag::4:: Get mySQL server details and assert them against the terraform output
	actualMYSQLFlexibleServer := azure.GetMySQLFlexibleServer(t, expectedResourceGroupName, expectedMYSQLServerFlexibleName, "")

	assert.Equal(t, expectedServerSkuName, *actualMYSQLServerFlexible.Sku.Name)
	assert.Equal(t, expectedServerStoragemGb, fmt.Sprint(*actualMYSQLFlexibleServer.ServerProperties.StorageProfile.StorageMB))

	assert.Equal(t, mysql.ServerStateReady, actualMYSQLFlexibleServer.ServerProperties.UserVisibleState)

	// website::tag::5:: Get  mySQL server DB details and assert them against the terraform output
	actualDatabase := azure.GetMyFlexibleSQLDB(t, expectedResourceGroupName, expectedMYSQLServerFlexibleName, expectedMYSQLDBFlexibleName, "")

	assert.Equal(t, expectedDatabaseCharSet, *actualDatabase.DatabaseProperties.Charset)
	assert.Equal(t, expectedDatabaseCollation, *actualDatabase.DatabaseProperties.Collation)
}
