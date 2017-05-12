drop table if exists tb_terminal_delivery;
CREATE TABLE `tb_terminal_delivery` (
  `terminal_id` varchar(32)  COMMENT '终端ID',
  `distributor_id` varchar(32)  COMMENT '配送点ID',
  `distributor_name` varchar(32)  COMMENT '配送点名称',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`terminal_id`,`distributor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='终端的配送关系表'