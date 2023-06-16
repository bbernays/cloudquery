INSERT INTO gcp_policy_results (resource_id, execution_time, framework, check_id, title, project_id, status)
SELECT gsi.name                                                                                              AS resource_id,
       :'execution_time'::timestamp                                                                          AS execution_time,
       :'framework'                                                                                          AS framework,
       :'check_id'                                                                                           AS check_id,
       'Ensure "user options" database flag for Cloud SQL SQL Server instance is not configured (Automated)' AS title,
       gsi.project_id                                                                                        AS project_id,
       CASE
           WHEN
                           gsi.database_version LIKE 'SQLSERVER%'
                       AND f->>'value' IS NOT NULL
               THEN 'fail'
           ELSE 'pass'
           END                                                                                               AS status
FROM gcp_sql_instances gsi LEFT JOIN JSONB_ARRAY_ELEMENTS(gsi.settings->'databaseFlags') AS f ON f->>'name'='user options';
