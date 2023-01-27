# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# ARM_CLIENT_ID
# ARM_CLIENT_SECRET
# ARM_SUBSCRIPTION_ID
# ARM_TENANT_ID

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "location" {
  description = "The Azure Region where the MySQL Flexible Server should exist."
  type        = string
  default     = "West Europe"
}

variable "mysql_server_administrator_login" {
  description = "The Administrator login for the MySQL Flexible Server."
  type        = string
  default     = "mysqlflexibleadmin"
}

variable "mysql_server_backup_retention_days" {
  description = "The backup retention days for the MySQL Flexible Server."
  type        = number
  default     = 7
}

variable "mysql_server_flexible_sku_name" {
  description = "The SKU Name for the MySQL Flexible Server"
  type        = string
  default     = "B_Standard_B1Ms"
}

variable "mysql_server_storage_size_gb" {
  description = "The max storage allowed for the MySQL Flexible Server."
  type        = string
  default     = "20"
}

variable "mysqldb_flexible_charset" {
  description = "Specifies the Charset for the MySQL Database."
  type        = string
  default     = "utf8"
}

variable "mysqldb_flexible_collation" {
  description = "Specifies the Collation for the MySQL Database"
  type        = string
  default     = "utf8_unicode_ci"
}

variable "postfix" {
  description = "A postfix string to centrally mitigate resource name collisions."
  type        = string
  default     = "example"
}