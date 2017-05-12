drop table if exists tb_occupied;
CREATE TABLE `tb_occupied` (
  `occupied_id` varchar(32)  COMMENT '试点ID',
  `occupied_range` double  COMMENT '试点范围',
  `status` int(11)  COMMENT '状态',
  `longitude` double  DEFAULT '0' COMMENT '经度',
  `latitude` double  DEFAULT '0' COMMENT '纬度',
  `description` varchar(1024)  DEFAULT '0' COMMENT 'geo信息',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`occupied_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='试点范围设置表'