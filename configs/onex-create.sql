-- Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file. The original repo for
-- this file is https://github.com/onexstack/onex.
--

-- usercenter_user

CREATE TABLE `usercenter_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `username` varchar(253) NOT NULL DEFAULT '' COMMENT '用户名称',
  `status` varchar(64) NOT NULL DEFAULT '' COMMENT '用户状态：registered,active,disabled,blacklisted,locked,deleted',
  `nickname` varchar(253) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '用户加密后的密码',
  `email` varchar(253) NOT NULL DEFAULT '' COMMENT '用户电子邮箱',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '用户手机号',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.usercenter_user.username` (`username`),
  UNIQUE KEY `uniq.usercenter_user.userID` (`userID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- usercenter_secret

CREATE TABLE `usercenter_secret` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `userID` varchar(253) NOT NULL DEFAULT '' COMMENT '用户 ID',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '密钥名称',
  `secretID` varchar(36) NOT NULL DEFAULT '' COMMENT '密钥 ID',
  `secretKey` varchar(36) NOT NULL DEFAULT '' COMMENT '密钥 Key',
  `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '密钥状态，0-禁用；1-启用',
  `expires` bigint(64) NOT NULL DEFAULT 0 COMMENT '0 永不过期',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥描述',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.usercenter_secret.secretID` (`secretID`),
  KEY `idx.usercenter_secret.userID` (`userID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='密钥表';

-- gateway_chain

CREATE TABLE `gateway_chain` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链名',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链展示名',
  `minerType` varchar(16) NOT NULL DEFAULT '' COMMENT '区块链矿机机型',
  `image` varchar(253) NOT NULL DEFAULT '' COMMENT '区块链镜像 ID',
  `minMineIntervalSeconds` int(8) NOT NULL DEFAULT 0 COMMENT '矿机挖矿间隔',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_chain.namespace_name` (`namespace`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='区块链表';


-- gateway_minerset

CREATE TABLE `gateway_minerset` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
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
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_minerset.namespace_name` (`namespace`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='矿机池表';

-- gateway_miner

CREATE TABLE `gateway_miner` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
  `namespace` varchar(253) NOT NULL DEFAULT '' COMMENT '命名空间',
  `name` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机名',
  `displayName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机展示名',
  `phase` varchar(45) NOT NULL DEFAULT '' COMMENT '矿机状态',
  `minerType` varchar(16) NOT NULL DEFAULT '' COMMENT '矿机机型',
  `chainName` varchar(253) NOT NULL DEFAULT '' COMMENT '矿机所属的区块链名',
  `cpu` int(8) NOT NULL DEFAULT 0 COMMENT '矿机 CPU 规格',
  `memory` int(8) NOT NULL DEFAULT 0 COMMENT '矿机内存规格',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.gateway_miner.namespace_name` (`namespace`,`name`),
  KEY `idx.gateway_miner.chainName` (`chainName`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='矿机表';

-- fs_order

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
  INDEX `idx.fakeserver_order.userID` (`userID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

CREATE TABLE `locks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) DEFAULT NULL,
  `ownerID` longtext DEFAULT NULL,
  `expiredAt` datetime(3) DEFAULT NULL,
  `createdAt` datetime(3) DEFAULT NULL,
  `updatedAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq.locks.name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='分布式锁表';


CREATE TABLE `ratelimit_leader_election` (
  `id` varchar(191) NOT NULL,
  `expiredAt` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx.ratelimit_leader_election.expiredAt` (`expiredAt`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

