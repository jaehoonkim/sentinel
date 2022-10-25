UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE uuid='50000000000000000000000000000001';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE uuid='50000000000000000000000000000002';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE uuid='50000000000000000000000000000003';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE uuid='50000000000000000000000000000004';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE uuid='50000000000000000000000000000005';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE uuid='50000000000000000000000000000006';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE uuid='50000000000000000000000000000007';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE uuid='50000000000000000000000000000008';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE uuid='50000000000000000000000000000009';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE uuid='50000000000000000000000000000010';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE uuid='50000000000000000000000000000011';
UPDATE `template_command` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE uuid='50000000000000000000000000000012';

UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE method='openstack.identity.projects.get';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE method='openstack.identity.projects.list';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE method='openstack.compute.servers.get';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE method='openstack.compute.servers.list';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE method='openstack.networking.networks.get';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE method='openstack.networking.networks.list';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE method='openstack.networking.routers.get';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE method='openstack.networking.routers.list';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE method='openstack.networking.subnets.get';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE method='openstack.networking.subnets.list';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"id":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key","id"]}' WHERE method='openstack.compute.hypervisors.get';
UPDATE `template_recipe` SET `args`='{"type":"object","properties":{"url":{"type":"string","pattern":"^."},"credential_key":{"type":"string","pattern":"^."},"query":{"type":"object","additionalProperties":{"type":"string"}}},"required":["url","credential_key"]}' WHERE method='openstack.compute.hypervisors.list';
