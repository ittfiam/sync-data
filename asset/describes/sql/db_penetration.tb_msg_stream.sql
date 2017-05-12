drop table if exists tb_msg_stream;
CREATE TABLE `tb_msg_stream` (
  `event_uuid` varchar(32)  COMMENT 'äº‹ä»¶ID',
  `uid` varchar(32)  COMMENT 'è§¦å‘äº‹ä»¶è€…uid',
  `chain_id` varchar(32)  COMMENT 'chain id',
  `msg` text  COMMENT 'äº‹ä»¶æ¶ˆæ¯',
  `create_time` datetime  COMMENT 'åˆ›å»ºæ—¶é—´',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'æ›´æ–°æ—¶é—´',
  PRIMARY KEY (`event_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='æ¶ˆæ¯æµæ°´è¡¨'