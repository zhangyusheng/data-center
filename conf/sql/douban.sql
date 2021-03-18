CREATE TABLE `center_movie` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`title` varchar(100) DEFAULT '' COMMENT '标题',
`subtitle` varchar(200) DEFAULT '' COMMENT '副标题',
`other` varchar(200) DEFAULT '' COMMENT '其他',
`desc` varchar(2000) DEFAULT '' COMMENT '简述',
`year` int(10) unsigned DEFAULT '0' COMMENT '年份',
`area` varchar(50) DEFAULT '' COMMENT '地区',
`tag` varchar(50) DEFAULT '' COMMENT '标签',
`star` float(10) unsigned DEFAULT '0' COMMENT 'star',
`comment` int(10) unsigned DEFAULT '0' COMMENT '评分',
`quote` varchar(300) DEFAULT '' COMMENT '引用',
`created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
`deleted_on` int(10) unsigned DEFAULT '0',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='豆瓣电影Top250';