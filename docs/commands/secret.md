<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw secret`
This API allows you to manage your Secret Manager services, for storing, accessing and sharing sensitive data such as passwords, API keys and certificates.
  
- [Secret management commands](#secret-management-commands)
  - [Allow a product to use the secret](#allow-a-product-to-use-the-secret)
  - [Create a secret](#create-a-secret)
  - [Delete a secret](#delete-a-secret)
  - [Get metadata using the secret's ID](#get-metadata-using-the-secret's-id)
  - [List secrets](#list-secrets)
  - [Enable secret protection](#enable-secret-protection)
  - [Disable secret protection](#disable-secret-protection)
  - [Update metadata of a secret](#update-metadata-of-a-secret)
- [Secret Version management commands](#secret-version-management-commands)
  - [Access a secret's version using the secret's ID](#access-a-secret's-version-using-the-secret's-id)
  - [Access a secret's version using the secret's name and path](#access-a-secret's-version-using-the-secret's-name-and-path)
  - [Create a version](#create-a-version)
  - [Delete a version](#delete-a-version)
  - [Disable a version](#disable-a-version)
  - [Enable a version](#enable-a-version)
  - [Get metadata of a secret's version using the secret's ID](#get-metadata-of-a-secret's-version-using-the-secret's-id)
  - [List versions of a secret using the secret's ID](#list-versions-of-a-secret-using-the-secret's-id)
  - [Update metadata of a version](#update-metadata-of-a-version)

  
## Secret management commands

Secrets are logical containers made up of zero or more immutable versions, that contain sensitive data.


### Allow a product to use the secret

Allow a product to use the secret.

**Usage:**

```
scw secret secret add-owner <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| product | One of: `unknown_product`, `edge_services` | ID of the product to add |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Create a secret

Create a secret in a given region specified by the `region` parameter.

**Usage:**

```
scw secret secret create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| name |  | Name of the secret |
| tags.{index} |  | List of the secret's tags |
| description |  | Description of the secret |
| type | One of: `unknown_type`, `opaque`, `certificate`, `key_value`, `basic_credentials`, `database_credentials`, `ssh_key` | Type of the secret |
| path |  | Path of the secret |
| ephemeral-policy.time-to-live |  | Time frame, from one second and up to one year, during which the secret's versions are valid. |
| ephemeral-policy.expires-once-accessed |  | Returns `true` if the version expires after a single user access. |
| ephemeral-policy.action | One of: `unknown_action`, `delete`, `disable` | Action to perform when the version of a secret expires |
| protected |  | Returns `true` if secret protection is applied to a given secret |
| key-id |  | ID of the Scaleway Key Manager key |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Create a given secret
```
scw secret secret create name=foobar description="$(cat <path/to/your/secret>)"
```




### Delete a secret

Delete a given secret specified by the `region` and `secret_id` parameters.

**Usage:**

```
scw secret secret delete <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Delete a given secret
```
scw secret secret delete 11111111-1111-1111-1111-111111111111
```




### Get metadata using the secret's ID

Retrieve the metadata of a secret specified by the `region` and `secret_id` parameters.

**Usage:**

```
scw secret secret get <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List secrets

Retrieve the list of secrets created within an Organization and/or Project. You must specify either the `organization_id` or the `project_id` and the `region`.

**Usage:**

```
scw secret secret list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| project-id |  | Filter by Project ID (optional) |
| order-by | One of: `name_asc`, `name_desc`, `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc` |  |
| tags.{index} |  | List of tags to filter on (optional) |
| name |  | Filter by secret name (optional) |
| path |  | Filter by exact path (optional) |
| ephemeral |  | Filter by ephemeral / not ephemeral (optional) |
| type | One of: `unknown_type`, `opaque`, `certificate`, `key_value`, `basic_credentials`, `database_credentials`, `ssh_key` | Filter by secret type (optional) |
| scheduled-for-deletion |  | Filter by whether the secret was scheduled for deletion / not scheduled for deletion. By default, it will display only not scheduled for deletion secrets. |
| organization-id |  | Filter by Organization ID (optional) |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



### Enable secret protection

Enable secret protection for a given secret specified by the `secret_id` parameter. Enabling secret protection means that your secret can be read and modified, but it cannot be deleted.

**Usage:**

```
scw secret secret protect <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret to enable secret protection for |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Enable secret protection
```
scw secret secret protect 11111111-1111-1111-1111-111111111111
```




### Disable secret protection

Disable secret protection for a given secret specified by the `secret_id` parameter. Disabling secret protection means that your secret can be read, modified and deleted.

**Usage:**

```
scw secret secret unprotect <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret to disable secret protection for |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Disable secret protection
```
scw secret secret unprotect 11111111-1111-1111-1111-111111111111
```




### Update metadata of a secret

Edit a secret's metadata such as name, tag(s), description and ephemeral policy. The secret to update is specified by the `secret_id` and `region` parameters.

**Usage:**

```
scw secret secret update <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| name |  | Secret's updated name (optional) |
| tags.{index} |  | Secret's updated list of tags (optional) |
| description |  | Description of the secret |
| path |  | Path of the folder |
| ephemeral-policy.time-to-live |  | Time frame, from one second and up to one year, during which the secret's versions are valid. |
| ephemeral-policy.expires-once-accessed |  | Returns `true` if the version expires after a single user access. |
| ephemeral-policy.action | One of: `unknown_action`, `delete`, `disable` | Action to perform when the version of a secret expires |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



## Secret Version management commands

Versions store the sensitive data contained in your secrets (API keys, passwords, or certificates).


### Access a secret's version using the secret's ID

Access sensitive data in a secret's version specified by the `region`, `secret_id` and `revision` parameters.

**Usage:**

```
scw secret version access <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| revision | Required | Version number |
| field |  | Return only the JSON field of the given name |
| raw |  | Return only the raw payload |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Get a raw json value from a secret version
```
scw secret version access 11111111-1111-1111-111111111111 revision=1 field=key raw=true
```




### Access a secret's version using the secret's name and path

Access sensitive data in a secret's version specified by the `region`, `secret_name`, `secret_path` and `revision` parameters.

**Usage:**

```
scw secret version access-by-path [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-path |  | Secret's path |
| secret-name |  | Secret's name |
| revision | Required | Version number |
| project-id |  | Project ID to use. If none is passed the default project ID will be used |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Create a version

Create a version of a given secret specified by the `region` and `secret_id` parameters.

**Usage:**

```
scw secret version create <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| data | Required | Content of the secret version. |
| description |  | Description of the version |
| disable-previous |  | Disable the previous secret version |
| data-crc32 |  | (Optional.) The CRC32 checksum of the data as a base-10 integer |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Create a json secret version
```
scw secret version create 11111111-1111-1111-111111111111 data={"key":"value"}
```




### Delete a version

Delete a secret's version and the sensitive data contained in it. Deleting a version is permanent and cannot be undone.

**Usage:**

```
scw secret version delete <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| revision | Required | Version number |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |


**Examples:**


Delete a given Secret Version
```
scw secret version delete 11111111-1111-1111-1111-111111111111 revision=1
```




### Disable a version

Make a specific version inaccessible. You must specify the `region`, `secret_id` and `revision` parameters.

**Usage:**

```
scw secret version disable <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| revision | Required | Version number |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Enable a version

Make a specific version accessible. You must specify the `region`, `secret_id` and `revision` parameters.

**Usage:**

```
scw secret version enable <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| revision | Required | Version number |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### Get metadata of a secret's version using the secret's ID

Retrieve the metadata of a secret's given version specified by the `region`, `secret_id` and `revision` parameters.

**Usage:**

```
scw secret version get <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| revision | Required | Version number |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



### List versions of a secret using the secret's ID

Retrieve the list of a given secret's versions specified by the `secret_id` and `region` parameters.

**Usage:**

```
scw secret version list <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| status.{index} | One of: `unknown_status`, `enabled`, `disabled`, `deleted`, `scheduled_for_deletion` | Filter results by status |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw`, `all` | Region to target. If none is passed will use default region from the config |



### Update metadata of a version

Edit the metadata of a secret's given version, specified by the `region`, `secret_id` and `revision` parameters.

**Usage:**

```
scw secret version update <secret-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| secret-id | Required | ID of the secret |
| revision | Required | Version number |
| description |  | Description of the version |
| ephemeral-properties.expires-at |  | The version's expiration date |
| ephemeral-properties.expires-once-accessed |  | Returns `true` if the version expires after a single user access. |
| ephemeral-properties.action | One of: `unknown_action`, `delete`, `disable` | Action to perform when the version of a secret expires |
| region | Default: `fr-par`<br />One of: `fr-par`, `nl-ams`, `pl-waw` | Region to target. If none is passed will use default region from the config |



