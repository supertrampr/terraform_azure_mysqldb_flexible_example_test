output "resource_group_name" {
  value = azurerm_resource_group.example.name
}

output "mysql_server_flexible_name" {
  value = azurerm_mysql_flexible_server.example.name
}

output "mysql_server_flexible_full_domain_name" {
  value = azurerm_mysql_flexible_server.example.fqdn
}

output "mysql_server_flexible_admin_login" {
  value = azurerm_mysql_flexible_server.example.administrator_login
}

output "mysql_server_flexible_admin_login_pass" {
  value     = azurerm_mysql_flexible_server.example.administrator_password
  sensitive = true
}

output "mysql_database_flexible_name" {
  value = azurerm_mysql_flexible_database.example.name
}