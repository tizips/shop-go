-- +goose Up
-- +goose StatementBegin
create table `shp_service`
(
    `id`                    varchar(64)       not null,
    `platform`              smallint unsigned not null default 0 comment '平台',
    `clique_id`             varchar(64)                default null comment '集团ID',
    `organization_id`       varchar(64)                default null comment '组织ID',
    `user_id`               varchar(64)       not null comment '用户id',
    `order_id`              varchar(64)       not null comment '订单ID',
    `detail_id`             varchar(64)                default null comment '明细ID',
    `type`                  varchar(10)       not null default '' comment '类型：un_receipt=未收到货；refund=退货退款；exchange=换货',
    `status`                varchar(16)       not null default '' comment '状态：pending=待处理；user=等待用户发货；org=等待商家发货；confirm_user=等待用户确认；confirm_org=等待商家确认；finish=完成；closed=已关闭',
    `result`                varchar(6)        not null default '' comment '结果：agree=同意；refuse=拒绝',
    `reason`                varchar(255)      not null default '' comment '原因',
    `pictures`              text                       default null comment '证据图',
    `subtotal`              int unsigned      not null default 0 comment '商品退款',
    `shipping`              int unsigned      not null default 0 comment '运费退款',
    `refund`                int unsigned      not null default 0 comment '退款金额',
    `shipment_user`         text                       default null comment '客户快递信息',
    `shipment_organization` text                       default null comment '商户快递信息',
    `created_at`            timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`            timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`            timestamp                  default null,
    primary key (`id`),
    key (`user_id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`order_id`),
    key (`detail_id`),
    key (`user_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-售后表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_service`;
-- +goose StatementEnd
