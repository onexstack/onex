-- MariaDB dump 10.19-11.2.2-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: 10.37.91.93    Database: onex
-- ------------------------------------------------------
-- Server version	10.11.6-MariaDB-0+deb12u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `onex`
--

/*!40000 DROP DATABASE IF EXISTS `onex`*/;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `onex` /*!40100 DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci */;

USE `onex`;

--
-- Table structure for table `fakeserver_order`
--

DROP TABLE IF EXISTS `fakeserver_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fakeserver_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `orderID` longtext DEFAULT NULL,
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `customer` longtext DEFAULT NULL,
  `product` longtext DEFAULT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx.fakeserver_order.userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway_chain`
--

DROP TABLE IF EXISTS `gateway_chain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gateway_chain` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链名',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链展示名',
  `minerType` varchar(16) NOT NULL DEFAULT '' COMMENT '区块链矿机机型',
  `image` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链镜像 ID',
  `minMineIntervalSeconds` int(8) NOT NULL DEFAULT 0 COMMENT '矿机挖矿间隔',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_chain.namespace_name` (`namespace`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='区块链表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway_miner`
--

DROP TABLE IF EXISTS `gateway_miner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gateway_miner` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机名',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机展示名',
  `phase` varchar(45) NOT NULL DEFAULT '' COMMENT '矿机状态',
  `minerType` varchar(16) NOT NULL DEFAULT '' COMMENT '矿机机型',
  `chainName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机所属的区块链名',
  `cpu` int(8) NOT NULL DEFAULT 0 COMMENT '矿机 CPU 规格',
  `memory` int(8) NOT NULL DEFAULT 0 COMMENT '矿机内存规格',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_miner.namespace_name` (`namespace`,`name`),
  KEY `idx.gateway_miner.chainName` (`chainName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='矿机表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway_minerset`
--

DROP TABLE IF EXISTS `gateway_minerset`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gateway_minerset` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机池名',
  `replicas` int(8) NOT NULL DEFAULT 0 COMMENT '矿机副本数',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机池展示名',
  `deletePolicy` varchar(32) NOT NULL DEFAULT '' COMMENT '矿机池缩容策略',
  `minReadySeconds` int(8) NOT NULL DEFAULT 0 COMMENT '矿机 Ready 最小等待时间',
  `fullyLabeledReplicas` int(8) NOT NULL DEFAULT 0 COMMENT '所有标签匹配的副本数',
  `readyReplicas` int(8) NOT NULL DEFAULT 0 COMMENT 'Ready 副本数',
  `availableReplicas` int(8) NOT NULL DEFAULT 0 COMMENT '可用副本数',
  `failureReason` longtext DEFAULT NULL COMMENT '失败原因',
  `failureMessage` longtext DEFAULT NULL COMMENT '失败信息',
  `conditions` longtext DEFAULT NULL COMMENT '矿机池状态',
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_minerset.namespace_name` (`namespace`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='矿机池表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `locks`
--

DROP TABLE IF EXISTS `locks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `locks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) DEFAULT NULL,
  `ownerID` longtext DEFAULT NULL,
  `expiredAt` datetime(3) DEFAULT NULL,
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.locks.name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='分布式锁表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ratelimit_leader_election`
--

DROP TABLE IF EXISTS `ratelimit_leader_election`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ratelimit_leader_election` (
  `id` varchar(191) NOT NULL,
  `expiredAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx.ratelimit_leader_election.expiredAt` (`expiredAt`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `usercenter_secret`
--

DROP TABLE IF EXISTS `usercenter_secret`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usercenter_secret` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '密钥名称',
  `secretID` varchar(36) NOT NULL DEFAULT '' COMMENT '密钥 ID',
  `secretKey` varchar(36) NOT NULL DEFAULT '' COMMENT '密钥 Key',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '密钥状态，0-禁用；1-启用',
  `expires` bigint(64) NOT NULL DEFAULT 0 COMMENT '0 永不过期',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥描述',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.usercenter_secret.secretID` (`secretID`),
  KEY `idx.usercenter_secret.userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='密钥表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `usercenter_user`
--

DROP TABLE IF EXISTS `usercenter_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usercenter_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `username` varchar(253) NOT NULL DEFAULT '' COMMENT '用户名称',
  `status` varchar(64) NOT NULL DEFAULT '' COMMENT '用户状态：registered,active,disabled,blacklisted,locked,deleted',
  `nickname` varchar(253) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '用户加密后的密码',
  `email` varchar(253) NOT NULL DEFAULT '' COMMENT '用户电子邮箱',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.usercenter_user.username` (`username`),
  UNIQUE KEY `uniq.usercenter_user.userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Current Database: `onex`
--

/*!40000 DROP DATABASE IF EXISTS `onex`*/;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `onex` /*!40100 DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci */;

USE `onex`;

--
-- Table structure for table `fakeserver_order`
--

DROP TABLE IF EXISTS `fakeserver_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fakeserver_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `orderID` longtext DEFAULT NULL,
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `customer` longtext DEFAULT NULL,
  `product` longtext DEFAULT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx.fakeserver_order.userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway_chain`
--

DROP TABLE IF EXISTS `gateway_chain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gateway_chain` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链名',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链展示名',
  `minerType` varchar(16) NOT NULL DEFAULT '' COMMENT '区块链矿机机型',
  `image` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链镜像 ID',
  `minMineIntervalSeconds` int(8) NOT NULL DEFAULT 0 COMMENT '矿机挖矿间隔',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_chain.namespace_name` (`namespace`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='区块链表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway_miner`
--

DROP TABLE IF EXISTS `gateway_miner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gateway_miner` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机名',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机展示名',
  `phase` varchar(45) NOT NULL DEFAULT '' COMMENT '矿机状态',
  `minerType` varchar(16) NOT NULL DEFAULT '' COMMENT '矿机机型',
  `chainName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机所属的区块链名',
  `cpu` int(8) NOT NULL DEFAULT 0 COMMENT '矿机 CPU 规格',
  `memory` int(8) NOT NULL DEFAULT 0 COMMENT '矿机内存规格',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_miner.namespace_name` (`namespace`,`name`),
  KEY `idx.gateway_miner.chainName` (`chainName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='矿机表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway_minerset`
--

DROP TABLE IF EXISTS `gateway_minerset`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gateway_minerset` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机池名',
  `replicas` int(8) NOT NULL DEFAULT 0 COMMENT '矿机副本数',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机池展示名',
  `deletePolicy` varchar(32) NOT NULL DEFAULT '' COMMENT '矿机池缩容策略',
  `minReadySeconds` int(8) NOT NULL DEFAULT 0 COMMENT '矿机 Ready 最小等待时间',
  `fullyLabeledReplicas` int(8) NOT NULL DEFAULT 0 COMMENT '所有标签匹配的副本数',
  `readyReplicas` int(8) NOT NULL DEFAULT 0 COMMENT 'Ready 副本数',
  `availableReplicas` int(8) NOT NULL DEFAULT 0 COMMENT '可用副本数',
  `failureReason` longtext DEFAULT NULL COMMENT '失败原因',
  `failureMessage` longtext DEFAULT NULL COMMENT '失败信息',
  `conditions` longtext DEFAULT NULL COMMENT '矿机池状态',
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_minerset.namespace_name` (`namespace`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='矿机池表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `locks`
--

DROP TABLE IF EXISTS `locks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `locks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) DEFAULT NULL,
  `ownerID` longtext DEFAULT NULL,
  `expiredAt` datetime(3) DEFAULT NULL,
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.locks.name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='分布式锁表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ratelimit_leader_election`
--

DROP TABLE IF EXISTS `ratelimit_leader_election`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ratelimit_leader_election` (
  `id` varchar(191) NOT NULL,
  `expiredAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx.ratelimit_leader_election.expiredAt` (`expiredAt`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `usercenter_secret`
--

DROP TABLE IF EXISTS `usercenter_secret`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usercenter_secret` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '密钥名称',
  `secretID` varchar(36) NOT NULL DEFAULT '' COMMENT '密钥 ID',
  `secretKey` varchar(36) NOT NULL DEFAULT '' COMMENT '密钥 Key',
  `status` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '密钥状态，0-禁用；1-启用',
  `expires` bigint(64) NOT NULL DEFAULT 0 COMMENT '0 永不过期',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥描述',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.usercenter_secret.secretID` (`secretID`),
  KEY `idx.usercenter_secret.userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='密钥表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `usercenter_user`
--

DROP TABLE IF EXISTS `usercenter_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usercenter_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `username` varchar(253) NOT NULL DEFAULT '' COMMENT '用户名称',
  `status` varchar(64) NOT NULL DEFAULT '' COMMENT '用户状态：registered,active,disabled,blacklisted,locked,deleted',
  `nickname` varchar(253) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '用户加密后的密码',
  `email` varchar(253) NOT NULL DEFAULT '' COMMENT '用户电子邮箱',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `createdAt` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.usercenter_user.username` (`username`),
  UNIQUE KEY `uniq.usercenter_user.userID` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-27  8:26:38
