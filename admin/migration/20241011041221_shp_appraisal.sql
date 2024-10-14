-- +goose Up
-- +goose StatementBegin
create table `shp_appraisal`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户id',
    `order_id`        varchar(64)       not null comment '订单ID',
    `star_product`    tinyint unsigned  not null default 0 comment '商品评分：1-5',
    `star_shipment`   tinyint unsigned  not null default 0 comment '物流评分：1-5',
    `remark`          varchar(255)      not null default '' comment '备注',
    `pictures`        text                       default null comment '证据图',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`user_id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`order_id`),
    key (`user_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-评价表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_appraisal`;
-- +goose StatementEnd
