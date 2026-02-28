# Azure CAF Resources

Comprehensive Azure defaults used by `cloud = "azure"`.
All resource identifiers come from the Azure CAF provider resource catalog (`resourceDefinition.json`).

- Total resource types: 395
- Acronym policy in Sigil Azure mode: CAF acronyms from the Azure CAF catalog
- Constraints source: Azure CAF definitions (which mirror Azure ARM naming restrictions)

References:
- Azure CAF provider catalog: https://github.com/aztfmod/terraform-provider-azurecaf/blob/main/resourceDefinition.json
- Azure ARM naming rules: https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules
- CAF naming guidance and abbreviations: https://learn.microsoft.com/en-us/azure/cloud-adoption-framework/ready/azure-best-practices/resource-abbreviations

| Resource | Acronym | CAF Scope | Sigil Scope | Min | Max | Dashes | Lowercase Only | Regex | Allowed Styles |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `aks_node_pool_linux` | `npl` | `parent` | `regional` | 1 | 12 | `false` | `false` | `^[a-z][0-9a-z]{0,11}$` | `pascal, camel, straight` |
| `aks_node_pool_windows` | `npw` | `parent` | `regional` | 1 | 6 | `false` | `false` | `^[a-z][0-9a-z]{0,5}$` | `pascal, camel, straight` |
| `azurerm_aadb2c_directory` | `aadb2c` | `global` | `non-regional` | 1 | 75 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,73}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_analysis_services_server` | `as` | `resourceGroup` | `regional` | 3 | 63 | `false` | `true` | `^[a-z][a-z0-9]{2,62}$` | `straight` |
| `azurerm_api_management` | `apim` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,48}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_api` | `apimapi` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_api_operation_tag` | `apimapiopt` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_backend` | `apimbe` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_certificate` | `apimcer` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_gateway` | `apimgw` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_group` | `apimgr` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_logger` | `apimlg` | `global` | `non-regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,78}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_api_management_service` | `apim` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-z][a-zA-Z0-9-]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_configuration` | `appcg` | `resourceGroup` | `regional` | 5 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{5,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_service` | `app` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_service_environment` | `ase` | `resourceGroup` | `regional` | 2 | 36 | `true` | `false` | `^[0-9A-Za-z-]{2,36}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_app_service_plan` | `plan` | `resourceGroup` | `regional` | 1 | 40 | `true` | `false` | `^[0-9A-Za-z-]{1,40}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_gateway` | `agw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_insights` | `appi` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^%&\\?/. ][^%&\\?/]{0,258}[^%&\\?/. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_insights_web_test` | `appiwt` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9- ]{0,62}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_application_security_group` | `asg` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_account` | `aa` | `resourceGroup` | `regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_certificate` | `aacert` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_credential` | `aacred` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_hybrid_runbook_worker_group` | `aahwg` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^([^<>*%&:\\?.+/#\\s]?[ ]?){0,127}[^<>*%&:\\?.+/#\\s]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_job_schedule` | `aajs` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_runbook` | `aarun` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,62}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_schedule` | `aasched` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_automation_variable` | `aavar` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_availability_set` | `avail` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bastion_host` | `bast` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_batch_account` | `ba` | `region` | `regional` | 3 | 24 | `false` | `true` | `^[a-z0-9]{3,24}$` | `straight` |
| `azurerm_batch_application` | `baapp` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9_-]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_batch_certificate` | `bacert` | `parent` | `regional` | 5 | 45 | `true` | `false` | `^[a-zA-Z0-9_-]{5,45}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_batch_pool` | `bapool` | `parent` | `regional` | 3 | 24 | `true` | `false` | `^[a-zA-Z0-9_-]{1,24}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_Email` | `botmail` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_directline` | `botline` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_ms_teams` | `botteams` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channel_slack` | `botslack` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_channels_registration` | `botchan` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_connection` | `botcon` | `parent` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_service_azure_bot` | `botaz` | `global` | `non-regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_bot_web_app` | `bot` | `global` | `non-regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_endpoint` | `cdn` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_custom_domain` | `cfdcd` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,258}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_endpoint` | `cfde` | `global` | `non-regional` | 1 | 46 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,44}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_firewall_policy` | `cfdfp` | `resourceGroup` | `regional` | 1 | 128 | `false` | `false` | `^[a-zA-Z][0-9a-zA-Z]{0,127}$` | `pascal, camel, straight` |
| `azurerm_cdn_frontdoor_origin` | `cfdo` | `parent` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_origin_group` | `cfdog` | `parent` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_profile` | `cfdp` | `resourceGroup` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_route` | `cfdroute` | `parent` | `regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,88}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_rule` | `cfdr` | `parent` | `regional` | 1 | 60 | `false` | `false` | `^[a-zA-Z][a-zA-Z0-9]{0,59}$` | `pascal, camel, straight` |
| `azurerm_cdn_frontdoor_rule_set` | `cfdrs` | `parent` | `regional` | 1 | 60 | `false` | `false` | `^[a-zA-Z][a-zA-Z0-9]{0,59}$` | `pascal, camel, straight` |
| `azurerm_cdn_frontdoor_secret` | `cfds` | `parent` | `regional` | 2 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_frontdoor_security_policy` | `cfdsp` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[0-9A-Za-z_-]{1,260}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cdn_profile` | `cdnprof` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cognitive_account` | `cog` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_cognitive_deployment` | `cog` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_communication_service` | `acs` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9_-]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_consumption_budget_resource_group` | `acbrg` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_consumption_budget_subscription` | `acbs` | `subscription` | `non-regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_containerGroups` | `cg` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_container_app` | `ca` | `resourceGroup` | `regional` | 1 | 32 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,30}[a-z0-9]$` | `dashed, straight` |
| `azurerm_container_app_environment` | `cae` | `resourceGroup` | `regional` | 1 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_container_registry` | `cr` | `resourceGroup` | `regional` | 1 | 63 | `false` | `true` | `^[a-zA-Z0-9]{1,63}$` | `straight` |
| `azurerm_container_registry_webhook` | `crwh` | `resourceGroup` | `regional` | 1 | 50 | `false` | `false` | `^[a-zA-Z0-9]{1,50}$` | `pascal, camel, straight` |
| `azurerm_cosmosdb_account` | `cosmos` | `resourceGroup` | `regional` | 1 | 44 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,42}[a-z0-9]$` | `dashed, straight` |
| `azurerm_custom_provider` | `prov` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[^&%?\\/]{2,63}[^&%.?\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dashboard` | `dsb` | `parent` | `regional` | 3 | 160 | `true` | `false` | `^[a-zA-Z0-9-]{3,160}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory` | `adf` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_azure_blob` | `adfblob` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_cosmosdb_sqlapi` | `adfsqlapi` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_delimited_text` | `adfdtext` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_http` | `adfhttp` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_json` | `adfjson` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_mysql` | `adfmysql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_postgresql` | `adfpsql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_dataset_sql_server_table` | `adfmssql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_integration_runtime_managed` | `adfir` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_blob_storage` | `adflsabs` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_databricks` | `adflsadb` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_function` | `adflsaf` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_azure_sql_database` | `adflsasdb` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_cosmosdb` | `adflsacdb` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_data_lake_storage_gen2` | `adfsvst` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_key_vault` | `adfsvkv` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_mysql` | `adfsvmysql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_postgresql` | `adfsvpsql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_sftp` | `adflsaftp` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_sql_server` | `adfsvmssql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_linked_service_web` | `adfsvweb` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_pipeline` | `adfpl` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_factory_trigger_schedule` | `adftg` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][^<>*%:.?\\+\\/]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_lake_analytics_account` | `dla` | `global` | `non-regional` | 3 | 24 | `false` | `false` | `^[a-z0-9]{3,24}$` | `pascal, camel, straight` |
| `azurerm_data_lake_analytics_firewall_rule` | `dlfw` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-z0-9-_]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_lake_store` | `dls` | `parent` | `regional` | 3 | 24 | `false` | `false` | `^[a-z0-9]{3,24}$` | `pascal, camel, straight` |
| `azurerm_data_lake_store_firewall_rule` | `dlsfw` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-_]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_blob_storage` | `dpbpb` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_disk` | `dpbpd` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_postgresql` | `dpbpp` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_policy_postgresql_flexible_server` | `dpbppf` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_data_protection_backup_vault` | `dpbv` | `resourceGroup` | `regional` | 2 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{1,49}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_database_migration_project` | `migr` | `parent` | `regional` | 2 | 57 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,56}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_database_migration_service` | `dms` | `resourceGroup` | `regional` | 2 | 62 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{1,61}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_databricks_workspace` | `dbw` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[a-zA-Z0-9-_]{3,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dedicated_host` | `dh` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dedicated_host_group` | `dhg` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center` | `dc` | `resourceGroup` | `regional` | 3 | 26 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,24}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_catalog` | `dcc` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_dev_box_definition` | `dcdb` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_environment_type` | `dcet` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_gallery` | `dcg` | `parent` | `regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9_.]{0,78}[a-zA-Z0-9])?$` | `pascal, camel, straight` |
| `azurerm_dev_center_network_connection` | `dcnc` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_project` | `dcp` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_center_project_environment_type` | `dcpet` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_test_lab` | `lab` | `resourceGroup` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9-_]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_test_linux_virtual_machine` | `labvm` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dev_test_windows_virtual_machine` | `labvm` | `parent` | `regional` | 1 | 15 | `true` | `false` | `^[a-zA-Z0-9-]{1,15}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_endpoint_eventgrid` | `adteg` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9_-]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_endpoint_eventhub` | `adteh` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9_-]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_endpoint_servicebus` | `adtsb` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9_-]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_digital_twins_instance` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_disk_encryption_set` | `des` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9-_]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_a_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_aaaa_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_caa_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_cname_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_mx_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_ns_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_ptr_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_txt_record` | `dnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_dns_zone` | `dns` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_domain` | `egd` | `resourceGroup` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_domain_topic` | `egdt` | `parent` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_event_subscription` | `egs` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[a-zA-Z0-9-]{3,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventgrid_topic` | `egt` | `resourceGroup` | `regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9-]{3,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub` | `evh` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_authorization_rule` | `ehar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_consumer_group` | `ehcg` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_namespace` | `ehn` | `global` | `non-regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_namespace_authorization_rule` | `ehnar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_eventhub_namespace_disaster_recovery_config` | `ehdr` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_express_route_circuit` | `erc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_express_route_gateway` | `ergw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_federated_identity_credential` | `fedcred` | `parent` | `regional` | 3 | 120 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_-]{2,119}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall` | `fw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_application_rule_collection` | `fwapp` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_ip_configuration` | `fwipconf` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_nat_rule_collection` | `fwnatrc` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_network_rule_collection` | `fwnetrc` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_firewall_policy` | `afwp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_frontdoor` | `fd` | `global` | `non-regional` | 5 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{3,62}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_frontdoor_firewall_policy` | `fdfw` | `global` | `non-regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9]{0,78}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_function_app` | `fa` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_function_app_slot` | `fas` | `global` | `non-regional` | 2 | 59 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,57}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_hadoop_cluster` | `hadoop` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_hbase_cluster` | `hbase` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_interactive_query_cluster` | `iqr` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_kafka_cluster` | `kafka` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_ml_services_cluster` | `mls` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_rserver_cluster` | `rser` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_spark_cluster` | `spark` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_hdinsight_storm_cluster` | `storm` | `global` | `non-regional` | 3 | 59 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,57}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_healthcare_dicom_service` | `dicom` | `parent` | `regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_fhir_service` | `fhir` | `parent` | `regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_medtech_service` | `medtech` | `parent` | `regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_service` | `hcasvc` | `global` | `non-regional` | 3 | 24 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,22}[a-z0-9]$` | `dashed, straight` |
| `azurerm_healthcare_workspace` | `hcw` | `global` | `non-regional` | 3 | 24 | `false` | `true` | `^[a-z0-9]{3,24}$` | `straight` |
| `azurerm_image` | `img` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_integration_service_environment` | `lappise` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\-\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iot_security_device_group` | `iotdg` | `parent` | `regional` | 1 | 32 | `true` | `false` | `^[a-zA-Z0-9-._]{1,32}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iot_security_solution` | `iotss` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9-_]{1,260}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iotcentral_application` | `iotapp` | `global` | `non-regional` | 2 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_iothub` | `iot` | `global` | `non-regional` | 3 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,48}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_certificate` | `iotcert` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_consumer_group` | `iotcg` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9-._]{1,50}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_dps` | `dps` | `resourceGroup` | `regional` | 3 | 64 | `true` | `false` | `^[a-zA-Z0-9-]{1,63}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_dps_certificate` | `dpscert` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_dps_shared_access_policy` | `dpssap` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_iothub_shared_access_policy` | `iotsap` | `parent` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_ip_group` | `ipgr` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault` | `kv` | `global` | `non-regional` | 3 | 24 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{1,22}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault_certificate` | `kvc` | `parent` | `regional` | 1 | 127 | `true` | `false` | `^[a-zA-Z0-9-]{1,127}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault_key` | `kvk` | `parent` | `regional` | 1 | 127 | `true` | `false` | `^[a-zA-Z0-9-]{1,127}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_key_vault_secret` | `kvs` | `parent` | `regional` | 1 | 127 | `true` | `false` | `^[a-zA-Z0-9-]{1,127}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_kubernetes_cluster` | `aks` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_kubernetes_fleet_manager` | `fleet` | `resourceGroup` | `regional` | 1 | 63 | `true` | `true` | `^[0-9a-z]([0-9a-z-]{0,61}[0-9a-z])?$` | `dashed, straight` |
| `azurerm_kusto_cluster` | `kc` | `global` | `non-regional` | 4 | 22 | `false` | `false` | `^[a-z][a-z0-9]{3,21}$` | `pascal, camel, straight` |
| `azurerm_kusto_database` | `kdb` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9- .]{1,260}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_kusto_eventhub_data_connection` | `kehc` | `parent` | `regional` | 1 | 40 | `true` | `false` | `^[a-zA-Z0-9- .]{1,40}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb` | `lb` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_backend_address_pool` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_backend_pool` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_nat_pool` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_nat_rule` | `lbnatrl` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_outbound_rule` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_probe` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_lb_rule` | `adt` | `subscription` | `non-regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9_-]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_linux_virtual_machine` | `vm` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,62}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_linux_virtual_machine_scale_set` | `vmss` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,62}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_linux_web_app` | `lwapp` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_load_test` | `load` | `global` | `non-regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{0,62}[a-zA-Z0-9|]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_local_network_gateway` | `lgw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_cluster` | `logc` | `resourceGroup` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_query_pack` | `laqp` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_solution` | `las` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_storage_insights` | `lasi` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_log_analytics_workspace` | `log` | `parent` | `regional` | 4 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{2,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_action_custom` | `lappac` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_action_http` | `lappah` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_integration_account` | `lappia` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_trigger_custom` | `lapptc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_trigger_http_request` | `lappth` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_trigger_recurrence` | `lapptc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_logic_app_workflow` | `lapp` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[0-9A-Za-z\\(\\-\\)\\_\\.]{1,80}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_machine_learning_compute_instance` | `amlci` | `parent` | `regional` | 1 | 16 | `true` | `false` | `^[a-zA-Z0-9][a-z0-9-]{0,14}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_machine_learning_workspace` | `mlw` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,259}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_maintenance_configuration` | `mcf` | `resourceGroup` | `regional` | 1 | 60 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,58}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_managed_disk` | `dsk` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_maps_account` | `map` | `resourceGroup` | `regional` | 1 | 98 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,97}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_database` | `mariadb` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_firewall_rule` | `mariafw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_server` | `maria` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mariadb_virtual_network_rule` | `mariavn` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_action_group` | `amag` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^|:<>+#%&\\?/]{0,259}[^|:<>+#%&\\?/. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_activity_log_alert` | `adfmysql` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:&?#\\+\\/]{0,259}[^<>*%:&.?#\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_autoscale_setting` | `amas` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,62}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_data_collection_endpoint` | `dce` | `resourceGroup` | `regional` | 3 | 44 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,42}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_data_collection_rule` | `dcr` | `resourceGroup` | `regional` | 3 | 44 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,42}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_diagnostic_setting` | `amds` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_metric_alert` | `ma` | `resourceGroup` | `regional` | 1 | 251 | `true` | `false` | `^[^<>*%&:\\?+/#@{}]{0,250}[^<>*%&:\\?+/#@{}. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_private_link_scope` | `ampls` | `resourceGroup` | `regional` | 1 | 255 | `true` | `false` | `^[0-9A-Za-z-._()]{0,254}[0-9A-Za-z-_()]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_monitor_scheduled_query_rules_alert` | `schqra` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%&:\\?/#{}]{0,259}[^<>*%&:\\?/#{}. ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mssql_database` | `sqldb` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{1,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mssql_elasticpool` | `sqlep` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{1,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mssql_mi` | `sqlmi` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_mssql_server` | `sql` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_mysql_database` | `mysqldb` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_firewall_rule` | `mysqlfw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_flexible_server` | `mysqlf` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_flexible_server_database` | `mysqlfdb` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_flexible_server_firewall_rule` | `mysqlffw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_server` | `mysql` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_mysql_virtual_network_rule` | `mysqlvn` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_account` | `ana` | `resourceGroup` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,126}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_pool` | `anp` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_snapshot` | `ans` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_netapp_volume` | `anv` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_ddos_protection_plan` | `ddospp` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_interface` | `nic` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_security_group` | `nsg` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_security_group_rule` | `nsgr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_security_rule` | `nsgr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_network_watcher` | `nw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_nginx_deployment` | `nginx` | `resourceGroup` | `regional` | 1 | 30 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,28}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_notification_hub` | `nh` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,259}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_notification_hub_authorization_rule` | `dnsrec` | `parent` | `regional` | 1 | 256 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,255}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_notification_hub_namespace` | `dnsrec` | `global` | `non-regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_point_to_site_vpn_gateway` | `vpngw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_portal_dashboard` | `dsb` | `parent` | `regional` | 3 | 160 | `true` | `false` | `^[a-zA-Z0-9-]{3,160}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_database` | `psqldb` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_firewall_rule` | `psqlfw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_flexible_server` | `psqlf` | `global` | `non-regional` | 3 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_postgresql_flexible_server_database` | `psqlfdb` | `parent` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{1,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_flexible_server_firewall_rule` | `psqlffw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_server` | `psql` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-zA-Z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_postgresql_virtual_network_rule` | `psqlvn` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9-_]{1,128}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_powerbi_embedded` | `pbi` | `region` | `regional` | 3 | 63 | `false` | `false` | `^[a-z0-9][a-z0-9]{2,62}$` | `pascal, camel, straight` |
| `azurerm_private_dns_a_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_aaaa_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_cname_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_mx_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_ptr_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver` | `dnspr` | `resourceGroup` | `regional` | 3 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{1,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_dns_forwarding_ruleset` | `dnsfwrs` | `resourceGroup` | `regional` | 2 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{0,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_forwarding_rule` | `dnsfwr` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-_]{0,78}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_inbound_endpoint` | `dnsprie` | `parent` | `regional` | 3 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{1,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_outbound_endpoint` | `dnsproe` | `parent` | `regional` | 3 | 80 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-_]{1,78}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_resolver_virtual_network_link` | `dnsfwrsvnetl` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9]([a-zA-Z0-9-_]{0,78}[a-zA-Z0-9])?$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_srv_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_txt_record` | `pdnsrec` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_zone` | `pdns` | `resourceGroup` | `regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_zone_group` | `pdnszg` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_dns_zone_virtual_network_link` | `pnetlk` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_endpoint` | `pe` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,62}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_link_service` | `pls` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_private_service_connection` | `psc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_proximity_placement_group` | `ppg` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_public_ip` | `pip` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_public_ip_prefix` | `pippf` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_purview_account` | `purv` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9_-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_recovery_services_vault` | `rsv` | `resourceGroup` | `regional` | 2 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{1,49}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_recovery_services_vault_backup_police` | `rsvbp` | `resourceGroup` | `regional` | 3 | 150 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9\\-]{1,148}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_redhat_openshift_cluster` | `aroc` | `resourceGroup` | `regional` | 1 | 30 | `false` | `false` | `^[a-zA-Z0-9]{1,30}$` | `pascal, camel, straight` |
| `azurerm_redhat_openshift_domain` | `arod` | `resourceGroup` | `regional` | 1 | 30 | `false` | `false` | `^[a-zA-Z0-9]{1,30}$` | `pascal, camel, straight` |
| `azurerm_redis_cache` | `redis` | `global` | `non-regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_redis_firewall_rule` | `redisfw` | `parent` | `regional` | 1 | 256 | `false` | `false` | `^[a-zA-Z0-9]{1,256}$` | `pascal, camel, straight` |
| `azurerm_relay_hybrid_connection` | `rlhc` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_relay_namespace` | `rln` | `global` | `non-regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_resource_group` | `rg` | `subscription` | `non-regional` | 1 | 90 | `true` | `false` | `^[a-zA-Z0-9-._\\(\\)]{0,89}[a-zA-Z0-9-_\\(\\)]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_resource_group_policy_assignment` | `argpa` | `resourceGroup` | `regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,126}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_role_assignment` | `ra` | `assignment` | `non-regional` | 1 | 64 | `true` | `false` | `^[^%]{0,63}[^ %.]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_role_definition` | `rd` | `definition` | `non-regional` | 1 | 64 | `true` | `false` | `^[^%]{0,63}[^ %.]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_route` | `rt` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_route_server` | `rts` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_route_table` | `route` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_search_service` | `srch` | `global` | `non-regional` | 2 | 60 | `true` | `true` | `^[a-z0-9](?:[a-z0-9-]{0,58}[a-z0-9])?$` | `dashed, straight` |
| `azurerm_service_fabric_cluster` | `sf` | `region` | `regional` | 4 | 23 | `true` | `true` | `^[a-z][a-z0-9-]{2,21}[a-z0-9]$` | `dashed, straight` |
| `azurerm_servicebus_namespace` | `sb` | `global` | `non-regional` | 6 | 50 | `true` | `false` | `^[a-zA-Z][a-zA-Z0-9-]{4,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_namespace_authorization_rule` | `sbar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_namespace_disaster_recovery_config` | `sbdr` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_queue` | `sbq` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_queue_authorization_rule` | `sbqar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_subscription` | `sbs` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_subscription_rule` | `sbsr` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_topic` | `sbt` | `parent` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,258}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_servicebus_topic_authorization_rule` | `sbtar` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_shared_image` | `si` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_shared_image_gallery` | `sig` | `resourceGroup` | `regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9._]{0,78}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_signalr_service` | `sgnlr` | `global` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_snapshots` | `snap` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_sql_elasticpool` | `sqlep` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{1,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_sql_failover_group` | `sqlfg` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_sql_firewall_rule` | `sqlfw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:?\\+\\/]{1,127}[^<>*%:.?\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_sql_server` | `sql` | `global` | `non-regional` | 1 | 63 | `true` | `true` | `^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$` | `dashed, straight` |
| `azurerm_static_site` | `stapp` | `resourceGroup` | `regional` | 1 | 40 | `true` | `false` | `^[a-zA-Z0-9-]{1,40}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_account` | `st` | `global` | `non-regional` | 3 | 24 | `false` | `true` | `^[a-z0-9]{3,24}$` | `straight` |
| `azurerm_storage_blob` | `blob` | `parent` | `regional` | 1 | 1024 | `true` | `false` | `^[^\\s\\/$#&]{1,1000}[^\\s\\/$#&]{0,24}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_container` | `stct` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{2,62}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_data_lake_gen2_filesystem` | `stdl` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_queue` | `stq` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_share` | `sts` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_share_directory` | `sts` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_sync` | `stsy` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,259}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_sync_group` | `stsg` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,259}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_storage_table` | `stt` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-z0-9][a-z0-9-]{1,61}[a-z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_function_javascript_udf` | `asafunc` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_job` | `asa` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_blob` | `asaoblob` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_eventhub` | `asaoeh` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_mssql` | `asaomssql` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_servicebus_queue` | `asaosbq` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_output_servicebus_topic` | `asaosbt` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_reference_input_blob` | `asarblob` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_stream_input_blob` | `asaiblob` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_stream_input_eventhub` | `asaieh` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_stream_analytics_stream_input_iothub` | `asaiiot` | `parent` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9-_]{3,63}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_subnet` | `snet` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_subscription_policy_assignment` | `aspa` | `subscription` | `non-regional` | 1 | 128 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,126}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_firewall_rule` | `syfw` | `parent` | `regional` | 1 | 128 | `true` | `false` | `^[^<>*%:?\\+\\/]{1,127}[^<>*%:.?\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_integration_runtime_azure` | `synira` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_integration_runtime_self_hosted` | `synirsh` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9_-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_linked_service` | `synls` | `subscription` | `non-regional` | 1 | 140 | `false` | `false` | `^[a-zA-Z0-9_]{1,140}$` | `pascal, camel, straight` |
| `azurerm_synapse_managed_private_endpoint` | `synmpe` | `subscription` | `non-regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_.-]{1,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_private_link_hub` | `synplh` | `subscription` | `non-regional` | 1 | 45 | `false` | `true` | `^[a-z0-9]{1,45}$` | `straight` |
| `azurerm_synapse_spark_pool` | `sysp` | `parent` | `regional` | 1 | 15 | `false` | `true` | `^[0-9a-zA-Z]{1,15}$` | `straight` |
| `azurerm_synapse_sql_pool` | `synsp` | `subscription` | `non-regional` | 3 | 15 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9_]{1,13}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_synapse_sql_pool_vulnerability_assessment_baseline` | `synspvab` | `subscription` | `non-regional` | 3 | 63 | `false` | `false` | `^[a-zA-Z0-9]{1,63}$` | `pascal, camel, straight` |
| `azurerm_synapse_sql_pool_workload_classifier` | `synspwc` | `subscription` | `non-regional` | 1 | 128 | `true` | `false` | `^[^<>*%:?\\+\\/]{1,127}[^<>*%:.?\\+\\/]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_sql_pool_workload_group` | `synspwg` | `subscription` | `non-regional` | 1 | 128 | `true` | `false` | `^[^<>*%:.?\\+\\/]{0,127}[^<>*%:.?\\+\\/ ]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_synapse_workspace` | `syws` | `resourceGroup` | `regional` | 1 | 45 | `false` | `true` | `^[0-9a-z]{1,45}$` | `straight` |
| `azurerm_template_deployment` | `deploy` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[a-zA-Z0-9-._\\(\\)]{1,64}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_traffic_manager_profile` | `traf` | `global` | `non-regional` | 1 | 63 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-.]{0,61}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_user_assigned_identity` | `msi` | `parent` | `regional` | 3 | 128 | `true` | `true` | `^[0-9a-zA-Z][0-9a-zA-Z-_]{2,127}$` | `dashed, straight` |
| `azurerm_virtual_desktop_application_group` | `dag` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_desktop_host_pool` | `hpool` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_desktop_workspace` | `wvdws` | `resourceGroup` | `regional` | 1 | 260 | `true` | `false` | `^[a-zA-Z0-9 ][a-zA-Z0-9-._ ]{0,258}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_hub` | `vhub` | `parent` | `regional` | 1 | 50 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,48}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_hub_connection` | `vhcon` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine` | `vm` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_extension` | `vmx` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_portal_name` | `vm` | `resourceGroup` | `regional` | 1 | 64 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,62}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_scale_set` | `vmss` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_machine_scale_set_extension` | `vmssx` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9\\-\\._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_network` | `vnet` | `resourceGroup` | `regional` | 2 | 64 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,62}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_network_gateway` | `vgw` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_network_peering` | `vpeer` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_virtual_wan` | `vwan` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vm_windows_computer_name_prefix` | `cn` | `resourceGroup` | `regional` | 1 | 9 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,7}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vmware_cluster` | `vwc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vmware_express_route_authorization` | `vwera` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vmware_private_cloud` | `vwpc` | `resourceGroup` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vpn_gateway_connection` | `vcn` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_vpn_site` | `vst` | `parent` | `regional` | 1 | 80 | `true` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9-._]{0,78}[a-zA-Z0-9_]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_web_application_firewall_policy` | `wafw` | `global` | `non-regional` | 1 | 80 | `false` | `false` | `^[a-zA-Z0-9][a-zA-Z0-9]{0,78}[a-zA-Z0-9]$` | `pascal, camel, straight` |
| `azurerm_web_pubsub` | `ps` | `resourceGroup` | `regional` | 3 | 63 | `true` | `false` | `^[a-zA-Z][-a-zA-Z0-9]{1,61}[a-zA-Z0-9]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_web_pubsub_hub` | `pshub` | `parent` | `regional` | 1 | 128 | `false` | `false` | `^[a-zA-Z][a-zA-Z0-9_`,.\\[\\]]{0,127}$` | `pascal, camel, straight` |
| `azurerm_windows_virtual_machine` | `vm` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_windows_virtual_machine_scale_set` | `vmss` | `resourceGroup` | `regional` | 1 | 15 | `true` | `false` | `^[^\\/\"\\[\\]:|<>+=;,?*@&_][^\\/\"\\[\\]:|<>+=;,?*@&]{0,13}[^\\/\"\\[\\]:|<>+=;,?*@&.-]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `azurerm_windows_web_app` | `wwapp` | `global` | `non-regional` | 2 | 60 | `true` | `false` | `^[0-9A-Za-z][0-9A-Za-z-]{0,58}[0-9a-zA-Z]$` | `dashed, pascaldashed, pascal, camel, straight` |
| `databricks_cluster` | `dbc` | `parent` | `regional` | 3 | 30 | `true` | `false` | `^[a-zA-Z0-9-_]{3,30}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `databricks_high_concurrency_cluster` | `dbhcc` | `parent` | `regional` | 3 | 30 | `true` | `false` | `^[a-zA-Z0-9-_]{3,30}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `databricks_standard_cluster` | `dbsc` | `parent` | `regional` | 3 | 30 | `true` | `false` | `^[a-zA-Z0-9-_]{3,30}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `general` | `` | `global` | `non-regional` | 1 | 250 | `true` | `false` | `^[a-zA-Z0-9-_]{1,250}$` | `dashed, pascaldashed, pascal, camel, straight` |
| `general_safe` | `` | `global` | `non-regional` | 1 | 250 | `false` | `true` | `^[a-z]{1,250}$` | `straight` |
