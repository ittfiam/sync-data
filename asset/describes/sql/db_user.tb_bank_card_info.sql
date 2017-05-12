drop table if exists tb_bank_card_info;
CREATE TABLE `tb_bank_card_info` (
  `uid` varchar(32)  COMMENT '用户ID',
  `name` varchar(128)  COMMENT '真实姓名',
  `name_pys` varchar(32)  COMMENT '真实姓名拼音首字母缩写',
  `bank_name` varchar(256)  COMMENT '银行名称',
  `bank_code` varchar(128)  COMMENT '银行编码',
  `province_code` varchar(32)  DEFAULT '00' COMMENT '开户行省代码',
  `city_code` varchar(32)  DEFAULT '00000' COMMENT '开户行市代码',
  `card` varchar(32)  COMMENT '银行卡号',
  `id` varchar(32)  COMMENT '身份证号',
  `status` int(11)  DEFAULT '0' COMMENT '状态:0=正常，1=禁用',
  `create_time` datetime  COMMENT '创建记录时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`uid`),
  KEY `bank_code` (`bank_code`),
  KEY `id` (`id`),
  KEY `status` (`status`),
  KEY `card` (`card`),
  KEY `province_code` (`province_code`) USING BTREE,
  KEY `city_code` (`city_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='银行卡信息表'