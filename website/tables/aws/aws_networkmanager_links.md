# Table: aws_networkmanager_links

This table shows data for Networkmanager Links.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_Link.html

The composite primary key for this table is (**region**, **global_network_id**).

## Relations

This table depends on [aws_networkmanager_global_networks](aws_networkmanager_global_networks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region (PK)|`utf8`|
|arn|`utf8`|
|tags|`json`|
|bandwidth|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|global_network_id (PK)|`utf8`|
|link_arn|`utf8`|
|link_id|`utf8`|
|provider|`utf8`|
|site_id|`utf8`|
|state|`utf8`|
|type|`utf8`|