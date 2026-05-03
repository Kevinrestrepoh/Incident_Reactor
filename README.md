## 🔐 GitHub Actions → AWS Authentication (OIDC)

This project uses **OIDC (OpenID Connect)** to allow GitHub Actions to authenticate with AWS **without storing credentials**.

---

### ⚙️ Step 1: Create OIDC Provider in AWS

In **Amazon Web Services**:

* Go to: `IAM → Identity providers → Add provider`
* Provider type: `OpenID Connect`
* Provider URL:

  ```
  https://token.actions.githubusercontent.com
  ```
* Audience:

  ```
  sts.amazonaws.com
  ```

---

### ⚙️ Step 2: Create IAM Role for GitHub Actions

* Go to: `IAM → Roles → Create role`
* Select: `Web identity`
* Choose the GitHub OIDC provider
* Audience: `sts.amazonaws.com`

---

### 🛡️ Step 3: Attach Permissions (ECR push)

Example policy:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:CompleteLayerUpload",
        "ecr:InitiateLayerUpload",
        "ecr:UploadLayerPart",
        "ecr:PutImage"
      ],
      "Resource": "*"
    }
  ]
}
```

---

### 🔑 Step 4: Add GitHub Secret

In your repository:

`Settings → Secrets and variables → Actions`

Add:

* **Name**: `AWS_ROLE_ARN`
* **Value**:

  ```
  arn:aws:iam::YOUR_ACCOUNT_ID:role/github-actions-role
  ```
