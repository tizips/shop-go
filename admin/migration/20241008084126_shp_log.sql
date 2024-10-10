-- +goose Up
-- +goose StatementBegin
create table `shp_log`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户ID',
    `order_id`        varchar(64)       not null comment '订单ID',
    `detail_id`       varchar(64)                default null comment '明细ID',
    `service_id`      varchar(64)                default null comment '售后ID',
    `action`          varchar(32)       not null default '' comment '操作',
    `content`         varchar(255)      not null default '' comment '内容',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`order_id`),
    key (`detail_id`),
    key (`service_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-订单日志表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_log`;
-- +goose StatementEnd
