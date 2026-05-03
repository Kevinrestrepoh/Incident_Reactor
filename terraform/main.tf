module "vpc" {
  source = "./modules/vpc"

  name       = "incident-reactor"
  cidr_block = "10.0.0.0/16"
  azs        = ["us-east-2a", "us-east-2b"]

  public_subnets = [
    "10.0.1.0/24",
    "10.0.2.0/24"
  ]

  private_subnets = [
    "10.0.101.0/24",
    "10.0.102.0/24"
  ]
}

module "ecr" {
  source = "./modules/ecr"

  repository_name = var.project_name
}

module "iam_oidc" {
  source = "./modules/iam_oidc"

  github_repo = var.github_repo
}
