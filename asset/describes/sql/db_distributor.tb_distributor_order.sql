drop table if exists tb_distributor_order;
CREATE TABLE `tb_distributor_order` (
  `distributor_id` varchar(64)  COMMENT '分销商ID',
  `user_id` varchar(32)  COMMENT '分销商管理员UID',
  `expiring_order_amount` int(11)  DEFAULT '0' COMMENT '即将过期的订单数量',
  `unfilled_order_amount` int(11)  DEFAULT '0' COMMENT '未发货的定单数量',
  `statistic_day` date  COMMENT '统计日期',
  `create_time` datetime  COMMENT '创建时间',
  `send_status` int(11)  DEFAULT '0' COMMENT '发送状态. 0: 未发送 1: 已发送',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`distributor_id`,`statistic_day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分销商每天的即将过期的订单数量'