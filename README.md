# zenas

## Getting Started

This section will guide you through getting @1sl4nds/zenas up and running on your local development environment.

### Prerequisites

Before you begin, ensure you have the following installed:

- Docker: For containerization and easy deployment of the application and its dependencies.
- Go: Required to build the application from the source.
- Python: Used for certain scripts and data processing tasks.
- Terraform: For provisioning and managing the infrastructure as code.

### Installation

Follow these steps to get your development environment set up:

1\. Clone the repository to your local machine:

```bash
git clone https://github.com/1sl4nds/zenas.git
cd zenas
```

2\. Build the Docker images:

```bash
docker-compose build
```

3\. Initialize the infrastructure with Terraform (ensure you're in the `depoy/terraform` directory):

```bash
terraform -chdir ./deploy/terraform init
terraform -chdir ./deploy/terraform apply
```

4\. Start services using Docker Compose:

```bash
docker-compose up -d
```

This will start all the required services as defined in the `docker-compose.yml` file.

## Usage

After successfully installing the application, you can start using it by:

- Accessing the OpenAPI documentation for the API at `http://localhost:8080/api-docs` to understand how to interact with the API.
- Exploring the Grafana dashboards at `http://localhost:3000` for monitoring and analytics.
- Checking the logs and metrics collected by Prometheus and Jaeger for performance and tracing.

## Contributing

We welcome contributions to @1sl4nds/zenas! Please read `CONTRIBUTING.md` for more information on how to get started with contributing.

## Built-with

- [Apache OpenNLP](https://opennlp.apache.org/)
- [Apache Tika](https://tika.apache.org/)
- [Babel](https://babeljs.io/)
- [Cobra](https://cobra.dev/)
- [Docker](https://www.docker.com/)
- [GitHub](https://docs.github.com/en/)
- [Go](https://go.dev/)
- [Grafana](https://grafana.com/)
- [Helm](https://helm.sh/)
- [Jaeger](https://www.jaegertracing.io/)
- [Jest](https://jestjs.io/)
- [Keycloak](https://www.keycloak.org/)
- [MinIO](https://min.io/)
- [OpenAPI](https://www.openapis.org/)
- [OpenSearch](https://opensearch.org/)
- [OpenTelemetry](https://opentelemetry.io/)
- [Postgres](https://www.postgresql.org/)
- [Prometheus](https://prometheus.io/)
- [Python](https://www.python.org/)
- [React](https://react.dev/)
- [scikit-learn](https://scikit-learn.org/stable/)
- [Terraform](https://www.terraform.io/)
- [TypeScript](https://www.typescriptlang.org/)

## References

- [What is Computational Law? | Stanford](https://law.stanford.edu/2021/03/10/what-is-computational-law/)
