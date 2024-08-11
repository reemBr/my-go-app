# Go Web Application with CI/CD Automation

This repository contains a simple Go web application with CI/CD automation using Ansible and Kubernetes. The project is structured into three main folders:

- **`app/`**: Contains the source code and Dockerfile for the Go web application.
- **`k8s/`**: Contains Kubernetes manifests for deploying the application.
- **`ansible/`**: Contains Ansible playbooks and requirements for automating the build and deployment process.
- **`.github/workflow/`** Contains Github workflow for automating running the ansible playbook upon code changes
## Project Structure

### `app/`

This folder includes:

- **`Dockerfile`**: Dockerfile to build the Go web application image.
- **`main.go`**: The Go source code for the web application.
- **`index.html`**: The HTML file served by the Go application.

### `k8s/`

This folder contains:

- **`deployment.yaml`**: Kubernetes Deployment manifest for deploying the Go application.
- **`service.yaml`**: Kubernetes Service manifest for exposing the Go application.

### `ansible/`

This folder includes:

- **`build_and_deploy.yml`**: Ansible playbook to build the Docker image, update the Kubernetes deployment, and apply changes to the Kubernetes cluster.
- **`requirements.yml`**: Ansible requirements file for managing dependencies.

## Setup and Usage

### 1. **Build and Run the Application Locally**

To build and run the application locally, follow (./app/README)[instructions]

### 2. **Deploy to Kubernetes**

To deploy the application to a Kubernetes cluster using Ansible, follow (./app/README)[instructions]

5. **Run the Ansible Playbook:**

To automaticlly do all the previous steps, follow (./app/README)[instructions]

### 3. **Automate Deployment with GitHub Actions**

This repository includes a GitHub Actions workflow to automate the build and deployment process upon code changes.

1. **Verify GitHub Actions Workflow:**

   The workflow file is located at `.github/workflows/deploy.yml`.

2. **Configure GitHub Secrets:**

   - Add `DOCKER_REGISTRY` with your Docker registry URL.
   - Add `KUBECONFIG` with the content of your Kubernetes configuration file.

3. **Push Changes:**

   Commit and push your changes to the `main` branch. This will trigger the GitHub Actions workflow to build and deploy the application.

   ```bash
   git add .
   git commit -m "Add GitHub Actions workflow for automated build and deploy"
   git push origin main
   ```
