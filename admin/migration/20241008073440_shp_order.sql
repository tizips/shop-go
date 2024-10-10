-- +goose Up
-- +goose StatementBegin
create table `shp_order`
(
    `id`              varchar(64)       not null,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户id',
    `cost_shipping`   int unsigned      not null default 0 comment '运费',
    `total_price`     int unsigned      not null default 0 comment '总价',
    `coupon_price`    int unsigned      not null default 0 comment '优惠',
    `cost_prices`     int unsigned      not null default 0 comment '成本合计',
    `prices`          int unsigned      not null default 0 comment '合计',
    `refund`          int unsigned      not null default 0 comment '退款',
    `status`          varchar(10)       not null default '' comment '订单状态：pay=待支付；shipment=待发货；receipt=待收货；evaluate=待评价；completed=已完成；cancel=已取消；close=已关闭；refund=已退款',
    `remark`          varchar(255)      not null default '' comment '备注',
    `payment_id`      varchar(64)                default null comment '支付ID',
    `is_paid`         tinyint unsigned  not null default 0 comment '是否支付：1=是；2=否',
    `is_invoice`      tinyint unsigned  not null default 0 comment '是否开发票：1=是；2=否',
    `is_appraisal`    tinyint unsigned  not null default 0 comment '是否评价：1=是；2=否',
    `completed_at`    timestamp                  default null comment '完成时间',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`payment_id`)
) default collate = utf8mb4_unicode_ci comment ='商城-订单表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_order`;
-- +goose StatementEnd
