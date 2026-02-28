# sigil

Terraform provider for consistent resource naming across multiple clouds. `aws` is the default cloud profile, and `azure` is supported with Azure CAF resource coverage.

## Provider Configuration

```hcl
terraform {
  required_providers {
    sigil = {
      source  = "jesinity/sigil"
      version = "1.0.0"
    }
  }
}

provider "sigil" {
  # Optional, defaults to "aws"
  cloud = "aws"

  org_prefix = "acme"
  project    = "iac"
  env        = "dev"
  region     = "ap-southeast-2"
  # Optional: omit region for resources marked regional (default true)
  ignore_region_for_regional_resources = false

  # Optional: override just one region
  region_overrides = {
    "us-east-1" = "ueue1"
  }

  # Optional: override the full region map
  # region_map = {
  #   "us-east-1" = "use1"
  # }

  # Optional: override the default recipe
  # recipe = ["org", "proj", "env", "region", "resource", "qualifier"]

  # Optional: override style priority
  # style_priority = ["dashed", "pascal", "pascaldashed", "camel", "straight", "underscore"]
}
```

For reuse across multiple provider aliases, you can supply a base `config` object and apply `overrides`. Precedence is: `config` -> top-level attributes -> `overrides`. Top-level attributes are a shorthand for the common case.

```hcl
locals {
  sigil_config = {
    cloud      = "aws"
    org_prefix = "acme"
    project    = "iac"
    env        = "dev"
    region     = "ap-southeast-2"

    ignore_region_for_regional_resources = false
    region_overrides = {
      "us-east-1" = "ueue1"
    }
  }
}

provider "sigil" {
  config = local.sigil_config
}

provider "sigil" {
  alias  = "secondary"
  config = local.sigil_config

  overrides = {
    region = "us-east-1"
  }
}
```

Azure example (`cloud = "azure"`):

```hcl
provider "sigil" {
  cloud      = "azure"
  org_prefix = "acme"
  project    = "payments"
  env        = "prod"
  region     = "eastus2"

  # Optional Azure-specific overrides
  # resource_acronyms = {
  #   azurerm_storage_account = "st" # Use official CAF abbreviation if preferred.
  # }
}
```

**Why This Design**
Top-level attributes keep the provider fast to configure for the common single-provider case. The optional `config` + `overrides` pattern reduces repetition when you need multiple provider aliases with small differences (like region), without forcing everyone into extra nesting. The merge order is explicit so it is easy to reason about which values win.

## Data Source `sigil_mark`

`what` identifies the resource type (formerly `resource`) and drives acronyms, style overrides, and constraints. The `resource` argument is still accepted but deprecated. In recipes and outputs, the component key remains `resource` (alias `what`).

### Basic name

```hcl
data "sigil_mark" "bucket" {
  what      = "s3"
  qualifier = "mydata"
}

output "bucket_name" {
  value = data.sigil_mark.bucket.name
}
```

### Override a component

```hcl
data "sigil_mark" "bucket" {
  what      = "s3"
  qualifier = "mydata"

  overrides = {
    env    = "prod"
    region = "use1"
  }
}
```

### Custom recipe and style priority

```hcl
data "sigil_mark" "bucket" {
  what      = "s3"
  qualifier = "mydata"

  recipe         = ["org", "proj", "env", "resource", "qualifier"]
  style_priority = ["pascal", "camel", "straight", "dashed", "underscore"]
}
```

### More examples

```hcl
data "sigil_mark" "iam_role" {
  what           = "iam_role"
  qualifier      = "app"
  style_priority = ["pascal", "camel", "dashed"]
}

output "iam_role_name" {
  value = data.sigil_mark.iam_role.name
  # Example: "AcmeIacDevApse2RoleApp"
}

output "iam_role_style" {
  value = data.sigil_mark.iam_role.style
  # Example: "pascal"
}
```

```hcl
data "sigil_mark" "lambda" {
  what           = "lambda"
  qualifier      = "ingest"
  style_priority = ["underscore", "dashed"]
}

output "lambda_name" {
  value = data.sigil_mark.lambda.name
  # Example: "acme_iac_dev_apse2_lmbd_ingest"
}

output "lambda_resource_acronym" {
  value = data.sigil_mark.lambda.resource_acronym
  # Example: "lmbd"
}
```

```hcl
data "sigil_mark" "queue" {
  what      = "sqs"
  qualifier = "jobs"
}

output "queue_name" {
  value = data.sigil_mark.queue.name
  # Example: "acme-iac-dev-apse2-sqs-jobs"
}

output "queue_style" {
  value = data.sigil_mark.queue.style
  # Example: "dashed"
}
```

```hcl
data "sigil_mark" "azure_storage_account" {
  what      = "azurerm_storage_account"
  qualifier = "raw"

  # Azure storage accounts are lowercase/no-dash constrained.
  # The Azure cloud defaults select an allowed style automatically.
  recipe = ["org", "proj", "env", "resource", "qualifier"]
}

output "azure_storage_account_name" {
  value = data.sigil_mark.azure_storage_account.name
  # Example: "acmepaymentsprodstazraw" ("staz" is Sigil's 4-char normalized form of CAF "st")
}

output "azure_storage_account_style" {
  value = data.sigil_mark.azure_storage_account.style
  # Example: "straight"
}
```

## Outputs

The data source returns:
- `name`
- `style`
- `region_code`
- `resource_acronym`
- `components`
- `parts`

### Output Examples

```hcl
output "bucket_name" {
  value = data.sigil_mark.bucket.name
  # Example: "acme-iac-dev-apse2-s3b-mydata"
}

output "bucket_style" {
  value = data.sigil_mark.bucket.style
  # Example: "dashed"
}

output "bucket_region_code" {
  value = data.sigil_mark.bucket.region_code
  # Example: "apse2"
}

output "bucket_resource_acronym" {
  value = data.sigil_mark.bucket.resource_acronym
  # Example: "s3b"
}

output "bucket_parts" {
  value = data.sigil_mark.bucket.parts
  # Example: ["acme", "iac", "dev", "apse2", "s3b", "mydata"]
}

output "bucket_components" {
  value = data.sigil_mark.bucket.components
  # Example:
  # {
  #   org       = "acme"
  #   proj      = "iac"
  #   env       = "dev"
  #   region    = "apse2"
  #   resource  = "s3b"
  #   qualifier = "mydata"
  # }
}
```

## Recipe and Optional Components

The recipe is an ordered list of components. Components are only included when they have a non-empty value, so you can omit any component by removing it from the recipe or leaving it empty. Configure a default recipe at the provider level and override per data source with `recipe`.

Example: omit `env` from the name:

```hcl
provider "sigil" {
  org_prefix = "acme"
  project    = "iac"
  env        = "dev"
  region     = "ap-southeast-2"

  recipe = ["org", "proj", "region", "resource", "qualifier"]
}
```

## Region Handling

When `ignore_region_for_regional_resources` is `true` (default), the `region` component is omitted for resources marked as `regional` in the table below. Resources marked `global` keep the region component even when the flag is enabled. Set it to `false` to always include the region in names. You can still force a region per name via `overrides`. When the region is omitted, `region_code` will be empty unless overridden.

## Resource Acronyms and Scope

Default resource acronyms and scope for `cloud = "aws"`. Scope is used by `ignore_region_for_regional_resources`. You can override acronyms with `resource_acronyms`.

| Resource | Acronym | Scope |
| --- | --- | --- |
| `acm_cert` | `acmc` | `regional` |
| `alb` | `albl` | `regional` |
| `api_gateway_model` | `agmd` | `regional` |
| `api_gateway_rest_api` | `agra` | `regional` |
| `api_gateway_v2` | `agv2` | `regional` |
| `appsync` | `apsy` | `regional` |
| `athena` | `athn` | `regional` |
| `aurora_cluster` | `arcl` | `regional` |
| `autoscaling_group` | `asgr` | `regional` |
| `cloudformation_stack` | `cfst` | `regional` |
| `cloudfront` | `clfr` | `global` |
| `cloudtrail` | `ctra` | `regional` |
| `cloudwatch_alarm` | `cwal` | `regional` |
| `cloudwatch_log_group` | `cwlg` | `regional` |
| `codebuild` | `cdbd` | `regional` |
| `codedeploy` | `cddp` | `regional` |
| `codepipeline` | `cdpl` | `regional` |
| `config_rule` | `cfrl` | `regional` |
| `dynamodb` | `dydb` | `regional` |
| `dynamodb_table` | `dybt` | `regional` |
| `ebs` | `ebs` | `regional` |
| `ec2_instance` | `ec2i` | `regional` |
| `ecr` | `ecr` | `regional` |
| `ecs` | `ecs` | `regional` |
| `ecs_cluster` | `ecsc` | `regional` |
| `ecs_service` | `ecss` | `regional` |
| `ecs_task` | `ecst` | `regional` |
| `efs` | `efs` | `regional` |
| `eks` | `eks` | `regional` |
| `eks_cluster` | `eksc` | `regional` |
| `eks_node_group` | `ekng` | `regional` |
| `elastic_ip` | `elip` | `regional` |
| `elasticache` | `elch` | `regional` |
| `elasticsearch` | `elsr` | `regional` |
| `elb` | `elbl` | `regional` |
| `eventbridge_bus` | `evbb` | `regional` |
| `eventbridge_rule` | `evbr` | `regional` |
| `glue` | `glue` | `regional` |
| `guardduty` | `gdty` | `regional` |
| `iam_group` | `iamg` | `global` |
| `iam_policy` | `iamp` | `global` |
| `iam_role` | `role` | `global` |
| `iam_user` | `iamu` | `global` |
| `igw` | `igtw` | `regional` |
| `kms_key` | `kmsk` | `regional` |
| `lambda` | `lmbd` | `regional` |
| `launch_template` | `lcht` | `regional` |
| `log_group` | `logg` | `regional` |
| `msk_cluster` | `mskc` | `regional` |
| `nacl` | `nacl` | `regional` |
| `nat_gw` | `ngtw` | `regional` |
| `nlb` | `nlbl` | `regional` |
| `opensearch` | `opsr` | `regional` |
| `rds` | `rds` | `regional` |
| `rds_cluster` | `rdsc` | `regional` |
| `redshift` | `rdsh` | `regional` |
| `role` | `role` | `global` |
| `role_policy` | `rlpl` | `global` |
| `route53_record` | `r53r` | `global` |
| `route53_zone` | `rt53` | `global` |
| `route_table` | `rttb` | `regional` |
| `s3` | `s3b` | `regional` |
| `s3_access_point` | `s3ap` | `regional` |
| `s3_bucket` | `s3bk` | `regional` |
| `s3_dir` | `s3dr` | `regional` |
| `s3_object` | `s3ob` | `regional` |
| `s3_table` | `s3tb` | `regional` |
| `sagemaker` | `sgmk` | `regional` |
| `sec_group` | `scgp` | `regional` |
| `secretsmanager_secret` | `smse` | `regional` |
| `sfn` | `stfn` | `regional` |
| `snow_notification_integration` | `snti` | `regional` |
| `sns` | `sns` | `regional` |
| `sqs` | `sqs` | `regional` |
| `ssm_parameter` | `ssmp` | `regional` |
| `step_function` | `stfn` | `regional` |
| `subnet` | `subn` | `regional` |
| `target_group` | `tgpt` | `regional` |
| `vpc` | `vpcn` | `regional` |
| `wafv2_ip_set` | `wfis` | `regional` |
| `wafv2_web_acl` | `wfac` | `regional` |
| `wafv2_web_acl_rule` | `wfar` | `regional` |

## Azure CAF Acronyms and Constraints

For `cloud = "azure"`, Sigil loads **all Azure CAF resource types** from `resourceDefinition.json` and applies:
- A 4-character acronym per resource identifier.
- Acronyms are derived from CAF slugs, then normalized to a fixed 4-character token.
- Per-resource min/max/regex constraints.
- Per-resource style allowances derived from CAF dash/lowercase metadata.

Comprehensive reference (395 resource types):
- `docs/azure-caf-resources.md`
- Azure naming rules: https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules
- CAF abbreviations: https://learn.microsoft.com/en-us/azure/cloud-adoption-framework/ready/azure-best-practices/resource-abbreviations

### Supported Azure Resources and Acronyms

Supported Azure `what` values are the Azure CAF resource identifiers listed in `docs/azure-caf-resources.md`. The `Acronym` column in that table is the value returned by `resource_acronym`.

Quick reference:

| Azure Resource (`what`) | Acronym |
| --- | --- |
| `azurerm_resource_group` | `rgaz` |
| `azurerm_storage_account` | `staz` |
| `azurerm_virtual_network` | `vnet` |
| `azurerm_subnet` | `snet` |
| `azurerm_kubernetes_cluster` | `aksa` |
| `azurerm_container_registry` | `craz` |
| `azurerm_key_vault` | `kvaz` |
| `azurerm_linux_virtual_machine` | `vmaz` |

Example: CAF lists storage account as `st`; Sigil's default is `staz` because Azure mode normalizes acronyms to 4 characters. Use `resource_acronyms` overrides when you need exact CAF abbreviations.

For the complete list of all 395 supported Azure resources and acronyms, see `docs/azure-caf-resources.md`.

## Naming Styles

Style priority determines how names are formatted. If a resource has style constraints, the provider selects the first allowed style in the priority list.

Valid styles:
- `dashed`
- `underscore`
- `straight`
- `pascal`
- `pascaldashed`
- `camel`

Style behaviors:
- `dashed` Lowercase words joined by `-`.
- `underscore` Lowercase words joined by `_`.
- `straight` Lowercase words concatenated.
- `pascal` Words in `PascalCase`.
- `pascaldashed` Words in `Pascal-Case` joined by `-`.
- `camel` Words in `camelCase`.

Words are extracted from each component using the pattern `[A-Za-z0-9]+`, so punctuation or separators become word boundaries. If no valid style matches the priority list and any resource overrides, the provider falls back to `dashed`.

Cloud-specific style overrides are applied automatically:
- `aws`: `s3` and `s3_bucket` are restricted to `dashed` and `straight`.
- `azure`: each CAF resource inherits style limits from CAF dash/lowercase metadata.

## Resource Constraints

Some resources have naming constraints enforced after formatting. The constraint name matches the `what` input (case-insensitive).

The table below lists built-in `aws` constraints. Azure constraints are listed in `docs/azure-caf-resources.md`.

| Resource | Min | Max | Pattern | Notes |
| --- | --- | --- | --- | --- |
| `s3` | 3 | 63 | lowercase letters, numbers, dots, and hyphens; must start and end with a letter or number | Forbidden prefixes: `xn--`, `sthree-`, `amzn-s3-demo-`; forbidden suffixes: `-s3alias`, `--ol-s3`; forbidden substrings: `..`; disallow IPv4 |
| `s3_bucket` | 3 | 63 | lowercase letters, numbers, dots, and hyphens; must start and end with a letter or number | Forbidden prefixes: `xn--`, `sthree-`, `amzn-s3-demo-`; forbidden suffixes: `-s3alias`, `--ol-s3`; forbidden substrings: `..`; disallow IPv4 |
| `role` | 1 | 64 | alphanumeric and the following: `+=,.@_-` | none |
| `iam_role` | 1 | 64 | alphanumeric and the following: `+=,.@_-` | none |
| `iam_user` | 1 | 64 | alphanumeric and the following: `+=,.@_-` | none |
| `iam_group` | 1 | 128 | alphanumeric and the following: `+=,.@_-` | none |
| `iam_policy` | 1 | 128 | alphanumeric and the following: `+=,.@_-` | none |
| `role_policy` | 1 | 128 | alphanumeric and the following: `+=,.@_-` | none |
| `sns` | 1 | 256 | letters, numbers, underscores, and hyphens; FIFO topics must end with `.fifo` | none |
| `sns_topic` | 1 | 256 | letters, numbers, underscores, and hyphens; FIFO topics must end with `.fifo` | none |
| `sqs` | 1 | 80 | letters, numbers, underscores, and hyphens; FIFO queues must end with `.fifo` | none |
| `sqs_queue` | 1 | 80 | letters, numbers, underscores, and hyphens; FIFO queues must end with `.fifo` | none |
| `lambda` | 1 | 64 | letters, numbers, hyphens, and underscores | none |
| `kms_alias` | 1 | 256 | must begin with `alias/` and contain only letters, numbers, slashes, underscores, and hyphens | Forbidden prefix: `alias/aws/` |
| `log_group` | 1 | 512 | letters, numbers, underscore, hyphen, slash, period, and `#` | Forbidden prefix: `aws/` |
| `cloudwatch_log_group` | 1 | 512 | letters, numbers, underscore, hyphen, slash, period, and `#` | Forbidden prefix: `aws/` |
| `sec_group` | 1 | 255 | letters, numbers, spaces, and `._-:/()#,@[]+=&;{}!$*` | Forbidden prefix: `sg-` (case-insensitive) |
| `security_group` | 1 | 255 | letters, numbers, spaces, and `._-:/()#,@[]+=&;{}!$*` | Forbidden prefix: `sg-` (case-insensitive) |

Constraint types include minimum or maximum length, required pattern, forbidden prefixes or suffixes, forbidden substrings, and checks that the name is not formatted as an IPv4 address.
