-- +goose Up
-- +goose StatementBegin
create table `shp_wishlist`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户ID',
    `product_id`      varchar(64)       not null comment '产品ID',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`product_id`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-愿望清单表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_wishlist`;
-- +goose StatementEnd
