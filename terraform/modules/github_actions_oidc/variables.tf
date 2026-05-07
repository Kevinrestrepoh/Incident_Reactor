variable "github_repo" {
  description = "GitHub repo (user/repo)"
  type        = string
}

variable "branch" {
  description = "Branch allowed to assume role"
  type        = string
  default     = "main"
}
