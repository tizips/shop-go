-- +goose Up
-- +goose StatementBegin
create table `shp_refund`
(
    `id`              varchar(64)       not null,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户id',
    `order_id`        varchar(64)                default null comment '订单ID',
    `detail_id`       varchar(64)                default null comment '明细ID',
    `payment_id`      varchar(64)                default null comment '支付ID',
    `service_id`      varchar(64)                default null comment '售后ID',
    `no`              varchar(64)       not null comment '第三方支付单号',
    `channel`         varchar(10)       not null comment '支付渠道：paypal=贝宝',
    `money`           int unsigned      not null default 0 comment '退款金额',
    `currency`        varchar(16)       not null default '' comment '币种',
    `is_confirmed`    tinyint unsigned  not null default 0 comment '是否确认：1=是；2=否',
    `remark`          varchar(255)      not null default '' comment '备注',
    `ext`             text              not null comment '扩展信息',
    `refunded_at`     timestamp                  default null comment '退款时间',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`order_id`),
    key (`detail_id`),
    key (`payment_id`),
    key (`service_id`),
    key (`no`)
) default collate = utf8mb4_unicode_ci comment ='商城-退款表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_refund`;
-- +goose StatementEnd
