-- +goose Up
-- +goose StatementBegin
create table `sys_secret`
(
    `id`              varchar(64)       not null comment 'id',
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `name`            varchar(120)      not null comment '名称',
    `secret`          varchar(255)      not null comment 'Secret',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='系统-开发密钥表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sys_secret`;
-- +goose StatementEnd
