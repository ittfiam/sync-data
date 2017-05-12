drop table if exists tbl_schedule_task;
CREATE TABLE `tbl_schedule_task` (
  `schedule_key` varchar(64)  COMMENT 'key',
  `create_time` varchar(32)  COMMENT 'æ—¶é—´',
  `db_record_update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '鏇存柊鏃堕棿',
  PRIMARY KEY (`schedule_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8