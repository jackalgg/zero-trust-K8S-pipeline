# Zero Trust Kubernetes Pipeline

### Overview
This repository hosts a reference implementation for a hardened software delivery lifecycle. The goal is to eliminate common attack vectors in containerized environments by enforcing zero CVE base images and cryptographic identity verification.

### Core Security Principles
Modern infrastructure requires more than just perimeter defense. This pipeline implements three specific pillars of the zero trust model.

**Minimal Attack Surface**
The application is built using Wolfi based Chainguard images. By utilizing a multi stage build process, the final runtime environment contains no shell, no package manager, and no unnecessary binaries. This distroless approach reduces the vulnerability count to zero at the point of build.

**Least Privilege Execution**
Standard container defaults often lead to over privileged processes. This implementation explicitly enforces execution as a non root user with UID 65532. The Go application includes internal logic to verify its own security context and will log a critical alert if it detects an elevated privilege state.

**Cryptographic Provenance**
Security is maintained through the build process using Sigstore Cosign. Every image produced by this pipeline is signed with an OIDC identity. This creates a verifiable link between the source code and the running artifact, ensuring that only trusted code reaches the cluster.



### Technical Stack
* Language: Go 1.22
* Base Images: Chainguard Static (Wolfi)
* Orchestration: Kubernetes
* Security Tooling: Sigstore Cosign and Kyverno
* CI Platform: GitHub Actions

### Implementation Details
The application logic focuses on system level introspection. It monitors for SIGTERM and SIGINT signals to ensure graceful shutdowns within Kubernetes pods, preventing orphaned processes and state inconsistency. The logging is structured to be consumed by enterprise observability platforms, prioritizing clarity and system metadata.

### About Michael
I am a Computer Scientist and Navy Veteran with a focus on systems security and cloud infrastructure. My approach to engineering is shaped by years of maintaining mission critical avionics where the margin for error was zero. I apply that same rigor to building secure, resilient distributed systems.
