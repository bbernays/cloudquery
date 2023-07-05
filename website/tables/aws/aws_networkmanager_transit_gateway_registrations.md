# Table: aws_networkmanager_transit_gateway_registrations

This table shows data for Networkmanager Transit Gateway Registrations.

https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_TransitGatewayRegistration.html

The composite primary key for this table is (**region**, **global_network_id**, **transit_gateway_arn**).

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
|global_network_id (PK)|`utf8`|
|state|`json`|
|transit_gateway_arn (PK)|`utf8`|