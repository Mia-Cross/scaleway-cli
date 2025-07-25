<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw container`
This API allows you to manage your Serverless Containers.
  
- [Container management commands](#container-management-commands)
  - [Create a new container](#create-a-new-container)
  - [Delete a container](#delete-a-container)
  - [Deploy a container](#deploy-a-container)
  - [Get a container](#get-a-container)
  - [List all your containers](#list-all-your-containers)
  - [Update an existing container](#update-an-existing-container)
- [Cron management commands](#cron-management-commands)
  - [Create a new cron](#create-a-new-cron)
  - [Delete an existing cron](#delete-an-existing-cron)
  - [Get a cron](#get-a-cron)
  - [List all your crons](#list-all-your-crons)
  - [Update an existing cron](#update-an-existing-cron)
- [Deploy a container](#deploy-a-container)
- [Domain management commands](#domain-management-commands)
  - [Create a custom domain](#create-a-custom-domain)
  - [Delete a custom domain](#delete-a-custom-domain)
  - [Get a custom domain](#get-a-custom-domain)
  - [List all custom domains](#list-all-custom-domains)
- [Namespace management commands](#namespace-management-commands)
  - [Create a new namespace](#create-a-new-namespace)
  - [Delete an existing namespace](#delete-an-existing-namespace)
  - [Get a namespace](#get-a-namespace)
  - [List all your namespaces](#list-all-your-namespaces)
  - [Update an existing namespace](#update-an-existing-namespace)
- [Token management commands](#token-management-commands)
  - [Create a new revocable token](#create-a-new-revocable-token)
  - [Delete a token](#delete-a-token)
  - [Get a token](#get-a-token)
  - [List all tokens](#list-all-tokens)
- [Trigger management commands](#trigger-management-commands)
  - [Create a trigger](#create-a-trigger)
  - [Delete a trigger](#delete-a-trigger)
  - [Get a trigger](#get-a-trigger)
  - [List all triggers](#list-all-triggers)
  - [Update a trigger](#update-a-trigger)

  
## Container management commands

Container management commands.


### Create a new container

Create a new container in the specified region.

**Usage:**

```
scw container container create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id |  | UUID of the namespace the container belongs to |
| name |  | Name of the container |
| environment-variables.{key} |  | Environment variables of the container |
| min-scale |  | Minimum number of instances to scale the container to |
| max-scale |  | Maximum number of instances to scale the container to |
| memory-limit |  | Memory limit of the container in MB |
| cpu-limit |  | CPU limit of the container in mvCPU |
| timeout |  | Processing time limit for the container |
| privacy | One of: `unknown_privacy`, `public`, `private` | Privacy setting of the container |
| description |  | Description of the container |
| registry-image |  | Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag"). |
| ~~max-concurrency~~ | Deprecated | Number of maximum concurrent executions of the container |
| protocol | One of: `unknown_protocol`, `http1`, `h2c` | Protocol the container uses |
| port |  | Port the container listens on |
| secret-environment-variables.{index}.key |  |  |
| secret-environment-variables.{index}.value |  |  |
| http-option | Default: `enabled`<br />One of: `unknown_http_option`, `enabled`, `redirected` | Configure how HTTP and HTTPS requests are handled |
| sandbox | One of: `unknown_sandbox`, `v1`, `v2` | Execution environment of the container |
| local-storage-limit |  | Local storage limit of the container (in MB) |
| scaling-option.concurrent-requests-threshold |  |  |
| scaling-option.cpu-usage-threshold |  |  |
| scaling-option.memory-usage-threshold |  |  |
| health-check.http.path |  | Path to use for the HTTP health check. |
| health-check.failure-threshold |  | Number of consecutive health check failures before considering the container unhealthy. |
| health-check.interval |  | Period between health checks. |
| tags.{index} |  | Tags of the Serverless Container |
| private-network-id |  | ID of the Private Network the container is connected to. |
| command.{index} |  | Container command |
| args.{index} |  | Container arguments |
| deploy | Default: `true` | Deploy container after creation |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Delete a container

Delete the container associated with the specified ID.

**Usage:**

```
scw container container delete <container-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| container-id | Required | UUID of the container to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Deploy a container

Deploy a container associated with the specified ID.

**Usage:**

```
scw container container deploy <container-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| container-id | Required | UUID of the container to deploy |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get a container

Get the container associated with the specified ID.

**Usage:**

```
scw container container get <container-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| container-id | Required | UUID of the container to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List all your containers

List all containers for a specified region.

**Usage:**

```
scw container container list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc` | Order of the containers |
| namespace-id |  | UUID of the namespace the container belongs to |
| name |  | Name of the container |
| project-id |  | UUID of the Project the container belongs to |
| organization-id |  | UUID of the Organization the container belongs to |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



### Update an existing container

Update the container associated with the specified ID.

**Usage:**

```
scw container container update <container-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| container-id | Required | UUID of the container to update |
| environment-variables.{key} |  | Environment variables of the container |
| min-scale |  | Minimum number of instances to scale the container to |
| max-scale |  | Maximum number of instances to scale the container to |
| memory-limit |  | Memory limit of the container in MB |
| cpu-limit |  | CPU limit of the container in mvCPU |
| timeout |  | Processing time limit for the container |
| redeploy |  | Defines whether to redeploy failed containers |
| privacy | One of: `unknown_privacy`, `public`, `private` | Privacy settings of the container |
| description |  | Description of the container |
| registry-image |  | Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag"). |
| ~~max-concurrency~~ | Deprecated | Number of maximum concurrent executions of the container |
| protocol | One of: `unknown_protocol`, `http1`, `h2c` | Protocol the container uses |
| port |  | Port the container listens on |
| secret-environment-variables.{index}.key |  |  |
| secret-environment-variables.{index}.value |  |  |
| http-option | One of: `unknown_http_option`, `enabled`, `redirected` | Configure how HTTP and HTTPS requests are handled |
| sandbox | One of: `unknown_sandbox`, `v1`, `v2` | Execution environment of the container |
| local-storage-limit |  | Local storage limit of the container (in MB) |
| scaling-option.concurrent-requests-threshold |  |  |
| scaling-option.cpu-usage-threshold |  |  |
| scaling-option.memory-usage-threshold |  |  |
| health-check.http.path |  | Path to use for the HTTP health check. |
| health-check.failure-threshold |  | Number of consecutive health check failures before considering the container unhealthy. |
| health-check.interval |  | Period between health checks. |
| tags.{index} |  | Tags of the Serverless Container |
| private-network-id |  | ID of the Private Network the container is connected to. |
| command.{index} |  | Container command |
| args.{index} |  | Container arguments |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



## Cron management commands

Cron management commands.


### Create a new cron

Create a new cron.

**Usage:**

```
scw container cron create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| container-id |  | UUID of the container to invoke by the cron |
| schedule |  | UNIX cron shedule |
| args |  | Arguments to pass with the cron |
| name |  | Name of the cron to create |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Delete an existing cron

Delete the cron associated with the specified ID.

**Usage:**

```
scw container cron delete <cron-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| cron-id | Required | UUID of the cron to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get a cron

Get the cron associated with the specified ID.

**Usage:**

```
scw container cron get <cron-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| cron-id | Required | UUID of the cron to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List all your crons

List all your crons.

**Usage:**

```
scw container cron list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc` | Order of the crons |
| container-id |  | UUID of the container invoked by the cron |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



### Update an existing cron

Update the cron associated with the specified ID.

**Usage:**

```
scw container cron update <cron-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| cron-id | Required | UUID of the cron to update |
| container-id |  | UUID of the container invoked by the cron |
| schedule |  | UNIX cron schedule |
| args |  | Arguments to pass with the cron |
| name |  | Name of the cron |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



## Deploy a container

Automatically build and deploy a container.

Automatically build and deploy a container.

**Usage:**

```
scw container deploy [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name |  | Name of the application (defaults to build-source's directory name) |
| builder | Default: `paketobuildpacks/builder-jammy-base:latest` | Builder image to use |
| dockerfile | Default: `Dockerfile` | Path to the Dockerfile |
| force-builder | Default: `false` | Force the use of the builder image (even if a Dockerfile is present) |
| build-source | Default: `.` | Path to the build context |
| cache | Default: `true` | Use cache when building the image |
| build-args.{key} |  | Build-time variables |
| port | Default: `8080` | Port to expose |
| namespace-id |  | Container Namespace ID to deploy to |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



## Domain management commands

Domain management commands.


### Create a custom domain

Create a custom domain for the container with the specified ID.

**Usage:**

```
scw container domain create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| hostname |  | Domain to assign |
| container-id |  | UUID of the container to assign the domain to |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Delete a custom domain

Delete the custom domain with the specific ID.

**Usage:**

```
scw container domain delete <domain-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| domain-id | Required | UUID of the domain to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get a custom domain

Get a custom domain for the container with the specified ID.

**Usage:**

```
scw container domain get <domain-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| domain-id | Required | UUID of the domain to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List all custom domains

List all custom domains in a specified region.

**Usage:**

```
scw container domain list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc`, `hostname_asc`, `hostname_desc` | Order of the domains |
| container-id |  | UUID of the container the domain belongs to |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



## Namespace management commands

Namespace management commands.


### Create a new namespace

Create a new namespace in a specified region.

**Usage:**

```
scw container namespace create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name | Default: `<generated>` | Name of the namespace to create |
| environment-variables.{key} |  | Environment variables of the namespace to create |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| description |  | Description of the namespace to create |
| secret-environment-variables.{index}.key |  |  |
| secret-environment-variables.{index}.value |  |  |
| tags.{index} |  | Tags of the Serverless Container Namespace |
| activate-vpc-integration |  | Activate VPC integration for the namespace. |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Delete an existing namespace

Delete the namespace associated with the specified ID.

**Usage:**

```
scw container namespace delete <namespace-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | UUID of the namespace to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get a namespace

Get the namespace associated with the specified ID.

**Usage:**

```
scw container namespace get <namespace-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | UUID of the namespace to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List all your namespaces

List all namespaces in a specified region.

**Usage:**

```
scw container namespace list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc` | Order of the namespaces |
| name |  | Name of the namespaces |
| project-id |  | UUID of the Project the namespace belongs to |
| organization-id |  | UUID of the Organization the namespace belongs to |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



### Update an existing namespace

Update the space associated with the specified ID.

**Usage:**

```
scw container namespace update <namespace-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| namespace-id | Required | UUID of the namespace to update |
| environment-variables.{key} |  | Environment variables of the namespace to update |
| description |  | Description of the namespace to update |
| secret-environment-variables.{index}.key |  |  |
| secret-environment-variables.{index}.value |  |  |
| tags.{index} |  | Tags of the Serverless Container Namespace |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



## Token management commands

Token management commands.


### Create a new revocable token

Create a new revocable token.

**Usage:**

```
scw container token create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| container-id |  | UUID of the container to create the token for |
| namespace-id |  | UUID of the namespace to create the token for |
| description |  | Description of the token |
| expires-at |  | Expiry date of the token |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Delete a token

Delete a token with a specified ID.

**Usage:**

```
scw container token delete <token-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| token-id | Required | UUID of the token to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get a token

Get a token with a specified ID.

**Usage:**

```
scw container token get <token-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| token-id | Required | UUID of the token to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List all tokens

List all tokens belonging to a specified Organization or Project.

**Usage:**

```
scw container token list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc` | Order of the tokens |
| container-id |  | UUID of the container the token belongs to |
| namespace-id |  | UUID of the namespace the token belongs to |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



## Trigger management commands

Trigger management commands.


### Create a trigger

Create a new trigger for a specified container.

**Usage:**

```
scw container trigger create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name | Required | Name of the trigger |
| container-id | Required | ID of the container to trigger |
| description |  | Description of the trigger |
| scw-sqs-config.queue |  | Name of the SQS queue the trigger should listen to |
| scw-sqs-config.mnq-project-id |  | ID of the Messaging and Queuing project |
| scw-sqs-config.mnq-region |  | Region in which the Messaging and Queuing project is activated. |
| scw-nats-config.subject |  | Name of the NATS subject the trigger should listen to |
| scw-nats-config.mnq-nats-account-id |  | ID of the Messaging and Queuing NATS account |
| scw-nats-config.mnq-project-id |  | ID of the Messaging and Queuing project |
| scw-nats-config.mnq-region |  | Region in which the Messaging and Queuing project is activated. |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Delete a trigger

Delete a trigger with a specified ID.

**Usage:**

```
scw container trigger delete <trigger-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| trigger-id | Required | ID of the trigger to delete |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get a trigger

Get a trigger with a specified ID.

**Usage:**

```
scw container trigger get <trigger-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| trigger-id | Required | ID of the trigger to get |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List all triggers

List all triggers belonging to a specified Organization or Project.

**Usage:**

```
scw container trigger list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc` | Order in which to return results |
| container-id |  | ID of the container the triggers belongs to |
| namespace-id |  | ID of the namespace the triggers belongs to |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



### Update a trigger

Update a trigger with a specified ID.

**Usage:**

```
scw container trigger update <trigger-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| trigger-id | Required | ID of the trigger to update |
| name |  | Name of the trigger |
| description |  | Description of the trigger |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



