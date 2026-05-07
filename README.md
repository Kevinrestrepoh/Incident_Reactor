# Incident Reactor

Cloud-native event simulation platform built with Go, AWS, Terraform, Kubernetes, and GitHub Actions.

---

# Architecture

## Stack

* Go
* Docker
* Kubernetes (EKS)
* Terraform
* AWS ECR
* AWS IAM / OIDC / IRSA
* AWS ALB Ingress Controller
* GitHub Actions
* Prometheus
* Grafana

---

# Infrastructure

## AWS

### Networking

* Custom VPC
* Public and private subnets
* Internet Gateway
* Route Tables

### Compute

* Amazon EKS managed cluster
* Managed node groups

### Container Registry

* Amazon ECR

### IAM

* GitHub Actions OIDC role
* IRSA role for AWS Load Balancer Controller

### Ingress

* AWS Application Load Balancer (ALB)
* Kubernetes Ingress

---

# CI/CD

GitHub Actions pipeline:

* Runs `go vet`
* Runs tests
* Builds Docker image
* Pushes image to ECR using OIDC authentication

No static AWS credentials are stored in GitHub.

---

# Observability

## Prometheus

Scrapes application metrics from `/metrics`.

## Grafana

Visualizes Prometheus metrics.

---

# GitHub Actions OIDC Setup

AWS IAM role is configured to trust GitHub's OIDC provider.

This allows GitHub Actions to authenticate securely to AWS without long-lived credentials.

## Required GitHub Secret

```text
AWS_ROLE_ARN
```

---

# AWS Load Balancer Controller

## Add Helm repository

```bash
helm repo add eks https://aws.github.io/eks-charts
helm repo update
```

## Install controller

```bash
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n kube-system \
  --set clusterName=incident-reactor \
  --set serviceAccount.create=true \
  --set serviceAccount.name=aws-load-balancer-controller \
  --set serviceAccount.annotations."eks\.amazonaws\.com/role-arn"=<ALB_ROLE_ARN> \
  --set vpcId=<VPC_ID>
```
