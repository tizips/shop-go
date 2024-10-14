-- +goose Up
-- +goose StatementBegin
create table `shp_service_detail`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null default '' comment '用户id',
    `order_id`        varchar(64)       not null default '' comment '订单ID',
    `product_id`      varchar(64)       not null default '' comment '产品ID',
    `service_id`      varchar(64)       not null default '' comment '售后ID',
    `detail_id`       varchar(64)       not null default '' comment '明细ID',
    `quantity`        int unsigned      not null default 0 comment '数量',
    `refund`          int unsigned      not null default 0 comment '退款',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`order_id`),
    key (`product_id`),
    key (`service_id`),
    key (`detail_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-售后明细表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_service_detail`;
-- +goose StatementEnd
