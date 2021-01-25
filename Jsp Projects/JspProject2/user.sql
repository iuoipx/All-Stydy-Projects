/*
Navicat MySQL Data Transfer

Source Server         : Test
Source Server Version : 50544
Source Host           : localhost:3306
Source Database       : gxsb

Target Server Type    : MYSQL
Target Server Version : 50544
File Encoding         : 65001

Date: 2020-12-11 14:30:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(20) NOT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `sex` varchar(4) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('3', 'gxsb', 'gxsb', 'gxsb', '20', '11123', 'gxsb');
