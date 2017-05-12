drop table if exists tb_phones;
CREATE TABLE `tb_phones` (
  `distributor_id` varchar(64)  COMMENT '分销商ID',
  `phone` varchar(128)  COMMENT '业务电话',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`distributor_id`,`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分销商业务电话列表'