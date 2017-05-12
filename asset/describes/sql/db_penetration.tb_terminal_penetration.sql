drop table if exists tb_terminal_penetration;
CREATE TABLE `tb_terminal_penetration` (
  `terminal_id` varchar(32)  COMMENT 'ç»ˆç«¯id',
  `goods_id` varchar(32)  COMMENT 'å•†å“id',
  `penetration_uid` varchar(32)  COMMENT 'é“ºå¸‚è€…uid',
  `terminal_name` varchar(256)  DEFAULT '' COMMENT '终端名称',
  `penetration_name` varchar(128)  DEFAULT '' COMMENT '铺市者用户名',
  `distributor_id` varchar(32)  DEFAULT '' COMMENT '分销商ID',
  `distributor_name` varchar(128)  DEFAULT '' COMMENT '分销商名称',
  `business_id` varchar(64)  DEFAULT '' COMMENT '资金交易业务ID',
  `order_id` varchar(32)  COMMENT 'è¾¾æˆé“ºå¸‚çš„è®¢å•ID',
  `send_time` datetime ,
  `place_time` datetime ,
  `goods_name` varchar(32)  COMMENT 'å•†å“åç§°',
  `money` decimal(16,2) ,
  `fee_distributor` decimal(16,2)  DEFAULT '0.00' COMMENT '平台支付金额',
  `amount_plateform` decimal(16,2)  DEFAULT '0.00' COMMENT '平台支付金额',
  `amount_distributor` decimal(16,2)  DEFAULT '0.00' COMMENT '分销商支付金额',
  `year` int(11)  DEFAULT '0' COMMENT 'å¹´ä»½',
  `month` int(11)  DEFAULT '0' COMMENT 'æœˆä»½',
  `week` int(11)  DEFAULT '0' COMMENT 'ä¸€å¹´ä¸­çš„ç¬¬å‡ å‘¨',
  `penetration_status` int(11)  DEFAULT '0' COMMENT 'é“ºå¸‚çŠ¶æ€(0:æœªé“ºå¸‚  1:å·²é“ºå¸‚)',
  `business_status` int(11)  DEFAULT '0' COMMENT '到帐状态(0:未到帐  1:已到帐)',
  `business_success_time` datetime  COMMENT '帐务处理成功时间',
  `data_version` int(11)  DEFAULT '0' COMMENT 'æ•°æ®ç‰ˆæœ¬',
  `create_time` datetime  COMMENT 'åˆ›å»ºæ—¶é—´',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'æ›´æ–°æ—¶é—´',
  `award_time` datetime  DEFAULT '0000-00-00 00:00:00',
  `order_status` int(11) ,
  PRIMARY KEY (`terminal_id`,`goods_id`),
  KEY `index_year` (`year`),
  KEY `index_month` (`month`),
  KEY `index_week` (`week`),
  KEY `index_order_id` (`order_id`),
  KEY `index_penetration_uid` (`penetration_uid`),
  KEY `index_place_time` (`place_time`),
  KEY `index_business` (`business_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='å•†å“åœ¨ç»ˆç«¯çš„é“ºå¸‚çŠ¶æ€è®°å½•è¡¨'