# Notes for Chapter 1: Introduction to Docker

## ✅ Benefits of Docker

- Makes tooling easier to write and share between applications.
- **Packages software** in a way that leverages the skills developers already have:
  - We do not need special release engineers to have knowledge of the tooling required to package software.
  - We can use the same tools we use for other projects, such as Docker.
  - It packages the knowledge of tools like `rpm`, `mock`, `dpkg`, and `pbuilder` into one packaging format known as the **Open Container Initiative (OCI)**.
- Builds application software and the required OS filesystem together in a single **standardized image format**.
  - Docker's layered images make this an efficient process that ensures your application is running in the expected environment.
  - It allows you to package the entire environment, including all dependencies, into one file.
- Uses the packaged artifacts to test and **deliver the exact same artifact** to all systems in all environments.
  - This makes it easier for developers to share their work with other teams.
  - It is a great way to ensure that your application works on any system.
- **Abstracts the software application** from the hardware without sacrificing resources.

## ❌ What Docker Isn't

- **It is not an enterprise virtualization platform.**
  - A virtual machine (VM) is a complete operating system on top of a hypervisor that is managed by the underlying host operating system.
  - With containers, both the host and the container share the same kernel, meaning the container utilizes fewer system resources compared to a VM.
- **It is not a cloud platform** (like OpenStack, CloudStack, etc.).
  - On the surface, it looks the same when it comes to horizontal scaling based on demand.
  - But there is a small difference: a cloud platform will add a new host, whereas Docker will add a new container to the existing Docker host.
  - Hence, systems like object storage and block storage are deployed on cloud platforms and not in containers.
- **It is not a configuration management tool** (like Puppet, Chef, etc.).
  - Dockerfiles are used to define how a container should look at build time, but do not manage the container's ongoing state.
- **It is not a deployment framework** (like Capistrano, Fabric, etc.).
  - Docker cannot be used to automate a complex deployment process. It is more of a tool for building and running containers.
- **It is not a development environment** (like Vagrant, etc.).
  - Tools like Vagrant help reduce the common issue of "it works on my machine" because they are able to create an identical environment in which you can test your code.
- **It is not a workload management tool** (like Kubernetes, Swarm, etc.).
  - We need an orchestration layer to coordinate work across a pool of Linux container hosts, and to track the current state of all the hosts and their resources.

## 📚 Important Terminologies

- **Docker Client:** The command-line interface for interacting with the Docker daemon. It is used to run containers and manage images.
- **Docker Server:** A process that runs on each host in a cluster which manages the lifecycle of the container.
- **Docker or OCI Images:** A set of instructions and files that describe how to build an image. The file format for Docker images is called a `Dockerfile`.
- **Linux Containers:** A lightweight, isolated environment that can run any application on top of it.
- **Atomic or Immutable Host:** An operating system that has a single process running at a time. It means that the state of the host machine cannot be changed by other processes.
