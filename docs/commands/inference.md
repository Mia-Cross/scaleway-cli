<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw inference`
This API allows you to manage your Inference services.
  
- [Access Control List (ACL) management commands](#access-control-list-(acl)-management-commands)
  - [Add new ACLs](#add-new-acls)
  - [Delete an existing ACL](#delete-an-existing-acl)
  - [List your ACLs](#list-your-acls)
  - [Set new ACL](#set-new-acl)
- [Deployment commands](#deployment-commands)
  - [Create a deployment](#create-a-deployment)
  - [Delete a deployment](#delete-a-deployment)
  - [Get a deployment](#get-a-deployment)
  - [Get the CA certificate](#get-the-ca-certificate)
  - [List inference deployments](#list-inference-deployments)
  - [Update a deployment](#update-a-deployment)
- [Endpoint management commands](#endpoint-management-commands)
  - [Create an endpoint](#create-an-endpoint)
  - [Delete an endpoint](#delete-an-endpoint)
  - [Update an endpoint](#update-an-endpoint)
- [Models commands](#models-commands)
  - [Get a model](#get-a-model)
  - [List models](#list-models)
- [Node types management commands](#node-types-management-commands)
  - [List available node types](#list-available-node-types)

  
## Access Control List (ACL) management commands

Access Control List (ACL) management commands.


### Add new ACLs

Add new ACL rules for a specific deployment.

**Usage:**

```
scw inference acl add <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to add ACL rules to |
| acls.{index}.ip |  | IP address to be allowed |
| acls.{index}.description |  | Description of the ACL rule |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Delete an existing ACL

Delete an existing ACL.

**Usage:**

```
scw inference acl delete <acl-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| acl-id | Required | ID of the ACL rule to delete |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List your ACLs

List ACLs for a specific deployment.

**Usage:**

```
scw inference acl list <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to list ACL rules for |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



### Set new ACL

Set new ACL rules for a specific deployment.

**Usage:**

```
scw inference acl set <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to set ACL rules for |
| acls.{index}.ip |  | IP address to be allowed |
| acls.{index}.description |  | Description of the ACL rule |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



## Deployment commands

Deployment commands.


### Create a deployment

Create a new inference deployment related to a specific model.

**Usage:**

```
scw inference deployment create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name | Required<br />Default: `<generated>` | Name of the deployment |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| model-name | Required | Name of the model to use |
| accept-eula |  | Accept the model's End User License Agreement (EULA). |
| node-type | Required | Name of the node type to use |
| tags.{index} |  | List of tags to apply to the deployment |
| min-size |  | Defines the minimum size of the pool |
| max-size |  | Defines the maximum size of the pool |
| endpoints.{index}.is-public | Default: `false` | Will configure your public endpoint if true |
| endpoints.{index}.private-network.private-network-id |  | ID of the Private Network |
| endpoints.{index}.disable-auth | Default: `false` | Disable the authentication on the endpoint. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Delete a deployment

Delete an existing inference deployment.

**Usage:**

```
scw inference deployment delete <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to delete |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get a deployment

Get the deployment for the given ID.

**Usage:**

```
scw inference deployment get <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to get |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Get the CA certificate

Get the CA certificate used for the deployment of private endpoints.
The CA certificate will be returned as a PEM file.

**Usage:**

```
scw inference deployment get-certificate [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required |  |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List inference deployments

List all your inference deployments.

**Usage:**

```
scw inference deployment list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_desc`, `created_at_asc`, `name_asc`, `name_desc` | Order in which to return results |
| project-id |  | Filter by Project ID |
| name |  | Filter by deployment name |
| tags.{index} |  | Filter by tags |
| organization-id |  | Filter by Organization ID |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



### Update a deployment

Update an existing inference deployment.

**Usage:**

```
scw inference deployment update <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to update |
| name |  | Name of the deployment |
| tags.{index} |  | List of tags to apply to the deployment |
| min-size |  | Defines the new minimum size of the pool |
| max-size |  | Defines the new maximum size of the pool |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



## Endpoint management commands

Endpoint management commands.


### Create an endpoint

Create a new Endpoint related to a specific deployment.

**Usage:**

```
scw inference endpoint create <deployment-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| deployment-id | Required | ID of the deployment to create the endpoint for |
| endpoint.is-public | Default: `false` | Will configure your public endpoint if true |
| endpoint.private-network.private-network-id |  | ID of the Private Network |
| endpoint.disable-auth | Default: `false` | Disable the authentication on the endpoint. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Delete an endpoint

Delete an existing Endpoint.

**Usage:**

```
scw inference endpoint delete <endpoint-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| endpoint-id | Required | ID of the endpoint to delete |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### Update an endpoint

Update an existing Endpoint.

**Usage:**

```
scw inference endpoint update <endpoint-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| endpoint-id | Required | ID of the endpoint to update |
| disable-auth |  | Disable the authentication on the endpoint. |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



## Models commands

Models commands.


### Get a model

Get the model for the given ID.

**Usage:**

```
scw inference model get <model-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| model-id | Required | ID of the model to get |
| region | Default: `fr-par`<br />One of: `fr-par` | Region to target. If none is passed will use default region from the config |



### List models

List all available models.

**Usage:**

```
scw inference model list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `display_rank_asc`, `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc` | Order in which to return results |
| project-id |  | Filter by Project ID |
| name |  | Filter by model name |
| tags.{index} |  | Filter by tags |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



## Node types management commands

Node types management commands.


### List available node types

List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.

**Usage:**

```
scw inference node-type list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| include-disabled-types |  | Include disabled node types in the response |
| region | Default: `fr-par`<br />One of: `fr-par`, `all` | Region to target. If none is passed will use default region from the config |



