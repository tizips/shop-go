-- +goose Up
-- +goose StatementBegin
create table `com_setting`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `organization_id` varchar(64)                default null comment '组织ID',
    `module`          varchar(20)       not null default '' comment '模块',
    `type`            varchar(10)       not null default '' comment '字段类型：input/textarea/enable/url/email/picture',
    `label`           varchar(10)       not null default '' comment '名称',
    `key`             varchar(20)       not null default '' comment '键',
    `val`             text                       default null comment '值',
    `is_required`     tinyint unsigned  not null default 0 comment '是否必填：1=是；2=否',
    `order`           tinyint unsigned  not null default 50 comment '序号',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default NULL,
    primary key (`id`),
    key (`organization_id`)
) auto_increment = 1000
  collate = utf8mb4_unicode_ci comment ='公共-设置表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `com_setting`;
-- +goose StatementEnd
