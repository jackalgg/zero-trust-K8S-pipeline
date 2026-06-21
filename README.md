# Zero Trust Kubernetes Pipeline

### Overview
This repository hosts a reference implementation for a hardened software delivery lifecycle. The goal is to eliminate common attack vectors in containerized environments by using minimal base images, least-privilege runtime defaults, and credential-less CI/CD authentication.

### Core Security Principles
Modern infrastructure requires more than just perimeter defense. This pipeline implements three specific pillars of the zero trust model.

**Minimal Attack Surface**
The application is built using Wolfi based Chainguard images. By utilizing a multi stage build process, the final runtime environment contains no shell, no package manager, and no unnecessary binaries. This distroless approach reduces the vulnerability count to zero at the point of build.

**Least Privilege Execution**
Standard container defaults often lead to over privileged processes. This implementation explicitly enforces execution as a non root user with UID 65532. The Go application includes internal logic to verify its own security context and will log a critical alert if it detects an elevated privilege state.

**Credential-less CI/CD**
Images are built and pushed to Amazon ECR via GitHub Actions using AWS OIDC federation. No long-lived access keys are stored in the repository or CI environment. The workflow assumes a scoped IAM role, limiting blast radius if credentials are ever compromised.


### Technical Stack
* Language: Go 1.25
* Base Images: Chainguard Static (Wolfi)
* Orchestration: Kubernetes (EKS)
* Infrastructure: Terraform (VPC, ECR, EKS)
* CI Platform: GitHub Actions (OIDC → ECR)

### Implementation Details
The application logic focuses on system level introspection. It monitors for SIGTERM and SIGINT signals to ensure graceful shutdowns within Kubernetes pods, preventing orphaned processes and state inconsistency. The logging is structured to be consumed by enterprise observability platforms, prioritizing clarity and system metadata.

### About Michael
I am a Computer Scientist and Navy Veteran with a focus on systems security and cloud infrastructure. My approach to engineering is shaped by years of maintaining mission critical avionics where the margin for error was zero. I apply that same rigor to building secure, resilient distributed systems.
