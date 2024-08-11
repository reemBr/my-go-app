# My Go Web App on Kubernetes with Minikube

## Overview

This project demonstrates how to deploy a simple containerized Go web application to a local Kubernetes cluster using Minikube. The application serves a static HTML page with a fancy "Hi" message. 

## Prerequisites

- **Docker**: Installed for building the Docker image.
- **Minikube**: Installed for running a local Kubernetes cluster.
- **Kubectl**: Installed for interacting with your Kubernetes cluster.

## Getting Started


### 1. Start Minikube

Start Minikube if it is not already running:

```bash
minikube start
```

### 2. Configure Docker to Use Minikube’s Docker Daemon

To build Docker images that can be used by Minikube, configure Docker to use Minikube’s Docker daemon:

```bash
eval $(minikube -p minikube docker-env)
```

### 3. Apply Kubernetes Manifests

Apply the Kubernetes manifests to deploy your application:

Apply the deployment:

```bash
kubectl apply -f deployment.yaml
```

Apply the service:

```bash
kubectl apply -f service.yaml
```

### 6. Access the Application

Check the status of your service to get the external IP address:

```bash
kubectl get services
```

If the `EXTERNAL-IP` is listed as `<pending>`, you can use Minikube’s service command to access the application:

```bash
minikube service my-go-app-service
```

This command will open your default web browser to the URL of your service. If you prefer to manually visit the URL, you can also use the NodePort service instead of LoadBalancer:

```bash
kubectl get services
```

Look for the `NodePort` value under `PORT(S)`. Access the application using `http://<minikube-ip>:<node-port>`. Find Minikube’s IP with:

```bash
minikube ip
```

### 7. Clean Up

To delete the resources created in your Kubernetes cluster:

```bash
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
```

Stop Minikube if you no longer need it:

```bash
minikube stop
```
