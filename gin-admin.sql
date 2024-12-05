-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.44 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.7.0.6856
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 gin-admin 的数据库结构
CREATE DATABASE IF NOT EXISTS `gin-admin` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `gin-admin`;

-- 导出  表 gin-admin.g_file 结构
CREATE TABLE IF NOT EXISTS `g_file` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `name` varchar(256) DEFAULT NULL,
  `url` varchar(256) DEFAULT NULL,
  `tag` varchar(256) DEFAULT NULL,
  `key` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  gin-admin.g_file 的数据：~0 rows (大约)

-- 导出  表 gin-admin.g_notice 结构
CREATE TABLE IF NOT EXISTS `g_notice` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '''创建时间''',
  `start_time` datetime DEFAULT NULL COMMENT '''开始时间''',
  `end_time` datetime DEFAULT NULL COMMENT '''结束时间''',
  `title` varchar(256) DEFAULT NULL COMMENT '''标题''',
  `content` varchar(256) DEFAULT NULL COMMENT '''内容''',
  `operator` varchar(256) DEFAULT NULL COMMENT '''操作者''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 正在导出表  gin-admin.g_notice 的数据：~0 rows (大约)

-- 导出  表 gin-admin.g_role 结构
CREATE TABLE IF NOT EXISTS `g_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `user_name` varchar(256) DEFAULT NULL,
  `value` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  gin-admin.g_role 的数据：~3 rows (大约)
INSERT INTO `g_role` (`id`, `user_id`, `user_name`, `value`) VALUES
	(1, 1, 'admin', 'admin'),
	(2, 1, 'admin', 'test'),
	(3, 2, 'test', 'test');

-- 导出  表 gin-admin.g_user 结构
CREATE TABLE IF NOT EXISTS `g_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `username` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG',
  `user_type` bigint(20) NOT NULL DEFAULT '0',
  `state` bigint(20) NOT NULL DEFAULT '1',
  `created_by` varchar(256) DEFAULT NULL,
  `modified_by` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  gin-admin.g_user 的数据：~2 rows (大约)
INSERT INTO `g_user` (`id`, `created_at`, `updated_at`, `username`, `password`, `avatar`, `user_type`, `state`, `created_by`, `modified_by`) VALUES
	(1, '2022-09-15 12:37:05', '2022-09-15 12:37:05', 'admin', '$2a$10$Bw7h/1hxOY5BJaG1nlpDLeRIDJoArmJZPjUsV3Jj.HCuK0Fgtckzi', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG', 1, 1, '', ''),
	(2, '2022-09-15 12:37:06', '2022-09-15 12:37:06', 'test', '$2a$10$Bw7h/1hxOY5BJaG1nlpDLeRIDJoArmJZPjUsV3Jj.HCuK0Fgtckzi', 'https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG', 2, 1, '', '');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
