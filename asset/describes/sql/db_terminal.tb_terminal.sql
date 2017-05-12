drop table if exists tb_terminal;
CREATE TABLE `tb_terminal` (
  `terminal_id` varchar(32)  COMMENT '终端ID',
  `terminal_name` varchar(256)  COMMENT '终端名字',
  `shopkeeper_mobile` varchar(16)  COMMENT '联系手机',
  `shopkeeper_phone` varchar(16)  DEFAULT '' COMMENT '联系座机',
  `shopkeeper_name` varchar(64)  COMMENT '老板姓名',
  `shop_picture` varchar(64)  COMMENT '店头照key',
  `province_code` int(11)  DEFAULT '0' COMMENT '省份编码',
  `city_code` int(11)  DEFAULT '0' COMMENT '城市编码',
  `region_code` int(11)  DEFAULT '0' COMMENT '区域编码',
  `province` varchar(64)  DEFAULT '' COMMENT '省名称',
  `city` varchar(64)  DEFAULT '' COMMENT '城市名称',
  `region` varchar(64)  DEFAULT '' COMMENT '区域名称',
  `details` varchar(128)  DEFAULT '' COMMENT '详细地址',
  `longitude` double  DEFAULT '0' COMMENT '经度',
  `latitude` double  DEFAULT '0' COMMENT '纬度',
  `description` varchar(1024)  DEFAULT '' COMMENT 'geo信息',
  `first_spell` varchar(128)  COMMENT '店名拼音首字母',
  `full_spell` varchar(600)  COMMENT '店名全部拼音',
  `add_uid` varchar(32)  COMMENT '添加者UID',
  `check_status` int(11)  DEFAULT '0' COMMENT '审核状态（0-待审核；1-审核通过；2-审核拒绝）',
  `status` int(11)  COMMENT '状态标识(0-正常；1-禁用）',
  `forbidden_reason` varchar(255) DEFAULT '' COMMENT '禁用理由',
  `main_type` int(11)  DEFAULT '0' COMMENT '大类型',
  `sub_type` int(11)  DEFAULT '0' COMMENT '小类型',
  `type_level` int(11)  DEFAULT '0' COMMENT '终端类型级别',
  `data_version` int(11)  DEFAULT '0' COMMENT '数据版本',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`terminal_id`),
  KEY `index_terminal_name` (`terminal_name`(255)),
  KEY `index_shopkeeper_name` (`shopkeeper_name`),
  KEY `index_shopkeeper_mobile` (`shopkeeper_mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='终端表'