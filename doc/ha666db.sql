/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : test.abc.red:3306
 Source Schema         : ha666db

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 27/08/2019 14:19:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for proj
-- ----------------------------
DROP TABLE IF EXISTS `proj`;
CREATE TABLE `proj` (
  `projId` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目编号',
  `projName` varchar(32) NOT NULL COMMENT '项目名称',
  `userCode` varchar(32) NOT NULL COMMENT '用户编码',
  `deleteStatus` tinyint(4) NOT NULL COMMENT '是否删除',
  `endTime` datetime NOT NULL COMMENT '结束时间',
  PRIMARY KEY (`projId`),
  KEY `idx_proj_projName` (`projName`) USING BTREE COMMENT '根据项目名称查询列表'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='项目表';

-- ----------------------------
-- Records of proj
-- ----------------------------
BEGIN;
INSERT INTO `proj` VALUES (1, '测试项目1', 'ha666', 1, '2019-08-26 14:32:48');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userCode` varchar(32) NOT NULL DEFAULT '' COMMENT '用户编码，取自钉钉编码',
  `realName` varchar(16) NOT NULL DEFAULT '' COMMENT '姓名',
  `jobNumber` varchar(16) NOT NULL DEFAULT '' COMMENT '员工的工号',
  `jobPosition` varchar(16) NOT NULL DEFAULT '' COMMENT '职位信息',
  `hiredDate` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '入职时间',
  `avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '头像url',
  `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别，0未知，1男，2女',
  `userType` tinyint(4) NOT NULL DEFAULT '0' COMMENT '用户类型',
  `deleteStatus` tinyint(4) NOT NULL DEFAULT '0' COMMENT '删除状态，0未知，1未删除，2删除',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`userCode`) USING BTREE,
  UNIQUE KEY `uk_user_userCode` (`userCode`) USING BTREE COMMENT '根据用户编码查询唯一用户信息',
  KEY `idx_user_realName` (`realName`) USING BTREE COMMENT '根据姓名查询多条记录'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('ha666', 'ha666HA666', 'T-000000', '暂无', '2019-04-30 12:37:23', '', 0, 1, 1, '2019-04-30 12:37:23', '2019-04-30 12:37:23');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
