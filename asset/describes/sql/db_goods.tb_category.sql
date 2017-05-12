drop table if exists tb_category;
CREATE TABLE `tb_category` (
  `category_id` varchar(32)  COMMENT 'ç±»ç›®ID',
  `category_name` varchar(32)  COMMENT 'ç±»ç›®åç§°',
  `category_level` tinyint(4)  COMMENT 'ç±»ç›®å±‚æ¬¡ï¼Œæœ€å¤š4å±‚',
  `first_category_id` int(10)  COMMENT 'ç±»ç›®å±‚çº§ä¸€IDï¼Œæ²¡æœ‰å¡«0ï¼Œæœ‰åˆ™å¡«å€¼',
  `first_category_name` varchar(32)  COMMENT 'ç±»ç›®å±‚çº§ä¸€åç§°ï¼Œæ²¡æœ‰å¡«ç©ºï¼Œæœ‰åˆ™å¡«å€¼',
  `second_category_id` int(10)  COMMENT 'ç±»ç›®å±‚çº§äºŒIDï¼Œæ²¡æœ‰å¡«0ï¼Œæœ‰åˆ™å¡«å€¼',
  `second_category_name` varchar(32)  COMMENT 'ç±»ç›®å±‚çº§äºŒåç§°ï¼Œæ²¡æœ‰å¡«ç©ºï¼Œæœ‰åˆ™å¡«å€¼',
  `third_category_id` int(10)  COMMENT 'ç±»ç›®å±‚çº§ä¸‰IDï¼Œæ²¡æœ‰å¡«0ï¼Œæœ‰åˆ™å¡«å€¼',
  `third_category_name` varchar(32)  COMMENT 'ç±»ç›®å±‚çº§ä¸‰åç§°ï¼Œæ²¡æœ‰å¡«ç©ºï¼Œæœ‰åˆ™å¡«å€¼',
  `fourth_category_id` int(10)  COMMENT 'ç±»ç›®å±‚çº§å››IDï¼Œæ²¡æœ‰å¡«0ï¼Œæœ‰åˆ™å¡«å€¼',
  `fourth_category_name` varchar(32)  COMMENT 'ç±»ç›®å±‚çº§å››åç§°ï¼Œæ²¡æœ‰å¡«ç©ºï¼Œæœ‰åˆ™å¡«å€¼',
  `sub_category_ids` varchar(512)  COMMENT 'ç±»ç›®ä¸‹é¢çš„å­ç±»ç›®é›†åˆ',
  `create_person` varchar(32)  COMMENT 'åˆ›å»ºäºº',
  `create_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'åˆ›å»ºæ—¶é—´',
  `update_person` varchar(32) DEFAULT NULL COMMENT 'æ›´æ–°äºº',
  `update_time` timestamp  DEFAULT '0000-00-00 00:00:00' COMMENT 'æ›´æ–°æ—¶é—´',
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品类目表'