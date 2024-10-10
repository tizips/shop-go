-- +goose Up
-- +goose StatementBegin
create table `shp_product`
(
    `id`               varchar(64)       not null,
    `platform`         smallint unsigned not null default 0 comment '平台',
    `clique_id`        varchar(64)                default null comment '集团ID',
    `organization_id`  varchar(64)                default null comment '组织ID',
    `i1_category_id`   int unsigned      not null default 0 comment '一级栏目ID',
    `i2_category_id`   int unsigned      not null default 0 comment '二级栏目ID',
    `i3_category_id`   int unsigned      not null default 0 comment '三级栏目ID',
    `name`             varchar(255)      not null default '' comment '名称',
    `summary`          text                       default null comment '简介',
    `is_hot`           tinyint unsigned  not null default 0 comment '是否热销：1=是；2=否',
    `is_recommend`     tinyint unsigned  not null default 0 comment '是否推荐：1=是；2=否',
    `is_multiple`      tinyint unsigned  not null default 0 comment '是否多规格：1=是；2=否',
    `is_free_shipping` tinyint unsigned  not null default 0 comment '是否包邮：1=是；2=否',
    `is_freeze`        tinyint unsigned  not null default 0 comment '是否冻结：1=是；2=否',
    `is_enable`        tinyint unsigned  not null default 0 comment '是否启用：1=是；2=否；',
    `created_at`       timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`       timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`       timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`i1_category_id`),
    key (`i2_category_id`),
    key (`i3_category_id`),
    key (`deleted_at`)
) default collate = utf8mb4_unicode_ci comment ='商城-产品表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_product`;
-- +goose StatementEnd
