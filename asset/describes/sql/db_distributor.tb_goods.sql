drop table if exists tb_goods;
CREATE TABLE `tb_goods` (
  `distributor_id` varchar(64)  COMMENT '分销商ID',
  `goods_id` varchar(64)  COMMENT '商品ID',
  `goods_price` double  DEFAULT '0' COMMENT '商品价格',
  `commission` double  DEFAULT '0' COMMENT '业务员可得提成',
  `sales_commission` double  DEFAULT '0' COMMENT '业务员可得提成',
  `penetration_amount` double  DEFAULT '0' COMMENT '商品铺市',
  `sales_penetration` double  DEFAULT '0' COMMENT '商品铺市',
  `goods_gift` text COMMENT '搭赠策略',
  `sellout_status` int(10)  DEFAULT '0' COMMENT '售罄状态,0:有货, 1:售罄',
  `create_time` datetime  COMMENT '创建时间',
  `update_time` timestamp  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `distributor_pay_commission` double  DEFAULT '0' COMMENT '配送点支付提成',
  `distributor_commission` double  DEFAULT '0' COMMENT '配送点提成费用',
  `platform_subsidy_commission` double  DEFAULT '0' COMMENT '平台提成补贴',
  `platform_commission_profit` double  DEFAULT '0' COMMENT '平台提成利润',
  `distributor_pay_penetration` double  DEFAULT '0' COMMENT '配送点支付铺市',
  `distributor_penetration` double  DEFAULT '0' COMMENT '配送点铺市费用',
  `platform_subsidy_penetration` double  DEFAULT '0' COMMENT '平台铺市补贴',
  `platform_penetration_profit` double  DEFAULT '0' COMMENT '平台铺市利润',
  PRIMARY KEY (`distributor_id`,`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分销商商品信息表'