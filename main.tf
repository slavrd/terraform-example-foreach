terraform {
  required_version = ">= 0.12.6"
}

variable "pets" {
  type        = map(number)
  description = "A map indicating the pets' names and their coresponding name length."
  default = {
    pet1 = 1
    pet2 = 2
    pet3 = 5
  }
}

resource "random_pet" "pets" {
  for_each = var.pets
  prefix   = each.key
  length   = each.value
}

output "pets" {
  description = "Created pets' names"
  value       = values(random_pet.pets)[*].id
}
