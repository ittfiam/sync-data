drop table if exists tb_user_distributor_deliver;
CREATE TABLE `tb_user_distributor_deliver` (
  `uid` varchar(32)  COMMENT '用户ID',
  `name` varchar(128)  COMMENT '姓名',
  `name_pys` varchar(32)  COMMENT '姓名拼音首字母缩写',
  `distributor_corp_id` varchar(32)  COMMENT '所属分销商企业ID',
  `create_time` datetime  COMMENT '创建记录时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新记录时间',
  PRIMARY KEY (`uid`),
  KEY `distributor_corp_id` (`distributor_corp_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分销商送货员个人信息表'