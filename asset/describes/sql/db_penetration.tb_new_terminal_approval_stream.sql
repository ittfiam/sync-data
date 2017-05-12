drop table if exists tb_new_terminal_approval_stream;
CREATE TABLE `tb_new_terminal_approval_stream` (
  `event_uuid` varchar(32)  COMMENT 'äº‹ä»¶ID',
  `check_id` varchar(32)  COMMENT 'å®¡æ ¸ID',
  `terminal_id` varchar(32)  COMMENT 'ç»ˆç«¯id',
  `auditor_uid` varchar(32)  COMMENT 'å®¡æ ¸è€…UID',
  `creator_uid` varchar(32)  COMMENT 'åˆ›å»ºè€…UID',
  `approve_result` int(11)  DEFAULT '0' COMMENT 'å®¡æ ¸ç»“æžœï¼š(0-é€šè¿‡ï¼›1-æ‹’ç»)',
  `terminal_create_time` datetime  COMMENT 'ç»ˆç«¯åˆ›å»ºæ—¶é—´',
  `terminal_approve_time` datetime  COMMENT 'ç»ˆç«¯å®¡æ ¸æ—¶é—´',
  `create_time` datetime  COMMENT 'åˆ›å»ºæ—¶é—´',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'æ›´æ–°æ—¶é—´',
  PRIMARY KEY (`event_uuid`),
  KEY `index_check_id` (`check_id`),
  KEY `index_terminal_id` (`terminal_id`),
  KEY `index_auditor_uid` (`auditor_uid`),
  KEY `index_creator_uid` (`creator_uid`),
  KEY `index_terminal_create_time` (`terminal_create_time`),
  KEY `index_terminal_approve_time` (`terminal_approve_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='æ–°å»ºç»ˆç«¯å®¡æ ¸è®°å½•è¡¨'