drop table if exists tb_user_internal_super_manager;
CREATE TABLE `tb_user_internal_super_manager` (
  `uid` varchar(32)  COMMENT '用户ID',
  `name` varchar(128)  COMMENT '姓名',
  `name_pys` varchar(32)  COMMENT '姓名拼音首字母缩写',
  `id_type` int(11)  DEFAULT '0' COMMENT '证件类型 0=默认身份证',
  `id` varchar(128)  COMMENT '证件号码',
  PRIMARY KEY (`uid`),
  KEY `id_type` (`id_type`),
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='超级管理员个人信息表'