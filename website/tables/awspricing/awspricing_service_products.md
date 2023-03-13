# Table: awspricing_service_products

The composite primary key for this table is (**sku**, **region_code**).

## Relations

This table depends on [awspricing_services](awspricing_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|sku (PK)|String|
|product_family|String|
|attributes|JSON|
|region_code (PK)|String|