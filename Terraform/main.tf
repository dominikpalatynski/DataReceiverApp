# We strongly recommend using the required_providers block to set the
# Azure Provider source and version being used
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}

# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {}
}

# Create a resource group
resource "azurerm_resource_group" "DataReceiverApp-resources" {
  name     = "DataReceiverApp-resources"
  location = "polandcentral"
  tags = {
    environment = "dev"
  }
}

# Create a virtual network within the resource group
resource "azurerm_virtual_network" "DataReceiverApp-network" {
  name                = "DataReceiverApp-network"
  resource_group_name = azurerm_resource_group.DataReceiverApp-resources.name
  location            = azurerm_resource_group.DataReceiverApp-resources.location
  address_space       = ["10.123.0.0/16"]
  tags = {
    environment = "dev"
  }
}

resource "azurerm_subnet" "DataReceiverApp-subnet" {
  name                 = "DataReceiverApp-subnet"
  resource_group_name  = azurerm_resource_group.DataReceiverApp-resources.name
  virtual_network_name = azurerm_virtual_network.DataReceiverApp-network.name
  address_prefixes     = ["10.123.1.0/24"]
}

resource "azurerm_network_security_group" "DataReceiverApp-security" {
  name                = "DataReceiverApp-security"
  location            = azurerm_resource_group.DataReceiverApp-resources.location
  resource_group_name = azurerm_resource_group.DataReceiverApp-resources.name
  tags = {
    environment = "dev"
  }
}

resource "azurerm_network_security_rule" "DataReceiverApp-dev-rule" {
  name                        = "DataReceiverApp-dev-rule"
  priority                    = 100
  direction                   = "Inbound"
  access                      = "Allow"
  protocol                    = "*"
  source_port_range           = "*"
  destination_port_range      = "*"
  source_address_prefix       = "*"
  destination_address_prefix  = "*"
  resource_group_name         = azurerm_resource_group.DataReceiverApp-resources.name
  network_security_group_name = azurerm_network_security_group.DataReceiverApp-security.name
}

resource "azurerm_subnet_network_security_group_association" "DataReceiverApp-subnet-security" {
  subnet_id                 = azurerm_subnet.DataReceiverApp-subnet.id
  network_security_group_id = azurerm_network_security_group.DataReceiverApp-security.id
}

resource "azurerm_public_ip" "DataReceiverApp-public-IP" {
  name                = "DataReceiverApp-public-IP"
  resource_group_name = azurerm_resource_group.DataReceiverApp-resources.name
  location            = azurerm_resource_group.DataReceiverApp-resources.location
  allocation_method   = "Dynamic"

  tags = {
    environment = "Dev"
  }
}

resource "azurerm_network_interface" "DataReceiverApp-network-interface" {
  name                = "DataReceiverApp-network-interface"
  location            = azurerm_resource_group.DataReceiverApp-resources.location
  resource_group_name = azurerm_resource_group.DataReceiverApp-resources.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.DataReceiverApp-subnet.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.DataReceiverApp-public-IP.id
  }

  tags = {
    environment = "Dev"
  }
}

resource "azurerm_linux_virtual_machine" "DataReceiverApp-linux-vm" {
  name                = "DataReceiverApp-linux-vm"
  location            = azurerm_resource_group.DataReceiverApp-resources.location
  resource_group_name = azurerm_resource_group.DataReceiverApp-resources.name
  size                = "Standard_DS1_v2"
  admin_username      = "adminuser"

  admin_password = "DataReceiverAppPassword2024!"

  custom_data = filebase64("customdata.tpl")

  network_interface_ids = [
    azurerm_network_interface.DataReceiverApp-network-interface.id,
  ]
  # 74.248.137.150
  # Authentication via SSH
  admin_ssh_key {
    username   = "adminuser"
    public_key = file("~/.ssh/DataReceiverApp_Azure_Key.pub") # Path to your SSH public key
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "19_10-daily-gen2"
    version   = "19.10.202007100"
  }

  disable_password_authentication = false


  computer_name = "DataReceiverAppVM"
  tags = {
    environment = "dev"
  }
}