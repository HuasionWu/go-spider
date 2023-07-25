# go-spider

用go写的简易爬虫  
用于爬取豆瓣排名前250的电影并入库  
1.启动脚本前需先配置好数据库信息  
2.到数据库里执行sql文件
```
CREATE TABLE `sp_douban_movie` (
                                   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                   `title` varchar(30) DEFAULT '' COMMENT '标题',
                                   `subtitle` varchar(20) DEFAULT '' COMMENT '副标题',
                                   `other` varchar(20) DEFAULT '' COMMENT '其他',
                                   `desc` varchar(30) DEFAULT '' COMMENT '简述',
                                   `year` int(10) unsigned DEFAULT '0' COMMENT '年份',
                                   `area` varchar(20) DEFAULT '' COMMENT '地区',
                                   `tag` varchar(20) DEFAULT '' COMMENT '标签',
                                   `star` int(10) unsigned DEFAULT '0' COMMENT 'star',
                                   `comment` int(10) unsigned DEFAULT '0' COMMENT '评分',
                                   `quote` varchar(30) DEFAULT '' COMMENT '引用',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='豆瓣电影Top250';
```

3.最后安装脚本所需包直接运行即可
```Go
go get //安装程序所需包
go run main.go //直接运程序即可
```

