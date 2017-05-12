drop table if exists tb_user_internal_district_manager;
CREATE TABLE `tb_user_internal_district_manager` (
  `uid` varchar(32)  COMMENT '用户ID',
  `name` varchar(128)  COMMENT '姓名',
  `name_pys` varchar(32)  COMMENT '姓名拼音首字母缩写',
  `id_type` int(11)  DEFAULT '0' COMMENT '证件类型 0=默认身份证',
  `id` varchar(128)  COMMENT '证件号码',
  `country_code` varchar(8)  COMMENT '国家编码',
  `country` varchar(256)  COMMENT '国家名称',
  `district_code` varchar(8)  COMMENT '大区编码',
  `district` varchar(256)  COMMENT '大区名称',
  `add_uid` varchar(32)  COMMENT '添加人uid',
  `add_time` datetime  COMMENT '添加时间',
  `set_uid` varchar(32)  COMMENT '修改人uid',
  `set_time` datetime  COMMENT '修改时间',
  `create_time` datetime  COMMENT '创建记录时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新记录时间',
  PRIMARY KEY (`uid`),
  KEY `id_type` (`id_type`),
  KEY `id` (`id`),
  KEY `country_code` (`country_code`),
  KEY `district_code` (`district_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='大区经理个人信息表'