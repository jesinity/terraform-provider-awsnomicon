# Azure CAF Resources

Comprehensive Azure defaults used by `cloud = "azure"`.
All resource identifiers come from the Azure CAF provider resource catalog (`resourceDefinition.json`).

- Total resource types: 395
- Acronym policy in Sigil Azure mode: defaults to fixed 4-character acronyms
- Default acronyms are derived from CAF slugs, then normalized to 4 characters (for example, storage account `st` becomes `stac`)
- Optional `use_azure_caf_acronyms = true` to use CAF slugs directly as defaults
- Constraints source: Azure CAF definitions (which mirror Azure ARM naming restrictions)

References:
- Azure CAF provider catalog: https://github.com/aztfmod/terraform-provider-azurecaf/blob/main/resourceDefinition.json
- Azure ARM naming rules: https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules
- CAF naming guidance and abbreviations: https://learn.microsoft.com/en-us/azure/cloud-adoption-framework/ready/azure-best-practices/resource-abbreviations

| Resource | Acronym | CAF Scope | Sigil Scope | Min | Max | Dashes | Lowercase Only | Regex | Allowed Styles |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `aks_node_pool_linux` | `npla` | `parent` | `regional` | 1 | 12 | `false` | `false` | `^[a-z][0-9a-z]{0,11}$` | `pascal, camel, straight` |
| `aks_node_pool_windows` | `npwa` | `parent` | `regional` | 1 | 6 | `false` | `false` | `^[a-z][0-9a-z]{0,5}$` | `pascal, camel, straight` |
| `azurerm_aadb2c_directory` | `aadb` | `global` | `non-regional` | 1 | 75 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,73}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_analysis_services_server` | `sese` | `resourceGroup` | `regional` | 3 | 63 | `false` | `true` | `^[a-z][a-z0-9]{2,62}$` | `straight` |
| `azurerm_api_management` | `apim` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,48}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_api` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_api_operation_tag` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_backend` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_certificate` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_gateway` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_group` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_logger` | `apim` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_service` | `apim` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-z][a-zA-Z0-9-]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_configuration` | `appc` | `resourceGroup` | `regional` | 5 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{5,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_service` | `appa` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_service_environment` | `asea` | `resourceGroup` | `regional` | 2 | 36 | `true` | `false` | `^[0-9A-Za-z-]{2,36}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_service_plan` | `plan` | `resourceGroup` | `regional` | 1 | 40 | `true` | `false` | `^[0-9A-Za-z-]{1,40}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_gateway` | `agwa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_insights` | `appi` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^%&\\?/. ][^%&\\?/]{0,258}[^%&\\?/. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_insights_web_test` | `appi` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9- ]{0,62}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_security_group` | `asga` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_account` | `auac` | `resourceGroup` | `regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_certificate` | `aace` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_credential` | `aacr` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_hybrid_runbook_worker_group` | `aahw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^([^<>*%&:\\?.+/#\\s]?[ ]?){0,127}[^<>*%&:\\?.+/#\\s]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_job_schedule` | `aajs` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_runbook` | `aaru` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,62}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_schedule` | `aasc` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_variable` | `aava` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_availability_set` | `avai` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bastion_host` | `bast` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_batch_account` | `baac` | `region` | `regional` | 3 | 24 | `false` | `true` | `^[a-z0-9]{3,24}$` | `straight` |
| `azurerm_batch_application` | `baap` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9_-]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_batch_certificate` | `bace` | `parent` | `regional` | 5 | 45 | `true` | `false` | `^[a-zA-Z0-9_-]{5,45}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_batch_pool` | `bapo` | `parent` | `regional` | 3 | 24 | `true` | `false` | `^[a-zA-Z0-9_-]{1,24}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_Email` | `botm` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_directline` | `botl` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_ms_teams` | `bott` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_slack` | `bots` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channels_registration` | `botc` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_connection` | `botc` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_service_azure_bot` | `bota` | `global` | `non-regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_web_app` | `bota` | `global` | `non-regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_endpoint` | `cdna` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_custom_domain` | `cfdc` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,258}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_endpoint` | `cfde` | `global` | `non-regional` | 1 | 46 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,44}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_firewall_policy` | `cfdf` | `resourceGroup` | `regional` | 1 | 128 | `false` | `false` | `^[a-zA-Z][0-9a-zA-Z]{0,127}$` | `pascal, camel, straight` |
| `azurerm_cdn_frontdoor_origin` | `cfdo` | `parent` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_origin_group` | `cfdo` | `parent` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_profile` | `cfdp` | `resourceGroup` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_route` | `cfdr` | `parent` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_rule` | `cfdr` | `parent` | `regional` | 1 | 60 | `false` | `false` | `^[a-zA-Z][a-zA-Z0-9]{0,59}$` | `pascal, camel, straight` |
| `azurerm_cdn_frontdoor_rule_set` | `cfdr` | `parent` | `regional` | 1 | 60 | `false` | `false` | `^[a-zA-Z][a-zA-Z0-9]{0,59}$` | `pascal, camel, straight` |
| `azurerm_cdn_frontdoor_secret` | `cfds` | `parent` | `regional` | 2 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_security_policy` | `cfds` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[0-9A-Za-z_-]{1,260}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_profile` | `cdnp` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cognitive_account` | `coga` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cognitive_deployment` | `coga` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_communication_service` | `acsa` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9_-]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_consumption_budget_resource_group` | `acbr` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_consumption_budget_subscription` | `acbs` | `subscription` | `non-regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_containerGroups` | `cont` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_container_app` | `coap` | `resourceGroup` | `regional` | 1 | 32 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,30}[a-z0-9]$` | `dashed, straight` |
| `azurerm_container_app_environment` | `caea` | `resourceGroup` | `regional` | 1 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_container_registry` | `core` | `resourceGroup` | `regional` | 1 | 63 | `false` | `true` | `^[a-zA-Z0-9]{1,63}$` | `straight` |
| `azurerm_container_registry_webhook` | `crwh` | `resourceGroup` | `regional` | 1 | 50 | `false` | `false` | `^[a-zA-Z0-9]{1,50}$` | `pascal, camel, straight` |
| `azurerm_cosmosdb_account` | `cosm` | `resourceGroup` | `regional` | 1 | 44 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,42}[a-z0-9]$` | `dashed, straight` |
| `azurerm_custom_provider` | `prov` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[^&%?\\/]{2,63}[^&%.?\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dashboard` | `dsba` | `parent` | `regional` | 3 | 160 | `true` | `false` | `^[a-zA-Z0-9-]{3,160}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory` | `adfa` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_azure_blob` | `adfb` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_cosmosdb_sqlapi` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_delimited_text` | `adfd` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_http` | `adfh` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_json` | `adfj` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_mysql` | `adfm` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_postgresql` | `adfp` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_sql_server_table` | `adfm` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_integration_runtime_managed` | `adfi` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_blob_storage` | `adfl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_databricks` | `adfl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_function` | `adfl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_sql_database` | `adfl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_cosmosdb` | `adfl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_data_lake_storage_gen2` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_key_vault` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_mysql` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_postgresql` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_sftp` | `adfl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_sql_server` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_web` | `adfs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_pipeline` | `adfp` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_trigger_schedule` | `adft` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_lake_analytics_account` | `dlaa` | `global` | `non-regional` | 3 | 24 | `false` | `false` | `^[a-z0-9]{3,24}$` | `pascal, camel, straight` |
| `azurerm_data_lake_analytics_firewall_rule` | `dlfw` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-z0-9-_]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_lake_store` | `dlsa` | `parent` | `regional` | 3 | 24 | `false` | `false` | `^[a-z0-9]{3,24}$` | `pascal, camel, straight` |
| `azurerm_data_lake_store_firewall_rule` | `dlsf` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-_]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_blob_storage` | `dpbp` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_disk` | `dpbp` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_postgresql` | `dpbp` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_postgresql_flexible_server` | `dpbp` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_vault` | `dpbv` | `resourceGroup` | `regional` | 2 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{1,49}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_database_migration_project` | `migr` | `parent` | `regional` | 2 | 57 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,56}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_database_migration_service` | `dmsa` | `resourceGroup` | `regional` | 2 | 62 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,61}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_databricks_workspace` | `dbwa` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[a-zA-Z0-9-_]{3,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dedicated_host` | `deho` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dedicated_host_group` | `dhga` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center` | `dece` | `resourceGroup` | `regional` | 3 | 26 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,24}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_catalog` | `dcca` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_dev_box_definition` | `dcdb` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_environment_type` | `dcet` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_gallery` | `dcga` | `parent` | `regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9_.]{0,78}[a-zA-Z0-9])?$` | `pascal, camel, straight` |
| `azurerm_dev_center_network_connection` | `dcnc` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_project` | `dcpa` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_project_environment_type` | `dcpe` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_test_lab` | `laba` | `resourceGroup` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9-_]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_test_linux_virtual_machine` | `labv` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_test_windows_virtual_machine` | `labv` | `parent` | `regional` | 1 | 15 | `true` | `false` | `^[a-zA-Z0-9-]{1,15}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_endpoint_eventgrid` | `adte` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9_-]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_endpoint_eventhub` | `adte` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9_-]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_endpoint_servicebus` | `adts` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9_-]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_instance` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_disk_encryption_set` | `desa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9-_]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_a_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_aaaa_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_caa_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_cname_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_mx_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_ns_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_ptr_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_txt_record` | `dnsr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_zone` | `dnsa` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_domain` | `egda` | `resourceGroup` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_domain_topic` | `egdt` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_event_subscription` | `egsa` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[a-zA-Z0-9-]{3,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_topic` | `egta` | `resourceGroup` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub` | `evha` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_authorization_rule` | `ehar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_consumer_group` | `ehcg` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_namespace` | `ehna` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_namespace_authorization_rule` | `ehna` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_namespace_disaster_recovery_config` | `ehdr` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_express_route_circuit` | `erca` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_express_route_gateway` | `ergw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_federated_identity_credential` | `fedc` | `parent` | `regional` | 3 | 120 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_-]{2,119}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall` | `fire` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_application_rule_collection` | `fwap` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_ip_configuration` | `fwip` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_nat_rule_collection` | `fwna` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_network_rule_collection` | `fwne` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_policy` | `afwp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_frontdoor` | `fron` | `global` | `non-regional` | 5 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{3,62}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_frontdoor_firewall_policy` | `fdfw` | `global` | `non-regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9]{0,78}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_function_app` | `fuap` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_function_app_slot` | `fasa` | `global` | `non-regional` | 2 | 59 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,57}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_hadoop_cluster` | `hado` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_hbase_cluster` | `hbas` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_interactive_query_cluster` | `iqra` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_kafka_cluster` | `kafk` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_ml_services_cluster` | `mlsa` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_rserver_cluster` | `rser` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_spark_cluster` | `spar` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_storm_cluster` | `stor` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_healthcare_dicom_service` | `dico` | `parent` | `regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_fhir_service` | `fhir` | `parent` | `regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_medtech_service` | `medt` | `parent` | `regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_service` | `hcas` | `global` | `non-regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_workspace` | `hcwa` | `global` | `non-regional` | 3 | 24 | `false` | `true` | `^[a-z0-9]{3,24}$` | `straight` |
| `azurerm_image` | `imga` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_integration_service_environment` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\-\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iot_security_device_group` | `iotd` | `parent` | `regional` | 1 | 32 | `true` | `false` | `^[a-zA-Z0-9-._]{1,32}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iot_security_solution` | `iots` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9-_]{1,260}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iotcentral_application` | `iota` | `global` | `non-regional` | 2 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_iothub` | `iota` | `global` | `non-regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,48}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_certificate` | `iotc` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_consumer_group` | `iotc` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9-._]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_dps` | `dpsa` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[a-zA-Z0-9-]{1,63}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_dps_certificate` | `dpsc` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_dps_shared_access_policy` | `dpss` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_shared_access_policy` | `iots` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_ip_group` | `ipgr` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault` | `keva` | `global` | `non-regional` | 3 | 24 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{1,22}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault_certificate` | `kvca` | `parent` | `regional` | 1 | 127 | `true` | `false` | `^[a-zA-Z0-9-]{1,127}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault_key` | `kvka` | `parent` | `regional` | 1 | 127 | `true` | `false` | `^[a-zA-Z0-9-]{1,127}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault_secret` | `kvsa` | `parent` | `regional` | 1 | 127 | `true` | `false` | `^[a-zA-Z0-9-]{1,127}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_kubernetes_cluster` | `aksa` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_kubernetes_fleet_manager` | `flee` | `resourceGroup` | `regional` | 1 | 63 | `true` | `true` | `^[0-9a-z]([0-9a-z-]{0,61}[0-9a-z])?$` | `dashed, straight` |
| `azurerm_kusto_cluster` | `kucl` | `global` | `non-regional` | 4 | 22 | `false` | `false` | `^[a-z][a-z0-9]{3,21}$` | `pascal, camel, straight` |
| `azurerm_kusto_database` | `kdba` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9- .]{1,260}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_kusto_eventhub_data_connection` | `kehc` | `parent` | `regional` | 1 | 40 | `true` | `false` | `^[a-zA-Z0-9- .]{1,40}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb` | `lbxx` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_backend_address_pool` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_backend_pool` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_nat_pool` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_nat_rule` | `lbna` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_outbound_rule` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_probe` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_rule` | `adta` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_linux_virtual_machine` | `vima` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,62}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_linux_virtual_machine_scale_set` | `vmss` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,62}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_linux_web_app` | `lwap` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_load_test` | `load` | `global` | `non-regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{0,62}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_local_network_gateway` | `lgwa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_cluster` | `logc` | `resourceGroup` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_query_pack` | `laqp` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_solution` | `lasa` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_storage_insights` | `lasi` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_workspace` | `loga` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_action_custom` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_action_http` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_integration_account` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_trigger_custom` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_trigger_http_request` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_trigger_recurrence` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_workflow` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_machine_learning_compute_instance` | `amlc` | `parent` | `regional` | 1 | 16 | `true` | `false` | `^[a-zA-Z0-9][a-z0-9-]{0,14}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_machine_learning_workspace` | `mlwa` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,259}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_maintenance_configuration` | `mcfa` | `resourceGroup` | `regional` | 1 | 60 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,58}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_managed_disk` | `dska` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_maps_account` | `mapa` | `resourceGroup` | `regional` | 1 | 98 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,97}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_database` | `mari` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_firewall_rule` | `mari` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_server` | `mari` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_virtual_network_rule` | `mari` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_action_group` | `amag` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^|:<>+#%&\\?/]{0,259}[^|:<>+#%&\\?/. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_activity_log_alert` | `adfm` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:&?#\\+\\/]{0,259}[^<>*%:&.?#\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_autoscale_setting` | `amas` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,62}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_data_collection_endpoint` | `dcea` | `resourceGroup` | `regional` | 3 | 44 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,42}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_data_collection_rule` | `dcra` | `resourceGroup` | `regional` | 3 | 44 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,42}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_diagnostic_setting` | `amds` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_metric_alert` | `meal` | `resourceGroup` | `regional` | 1 | 251 | `true` | `false` | `^[^<>*%&:\\?+/#@{}]{0,250}[^<>*%&:\\?+/#@{}. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_private_link_scope` | `ampl` | `resourceGroup` | `regional` | 1 | 255 | `true` | `false` | `^[0-9A-Za-z-._()]{0,254}[0-9A-Za-z-_()]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_scheduled_query_rules_alert` | `schq` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%&:\\?/#{}]{0,259}[^<>*%&:\\?/#{}. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mssql_database` | `sqld` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{1,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mssql_elasticpool` | `sqle` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{1,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mssql_mi` | `sqlm` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_mssql_server` | `sqla` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_mysql_database` | `mysq` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_firewall_rule` | `mysq` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_flexible_server` | `mysq` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_flexible_server_database` | `mysq` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_flexible_server_firewall_rule` | `mysq` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_server` | `mysq` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_virtual_network_rule` | `mysq` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_account` | `anaa` | `resourceGroup` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,126}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_pool` | `anpa` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_snapshot` | `ansa` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_volume` | `anva` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_ddos_protection_plan` | `ddos` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_interface` | `nica` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_security_group` | `nsga` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_security_group_rule` | `nsgr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_security_rule` | `nsgr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_watcher` | `newa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_nginx_deployment` | `ngin` | `resourceGroup` | `regional` | 1 | 30 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,28}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_notification_hub` | `nohu` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_notification_hub_authorization_rule` | `dnsr` | `parent` | `regional` | 1 | 256 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,255}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_notification_hub_namespace` | `dnsr` | `global` | `non-regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_point_to_site_vpn_gateway` | `vpng` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_portal_dashboard` | `dsba` | `parent` | `regional` | 3 | 160 | `true` | `false` | `^[a-zA-Z0-9-]{3,160}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_database` | `psql` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_firewall_rule` | `psql` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_flexible_server` | `psql` | `global` | `non-regional` | 3 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_postgresql_flexible_server_database` | `psql` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_flexible_server_firewall_rule` | `psql` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_server` | `psql` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_virtual_network_rule` | `psql` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_powerbi_embedded` | `pbia` | `region` | `regional` | 3 | 63 | `false` | `false` | `^[a-z0-9][a-z0-9]{2,62}$` | `pascal, camel, straight` |
| `azurerm_private_dns_a_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_aaaa_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_cname_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_mx_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_ptr_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver` | `dnsp` | `resourceGroup` | `regional` | 3 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{1,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_dns_forwarding_ruleset` | `dnsf` | `resourceGroup` | `regional` | 2 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{0,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_forwarding_rule` | `dnsf` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-_]{0,78}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_inbound_endpoint` | `dnsp` | `parent` | `regional` | 3 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{1,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_outbound_endpoint` | `dnsp` | `parent` | `regional` | 3 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{1,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_virtual_network_link` | `dnsf` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-_]{0,78}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_srv_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_txt_record` | `pdns` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_zone` | `pdns` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_zone_group` | `pdns` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_zone_virtual_network_link` | `pnet` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_endpoint` | `pren` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,62}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_link_service` | `plsa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_service_connection` | `psca` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_proximity_placement_group` | `ppga` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_public_ip` | `pipa` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_public_ip_prefix` | `pipp` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_purview_account` | `purv` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9_-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_recovery_services_vault` | `rsva` | `resourceGroup` | `regional` | 2 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{1,49}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_recovery_services_vault_backup_police` | `rsvb` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_redhat_openshift_cluster` | `aroc` | `resourceGroup` | `regional` | 1 | 30 | `false` | `false` | `^[a-zA-Z0-9]{1,30}$` | `pascal, camel, straight` |
| `azurerm_redhat_openshift_domain` | `arod` | `resourceGroup` | `regional` | 1 | 30 | `false` | `false` | `^[a-zA-Z0-9]{1,30}$` | `pascal, camel, straight` |
| `azurerm_redis_cache` | `redi` | `global` | `non-regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_redis_firewall_rule` | `redi` | `parent` | `regional` | 1 | 256 | `false` | `false` | `^[a-zA-Z0-9]{1,256}$` | `pascal, camel, straight` |
| `azurerm_relay_hybrid_connection` | `rlhc` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_relay_namespace` | `rlna` | `global` | `non-regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_resource_group` | `regr` | `subscription` | `non-regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9-._\\(\\)]{0,89}[a-zA-Z0-9-_\\(\\)]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_resource_group_policy_assignment` | `argp` | `resourceGroup` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,126}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_role_assignment` | `roas` | `assignment` | `non-regional` | 1 | 64 | `true` | `false` | `^[^%]{0,63}[^ %.]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_role_definition` | `rode` | `definition` | `non-regional` | 1 | 64 | `true` | `false` | `^[^%]{0,63}[^ %.]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_route` | `rout` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_route_server` | `rtsa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_route_table` | `rout` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_search_service` | `srch` | `global` | `non-regional` | 2 | 60 | `true` | `true` | `^[a-z0-9](?:[a-z0-9-]{0,58}[a-z0-9])?$` | `dashed, straight` |
| `azurerm_service_fabric_cluster` | `facl` | `region` | `regional` | 4 | 23 | `true` | `true` | `^[a-z][a-z0-9-]{2,21}[a-z0-9]$` | `dashed, straight` |
| `azurerm_servicebus_namespace` | `sena` | `global` | `non-regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_namespace_authorization_rule` | `sbar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_namespace_disaster_recovery_config` | `sbdr` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_queue` | `sbqa` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_queue_authorization_rule` | `sbqa` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_subscription` | `sbsa` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_subscription_rule` | `sbsr` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_topic` | `sbta` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_topic_authorization_rule` | `sbta` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_shared_image` | `shim` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_shared_image_gallery` | `siga` | `resourceGroup` | `regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9._]{0,78}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_signalr_service` | `sgnl` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_snapshots` | `snap` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_sql_elasticpool` | `sqle` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{1,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_sql_failover_group` | `sqlf` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_sql_firewall_rule` | `sqlf` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:?\\+\\/]{1,127}[^<>*%:.?\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_sql_server` | `sqla` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_static_site` | `stap` | `resourceGroup` | `regional` | 1 | 40 | `true` | `false` | `^[a-zA-Z0-9-]{1,40}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_account` | `stac` | `global` | `non-regional` | 3 | 24 | `false` | `true` | `^[a-z0-9]{3,24}$` | `straight` |
| `azurerm_storage_blob` | `blob` | `parent` | `regional` | 1 | 1024 | `true` | `false` | `^[^\\s\\/$#&]{1,1000}[^\\s\\/$#&]{0,24}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_container` | `stct` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{2,62}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_data_lake_gen2_filesystem` | `stdl` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_queue` | `stqa` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_share` | `stsa` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_share_directory` | `stsa` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_sync` | `stsy` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,259}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_sync_group` | `stsg` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,259}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_table` | `stta` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_function_javascript_udf` | `asaf` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_job` | `asaa` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_blob` | `asao` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_eventhub` | `asao` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_mssql` | `asao` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_servicebus_queue` | `asao` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_servicebus_topic` | `asao` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_reference_input_blob` | `asar` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_stream_input_blob` | `asai` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_stream_input_eventhub` | `asai` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_stream_input_iothub` | `asai` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_subnet` | `snet` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_subscription_policy_assignment` | `aspa` | `subscription` | `non-regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,126}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_firewall_rule` | `syfw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:?\\+\\/]{1,127}[^<>*%:.?\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_integration_runtime_azure` | `syni` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_integration_runtime_self_hosted` | `syni` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9_-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_linked_service` | `synl` | `subscription` | `non-regional` | 1 | 140 | `false` | `false` | `^[a-zA-Z0-9_]{1,140}$` | `pascal, camel, straight` |
| `azurerm_synapse_managed_private_endpoint` | `synm` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_.-]{1,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_private_link_hub` | `synp` | `subscription` | `non-regional` | 1 | 45 | `false` | `true` | `^[a-z0-9]{1,45}$` | `straight` |
| `azurerm_synapse_spark_pool` | `sysp` | `parent` | `regional` | 1 | 15 | `false` | `true` | `^[0-9a-zA-Z]{1,15}$` | `straight` |
| `azurerm_synapse_sql_pool` | `syns` | `subscription` | `non-regional` | 3 | 15 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_]{1,13}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_synapse_sql_pool_vulnerability_assessment_baseline` | `syns` | `subscription` | `non-regional` | 3 | 63 | `false` | `false` | `^[a-zA-Z0-9]{1,63}$` | `pascal, camel, straight` |
| `azurerm_synapse_sql_pool_workload_classifier` | `syns` | `subscription` | `non-regional` | 1 | 128 | `true` | `false` | `^[^<>*%:?\\+\\/]{1,127}[^<>*%:.?\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_sql_pool_workload_group` | `syns` | `subscription` | `non-regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_workspace` | `syws` | `resourceGroup` | `regional` | 1 | 45 | `false` | `true` | `^[0-9a-z]{1,45}$` | `straight` |
| `azurerm_template_deployment` | `depl` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._\\(\\)]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_traffic_manager_profile` | `traf` | `global` | `non-regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-.]{0,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_user_assigned_identity` | `msia` | `parent` | `regional` | 3 | 128 | `true` | `true` | `^[0-9a-zA-Z][0-9a-zA-Z-_]{2,127}$` | `dashed, straight` |
| `azurerm_virtual_desktop_application_group` | `daga` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_desktop_host_pool` | `hpoo` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_desktop_workspace` | `wvdw` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_hub` | `vhub` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_hub_connection` | `vhco` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine` | `vima` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_extension` | `vmxa` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_portal_name` | `pona` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,62}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_scale_set` | `vmss` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_scale_set_extension` | `vmss` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_network` | `vnet` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,62}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_network_gateway` | `vgwa` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_network_peering` | `vpee` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_wan` | `vwan` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vm_windows_computer_name_prefix` | `napr` | `resourceGroup` | `regional` | 1 | 9 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,7}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vmware_cluster` | `vwca` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vmware_express_route_authorization` | `vwer` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vmware_private_cloud` | `vwpc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vpn_gateway_connection` | `vcna` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vpn_site` | `vsta` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_web_application_firewall_policy` | `wafw` | `global` | `non-regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9]{0,78}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_web_pubsub` | `wepu` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z][-a-zA-Z0-9]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_web_pubsub_hub` | `pshu` | `parent` | `regional` | 1 | 128 | `false` | `false` | `^[a-zA-Z][a-zA-Z0-9_`,.\\[\\]]{0,127}$` | `pascal, camel, straight` |
| `azurerm_windows_virtual_machine` | `vima` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_windows_virtual_machine_scale_set` | `vmss` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_windows_web_app` | `wwap` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `databricks_cluster` | `dbcd` | `parent` | `regional` | 3 | 30 | `true` | `false` | `^[a-zA-Z0-9-_]{3,30}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `databricks_high_concurrency_cluster` | `dbhc` | `parent` | `regional` | 3 | 30 | `true` | `false` | `^[a-zA-Z0-9-_]{3,30}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `databricks_standard_cluster` | `dbsc` | `parent` | `regional` | 3 | 30 | `true` | `false` | `^[a-zA-Z0-9-_]{3,30}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `general` | `gene` | `global` | `non-regional` | 1 | 250 | `true` | `false` | `^[a-zA-Z0-9-_]{1,250}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `general_safe` | `gene` | `global` | `non-regional` | 1 | 250 | `false` | `true` | `^[a-z]{1,250}$` | `straight` |
