# GORCHASE

Gorchase is a minimalistic & distributed generic asset purchase service written in Go.

# Status

![Lint & Test](https://github.com/nmarcetic/gorchase/actions/workflows/test.yml/badge.svg)
![Image Publishe](https://github.com/nmarcetic/gorchase/actions/workflows/publish.yml/badge.svg)

**NOTE:** Project is currently under heavy development

# Motivation

TBA

# What's included

* User accounts (sign-in/signup) with Admin account
* REST full API for generic item purchase (gin, gorm, mysql/postgresql)
* High test coverage
* CI/CD with Github Actions
* Dockerized (multi-stage build, small footprint)
* docker-compose for local development
* K8s provisioning manifests for resilient AWS EKS cluster
* Helm Chart
* Logging (EFK stack) , Monitoring (Prometheus, Grafana) and Tracing (Jaeger)
* Swagger included
* Postman collection for E2E testing

# Getting Started

TBA

## Requirments

You'll need the following software installed to get started:

* [Docker][docker-install]  engine v20ˆ
* [Docker Compose][docker-compose-install] v3.7

### Development

* Go 1.17

## Installation

Gorchase is containerized with Docker and orchestrated with [Kubernetes][k8s-website]
docker-compose can be used for development and testing. For production, I suggest using Kubernetes.

Once the prerequisites are installed, execute the following commands from the project's root:

`docker-compose -f docker/docker-compose.yml up`


## Configuration

   Configure env variables.

## Usage

TBA

# Contributing 

TBA

## Getting Help

TBA

# Licencing 

TBA

<!-- Named links -->
[docker-install]: https://docs.docker.com/get-docker/
[docker-compose-install]: https://docs.docker.com/compose/install/
[k8s-website]: https://kubernetes.io/