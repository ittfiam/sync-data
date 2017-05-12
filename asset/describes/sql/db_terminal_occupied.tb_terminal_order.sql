drop table if exists tb_terminal_order;
CREATE TABLE `tb_terminal_order` (
  `terminal_id` varchar(32)  COMMENT '终端id',
  `uid` varchar(32)  COMMENT 'uid',
  `order_time` datetime  COMMENT '下单时间',
  `occupied_days` int(11)  COMMENT '店铺霸占天数',
  `data_version` int(11)  DEFAULT '0' COMMENT '数据版本',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`terminal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='终端最近下单记录'