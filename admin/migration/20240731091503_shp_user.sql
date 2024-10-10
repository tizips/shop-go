-- +goose Up
-- +goose StatementBegin
create table `shp_user`
(
    `id`              varchar(64)       not null,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `email`           varchar(64)       not null comment '邮箱',
    `first_name`      varchar(120)      not null default '' comment 'FIRST NAME',
    `last_name`       varchar(120)      not null default '' comment 'LAST NAME',
    `password`        varchar(64)       not null default '' comment '密码',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`email`),
    key (`deleted_at`)
) default collate = utf8mb4_unicode_ci comment ='商城-用户表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_user`;
-- +goose StatementEnd
