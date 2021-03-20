/*
SQLyog Community v13.1.6 (64 bit)
MySQL - 8.0.11 : Database - bhu
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`bhu` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;

USE `bhu`;

/*Table structure for table `articleinfo` */

DROP TABLE IF EXISTS `articleinfo`;

CREATE TABLE `articleinfo` (
  `articleid` int(8) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `articleicon` varchar(1024) DEFAULT NULL,
  `articletitle` varchar(100) DEFAULT NULL,
  `articledetail` varchar(10240) DEFAULT NULL,
  PRIMARY KEY (`articleid`),
  KEY `userid` (`userid`),
  CONSTRAINT `articleinfo_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `userinfo` (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `articleinfo` */

insert  into `articleinfo`(`articleid`,`userid`,`articleicon`,`articletitle`,`articledetail`) values 
(7,1,'','212121','21212121'),
(10,2,'','dddd','dsds'),
(11,3,'','ddsdsds','bfbfgfgfg'),
(12,4,'','rgrtrt','ererer'),
(13,5,'','dfdfdfdf','fdfd'),
(14,6,'','343434dfdfd','sssdsdsd');

/*Table structure for table `userinfo` */

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo` (
  `userid` int(8) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) DEFAULT NULL,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  `age` int(8) DEFAULT NULL,
  `icon` varchar(1024) DEFAULT NULL,
  `columnname` varchar(100) DEFAULT NULL,
  `columnintro` varchar(1024) DEFAULT NULL,
  `personalintro` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `userinfo` */

insert  into `userinfo`(`userid`,`email`,`username`,`password`,`age`,`icon`,`columnname`,`columnintro`,`personalintro`) values 
(1,'1323287928@qq.com','夏夏','xzh1323287928',18,'E:\\前端空间\\开课吧\\B乎项目\\userIcon\\userIcon\\img_2.png','夏夏的专栏','时间会改变很多，我在期待着111111','1'),
(2,'1063772281@qq.com','123','1',19,NULL,NULL,'时间会改变很多，我在期待着111111',NULL),
(3,'15854132785@qq.com','xzh1323287928','xzh1323287928',19,NULL,NULL,'时间会改变很多，我在期待着111111',NULL),
(4,'4759173@qq.com','邹邹','123',19,NULL,NULL,'时间会改变很多，我在期待着111111',NULL),
(5,'7788@qq.com','11','11',19,NULL,NULL,'时间会改变很多，我在期待着111111',NULL),
(6,'111','11','11',22,NULL,NULL,'时间会改变很多，我在期待着111111',NULL),
(7,'112121','21212','111',111,NULL,NULL,NULL,NULL),
(8,'21212121','212121','12121',11111,NULL,NULL,NULL,NULL),
(9,'21212','21212','111',111,NULL,NULL,NULL,NULL),
(10,'1111','1111','1111',1111,NULL,NULL,NULL,NULL),
(11,'','','',0,NULL,NULL,NULL,NULL),
(12,'','','',0,NULL,NULL,NULL,NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
