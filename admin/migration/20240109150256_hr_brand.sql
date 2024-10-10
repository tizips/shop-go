-- +goose Up
-- +goose StatementBegin
create table `hr_brand`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null      default 0 comment '平台',
    `organization_id` varchar(64)                     default null comment '组织ID',
    `name`            varchar(64)       not null      default '' comment '名称',
    `logo`            varchar(255)      not null null default '' comment 'LOGO',
    `order`           tinyint unsigned  not null      default 0 comment '序号：正序',
    `created_at`      timestamp         not null      default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null      default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                       default null,
    primary key (`id`),
    key (`organization_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '人资-品牌表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `hr_brand`;
-- +goose StatementEnd
