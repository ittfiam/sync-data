drop table if exists tb_order_stream;
CREATE TABLE `tb_order_stream` (
  `event_uuid` varchar(32)  COMMENT 'äº‹ä»¶ID',
  `year` int(11)  DEFAULT '0' COMMENT 'å¹´ä»½',
  `month` int(11)  DEFAULT '0' COMMENT 'æœˆä»½',
  `week` int(11)  DEFAULT '0' COMMENT 'ä¸€å¹´ä¸­çš„ç¬¬å‡ å‘¨',
  `terminal_id` varchar(32)  COMMENT 'ç»ˆç«¯id',
  `place_uid` varchar(32)  COMMENT 'ä¸‹å•äººuid',
  `order_id` varchar(32)  COMMENT 'è®¢å•id',
  `amount` decimal(16,2)  COMMENT 'è®¢å•é¢',
  `status` int(11)  DEFAULT '0' COMMENT 'è®¢å•çŠ¶æ€ï¼šï¼ˆ0ï¼šæ–°è®¢å• 1ï¼šè®¢å•å®Œæˆ 2ï¼šè®¢å•å–æ¶ˆï¼‰',
  `order_content` text  COMMENT 'è®¢å•å†…å®¹',
  `place_time` datetime  COMMENT 'ä¸‹å•æ—¶é—´',
  `create_time` datetime  COMMENT 'åˆ›å»ºæ—¶é—´',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'æ›´æ–°æ—¶é—´',
  PRIMARY KEY (`event_uuid`),
  KEY `index_year` (`year`),
  KEY `index_month` (`month`),
  KEY `index_week` (`week`),
  KEY `index_uid` (`place_uid`),
  KEY `index_order_id` (`order_id`),
  KEY `index_place_time` (`place_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='è®¢å•æ¶ˆæ¯è®°å½•è¡¨'