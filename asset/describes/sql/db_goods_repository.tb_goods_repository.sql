drop table if exists tb_goods_repository;
CREATE TABLE `tb_goods_repository` (
  `goods_code` varchar(64)  DEFAULT '' COMMENT '商品条码',
  `goods_name` varchar(255)  DEFAULT '' COMMENT '商品名称',
  `manu_name` varchar(255)  DEFAULT '' COMMENT '厂商名称',
  `goods_spec` varchar(32)  DEFAULT '' COMMENT '商品规格',
  `goods_price` double(32,2)  DEFAULT '0.00' COMMENT '商品价格',
  `goods_brand` varchar(255)  DEFAULT '' COMMENT '商品品牌',
  `goods_img` varchar(255)  DEFAULT '' COMMENT '商品图片',
  `goods_type` varchar(255)  DEFAULT '' COMMENT '商品分类',
  `goods_code_img` varchar(255)  DEFAULT '' COMMENT '商品条码图片',
  `remark` varchar(255) DEFAULT '' COMMENT '备注信息',
  `create_time` datetime  COMMENT '创建记录时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新记录时间',
  PRIMARY KEY (`goods_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品库信息表'