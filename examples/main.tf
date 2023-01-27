# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY AN AZURE Database For MySQL - Flexible Server
# This is an example of how to deploy an Azure Database Mysql - Flexible Server .
# See test/terraform_azure_example_test.go for how to write automated tests for this code.
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = ">=1.0.0, <2.0.0"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.41.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.4.3"
    }
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# CONFIGURE OUR AZURE CONNECTION
# ---------------------------------------------------------------------------------------------------------------------

provider "azurerm" {
  features {}
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A RESOURCE GROUP
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "example" {
  name     = "terratest-mysqlflexible-${var.postfix}"
  location = var.location
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY AZURE MySQL FLEXIBLE SERVER
# ---------------------------------------------------------------------------------------------------------------------

# Random password is used as an example to simplify the deployment and improve the security of the database.
# This is not as a production recommendation as the password is stored in the Terraform state file.
resource "random_password" "password" {
  length           = 16
  override_special = "_%@"
  min_upper        = "1"
  min_lower        = "1"
  min_numeric      = "1"
  min_special      = "1"
}

resource "azurerm_mysql_flexible_server" "example" {
  location            = var.location
  name                = "mysqlserver-flexible-${var.postfix}"
  resource_group_name = azurerm_resource_group.example.name

  administrator_login    = var.mysql_server_administrator_login
  administrator_password = random_password.password.result

  backup_retention_days = var.mysql_server_backup_retention_days
  sku_name              = var.mysql_server_flexible_sku_name
  version               = "5.7"

  storage {
    auto_grow_enabled = false
    size_gb           = var.mysql_server_storage_size_gb
  }
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY AZURE MySQL DATABASE
# ---------------------------------------------------------------------------------------------------------------------
resource "azurerm_mysql_flexible_database" "example" {
  name                = "mysqldb-flexible-${var.postfix}"
  resource_group_name = azurerm_resource_group.example.name
  server_name         = azurerm_mysql_flexible_server.example.name

  charset   = var.mysqldb_flexible_charset
  collation = var.mysqldb_flexible_collation
}