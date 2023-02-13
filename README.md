# DevOps Directive Docker Course

This is the companion repo to: `<INSERT VIDEO LINK>`

`<INSERT THUMBNAIL IMAGE>`

## 01 - History and Motivation

Examines the evolution of virtualization technologies from bare metal, virtual machines, and containers and the tradeoffs between them.

## 02 - Technology Overview

Explores the three core Linux features that enable containers to function (cgroups, namespaces, and union filesystems), as well as the architecture of the Docker components.

## 03 - Installation and Set Up

Covers the steps to install and configure Docker Desktop on your system.

## 04 - Using 3rd Party Containers

Before we build our own container images, we can familiarize ourselves with the technology by using publicly available container images. This section covers the nuances of data persistence with containers and then highlights some key use cases for using public container images.

## 05 - Example Web Application

Learning about containerization is interesting, but without a practical example it isn't very useful. In this section we create a 3 tier web application with a React front end client, two apis (node.js + golang), and a database. The application is as simple as possible while still providing a realistic microservice system to containerize. 

## 06 - Building Container Images

Demonstrates how to write Dockerfiles and build container images for the components of the example web app. Starting with a naive implementation, we then iterate towards a production ready container image.

## 07 - Container Registries

Explains what container registries are and how to use them to share and distribute container images.

## 08 - Running Containers

Using the containerized web application from sections 05 and 06, we craft the necessary commands to run our application with Docker and Docker Compose. We also cover the variety of runtime configuration options and when to use them.

## 09 - Container Security

Highlights best practices for container image and container runtime security.

## 10 - Interacting with Docker Objects

Describes how to use Docker to interact with containers, container images, volumes, and networks.

## 11 - Development Workflows

Establishes tooling and configuration to enable improved developer experience when working with containers.

## 12 - Deploying Containers

Demonstrates deploying container applications to production using three different approaches: railway.app, a single node Docker Swarm, and a Kubernetes cluster.
