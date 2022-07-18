// Code generated by Ice-cream-maker DO NOT EDIT.
package v2
 
func (ManagedChannel) ColumnNames() []string {
	return []string{
 		"name",
 		"summary",
 		"event_category",
 		"created",
 		"updated",
 		"deleted",
 		"uuid",
	}
}
 
func (ManagedChannel_property) ColumnNames() []string {
	return []string{
 		"name",
 		"summary",
 		"event_category",
	}
}
 
func (ManagedChannel_option) ColumnNames() []string {
	return []string{
 		"`name`",
 		"`summary`",
 		"IFNULL(`event_category`, 0) AS `event_category`",
 		"`created`",
 		"`updated`",
 		"`deleted`",
 		"`uuid`",
 		"IFNULL(`status_option.status_max_count`, 0) AS `status_option.status_max_count`",
 		"`status_option.created`",
 		"`status_option.updated`",
 		"IFNULL(`format.format_type`, 0) AS `format.format_type`",
 		"IFNULL(`format.format_data`, '') AS `format.format_data`",
 		"`format.created`",
 		"`format.updated`",
 		"IFNULL(`edge.notifier_type`, 0) AS `edge.notifier_type`",
 		"`edge.created`",
 		"`edge.updated`",
	}
}
 
func (ManagedChannel_tangled) ColumnNames() []string {
	return []string{
 		"`name`",
 		"`summary`",
 		"IFNULL(`event_category`, 0) AS `event_category`",
 		"`created`",
 		"`updated`",
 		"`deleted`",
 		"`uuid`",
 		"IFNULL(`status_option.status_max_count`, 0) AS `status_option.status_max_count`",
 		"`status_option.created`",
 		"`status_option.updated`",
 		"IFNULL(`format.format_type`, 0) AS `format.format_type`",
 		"IFNULL(`format.format_data`, '') AS `format.format_data`",
 		"`format.created`",
 		"`format.updated`",
 		"IFNULL(`notifier.edge.notifier_type`, 0) AS `notifier.edge.notifier_type`",
 		"`notifier.edge.created`",
 		"`notifier.edge.updated`",
 		"`notifier.console.created`",
 		"`notifier.console.updated`",
 		"IFNULL(`notifier.webhook.method`, '') AS `notifier.webhook.method`",
 		"IFNULL(`notifier.webhook.url`, '') AS `notifier.webhook.url`",
 		"`notifier.webhook.request_headers`",
 		"IFNULL(`notifier.webhook.request_timeout`, 0) AS `notifier.webhook.request_timeout`",
 		"`notifier.webhook.created`",
 		"`notifier.webhook.updated`",
 		"IFNULL(`notifier.rabbitmq.url`, 0) AS `notifier.rabbitmq.url`",
 		"`notifier.rabbitmq.exchange`",
 		"`notifier.rabbitmq.routing_key`",
 		"`notifier.rabbitmq.mandatory`",
 		"`notifier.rabbitmq.immediate`",
 		"`notifier.rabbitmq.message_headers`",
 		"`notifier.rabbitmq.message_content_type`",
 		"`notifier.rabbitmq.message_content_encoding`",
 		"`notifier.rabbitmq.message_delivery_mode`",
 		"`notifier.rabbitmq.message_priority`",
 		"`notifier.rabbitmq.message_correlation_id`",
 		"`notifier.rabbitmq.message_reply_to`",
 		"`notifier.rabbitmq.message_expiration`",
 		"`notifier.rabbitmq.message_message_id`",
 		"`notifier.rabbitmq.message_timestamp`",
 		"`notifier.rabbitmq.message_type`",
 		"`notifier.rabbitmq.message_user_id`",
 		"`notifier.rabbitmq.message_app_id`",
 		"`notifier.rabbitmq.created`",
 		"`notifier.rabbitmq.updated`",
	}
}
 
func (NotifierEdge) ColumnNames() []string {
	return []string{
 		"notifier_type",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (NotifierEdge_property) ColumnNames() []string {
	return []string{
 		"notifier_type",
 		"created",
 		"updated",
	}
}
 
func (NotifierEdge_option) ColumnNames() []string {
	return []string{
 		"IFNULL(`notifier_type`, 0) AS `notifier_type`",
 		"`created`",
 		"`updated`",
 		"`uuid`",
 		"`console.created`",
 		"`console.updated`",
 		"IFNULL(`webhook.method`, '') AS `webhook.method`",
 		"IFNULL(`webhook.url`, '') AS `webhook.url`",
 		"`webhook.request_headers`",
 		"IFNULL(`webhook.request_timeout`, 0) AS `webhook.request_timeout`",
 		"`webhook.created`",
 		"`webhook.updated`",
 		"IFNULL(`rabbitmq.url`, 0) AS `rabbitmq.url`",
 		"`rabbitmq.exchange`",
 		"`rabbitmq.routing_key`",
 		"`rabbitmq.mandatory`",
 		"`rabbitmq.immediate`",
 		"`rabbitmq.message_headers`",
 		"`rabbitmq.message_content_type`",
 		"`rabbitmq.message_content_encoding`",
 		"`rabbitmq.message_delivery_mode`",
 		"`rabbitmq.message_priority`",
 		"`rabbitmq.message_correlation_id`",
 		"`rabbitmq.message_reply_to`",
 		"`rabbitmq.message_expiration`",
 		"`rabbitmq.message_message_id`",
 		"`rabbitmq.message_timestamp`",
 		"`rabbitmq.message_type`",
 		"`rabbitmq.message_user_id`",
 		"`rabbitmq.message_app_id`",
 		"`rabbitmq.created`",
 		"`rabbitmq.updated`",
	}
}
 
func (NotifierConsole) ColumnNames() []string {
	return []string{
 		"uuid",
 		"created",
 		"updated",
	}
}
 
func (NotifierRabbitMq) ColumnNames() []string {
	return []string{
 		"url",
 		"exchange",
 		"routing_key",
 		"mandatory",
 		"immediate",
 		"message_headers",
 		"message_content_type",
 		"message_content_encoding",
 		"message_delivery_mode",
 		"message_priority",
 		"message_correlation_id",
 		"message_reply_to",
 		"message_expiration",
 		"message_message_id",
 		"message_timestamp",
 		"message_type",
 		"message_user_id",
 		"message_app_id",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (NotifierRabbitMq_property) ColumnNames() []string {
	return []string{
 		"url",
 		"exchange",
 		"routing_key",
 		"mandatory",
 		"immediate",
 		"message_headers",
 		"message_content_type",
 		"message_content_encoding",
 		"message_delivery_mode",
 		"message_priority",
 		"message_correlation_id",
 		"message_reply_to",
 		"message_expiration",
 		"message_message_id",
 		"message_timestamp",
 		"message_type",
 		"message_user_id",
 		"message_app_id",
 		"created",
 		"updated",
	}
}
 
func (NotifierWebhook) ColumnNames() []string {
	return []string{
 		"method",
 		"url",
 		"request_headers",
 		"request_timeout",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (NotifierWebhook_property) ColumnNames() []string {
	return []string{
 		"method",
 		"url",
 		"request_headers",
 		"request_timeout",
 		"created",
 		"updated",
	}
}
 
func (ChannelStatusOption) ColumnNames() []string {
	return []string{
 		"status_max_count",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (ChannelStatus) ColumnNames() []string {
	return []string{
 		"uuid",
 		"created",
 		"message",
	}
}
 
func (Filter) ColumnNames() []string {
	return []string{
 		"uuid",
 		"filter_op",
 		"filter_key",
 		"filter_value",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (Format) ColumnNames() []string {
	return []string{
 		"format_type",
 		"format_data",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (Format_property) ColumnNames() []string {
	return []string{
 		"format_type",
 		"format_data",
 		"created",
 		"updated",
	}
}
 
func (ManagedChannel) ColumnNamesWithAlias() []string {
	return []string{
 		"name",
 		"summary",
 		"event_category",
 		"created",
 		"updated",
 		"deleted",
 		"uuid",
	}
}
 
func (ManagedChannel_property) ColumnNamesWithAlias() []string {
	return []string{
 		"name",
 		"summary",
 		"event_category",
	}
}
 
func (ManagedChannel_option) ColumnNamesWithAlias() []string {
	return []string{
 		"CN.name AS `name`",
 		"CN.summary AS `summary`",
 		"IFNULL(CN.event_category, 0) AS `event_category`",
 		"CN.created AS `created`",
 		"CN.updated AS `updated`",
 		"CN.deleted AS `deleted`",
 		"CN.uuid AS `uuid`",
 		"IFNULL(SO.status_max_count, 0) AS `status_option.status_max_count`",
 		"SO.created AS `status_option.created`",
 		"SO.updated AS `status_option.updated`",
 		"IFNULL(FM.format_type, 0) AS `format.format_type`",
 		"IFNULL(FM.format_data, '') AS `format.format_data`",
 		"FM.created AS `format.created`",
 		"FM.updated AS `format.updated`",
 		"IFNULL(EG.notifier_type, 0) AS `edge.notifier_type`",
 		"EG.created AS `edge.created`",
 		"EG.updated AS `edge.updated`",
	}
}
 
func (ManagedChannel_tangled) ColumnNamesWithAlias() []string {
	return []string{
 		"CN.name AS `name`",
 		"CN.summary AS `summary`",
 		"IFNULL(CN.event_category, 0) AS `event_category`",
 		"CN.created AS `created`",
 		"CN.updated AS `updated`",
 		"CN.deleted AS `deleted`",
 		"CN.uuid AS `uuid`",
 		"IFNULL(SO.status_max_count, 0) AS `status_option.status_max_count`",
 		"SO.created AS `status_option.created`",
 		"SO.updated AS `status_option.updated`",
 		"IFNULL(FM.format_type, 0) AS `format.format_type`",
 		"IFNULL(FM.format_data, '') AS `format.format_data`",
 		"FM.created AS `format.created`",
 		"FM.updated AS `format.updated`",
 		"IFNULL(EG.notifier_type, 0) AS `notifier.edge.notifier_type`",
 		"EG.created AS `notifier.edge.created`",
 		"EG.updated AS `notifier.edge.updated`",
 		"CS.created AS `notifier.console.created`",
 		"CS.updated AS `notifier.console.updated`",
 		"IFNULL(WH.method, '') AS `notifier.webhook.method`",
 		"IFNULL(WH.url, '') AS `notifier.webhook.url`",
 		"WH.request_headers AS `notifier.webhook.request_headers`",
 		"IFNULL(WH.request_timeout, 0) AS `notifier.webhook.request_timeout`",
 		"WH.created AS `notifier.webhook.created`",
 		"WH.updated AS `notifier.webhook.updated`",
 		"IFNULL(RQ.url, 0) AS `notifier.rabbitmq.url`",
 		"RQ.exchange AS `notifier.rabbitmq.exchange`",
 		"RQ.routing_key AS `notifier.rabbitmq.routing_key`",
 		"RQ.mandatory AS `notifier.rabbitmq.mandatory`",
 		"RQ.immediate AS `notifier.rabbitmq.immediate`",
 		"RQ.message_headers AS `notifier.rabbitmq.message_headers`",
 		"RQ.message_content_type AS `notifier.rabbitmq.message_content_type`",
 		"RQ.message_content_encoding AS `notifier.rabbitmq.message_content_encoding`",
 		"RQ.message_delivery_mode AS `notifier.rabbitmq.message_delivery_mode`",
 		"RQ.message_priority AS `notifier.rabbitmq.message_priority`",
 		"RQ.message_correlation_id AS `notifier.rabbitmq.message_correlation_id`",
 		"RQ.message_reply_to AS `notifier.rabbitmq.message_reply_to`",
 		"RQ.message_expiration AS `notifier.rabbitmq.message_expiration`",
 		"RQ.message_message_id AS `notifier.rabbitmq.message_message_id`",
 		"RQ.message_timestamp AS `notifier.rabbitmq.message_timestamp`",
 		"RQ.message_type AS `notifier.rabbitmq.message_type`",
 		"RQ.message_user_id AS `notifier.rabbitmq.message_user_id`",
 		"RQ.message_app_id AS `notifier.rabbitmq.message_app_id`",
 		"RQ.created AS `notifier.rabbitmq.created`",
 		"RQ.updated AS `notifier.rabbitmq.updated`",
	}
}
 
func (NotifierEdge) ColumnNamesWithAlias() []string {
	return []string{
 		"notifier_type",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (NotifierEdge_property) ColumnNamesWithAlias() []string {
	return []string{
 		"notifier_type",
 		"created",
 		"updated",
	}
}
 
func (NotifierEdge_option) ColumnNamesWithAlias() []string {
	return []string{
 		"IFNULL(EG.notifier_type, 0) AS `notifier_type`",
 		"EG.created AS `created`",
 		"EG.updated AS `updated`",
 		"EG.uuid AS `uuid`",
 		"CS.created AS `console.created`",
 		"CS.updated AS `console.updated`",
 		"IFNULL(WH.method, '') AS `webhook.method`",
 		"IFNULL(WH.url, '') AS `webhook.url`",
 		"WH.request_headers AS `webhook.request_headers`",
 		"IFNULL(WH.request_timeout, 0) AS `webhook.request_timeout`",
 		"WH.created AS `webhook.created`",
 		"WH.updated AS `webhook.updated`",
 		"IFNULL(RQ.url, 0) AS `rabbitmq.url`",
 		"RQ.exchange AS `rabbitmq.exchange`",
 		"RQ.routing_key AS `rabbitmq.routing_key`",
 		"RQ.mandatory AS `rabbitmq.mandatory`",
 		"RQ.immediate AS `rabbitmq.immediate`",
 		"RQ.message_headers AS `rabbitmq.message_headers`",
 		"RQ.message_content_type AS `rabbitmq.message_content_type`",
 		"RQ.message_content_encoding AS `rabbitmq.message_content_encoding`",
 		"RQ.message_delivery_mode AS `rabbitmq.message_delivery_mode`",
 		"RQ.message_priority AS `rabbitmq.message_priority`",
 		"RQ.message_correlation_id AS `rabbitmq.message_correlation_id`",
 		"RQ.message_reply_to AS `rabbitmq.message_reply_to`",
 		"RQ.message_expiration AS `rabbitmq.message_expiration`",
 		"RQ.message_message_id AS `rabbitmq.message_message_id`",
 		"RQ.message_timestamp AS `rabbitmq.message_timestamp`",
 		"RQ.message_type AS `rabbitmq.message_type`",
 		"RQ.message_user_id AS `rabbitmq.message_user_id`",
 		"RQ.message_app_id AS `rabbitmq.message_app_id`",
 		"RQ.created AS `rabbitmq.created`",
 		"RQ.updated AS `rabbitmq.updated`",
	}
}
 
func (NotifierConsole) ColumnNamesWithAlias() []string {
	return []string{
 		"uuid",
 		"created",
 		"updated",
	}
}
 
func (NotifierRabbitMq) ColumnNamesWithAlias() []string {
	return []string{
 		"url",
 		"exchange",
 		"routing_key",
 		"mandatory",
 		"immediate",
 		"message_headers",
 		"message_content_type",
 		"message_content_encoding",
 		"message_delivery_mode",
 		"message_priority",
 		"message_correlation_id",
 		"message_reply_to",
 		"message_expiration",
 		"message_message_id",
 		"message_timestamp",
 		"message_type",
 		"message_user_id",
 		"message_app_id",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (NotifierRabbitMq_property) ColumnNamesWithAlias() []string {
	return []string{
 		"url",
 		"exchange",
 		"routing_key",
 		"mandatory",
 		"immediate",
 		"message_headers",
 		"message_content_type",
 		"message_content_encoding",
 		"message_delivery_mode",
 		"message_priority",
 		"message_correlation_id",
 		"message_reply_to",
 		"message_expiration",
 		"message_message_id",
 		"message_timestamp",
 		"message_type",
 		"message_user_id",
 		"message_app_id",
 		"created",
 		"updated",
	}
}
 
func (NotifierWebhook) ColumnNamesWithAlias() []string {
	return []string{
 		"method",
 		"url",
 		"request_headers",
 		"request_timeout",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (NotifierWebhook_property) ColumnNamesWithAlias() []string {
	return []string{
 		"method",
 		"url",
 		"request_headers",
 		"request_timeout",
 		"created",
 		"updated",
	}
}
 
func (ChannelStatusOption) ColumnNamesWithAlias() []string {
	return []string{
 		"status_max_count",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (ChannelStatus) ColumnNamesWithAlias() []string {
	return []string{
 		"uuid",
 		"created",
 		"message",
	}
}
 
func (Filter) ColumnNamesWithAlias() []string {
	return []string{
 		"uuid",
 		"filter_op",
 		"filter_key",
 		"filter_value",
 		"created",
 		"updated",
 		"deleted",
	}
}
 
func (Format) ColumnNamesWithAlias() []string {
	return []string{
 		"format_type",
 		"format_data",
 		"created",
 		"updated",
 		"uuid",
	}
}
 
func (Format_property) ColumnNamesWithAlias() []string {
	return []string{
 		"format_type",
 		"format_data",
 		"created",
 		"updated",
	}
}
 
func (row ManagedChannel) Values() []interface{} {
	return []interface{}{
		row.Name,
		row.Summary,
		row.EventCategory,
		row.Created,
		row.Updated,
		row.Deleted,
		row.Uuid,
	}
}
 
func (row ManagedChannel_property) Values() []interface{} {
	return []interface{}{
		row.Name,
		row.Summary,
		row.EventCategory,
	}
}
 
func (row ManagedChannel_option) Values() []interface{} {
	return []interface{}{
		row.Name,
		row.Summary,
		row.EventCategory,
		row.Created,
		row.Updated,
		row.Deleted,
		row.Uuid,
		row.StatusOption.StatusMaxCount,
		row.StatusOption.Created,
		row.StatusOption.Updated,
		row.Format.FormatType,
		row.Format.FormatData,
		row.Format.Created,
		row.Format.Updated,
		row.Edge.NotifierType,
		row.Edge.Created,
		row.Edge.Updated,
	}
}
 
func (row ManagedChannel_tangled) Values() []interface{} {
	return []interface{}{
		row.Name,
		row.Summary,
		row.EventCategory,
		row.Created,
		row.Updated,
		row.Deleted,
		row.Uuid,
		row.StatusOption.StatusMaxCount,
		row.StatusOption.Created,
		row.StatusOption.Updated,
		row.Format.FormatType,
		row.Format.FormatData,
		row.Format.Created,
		row.Format.Updated,
		row.Notifier.Edge.NotifierType,
		row.Notifier.Edge.Created,
		row.Notifier.Edge.Updated,
		row.Notifier.Console.Created,
		row.Notifier.Console.Updated,
		row.Notifier.Webhook.Method,
		row.Notifier.Webhook.Url,
		row.Notifier.Webhook.RequestHeaders,
		row.Notifier.Webhook.RequestTimeout,
		row.Notifier.Webhook.Created,
		row.Notifier.Webhook.Updated,
		row.Notifier.RabbitMq.Url,
		row.Notifier.RabbitMq.ChannelPublish.Exchange,
		row.Notifier.RabbitMq.ChannelPublish.RoutingKey,
		row.Notifier.RabbitMq.ChannelPublish.Mandatory,
		row.Notifier.RabbitMq.ChannelPublish.Immediate,
		row.Notifier.RabbitMq.Publishing.MessageHeaders,
		row.Notifier.RabbitMq.Publishing.MessageContentType,
		row.Notifier.RabbitMq.Publishing.MessageContentEncoding,
		row.Notifier.RabbitMq.Publishing.MessageDeliveryMode,
		row.Notifier.RabbitMq.Publishing.MessagePriority,
		row.Notifier.RabbitMq.Publishing.MessageCorrelationId,
		row.Notifier.RabbitMq.Publishing.MessageReplyTo,
		row.Notifier.RabbitMq.Publishing.MessageExpiration,
		row.Notifier.RabbitMq.Publishing.MessageMessageId,
		row.Notifier.RabbitMq.Publishing.MessageTimestamp,
		row.Notifier.RabbitMq.Publishing.MessageType,
		row.Notifier.RabbitMq.Publishing.MessageUserId,
		row.Notifier.RabbitMq.Publishing.MessageAppId,
		row.Notifier.RabbitMq.Created,
		row.Notifier.RabbitMq.Updated,
	}
}
 
func (row NotifierEdge) Values() []interface{} {
	return []interface{}{
		row.NotifierType,
		row.Created,
		row.Updated,
		row.Uuid,
	}
}
 
func (row NotifierEdge_property) Values() []interface{} {
	return []interface{}{
		row.NotifierType,
		row.Created,
		row.Updated,
	}
}
 
func (row NotifierEdge_option) Values() []interface{} {
	return []interface{}{
		row.NotifierType,
		row.Created,
		row.Updated,
		row.Uuid,
		row.Console.Created,
		row.Console.Updated,
		row.Webhook.Method,
		row.Webhook.Url,
		row.Webhook.RequestHeaders,
		row.Webhook.RequestTimeout,
		row.Webhook.Created,
		row.Webhook.Updated,
		row.RabbitMq.Url,
		row.RabbitMq.ChannelPublish.Exchange,
		row.RabbitMq.ChannelPublish.RoutingKey,
		row.RabbitMq.ChannelPublish.Mandatory,
		row.RabbitMq.ChannelPublish.Immediate,
		row.RabbitMq.Publishing.MessageHeaders,
		row.RabbitMq.Publishing.MessageContentType,
		row.RabbitMq.Publishing.MessageContentEncoding,
		row.RabbitMq.Publishing.MessageDeliveryMode,
		row.RabbitMq.Publishing.MessagePriority,
		row.RabbitMq.Publishing.MessageCorrelationId,
		row.RabbitMq.Publishing.MessageReplyTo,
		row.RabbitMq.Publishing.MessageExpiration,
		row.RabbitMq.Publishing.MessageMessageId,
		row.RabbitMq.Publishing.MessageTimestamp,
		row.RabbitMq.Publishing.MessageType,
		row.RabbitMq.Publishing.MessageUserId,
		row.RabbitMq.Publishing.MessageAppId,
		row.RabbitMq.Created,
		row.RabbitMq.Updated,
	}
}
 
func (row NotifierConsole) Values() []interface{} {
	return []interface{}{
		row.Uuid,
		row.Created,
		row.Updated,
	}
}
 
func (row NotifierRabbitMq) Values() []interface{} {
	return []interface{}{
		row.Url,
		row.ChannelPublish.Exchange,
		row.ChannelPublish.RoutingKey,
		row.ChannelPublish.Mandatory,
		row.ChannelPublish.Immediate,
		row.Publishing.MessageHeaders,
		row.Publishing.MessageContentType,
		row.Publishing.MessageContentEncoding,
		row.Publishing.MessageDeliveryMode,
		row.Publishing.MessagePriority,
		row.Publishing.MessageCorrelationId,
		row.Publishing.MessageReplyTo,
		row.Publishing.MessageExpiration,
		row.Publishing.MessageMessageId,
		row.Publishing.MessageTimestamp,
		row.Publishing.MessageType,
		row.Publishing.MessageUserId,
		row.Publishing.MessageAppId,
		row.Created,
		row.Updated,
		row.Uuid,
	}
}
 
func (row NotifierRabbitMq_property) Values() []interface{} {
	return []interface{}{
		row.Url,
		row.ChannelPublish.Exchange,
		row.ChannelPublish.RoutingKey,
		row.ChannelPublish.Mandatory,
		row.ChannelPublish.Immediate,
		row.Publishing.MessageHeaders,
		row.Publishing.MessageContentType,
		row.Publishing.MessageContentEncoding,
		row.Publishing.MessageDeliveryMode,
		row.Publishing.MessagePriority,
		row.Publishing.MessageCorrelationId,
		row.Publishing.MessageReplyTo,
		row.Publishing.MessageExpiration,
		row.Publishing.MessageMessageId,
		row.Publishing.MessageTimestamp,
		row.Publishing.MessageType,
		row.Publishing.MessageUserId,
		row.Publishing.MessageAppId,
		row.Created,
		row.Updated,
	}
}
 
func (row NotifierWebhook) Values() []interface{} {
	return []interface{}{
		row.Method,
		row.Url,
		row.RequestHeaders,
		row.RequestTimeout,
		row.Created,
		row.Updated,
		row.Uuid,
	}
}
 
func (row NotifierWebhook_property) Values() []interface{} {
	return []interface{}{
		row.Method,
		row.Url,
		row.RequestHeaders,
		row.RequestTimeout,
		row.Created,
		row.Updated,
	}
}
 
func (row ChannelStatusOption) Values() []interface{} {
	return []interface{}{
		row.StatusMaxCount,
		row.Created,
		row.Updated,
		row.Uuid,
	}
}
 
func (row ChannelStatus) Values() []interface{} {
	return []interface{}{
		row.Uuid,
		row.Created,
		row.Message,
	}
}
 
func (row Filter) Values() []interface{} {
	return []interface{}{
		row.Uuid,
		row.FilterOp,
		row.FilterKey,
		row.FilterValue,
		row.Created,
		row.Updated,
		row.Deleted,
	}
}
 
func (row Format) Values() []interface{} {
	return []interface{}{
		row.FormatType,
		row.FormatData,
		row.Created,
		row.Updated,
		row.Uuid,
	}
}
 
func (row Format_property) Values() []interface{} {
	return []interface{}{
		row.FormatType,
		row.FormatData,
		row.Created,
		row.Updated,
	}
}

type Scanner interface {
	Scan(dest ...interface{}) error
}
 
func (row *ManagedChannel) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Name,
		&row.Summary,
		&row.EventCategory,
		&row.Created,
		&row.Updated,
		&row.Deleted,
		&row.Uuid,
	)
}
 
func (row *ManagedChannel_property) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Name,
		&row.Summary,
		&row.EventCategory,
	)
}
 
func (row *ManagedChannel_option) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Name,
		&row.Summary,
		&row.EventCategory,
		&row.Created,
		&row.Updated,
		&row.Deleted,
		&row.Uuid,
		&row.StatusOption.StatusMaxCount,
		&row.StatusOption.Created,
		&row.StatusOption.Updated,
		&row.Format.FormatType,
		&row.Format.FormatData,
		&row.Format.Created,
		&row.Format.Updated,
		&row.Edge.NotifierType,
		&row.Edge.Created,
		&row.Edge.Updated,
	)
}
 
func (row *ManagedChannel_tangled) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Name,
		&row.Summary,
		&row.EventCategory,
		&row.Created,
		&row.Updated,
		&row.Deleted,
		&row.Uuid,
		&row.StatusOption.StatusMaxCount,
		&row.StatusOption.Created,
		&row.StatusOption.Updated,
		&row.Format.FormatType,
		&row.Format.FormatData,
		&row.Format.Created,
		&row.Format.Updated,
		&row.Notifier.Edge.NotifierType,
		&row.Notifier.Edge.Created,
		&row.Notifier.Edge.Updated,
		&row.Notifier.Console.Created,
		&row.Notifier.Console.Updated,
		&row.Notifier.Webhook.Method,
		&row.Notifier.Webhook.Url,
		&row.Notifier.Webhook.RequestHeaders,
		&row.Notifier.Webhook.RequestTimeout,
		&row.Notifier.Webhook.Created,
		&row.Notifier.Webhook.Updated,
		&row.Notifier.RabbitMq.Url,
		&row.Notifier.RabbitMq.ChannelPublish.Exchange,
		&row.Notifier.RabbitMq.ChannelPublish.RoutingKey,
		&row.Notifier.RabbitMq.ChannelPublish.Mandatory,
		&row.Notifier.RabbitMq.ChannelPublish.Immediate,
		&row.Notifier.RabbitMq.Publishing.MessageHeaders,
		&row.Notifier.RabbitMq.Publishing.MessageContentType,
		&row.Notifier.RabbitMq.Publishing.MessageContentEncoding,
		&row.Notifier.RabbitMq.Publishing.MessageDeliveryMode,
		&row.Notifier.RabbitMq.Publishing.MessagePriority,
		&row.Notifier.RabbitMq.Publishing.MessageCorrelationId,
		&row.Notifier.RabbitMq.Publishing.MessageReplyTo,
		&row.Notifier.RabbitMq.Publishing.MessageExpiration,
		&row.Notifier.RabbitMq.Publishing.MessageMessageId,
		&row.Notifier.RabbitMq.Publishing.MessageTimestamp,
		&row.Notifier.RabbitMq.Publishing.MessageType,
		&row.Notifier.RabbitMq.Publishing.MessageUserId,
		&row.Notifier.RabbitMq.Publishing.MessageAppId,
		&row.Notifier.RabbitMq.Created,
		&row.Notifier.RabbitMq.Updated,
	)
}
 
func (row *NotifierEdge) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.NotifierType,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	)
}
 
func (row *NotifierEdge_property) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.NotifierType,
		&row.Created,
		&row.Updated,
	)
}
 
func (row *NotifierEdge_option) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.NotifierType,
		&row.Created,
		&row.Updated,
		&row.Uuid,
		&row.Console.Created,
		&row.Console.Updated,
		&row.Webhook.Method,
		&row.Webhook.Url,
		&row.Webhook.RequestHeaders,
		&row.Webhook.RequestTimeout,
		&row.Webhook.Created,
		&row.Webhook.Updated,
		&row.RabbitMq.Url,
		&row.RabbitMq.ChannelPublish.Exchange,
		&row.RabbitMq.ChannelPublish.RoutingKey,
		&row.RabbitMq.ChannelPublish.Mandatory,
		&row.RabbitMq.ChannelPublish.Immediate,
		&row.RabbitMq.Publishing.MessageHeaders,
		&row.RabbitMq.Publishing.MessageContentType,
		&row.RabbitMq.Publishing.MessageContentEncoding,
		&row.RabbitMq.Publishing.MessageDeliveryMode,
		&row.RabbitMq.Publishing.MessagePriority,
		&row.RabbitMq.Publishing.MessageCorrelationId,
		&row.RabbitMq.Publishing.MessageReplyTo,
		&row.RabbitMq.Publishing.MessageExpiration,
		&row.RabbitMq.Publishing.MessageMessageId,
		&row.RabbitMq.Publishing.MessageTimestamp,
		&row.RabbitMq.Publishing.MessageType,
		&row.RabbitMq.Publishing.MessageUserId,
		&row.RabbitMq.Publishing.MessageAppId,
		&row.RabbitMq.Created,
		&row.RabbitMq.Updated,
	)
}
 
func (row *NotifierConsole) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Uuid,
		&row.Created,
		&row.Updated,
	)
}
 
func (row *NotifierRabbitMq) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Url,
		&row.ChannelPublish.Exchange,
		&row.ChannelPublish.RoutingKey,
		&row.ChannelPublish.Mandatory,
		&row.ChannelPublish.Immediate,
		&row.Publishing.MessageHeaders,
		&row.Publishing.MessageContentType,
		&row.Publishing.MessageContentEncoding,
		&row.Publishing.MessageDeliveryMode,
		&row.Publishing.MessagePriority,
		&row.Publishing.MessageCorrelationId,
		&row.Publishing.MessageReplyTo,
		&row.Publishing.MessageExpiration,
		&row.Publishing.MessageMessageId,
		&row.Publishing.MessageTimestamp,
		&row.Publishing.MessageType,
		&row.Publishing.MessageUserId,
		&row.Publishing.MessageAppId,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	)
}
 
func (row *NotifierRabbitMq_property) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Url,
		&row.ChannelPublish.Exchange,
		&row.ChannelPublish.RoutingKey,
		&row.ChannelPublish.Mandatory,
		&row.ChannelPublish.Immediate,
		&row.Publishing.MessageHeaders,
		&row.Publishing.MessageContentType,
		&row.Publishing.MessageContentEncoding,
		&row.Publishing.MessageDeliveryMode,
		&row.Publishing.MessagePriority,
		&row.Publishing.MessageCorrelationId,
		&row.Publishing.MessageReplyTo,
		&row.Publishing.MessageExpiration,
		&row.Publishing.MessageMessageId,
		&row.Publishing.MessageTimestamp,
		&row.Publishing.MessageType,
		&row.Publishing.MessageUserId,
		&row.Publishing.MessageAppId,
		&row.Created,
		&row.Updated,
	)
}
 
func (row *NotifierWebhook) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Method,
		&row.Url,
		&row.RequestHeaders,
		&row.RequestTimeout,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	)
}
 
func (row *NotifierWebhook_property) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Method,
		&row.Url,
		&row.RequestHeaders,
		&row.RequestTimeout,
		&row.Created,
		&row.Updated,
	)
}
 
func (row *ChannelStatusOption) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.StatusMaxCount,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	)
}
 
func (row *ChannelStatus) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Uuid,
		&row.Created,
		&row.Message,
	)
}
 
func (row *Filter) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.Uuid,
		&row.FilterOp,
		&row.FilterKey,
		&row.FilterValue,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	)
}
 
func (row *Format) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.FormatType,
		&row.FormatData,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	)
}
 
func (row *Format_property) Scan(scanner Scanner) error {
	return scanner.Scan(
		&row.FormatType,
		&row.FormatData,
		&row.Created,
		&row.Updated,
	)
}
 
func (row *ManagedChannel) Ptrs() []interface{} {
	return []interface{}{
		&row.Name,
		&row.Summary,
		&row.EventCategory,
		&row.Created,
		&row.Updated,
		&row.Deleted,
		&row.Uuid,
	}
}
 
func (row *ManagedChannel_property) Ptrs() []interface{} {
	return []interface{}{
		&row.Name,
		&row.Summary,
		&row.EventCategory,
	}
}
 
func (row *ManagedChannel_option) Ptrs() []interface{} {
	return []interface{}{
		&row.Name,
		&row.Summary,
		&row.EventCategory,
		&row.Created,
		&row.Updated,
		&row.Deleted,
		&row.Uuid,
		&row.StatusOption.StatusMaxCount,
		&row.StatusOption.Created,
		&row.StatusOption.Updated,
		&row.Format.FormatType,
		&row.Format.FormatData,
		&row.Format.Created,
		&row.Format.Updated,
		&row.Edge.NotifierType,
		&row.Edge.Created,
		&row.Edge.Updated,
	}
}
 
func (row *ManagedChannel_tangled) Ptrs() []interface{} {
	return []interface{}{
		&row.Name,
		&row.Summary,
		&row.EventCategory,
		&row.Created,
		&row.Updated,
		&row.Deleted,
		&row.Uuid,
		&row.StatusOption.StatusMaxCount,
		&row.StatusOption.Created,
		&row.StatusOption.Updated,
		&row.Format.FormatType,
		&row.Format.FormatData,
		&row.Format.Created,
		&row.Format.Updated,
		&row.Notifier.Edge.NotifierType,
		&row.Notifier.Edge.Created,
		&row.Notifier.Edge.Updated,
		&row.Notifier.Console.Created,
		&row.Notifier.Console.Updated,
		&row.Notifier.Webhook.Method,
		&row.Notifier.Webhook.Url,
		&row.Notifier.Webhook.RequestHeaders,
		&row.Notifier.Webhook.RequestTimeout,
		&row.Notifier.Webhook.Created,
		&row.Notifier.Webhook.Updated,
		&row.Notifier.RabbitMq.Url,
		&row.Notifier.RabbitMq.ChannelPublish.Exchange,
		&row.Notifier.RabbitMq.ChannelPublish.RoutingKey,
		&row.Notifier.RabbitMq.ChannelPublish.Mandatory,
		&row.Notifier.RabbitMq.ChannelPublish.Immediate,
		&row.Notifier.RabbitMq.Publishing.MessageHeaders,
		&row.Notifier.RabbitMq.Publishing.MessageContentType,
		&row.Notifier.RabbitMq.Publishing.MessageContentEncoding,
		&row.Notifier.RabbitMq.Publishing.MessageDeliveryMode,
		&row.Notifier.RabbitMq.Publishing.MessagePriority,
		&row.Notifier.RabbitMq.Publishing.MessageCorrelationId,
		&row.Notifier.RabbitMq.Publishing.MessageReplyTo,
		&row.Notifier.RabbitMq.Publishing.MessageExpiration,
		&row.Notifier.RabbitMq.Publishing.MessageMessageId,
		&row.Notifier.RabbitMq.Publishing.MessageTimestamp,
		&row.Notifier.RabbitMq.Publishing.MessageType,
		&row.Notifier.RabbitMq.Publishing.MessageUserId,
		&row.Notifier.RabbitMq.Publishing.MessageAppId,
		&row.Notifier.RabbitMq.Created,
		&row.Notifier.RabbitMq.Updated,
	}
}
 
func (row *NotifierEdge) Ptrs() []interface{} {
	return []interface{}{
		&row.NotifierType,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	}
}
 
func (row *NotifierEdge_property) Ptrs() []interface{} {
	return []interface{}{
		&row.NotifierType,
		&row.Created,
		&row.Updated,
	}
}
 
func (row *NotifierEdge_option) Ptrs() []interface{} {
	return []interface{}{
		&row.NotifierType,
		&row.Created,
		&row.Updated,
		&row.Uuid,
		&row.Console.Created,
		&row.Console.Updated,
		&row.Webhook.Method,
		&row.Webhook.Url,
		&row.Webhook.RequestHeaders,
		&row.Webhook.RequestTimeout,
		&row.Webhook.Created,
		&row.Webhook.Updated,
		&row.RabbitMq.Url,
		&row.RabbitMq.ChannelPublish.Exchange,
		&row.RabbitMq.ChannelPublish.RoutingKey,
		&row.RabbitMq.ChannelPublish.Mandatory,
		&row.RabbitMq.ChannelPublish.Immediate,
		&row.RabbitMq.Publishing.MessageHeaders,
		&row.RabbitMq.Publishing.MessageContentType,
		&row.RabbitMq.Publishing.MessageContentEncoding,
		&row.RabbitMq.Publishing.MessageDeliveryMode,
		&row.RabbitMq.Publishing.MessagePriority,
		&row.RabbitMq.Publishing.MessageCorrelationId,
		&row.RabbitMq.Publishing.MessageReplyTo,
		&row.RabbitMq.Publishing.MessageExpiration,
		&row.RabbitMq.Publishing.MessageMessageId,
		&row.RabbitMq.Publishing.MessageTimestamp,
		&row.RabbitMq.Publishing.MessageType,
		&row.RabbitMq.Publishing.MessageUserId,
		&row.RabbitMq.Publishing.MessageAppId,
		&row.RabbitMq.Created,
		&row.RabbitMq.Updated,
	}
}
 
func (row *NotifierConsole) Ptrs() []interface{} {
	return []interface{}{
		&row.Uuid,
		&row.Created,
		&row.Updated,
	}
}
 
func (row *NotifierRabbitMq) Ptrs() []interface{} {
	return []interface{}{
		&row.Url,
		&row.ChannelPublish.Exchange,
		&row.ChannelPublish.RoutingKey,
		&row.ChannelPublish.Mandatory,
		&row.ChannelPublish.Immediate,
		&row.Publishing.MessageHeaders,
		&row.Publishing.MessageContentType,
		&row.Publishing.MessageContentEncoding,
		&row.Publishing.MessageDeliveryMode,
		&row.Publishing.MessagePriority,
		&row.Publishing.MessageCorrelationId,
		&row.Publishing.MessageReplyTo,
		&row.Publishing.MessageExpiration,
		&row.Publishing.MessageMessageId,
		&row.Publishing.MessageTimestamp,
		&row.Publishing.MessageType,
		&row.Publishing.MessageUserId,
		&row.Publishing.MessageAppId,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	}
}
 
func (row *NotifierRabbitMq_property) Ptrs() []interface{} {
	return []interface{}{
		&row.Url,
		&row.ChannelPublish.Exchange,
		&row.ChannelPublish.RoutingKey,
		&row.ChannelPublish.Mandatory,
		&row.ChannelPublish.Immediate,
		&row.Publishing.MessageHeaders,
		&row.Publishing.MessageContentType,
		&row.Publishing.MessageContentEncoding,
		&row.Publishing.MessageDeliveryMode,
		&row.Publishing.MessagePriority,
		&row.Publishing.MessageCorrelationId,
		&row.Publishing.MessageReplyTo,
		&row.Publishing.MessageExpiration,
		&row.Publishing.MessageMessageId,
		&row.Publishing.MessageTimestamp,
		&row.Publishing.MessageType,
		&row.Publishing.MessageUserId,
		&row.Publishing.MessageAppId,
		&row.Created,
		&row.Updated,
	}
}
 
func (row *NotifierWebhook) Ptrs() []interface{} {
	return []interface{}{
		&row.Method,
		&row.Url,
		&row.RequestHeaders,
		&row.RequestTimeout,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	}
}
 
func (row *NotifierWebhook_property) Ptrs() []interface{} {
	return []interface{}{
		&row.Method,
		&row.Url,
		&row.RequestHeaders,
		&row.RequestTimeout,
		&row.Created,
		&row.Updated,
	}
}
 
func (row *ChannelStatusOption) Ptrs() []interface{} {
	return []interface{}{
		&row.StatusMaxCount,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	}
}
 
func (row *ChannelStatus) Ptrs() []interface{} {
	return []interface{}{
		&row.Uuid,
		&row.Created,
		&row.Message,
	}
}
 
func (row *Filter) Ptrs() []interface{} {
	return []interface{}{
		&row.Uuid,
		&row.FilterOp,
		&row.FilterKey,
		&row.FilterValue,
		&row.Created,
		&row.Updated,
		&row.Deleted,
	}
}
 
func (row *Format) Ptrs() []interface{} {
	return []interface{}{
		&row.FormatType,
		&row.FormatData,
		&row.Created,
		&row.Updated,
		&row.Uuid,
	}
}
 
func (row *Format_property) Ptrs() []interface{} {
	return []interface{}{
		&row.FormatType,
		&row.FormatData,
		&row.Created,
		&row.Updated,
	}
}
