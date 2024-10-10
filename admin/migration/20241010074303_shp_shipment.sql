-- +goose Up
-- +goose StatementBegin
create table `shp_shipment`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `order_id`        varchar(64)       not null comment '订单ID',
    `user_id`         varchar(64)       not null comment '用户id',
    `shipping_id`     int unsigned      not null default 0 comment '快递ID',
    `money`           int unsigned      not null default 0 comment '费用',
    `company`         varchar(255)      not null default '' comment '快递公司',
    `no`              varchar(255)      not null default '' comment '快递单号',
    `remark`          varchar(255)      not null default '' comment '备注',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`user_id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`order_id`),
    key (`shipping_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-快递表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_shipment`;
-- +goose StatementEnd
