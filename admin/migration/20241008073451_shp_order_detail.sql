-- +goose Up
-- +goose StatementBegin
create table `shp_order_detail`
(
    `id`              varchar(64)       not null,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户id',
    `order_id`        varchar(64)       not null comment '订单ID',
    `product_id`      varchar(64)       not null comment '产品ID',
    `sku_id`          varchar(64)       not null comment 'SkuID',
    `appraisal_id`    int unsigned      not null default 0 comment '评价ID',
    `name`            varchar(255)      not null default '' comment '名称',
    `specifications`  text              not null comment '规格',
    `picture`         varchar(255)      not null default '' comment '图片',
    `price`           int unsigned      not null default 0 comment '价格',
    `cost_price`      int unsigned      not null default 0 comment '成本',
    `quantity`        int unsigned      not null default 0 comment '数量',
    `total_price`     int unsigned      not null default 0 comment '总价',
    `coupon_price`    int unsigned      not null default 0 comment '优惠',
    `prices`          int unsigned      not null default 0 comment '合计',
    `cost_prices`     int unsigned      not null default 0 comment '成本合计',
    `weight`          int unsigned      not null default 0 comment '重量',
    `refund`          int unsigned      not null default 0 comment '退款',
    `returned`        int unsigned      not null default 0 comment '退货数量',
    `is_invoiced`     tinyint unsigned  not null default 0 comment '是否已开发票：1=是；2=否',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`order_id`),
    key (`product_id`),
    key (`sku_id`),
    key (`appraisal_id`)
) default collate = utf8mb4_unicode_ci comment ='商城-订单明细表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_order_detail`;
-- +goose StatementEnd
