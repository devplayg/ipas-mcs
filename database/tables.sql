-- MySQL dump 10.16  Distrib 10.1.24-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: ipasm
-- ------------------------------------------------------
-- Server version	10.1.24-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `ipasm`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `ipasm` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `ipasm`;

--
-- Table structure for table `adt_audit`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `adt_audit` (
  `audit_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int(11) NOT NULL,
  `category` varchar(32) NOT NULL,
  `ip` int(10) unsigned NOT NULL,
  `message` varchar(256) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`audit_id`),
  KEY `ix_member_id` (`member_id`),
  KEY `ix_created` (`created`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `adt_audit_detail`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `adt_audit_detail` (
  `audit_id` int(10) unsigned NOT NULL,
  `detail` mediumtext NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `ix_audit_id` (`audit_id`),
  KEY `ix_created` (`created`),
  CONSTRAINT `fk_adt_audit_detail_audit_id` FOREIGN KEY (`audit_id`) REFERENCES `adt_audit` (`audit_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ast_asset`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ast_asset` (
  `asset_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `class` int(10) unsigned NOT NULL COMMENT '1:Server, 2: Sensor, 3: Agent',
  `parent_id` int(10) unsigned NOT NULL,
  `name` varchar(128) NOT NULL,
  `type1` int(10) unsigned NOT NULL,
  `type2` int(10) unsigned NOT NULL,
  `seq` int(10) unsigned NOT NULL DEFAULT '0',
  `hostname` varchar(256) NOT NULL DEFAULT '',
  `cidr` tinyint(3) unsigned NOT NULL DEFAULT '32',
  `guid` varchar(64) NOT NULL DEFAULT '',
  `mac` varchar(32) NOT NULL DEFAULT '',
  `port` smallint(5) unsigned NOT NULL DEFAULT '0',
  `port_sub1` smallint(5) unsigned NOT NULL DEFAULT '0',
  `port_sub2` smallint(5) unsigned NOT NULL DEFAULT '0',
  `version` varchar(64) NOT NULL DEFAULT '',
  `username` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(128) NOT NULL DEFAULT '',
  `usage_cpu` float(3,1) NOT NULL DEFAULT '0.0',
  `usage_mem` float(3,1) NOT NULL DEFAULT '0.0',
  `usage_disk1` float(3,1) NOT NULL DEFAULT '0.0',
  `usage_disk2` float(3,1) NOT NULL DEFAULT '0.0',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `n1` int(11) NOT NULL DEFAULT '0',
  `n2` int(11) NOT NULL DEFAULT '0',
  `n3` int(11) NOT NULL DEFAULT '0',
  `s1` varchar(256) NOT NULL DEFAULT '',
  `s2` varchar(256) NOT NULL DEFAULT '',
  `s3` varchar(256) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`asset_id`),
  KEY `ix_parent_id` (`parent_id`),
  KEY `ix_class` (`class`),
  KEY `ix_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ast_ipcard`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ast_namecard` (
  `ip` int(10) unsigned NOT NULL,
  `network_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(128) DEFAULT NULL,
  `email` varchar(254) NOT NULL DEFAULT '',
  `nickname` varchar(128) DEFAULT NULL,
  `dept1` varchar(128) DEFAULT NULL,
  `dept2` varchar(128) DEFAULT NULL,
  `position` varchar(64) DEFAULT NULL,
  `phone1` varchar(64) DEFAULT NULL,
  `phone2` varchar(64) DEFAULT NULL,
  `address1` varchar(64) DEFAULT NULL,
  `address2` varchar(64) DEFAULT NULL,
  `country` varchar(64) DEFAULT NULL,
  `state` varchar(64) DEFAULT NULL,
  `city` varchar(64) DEFAULT NULL,
  `zipcode` varchar(64) DEFAULT NULL,
  `memo` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`ip`),
  KEY `ix_network_id` (`network_id`),
  KEY `ix_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


--
-- Table structure for table `mbr_config`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mbr_config` (
  `member_id` int(10) unsigned NOT NULL,
  `keyword` varchar(64) NOT NULL,
  `value_s` varchar(128) NOT NULL,
  `value_n` int(11) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `ix_member_id` (`member_id`),
  KEY `ix_keyword` (`keyword`),
  CONSTRAINT `fk_mbr_config_member_id` FOREIGN KEY (`member_id`) REFERENCES `mbr_member` (`member_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mbr_member`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mbr_member` (
  `member_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `org_id` varchar(16) NOT NULL DEFAULT '',
  `username` varchar(32) NOT NULL,
  `email` varchar(256) NOT NULL,
  `position` int(11) unsigned NOT NULL,
  `name` varchar(64) NOT NULL,
  `birth` date NOT NULL DEFAULT '1970-01-01',
  `nickname` varchar(64) NOT NULL DEFAULT '',
  `zipcode` varchar(16) NOT NULL DEFAULT '',
  `country` varchar(64) NOT NULL DEFAULT '',
  `state` varchar(64) NOT NULL DEFAULT '',
  `city` varchar(64) NOT NULL DEFAULT '',
  `address1` varchar(128) NOT NULL DEFAULT '',
  `address2` varchar(128) NOT NULL DEFAULT '',
  `phone1` varchar(64) NOT NULL DEFAULT '',
  `phone2` varchar(64) NOT NULL DEFAULT '',
  `login_count` int(11) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(3) NOT NULL DEFAULT '0',
  `timezone` varchar(64) NOT NULL DEFAULT '',
  `failed_login_count` int(11) unsigned NOT NULL DEFAULT '0',
  `last_sucess_login` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `last_failed_login` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `last_read_message` int(11) unsigned NOT NULL DEFAULT '0',
  `session_id` varchar(64) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`member_id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `mbr_password`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mbr_password` (
  `member_id` int(11) unsigned NOT NULL,
  `password` varchar(64) NOT NULL,
  `salt` varchar(32) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`member_id`),
  CONSTRAINT `fk_mbr_password_member_id` FOREIGN KEY (`member_id`) REFERENCES `mbr_member` (`member_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sys_config`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_config` (
  `section` varchar(64) NOT NULL,
  `keyword` varchar(64) NOT NULL,
  `value_s` varchar(256) NOT NULL,
  `value_n` int(11) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`section`,`keyword`),
  KEY `ix_section` (`section`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;