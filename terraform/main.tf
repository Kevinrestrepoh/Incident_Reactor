module "vpc" {
  source = "./modules/vpc"

  name       = var.project_name
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

module "github_actions_oidc" {
  source = "./modules/github_actions_oidc"

  github_repo = var.github_repo
}

module "eks" {
  source = "./modules/eks"

  name            = var.project_name
  vpc_id          = module.vpc.vpc_id
  private_subnets = module.vpc.private_subnets

  instance_types = ["t3.small"]

  desired_size = 1
  max_size     = 2
  min_size     = 1

  github_actions_role_arn = module.github_actions_oidc.role_arn
}

module "alb_irsa" {
  source = "./modules/irsa_alb_controller"

  name              = var.project_name
  oidc_provider_arn = module.eks.oidc_provider_arn
  oidc_provider_url = module.eks.oidc_provider_url
}
