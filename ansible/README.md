# Ansible Automation for Docker and Kubernetes Deployment


## Ansible Playbook

The Ansible playbook, `build_and_deploy.yml`, is used to automate the following tasks:

1. Extract the latest Git commit hash.
2. Build the Docker image tagged with the commit hash.
3. Push the Docker image to a Docker registry.
4. Update the Kubernetes deployment to use the new image tag.


## Prerequisites

- **ansible**: Installed for automating image build and push.

## Getting Started


### 1. Install Ansible Collections

Install the required Ansible collections specified in `requirements.yml`:

```bash
ansible-galaxy collection install -r requirements.yml
```

### 2. Configure Docker Registry

Ensure your Docker registry is accessible and replace `my-registry` in the playbook with your registry name.


### 3. Run the Ansible Playbook

Execute the Ansible playbook to build the Docker image and update the Kubernetes deployment:

```bash
ansible-playbook ansible/build_and_deploy.yml
```

