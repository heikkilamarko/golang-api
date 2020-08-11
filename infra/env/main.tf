provider "azurerm" {
  features {}
}

# variables

variable "acr_resource_group_name" {}
variable "acr_name" {}

variable "name" {}
variable "environment" {}
variable "suffix" {}

variable "db_name" {}
variable "db_login" {}
variable "db_password" {}

# data sources

data "azurerm_container_registry" "shared" {
  name                = var.acr_name
  resource_group_name = var.acr_resource_group_name
}

# resources

resource "azurerm_resource_group" "env" {
  name     = "rg-${var.name}-${var.environment}"
  location = "West Europe"
}

resource "azurerm_storage_account" "env" {
  name                     = "st${var.name}${var.environment}${var.suffix}"
  resource_group_name      = azurerm_resource_group.env.name
  location                 = azurerm_resource_group.env.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  static_website {
    index_document     = "index.html"
    error_404_document = "index.html"
  }
}

resource "azurerm_application_insights" "env" {
  name                = "appi-${var.name}-${var.environment}-${var.suffix}"
  location            = azurerm_resource_group.env.location
  resource_group_name = azurerm_resource_group.env.name
  application_type    = "web"
}

resource "azurerm_log_analytics_workspace" "env" {
  name                = "log-${var.name}-${var.environment}-${var.suffix}"
  location            = azurerm_resource_group.env.location
  resource_group_name = azurerm_resource_group.env.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

resource "azurerm_app_service_plan" "env" {
  name                = "plan-${var.name}-${var.environment}-${var.suffix}"
  location            = azurerm_resource_group.env.location
  resource_group_name = azurerm_resource_group.env.name
  kind                = "Linux"
  reserved            = true

  sku {
    tier = "Basic"
    size = "B1"
  }
}

resource "azurerm_app_service" "env" {
  name                = "app-${var.name}-${var.environment}-${var.suffix}"
  location            = azurerm_resource_group.env.location
  resource_group_name = azurerm_resource_group.env.name
  app_service_plan_id = azurerm_app_service_plan.env.id

  site_config {
    linux_fx_version = "DOCKER|"
  }

  app_settings = {
    DOCKER_REGISTRY_SERVER_URL      = "https://${data.azurerm_container_registry.shared.login_server}"
    DOCKER_REGISTRY_SERVER_USERNAME = data.azurerm_container_registry.shared.admin_username
    DOCKER_REGISTRY_SERVER_PASSWORD = data.azurerm_container_registry.shared.admin_password
    APPINSIGHTS_INSTRUMENTATIONKEY  = azurerm_application_insights.env.instrumentation_key
  }

  lifecycle {
    ignore_changes = [
      site_config,
      app_settings,
    ]
  }
}

resource "azurerm_postgresql_server" "env" {
  name                             = "psql-${var.name}-${var.environment}-${var.suffix}"
  location                         = azurerm_resource_group.env.location
  resource_group_name              = azurerm_resource_group.env.name
  administrator_login              = var.db_login
  administrator_login_password     = var.db_password
  sku_name                         = "B_Gen5_2"
  version                          = "10"
  storage_mb                       = 5120
  backup_retention_days            = 7
  ssl_enforcement_enabled          = true
  ssl_minimal_tls_version_enforced = "TLS1_2"
}

resource "azurerm_postgresql_firewall_rule" "env" {
  name                = "azure"
  resource_group_name = azurerm_resource_group.env.name
  server_name         = azurerm_postgresql_server.env.name
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "0.0.0.0"
}

resource "azurerm_postgresql_database" "env" {
  name                = var.db_name
  resource_group_name = azurerm_resource_group.env.name
  server_name         = azurerm_postgresql_server.env.name
  charset             = "UTF8"
  collation           = "English_United States.1252"
}

resource "azurerm_monitor_diagnostic_setting" "env" {
  name                       = "diag-${var.name}-${var.environment}-${var.suffix}"
  target_resource_id         = azurerm_app_service.env.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.env.id

  log {
    category = "AppServiceHTTPLogs"
    retention_policy {
      enabled = false
    }
  }

  log {
    category = "AppServiceConsoleLogs"
    retention_policy {
      enabled = false
    }
  }

  log {
    category = "AppServiceAppLogs"
    retention_policy {
      enabled = false
    }
  }

  log {
    category = "AppServiceFileAuditLogs"
    retention_policy {
      enabled = false
    }
  }

  log {
    category = "AppServiceAuditLogs"
    retention_policy {
      enabled = false
    }
  }

  log {
    category = "AppServiceIPSecAuditLogs"
    retention_policy {
      enabled = false
    }
  }

  log {
    category = "AppServicePlatformLogs"
    retention_policy {
      enabled = false
    }
  }

  metric {
    category = "AllMetrics"
    retention_policy {
      enabled = false
    }
  }
}
