# Table: aws_networkmanager_sites

This table shows data for Networkmanager Sites.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_Site.html

The composite primary key for this table is (**region**, **global_network_id**, **site_arn**).

## Relations

This table depends on [aws_networkmanager_global_networks](aws_networkmanager_global_networks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|global_network_id (PK)|`utf8`|
|location|`json`|
|site_arn (PK)|`utf8`|
|site_id|`utf8`|
|state|`utf8`|