drop table if exists tb_send_award_stream;
CREATE TABLE `tb_send_award_stream` (
  `event_uuid` varchar(32)  COMMENT '发送奖励消息ID',
  `chain_id` varchar(32)  COMMENT 'chain id',
  `order_event_uuid` varchar(32)  COMMENT '订单消息ID',
  `distributor_id` varchar(32)  DEFAULT '' COMMENT '分销商ID',
  `order_id` varchar(32)  DEFAULT '' COMMENT '订单ID',
  `business_id` varchar(64)  DEFAULT '' COMMENT '资金交易业务ID',
  `uid` varchar(32)  COMMENT '获得奖励者uid',
  `terminal_id` varchar(32)  COMMENT '终端id',
  `award_money` double(16,2)  COMMENT '奖励金额',
  `award_msg` text  COMMENT '事件消息',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`event_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='发送奖励消息流水表'