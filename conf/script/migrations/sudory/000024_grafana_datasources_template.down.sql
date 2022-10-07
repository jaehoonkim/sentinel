DELETE FROM `template` WHERE uuid IN ('40000000000000000000000000000001', '40000000000000000000000000000002', '40000000000000000000000000000003');
DELETE FROM `template_command` WHERE uuid IN ('40000000000000000000000000000001', '40000000000000000000000000000002', '40000000000000000000000000000003');

DELETE FROM `template_recipe` WHERE method IN ('grafana.datasources.get', 'grafana.datasources.list', 'grafana.datasources.delete');