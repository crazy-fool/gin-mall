
--  用户表
CREATE TABLE `hmh_customer` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'code',
    `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '姓名',
    `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '账号',
    `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密码',
    `mobile` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '手机号',
    `last_login_at` datetime DEFAULT NULL COMMENT '最近登录时间',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `un_account` (`account`),
    UNIQUE KEY `un_code` (`code`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类目名称',
    `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父id',
    `is_parent` tinyint(1) DEFAULT '0' COMMENT '是否是父id',
    `sort` int(11) DEFAULT '0' COMMENT '排序',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_sku` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'skuname',
    `sku` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'sku',
    `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '售价',
    `market_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '市场价',
    `img` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '图片',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_sku_stock` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `sku` varchar(20) COLLATE utf8mb4_general_ci DEFAULT '',
    `stock` int(11) DEFAULT '0',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_spec` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `category_id` int(11) DEFAULT '0' COMMENT '商品分类id',
    `group_id` int(11) DEFAULT '0' COMMENT '分组id',
    `name` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数名称',
    `is_generate` tinyint(1) DEFAULT '1' COMMENT '是否通用',
    `searching` tinyint(1) DEFAULT '0' COMMENT '是否用于搜索',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `hmh_spec_group` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品分类id',
    `name` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '规格组名称',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_spu` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `spu` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'spucode',
    `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
    `img` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '图片',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态1启用，0、禁用',
    `brand_bn` varchar(20) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '品牌code',
    `store_bn` varchar(20) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '商户号',
    `sale_cnt` int(11) DEFAULT '0' COMMENT '销量',
    `price_range` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '价格区间',
    `cate_id` int(11) DEFAULT '0' COMMENT '子分类id',
    `one_cate_id` int(11) DEFAULT '0' COMMENT '一级分类id',
    `two_cate_id` int(11) DEFAULT '0' COMMENT '二级分类id',
    `created_at` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_sku_detail` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `hmh_spu_detail` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;