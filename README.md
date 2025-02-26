# Custom KIE Sandbox

This project provides a customizable version of KIE Sandbox, allowing for branding modifications, custom accelerators, and configurable Git/Kubernetes integration.

## Overview

KIE Sandbox is an open-source business automation workbench that enables you to design, test, and deploy business processes, decision models, and other business automation assets. This project showcases how you to rebrand the sandbox and configure it for your specific needs.

## Getting Started

### Prerequisites

- Docker and Docker Compose installed on your system
- Basic familiarity with Docker and containerization concepts

### Running the Sandbox

1. Clone this repository:

   ```
   git clone https://github.com/aletyx-labs/kie-10.0.0-sandbox.git
   cd kie-10.0.0-sandbox
   ```

2. Start the containers, which will build the workspace Dockerfile to use for the kie_sandbox container. If you add a `-d` it will run it in detached mode so that the terminal is still accessible and not tied to the running of the compose file, but is purely optional:

   ```
   docker-compose up --build -d
   ```

3. Access the KIE Sandbox in your browser:
   ```
   http://localhost:9090
   ```

## Building a Custom Image

The project includes a Dockerfile that extends the base KIE Sandbox image with your customizations.

To build a custom image:

1. Modify the Dockerfile as needed
2. Run Docker Compose with the build flag:

   ```
   docker-compose up --build -d
   ```

## Customization Options

### Branding

You can customize the branding by replacing the SVG files in the `branding` directory:

- `branding/aletyx-light.svg` - Logo used in light mode
- `branding/aletyx-dark.svg` - Logo used in dark mode
- `branding/favicon.png` - Browser tab icon

The Dockerfile copies these files to the appropriate locations in the container.

### Custom Accelerators

Accelerators are templates that help you quickly create new projects. They are configured in the `docker-compose.yml` file under the `KIE_SANDBOX_ACCELERATORS` environment variable.

Each accelerator configuration includes:

- `name`: Display name for the accelerator
- `iconUrl`: URL to the icon image
- `gitRepositoryUrl`: Git repository containing the accelerator template
- `gitRepositoryGitRef`: Branch or tag to use
- `dmnDestinationFolder`: Destination folder for DMN files
- `bpmnDestinationFolder`: Destination folder for BPMN files
- `otherFilesDestinationFolder`: Destination folder for other files

To add or modify accelerators, edit the JSON array in the `KIE_SANDBOX_ACCELERATORS` environment variable.

```yaml
KIE_SANDBOX_ACCELERATORS: >
  [
                       {
     "name": "Process Kogito jBPM Profile",
     "iconUrl": "https://raw.githubusercontent.com/timwuthenow/sandbox-accelerators/refs/heads/main/quarkus-logo.png",
     "gitRepositoryUrl": "https://github.com/timwuthenow/sandbox-accelerators.git",
     "gitRepositoryGitRef": "main",
     "dmnDestinationFolder": "src/main/resources/dmn",
     "bpmnDestinationFolder": "src/main/resources/bpmn",
     "otherFilesDestinationFolder": "src/main/resources/others"
   },
   {
     "name": "Decisions Kogito jBPM Profile",
     "iconUrl": "https://raw.githubusercontent.com/timwuthenow/sandbox-accelerators/refs/heads/main/quarkus-logo.png",
     "gitRepositoryUrl": "https://github.com/timwuthenow/sandbox-accelerators.git",
     "gitRepositoryGitRef": "dmn",
     "dmnDestinationFolder": "src/main/resources/dmn",
     "bpmnDestinationFolder": "src/main/resources/bpmn",
     "otherFilesDestinationFolder": "src/main/resources/others"
   }
  ]
```

### Git Integration

You can configure which Git providers are available through the `KIE_SANDBOX_AUTH_PROVIDERS` environment variable. The default configuration includes:

- GitHub
- GitLab
- Bitbucket

Each provider configuration includes:

- `id`: Unique identifier
- `domain`: Provider domain
- `type`: Provider type
- `name`: Display name
- `enabled`: Whether the provider is enabled
- `iconPath`: Path to the provider icon
- `group`: Provider group (e.g., "git" or "cloud")

### Kubernetes/OpenShift Integration

Kubernetes and OpenShift integration is also configured through the `KIE_SANDBOX_AUTH_PROVIDERS` environment variable. The default configuration includes both Kubernetes and OpenShift providers.

### Custom Git Commit Message

You can require custom commit messages by setting the `KIE_SANDBOX_REQUIRE_CUSTOM_COMMIT_MESSAGE` environment variable to `true`.

## Environment Variables

Key environment variables:

- `KIE_SANDBOX_APP_NAME`: Sets the application name
- `KIE_SANDBOX_EXTENDED_SERVICES_URL`: URL for extended services
- `KIE_SANDBOX_CORS_PROXY_URL`: URL for CORS proxy
- `KIE_SANDBOX_ACCELERATORS`: JSON array of accelerator configurations
- `KIE_SANDBOX_AUTH_PROVIDERS`: JSON array of authentication provider configurations
- `KIE_SANDBOX_REQUIRE_CUSTOM_COMMIT_MESSAGE`: Requires custom commit messages when set to true

## Architecture

The project consists of three main services:

1. `kie_sandbox`: The main KIE Sandbox application with your customizations
2. `extended_services`: Provides additional functionality for the sandbox
3. `cors`: A CORS proxy to allow communication with external services

## Troubleshooting

If you encounter issues:

1. Check container logs:

   ```
   docker-compose logs kie_sandbox
   ```

2. Verify that all services are running:

   ```
   docker-compose ps
   ```

3. Ensure that the branding files are correctly formatted SVG files

## License

This project is distributed under the Apache License 2.0.
