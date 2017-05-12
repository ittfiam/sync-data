drop table if exists tb_order_goods;
CREATE TABLE `tb_order_goods` (
  `terminal_id` varchar(32)  COMMENT 'ç»ˆç«¯id',
  `goods_id` varchar(32)  COMMENT 'å•†å“id',
  `place_uid` varchar(32)  COMMENT 'ä¸‹å•äººuid',
  `order_id` varchar(32)  COMMENT 'è®¢å•id',
  `place_time` datetime  COMMENT 'ä¸‹å•æ—¶é—´',
  `goods_info` text  COMMENT 'å•†å“ä¿¡æ¯',
  `year` int(11)  DEFAULT '0' COMMENT 'å¹´ä»½',
  `month` int(11)  DEFAULT '0' COMMENT 'æœˆä»½',
  `week` int(11)  DEFAULT '0' COMMENT 'ä¸€å¹´ä¸­çš„ç¬¬å‡ å‘¨',
  `penetration_status` int(11)  DEFAULT '0' COMMENT 'é“ºå¸‚çŠ¶æ€(0:æœªé“ºå¸‚  1:å·²é“ºå¸‚)',
  `create_time` datetime  COMMENT 'åˆ›å»ºæ—¶é—´',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'æ›´æ–°æ—¶é—´',
  PRIMARY KEY (`terminal_id`,`goods_id`),
  KEY `index_year` (`year`),
  KEY `index_month` (`month`),
  KEY `index_week` (`week`),
  KEY `index_uid` (`place_uid`),
  KEY `index_order_id` (`order_id`),
  KEY `index_place_time` (`place_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='è®¢å•å•†å“ä¸Žç»ˆç«¯è®°å½•è¡¨'