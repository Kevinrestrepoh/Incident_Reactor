module "ecr" {
  source = "./modules/ecr"

  repository_name = var.project_name
}

module "iam_oidc" {
  source = "./modules/iam_oidc"

  github_repo = var.github_repo
}
