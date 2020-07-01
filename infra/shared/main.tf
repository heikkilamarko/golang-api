terraform {
  required_version = ">= 0.12.0"
  required_providers {
    azurerm = ">= 2.12.0"
  }
}

provider "azurerm" {
  features {}
}

# variables

variable "name" {}
variable "suffix" {}

# resources

resource "azurerm_resource_group" "shared" {
  name     = "rg-${var.name}"
  location = "West Europe"
}

resource "azurerm_container_registry" "shared" {
  name                = "acr${var.name}${var.suffix}"
  location            = azurerm_resource_group.shared.location
  resource_group_name = azurerm_resource_group.shared.name
  sku                 = "Standard"
  admin_enabled       = true
}
