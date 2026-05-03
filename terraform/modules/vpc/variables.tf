variable "name" {
  description = "VPC name"
  type        = string
}

variable "cidr_block" {
  description = "VPC CIDR"
  type        = string
}

variable "azs" {
  description = "Availability zones"
  type        = list(string)
}

variable "public_subnets" {
  description = "Public subnet CIDRs"
  type        = list(string)
  validation {
    condition     = length(var.public_subnets) == length(var.azs)
    error_message = "public_subnets must match number of AZs"
  }
}

variable "private_subnets" {
  description = "Private subnet CIDRs"
  type        = list(string)
  validation {
    condition     = length(var.private_subnets) == length(var.azs)
    error_message = "private_subnets must match number of AZs"
  }
}
