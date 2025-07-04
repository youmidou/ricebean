/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80042 (8.0.42)
 Source Host           : localhost:3307
 Source Schema         : ricebean_db

 Target Server Type    : MySQL
 Target Server Version : 80042 (8.0.42)
 File Encoding         : 65001

 Date: 04/07/2025 18:41:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apple_user_info
-- ----------------------------
DROP TABLE IF EXISTS `apple_user_info`;
CREATE TABLE `apple_user_info` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `aud` varchar(128) DEFAULT NULL,
  `iss` varchar(128) DEFAULT NULL,
  `email_verified` varchar(8) DEFAULT NULL,
  `exp` varchar(32) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `sub` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_apple_user_info_sub` (`sub`) USING BTREE,
  KEY `ix_apple_user_info_user_id` (`user_id`) USING BTREE,
  KEY `ix_apple_user_info_email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for b_system_gallery
-- ----------------------------
DROP TABLE IF EXISTS `b_system_gallery`;
CREATE TABLE `b_system_gallery` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `archer_gallery_data` varchar(1024) DEFAULT NULL,
  `bingo_gallery_data` varchar(4096) DEFAULT NULL,
  `coin_rush_gallery_data` varchar(1024) DEFAULT NULL,
  `cooking_gallery_data` varchar(1024) DEFAULT NULL,
  `journey_gallery_data` varchar(1024) DEFAULT NULL,
  `normal_quest_gallery_data` varchar(1024) DEFAULT NULL,
  `pick_party_gallery_data` varchar(1024) DEFAULT NULL,
  `rocket_gallery_data` varchar(1024) DEFAULT NULL,
  `reserve_data1` varchar(1024) DEFAULT NULL,
  `reserve_data2` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for behaviour_data
-- ----------------------------
DROP TABLE IF EXISTS `behaviour_data`;
CREATE TABLE `behaviour_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `last_update_date` datetime DEFAULT NULL,
  `country` varchar(16) DEFAULT NULL,
  `city` varchar(32) DEFAULT NULL,
  `city_level` smallint DEFAULT NULL,
  `first_login_date` datetime DEFAULT NULL,
  `last_login_date` datetime DEFAULT NULL,
  `platform_type` smallint DEFAULT NULL,
  `phone_brand` varchar(16) DEFAULT NULL,
  `phone_brand_level` smallint DEFAULT NULL,
  `vip_level` smallint DEFAULT NULL,
  `r_level` smallint DEFAULT NULL,
  `yy1` varchar(32) DEFAULT NULL,
  `yy2` varchar(16) DEFAULT NULL,
  `yy3` varchar(16) DEFAULT NULL,
  `xx1` smallint DEFAULT NULL,
  `xx2` smallint DEFAULT NULL,
  `xx3` smallint DEFAULT NULL,
  `xx4` smallint DEFAULT NULL,
  `xx5` smallint DEFAULT NULL,
  `xx6` int DEFAULT NULL,
  `xx7` int DEFAULT NULL,
  `xx8` int DEFAULT NULL,
  `first_purchase_date` datetime DEFAULT NULL,
  `last_purchase_date` datetime DEFAULT NULL,
  `purchase_count` int DEFAULT NULL,
  `total_purchase` int DEFAULT NULL,
  `max_purchase` smallint DEFAULT NULL,
  `min_purchase` smallint DEFAULT NULL,
  `last_7days_complete_mission_times_data` varchar(64) DEFAULT NULL,
  `last_7days_complete_mission_times` int DEFAULT NULL,
  `last_7days_complete_bc_times_data` varchar(64) DEFAULT NULL,
  `last_7days_complete_bc_times` int DEFAULT NULL,
  `last_7days_click_store_times_data` varchar(64) DEFAULT NULL,
  `last_7days_click_store_times` int DEFAULT NULL,
  `last_30days_login_flag_data` varchar(128) DEFAULT NULL,
  `last_7days_login_days` smallint DEFAULT NULL,
  `last_30days_login_days` smallint DEFAULT NULL,
  `last_30days_online_time_data` varchar(256) DEFAULT NULL,
  `last_7days_online_time` int DEFAULT NULL,
  `last_30days_online_time` int DEFAULT NULL,
  `last_30days_spin_times_data` varchar(256) DEFAULT NULL,
  `last_7days_spin_times` int DEFAULT NULL,
  `last_30days_spin_times` int DEFAULT NULL,
  `last_30days_purchase_times_data` varchar(256) DEFAULT NULL,
  `last_7days_purchase_times` int DEFAULT NULL,
  `last_30days_purchase_times` int DEFAULT NULL,
  `last_30days_purchase_data` varchar(256) DEFAULT NULL,
  `last_7days_purchase` int DEFAULT NULL,
  `last_30days_purchase` int DEFAULT NULL,
  `last_7days_xx1_data` varchar(256) DEFAULT NULL,
  `last_7days_xx1` int DEFAULT NULL,
  `last_7days_xx2_data` varchar(256) DEFAULT NULL,
  `last_7days_xx2` int DEFAULT NULL,
  `last_7days_xx3_data` varchar(256) DEFAULT NULL,
  `last_7days_xx3` int DEFAULT NULL,
  `last_7days_xx4_data` varchar(256) DEFAULT NULL,
  `last_7days_xx4` int DEFAULT NULL,
  `most_purchase_type_id` smallint DEFAULT NULL,
  `second_purchase_type_id` smallint DEFAULT NULL,
  `all_purchase_stats_data` varchar(1024) DEFAULT NULL,
  `last_7days_purchase_stats_data` varchar(1024) DEFAULT NULL,
  `last_7days_most_purchase_type_id` smallint DEFAULT NULL,
  `last_7days_second_purchase_type_id` smallint DEFAULT NULL,
  `store_purchase_count` int DEFAULT NULL,
  `promotion_purchase_count` int DEFAULT NULL,
  `money_bank_purchase_count` int DEFAULT NULL,
  `premium_pachinko_purchase_count` int DEFAULT NULL,
  `deluxe_slot_purchase_count` int DEFAULT NULL,
  `spinup_purchase_count` int DEFAULT NULL,
  `booster_purchase_count` int DEFAULT NULL,
  `special_purchase_count` int DEFAULT NULL,
  `pick_party_purchase_count` int DEFAULT NULL,
  `rocket_purchase_count` int DEFAULT NULL,
  `bingo_purchase_count` int DEFAULT NULL,
  `archer_purchase_count` int DEFAULT NULL,
  `cooking_purchase_count` int DEFAULT NULL,
  `coin_rush_purchase_count` int DEFAULT NULL,
  `journey_purchase_count` int DEFAULT NULL,
  `normal_quest_purchase_count` int DEFAULT NULL,
  `sail_purchase_count` int DEFAULT NULL,
  `xx1_purchase_count` int DEFAULT NULL,
  `xx2_purchase_count` int DEFAULT NULL,
  `xx3_purchase_count` int DEFAULT NULL,
  `xx4_purchase_count` int DEFAULT NULL,
  `xx5_purchase_count` int DEFAULT NULL,
  `xx6_purchase_count` int DEFAULT NULL,
  `xx7_purchase_count` int DEFAULT NULL,
  `xx8_purchase_count` int DEFAULT NULL,
  `xx9_purchase_count` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_behaviour_data_first_login_date` (`first_login_date`) USING BTREE,
  KEY `ix_behaviour_data_last_7days_login_days` (`last_7days_login_days`) USING BTREE,
  KEY `ix_behaviour_data_user_id` (`user_id`) USING BTREE,
  KEY `ix_behaviour_data_purchase_count` (`purchase_count`) USING BTREE,
  KEY `ix_behaviour_data_last_login_date` (`last_login_date`) USING BTREE,
  KEY `ix_behaviour_data_first_purchase_date` (`first_purchase_date`) USING BTREE,
  KEY `ix_behaviour_data_last_30days_login_days` (`last_30days_login_days`) USING BTREE,
  KEY `ix_behaviour_data_last_update_date` (`last_update_date`) USING BTREE,
  KEY `ix_behaviour_data_total_purchase` (`total_purchase`) USING BTREE,
  KEY `ix_behaviour_data_last_purchase_date` (`last_purchase_date`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3132 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for email_unsubscribe
-- ----------------------------
DROP TABLE IF EXISTS `email_unsubscribe`;
CREATE TABLE `email_unsubscribe` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `email_id` varchar(128) NOT NULL,
  `facebook_id` varchar(128) DEFAULT NULL,
  `email_address` varchar(128) DEFAULT NULL,
  `unsubscribe_reason` varchar(128) DEFAULT NULL,
  `else_reason` text,
  `feedback` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_address` (`email_address`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for fever_vault
-- ----------------------------
DROP TABLE IF EXISTS `fever_vault`;
CREATE TABLE `fever_vault` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `chips` int DEFAULT NULL,
  `a_data` varchar(1024) DEFAULT NULL,
  `b_data` varchar(1024) DEFAULT NULL,
  `c_data` varchar(1024) DEFAULT NULL,
  `a_type` varchar(1024) DEFAULT NULL,
  `b_type` varchar(1024) DEFAULT NULL,
  `c_type` varchar(1024) DEFAULT NULL,
  `sapphire_disconnection_data` varchar(2048) DEFAULT NULL,
  `blazing_spin_count` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `ix_fever_vault_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=865 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for flame_challenge
-- ----------------------------
DROP TABLE IF EXISTS `flame_challenge`;
CREATE TABLE `flame_challenge` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `exp` int DEFAULT NULL,
  `level` int DEFAULT NULL,
  `season` int DEFAULT NULL,
  `season_end` bigint DEFAULT NULL,
  `mission_data` varchar(2048) DEFAULT NULL,
  `mission_end` bigint DEFAULT NULL,
  `prizes_data` varchar(20000) DEFAULT NULL,
  `bonus_game_disconnection_data` varchar(2048) DEFAULT NULL,
  `ranking_prize_disconnection_data` varchar(2048) DEFAULT NULL,
  `mini_theme_next_time` bigint DEFAULT NULL,
  `boosters_data` varchar(64) DEFAULT NULL,
  `history_rank_data` varchar(20000) DEFAULT NULL,
  `mission_fresh` int DEFAULT NULL,
  `avg_bet` bigint DEFAULT NULL,
  `commodities_data` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_flame_challenge_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for game_king
-- ----------------------------
DROP TABLE IF EXISTS `game_king`;
CREATE TABLE `game_king` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `free_chest_ts` int DEFAULT NULL,
  `chest_data` varchar(1024) DEFAULT NULL,
  `quest_chest_data` varchar(1024) DEFAULT NULL,
  `items_data` varchar(4096) DEFAULT NULL,
  `game_king_start_ts` int DEFAULT NULL,
  `chest_exp` int DEFAULT NULL,
  `chest_disconnection_data` varchar(4096) DEFAULT NULL,
  `is_collected` int DEFAULT NULL,
  `wild_chest_data` varchar(1024) DEFAULT NULL,
  `game_king_gallery` varchar(4096) DEFAULT NULL,
  `current_season` int DEFAULT NULL,
  `history_rank_date` varchar(1024) DEFAULT NULL,
  `rank_building_date` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_game_king_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for hero_album
-- ----------------------------
DROP TABLE IF EXISTS `hero_album`;
CREATE TABLE `hero_album` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `reset_ts` int DEFAULT NULL,
  `game_idx` int DEFAULT NULL,
  `level` int DEFAULT NULL,
  `hero_album_data` varchar(2048) DEFAULT NULL,
  `hero_album_collected_data` varchar(1024) DEFAULT NULL,
  `game_disconnection_data` varchar(8192) DEFAULT NULL,
  `history_data` varchar(4096) DEFAULT NULL,
  `game_prize_data` varchar(2048) DEFAULT NULL,
  `extra_num_a` int DEFAULT NULL,
  `extra_num_b` int DEFAULT NULL,
  `extra_data_a` varchar(2048) DEFAULT NULL,
  `extra_data_b` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=534 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for inbox
-- ----------------------------
DROP TABLE IF EXISTS `inbox`;
CREATE TABLE `inbox` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` int DEFAULT NULL,
  `receiver` varchar(32) DEFAULT NULL,
  `sender` varchar(32) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `ts` datetime DEFAULT NULL,
  `gift_info` varchar(1024) DEFAULT NULL,
  `icon` varchar(32) DEFAULT NULL,
  `log_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_inbox_receiver_type` (`receiver`,`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for inbox_normal
-- ----------------------------
DROP TABLE IF EXISTS `inbox_normal`;
CREATE TABLE `inbox_normal` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` int DEFAULT NULL,
  `receiver` varchar(32) DEFAULT NULL,
  `sender` varchar(32) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `ts` datetime DEFAULT NULL,
  `gift_info` varchar(1024) DEFAULT NULL,
  `icon` varchar(32) DEFAULT NULL,
  `log_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for inbox_voucher
-- ----------------------------
DROP TABLE IF EXISTS `inbox_voucher`;
CREATE TABLE `inbox_voucher` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` int DEFAULT NULL,
  `receiver` varchar(32) DEFAULT NULL,
  `sender` varchar(32) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `ts` datetime DEFAULT NULL,
  `gift_info` varchar(1024) DEFAULT NULL,
  `icon` varchar(32) DEFAULT NULL,
  `log_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for mission_pass
-- ----------------------------
DROP TABLE IF EXISTS `mission_pass`;
CREATE TABLE `mission_pass` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `exp` int DEFAULT NULL,
  `level` int DEFAULT NULL,
  `season` bigint DEFAULT NULL,
  `season_store` int DEFAULT NULL,
  `total_exp` int DEFAULT NULL,
  `reset_ts` bigint DEFAULT NULL,
  `boosters` varchar(64) DEFAULT NULL,
  `exp_table` varchar(2048) DEFAULT NULL,
  `mission_pass_type` int DEFAULT NULL,
  `free_pass_prize` varchar(20000) DEFAULT NULL,
  `mission_pass_prize` varchar(20000) DEFAULT NULL,
  `available_free_reward` varchar(2048) DEFAULT NULL,
  `available_high_reward` varchar(2048) DEFAULT NULL,
  `chest_status` int DEFAULT NULL,
  `chest_prize` varchar(2048) DEFAULT NULL,
  `chest_level` int DEFAULT NULL,
  `chest_exp` int DEFAULT NULL,
  `chest_avg_bet` int DEFAULT NULL,
  `check_list_flag` int DEFAULT NULL,
  `free_player_flag` int DEFAULT NULL,
  `store_info` varchar(2048) DEFAULT NULL,
  `EXP_TIMES_IN_BOOSTERS` varchar(64) DEFAULT NULL,
  `commodities` varchar(2048) DEFAULT NULL,
  `current_season_rlv` int DEFAULT NULL,
  `change_rlv_flag` int DEFAULT '0',
  `new_season_flag` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `ix_mission_pass_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for mp_limit_prize
-- ----------------------------
DROP TABLE IF EXISTS `mp_limit_prize`;
CREATE TABLE `mp_limit_prize` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `time_limited_free_prize_data` varchar(4096) DEFAULT NULL,
  `time_limited_mission_prize_data` varchar(4096) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_mp_limit_prize_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for new_fb_info
-- ----------------------------
DROP TABLE IF EXISTS `new_fb_info`;
CREATE TABLE `new_fb_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `first_name` varchar(64) DEFAULT NULL,
  `last_name` varchar(64) DEFAULT NULL,
  `facebook_id` varchar(32) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `sex` varchar(16) DEFAULT NULL,
  `age` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_new_fb_info_user_id` (`user_id`) USING BTREE,
  UNIQUE KEY `ix_new_fb_info_email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for new_mission_pass
-- ----------------------------
DROP TABLE IF EXISTS `new_mission_pass`;
CREATE TABLE `new_mission_pass` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `exp` int DEFAULT NULL,
  `level` int DEFAULT NULL,
  `season` bigint DEFAULT NULL,
  `season_store` int DEFAULT NULL,
  `total_exp` int DEFAULT NULL,
  `reset_ts` bigint DEFAULT NULL,
  `boosters` varchar(64) DEFAULT NULL,
  `exp_table` varchar(2048) DEFAULT NULL,
  `mission_pass_type` int DEFAULT NULL,
  `free_pass_status` varchar(256) DEFAULT NULL,
  `mission_pass_status` varchar(256) DEFAULT NULL,
  `time_limited_mission_status_data` varchar(256) DEFAULT NULL,
  `time_limited_free_status_data` varchar(256) DEFAULT NULL,
  `chest_status` int DEFAULT NULL,
  `chest_prize` varchar(2048) DEFAULT NULL,
  `chest_level` int DEFAULT NULL,
  `chest_exp` int DEFAULT NULL,
  `chest_avg_bet` int DEFAULT NULL,
  `check_list_flag` int DEFAULT NULL,
  `free_player_flag` int DEFAULT NULL,
  `EXP_TIMES_IN_BOOSTERS` varchar(64) DEFAULT NULL,
  `commodities` varchar(2048) DEFAULT NULL,
  `current_season_rlv` int DEFAULT NULL,
  `change_rlv_flag` int DEFAULT NULL,
  `new_season_flag` int DEFAULT NULL,
  `migrate_flag` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_new_mission_pass_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for purchase_google_transaction
-- ----------------------------
DROP TABLE IF EXISTS `purchase_google_transaction`;
CREATE TABLE `purchase_google_transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identifier` varchar(128) DEFAULT NULL,
  `payment` int DEFAULT NULL,
  `app_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `payload` text,
  `item_id` int DEFAULT NULL,
  `state` int DEFAULT NULL,
  `receiptData` text,
  `createTime` datetime DEFAULT NULL,
  `verifyTime` datetime DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `price` float DEFAULT NULL,
  `currency` varchar(32) DEFAULT NULL,
  `ts` datetime DEFAULT NULL,
  `count` int DEFAULT NULL,
  `user_level` int DEFAULT NULL,
  `user_coins` float DEFAULT NULL,
  `user_spin_count` bigint DEFAULT NULL,
  `last_theme` int DEFAULT NULL,
  `user_first_login` datetime DEFAULT NULL,
  `user_last_login` datetime DEFAULT NULL,
  `last_bet` float DEFAULT NULL,
  `purchase_count` int DEFAULT NULL,
  `package` varchar(64) DEFAULT NULL,
  `tag` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for purchase_transaction
-- ----------------------------
DROP TABLE IF EXISTS `purchase_transaction`;
CREATE TABLE `purchase_transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identifier` varchar(128) DEFAULT NULL,
  `payment` int DEFAULT NULL,
  `app_id` int DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `payload` text,
  `item_id` int DEFAULT NULL,
  `state` int DEFAULT NULL,
  `receiptData` text,
  `createTime` datetime DEFAULT NULL,
  `verifyTime` datetime DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `price` float DEFAULT NULL,
  `currency` varchar(32) DEFAULT NULL,
  `ts` datetime DEFAULT NULL,
  `count` int DEFAULT NULL,
  `user_level` int DEFAULT NULL,
  `user_coins` float DEFAULT NULL,
  `user_spin_count` bigint DEFAULT NULL,
  `last_theme` int DEFAULT NULL,
  `user_first_login` datetime DEFAULT NULL,
  `user_last_login` datetime DEFAULT NULL,
  `last_bet` float DEFAULT NULL,
  `purchase_count` int DEFAULT NULL,
  `package` varchar(64) DEFAULT NULL,
  `tag` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for send_email_fail
-- ----------------------------
DROP TABLE IF EXISTS `send_email_fail`;
CREATE TABLE `send_email_fail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `email` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `email_id` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC COMMENT='Jf给所有玩家发送邮件表';

-- ----------------------------
-- Table structure for send_mail_list
-- ----------------------------
DROP TABLE IF EXISTS `send_mail_list`;
CREATE TABLE `send_mail_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `facebook_name` varchar(128) DEFAULT NULL,
  `first_name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`) USING BTREE,
  KEY `index_user_id` (`user_id`) USING BTREE,
  KEY `index_email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='Jf给所有玩家发送邮件的表';

-- ----------------------------
-- Table structure for stamp
-- ----------------------------
DROP TABLE IF EXISTS `stamp`;
CREATE TABLE `stamp` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `section` int DEFAULT NULL,
  `stars` int DEFAULT NULL,
  `frenzy_count` int DEFAULT NULL,
  `stars_ts` int DEFAULT NULL,
  `stamp_start_ts` int DEFAULT NULL,
  `albums_data` varchar(4096) DEFAULT '{}',
  `frenzy_data` varchar(1024) DEFAULT NULL,
  `collected_data` varchar(1024) DEFAULT NULL,
  `season_collected_data` varchar(1024) DEFAULT NULL,
  `wild_data` varchar(1024) DEFAULT NULL,
  `history_data` varchar(20000) DEFAULT '[]',
  `games_data` varchar(20000) DEFAULT '{}',
  `disconnection_data` varchar(1024) DEFAULT '{}',
  `starter_kit_data` varchar(512) DEFAULT '{}',
  `blazing_spin_count` int DEFAULT '0',
  `last_blazing_num` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `index_section` (`section`) USING BTREE,
  KEY `index_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7954 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for stamp_play_again
-- ----------------------------
DROP TABLE IF EXISTS `stamp_play_again`;
CREATE TABLE `stamp_play_again` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `section` int DEFAULT NULL,
  `season` int DEFAULT NULL,
  `end_ts` int DEFAULT NULL,
  `collected` int DEFAULT NULL,
  `album_data` varchar(2048) DEFAULT NULL,
  `xx1` int DEFAULT NULL,
  `xx2` int DEFAULT NULL,
  `yy1` varchar(1024) DEFAULT NULL,
  `yy2` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_stamp_play_again_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for theme_item
-- ----------------------------
DROP TABLE IF EXISTS `theme_item`;
CREATE TABLE `theme_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `theme_id` int DEFAULT NULL,
  `free_spins` int DEFAULT NULL,
  `bet` float DEFAULT NULL,
  `new_theme_bank` float DEFAULT NULL,
  `new_theme_bank_count` int DEFAULT NULL,
  `complex_data` text,
  `param_1` bigint DEFAULT NULL,
  `param_8` bigint DEFAULT NULL,
  `rtp_param_2` bigint DEFAULT NULL,
  `update_ts` int DEFAULT NULL,
  `r_idx` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_uid_tid` (`user_id`,`theme_id`) USING BTREE,
  KEY `index_update_time` (`update_ts`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11088 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for theme_map
-- ----------------------------
DROP TABLE IF EXISTS `theme_map`;
CREATE TABLE `theme_map` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `theme_map_id` int DEFAULT NULL,
  `map_credits` int DEFAULT NULL,
  `map_level` int DEFAULT NULL,
  `map_round` int DEFAULT NULL,
  `total_bet` bigint DEFAULT NULL,
  `spin_count` int DEFAULT NULL,
  `map_data` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for theme_statistical_data
-- ----------------------------
DROP TABLE IF EXISTS `theme_statistical_data`;
CREATE TABLE `theme_statistical_data` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `level` int DEFAULT NULL,
  `last_update_date` datetime DEFAULT NULL,
  `first_login_date` datetime DEFAULT NULL,
  `last_login_date` datetime DEFAULT NULL,
  `spin_count` int DEFAULT NULL,
  `last_7days_theme_statistical_spin_data` varchar(8000) DEFAULT NULL,
  `xxx1_data` varchar(2048) DEFAULT NULL,
  `xxx2_data` varchar(1024) DEFAULT NULL,
  `xxx3_data` varchar(1024) DEFAULT NULL,
  `last_7days_most_spin_tid` smallint DEFAULT NULL,
  `last_7days_second_spin_tid` smallint DEFAULT NULL,
  `xx1` smallint DEFAULT NULL,
  `xx2` smallint DEFAULT NULL,
  `xx3` smallint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_theme_statistical_data_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2874 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `device_pass` varchar(60) DEFAULT NULL,
  `facebook_id` varchar(64) DEFAULT NULL,
  `uuid` varchar(64) DEFAULT NULL,
  `coins` bigint DEFAULT NULL,
  `exp` bigint DEFAULT NULL,
  `level` int DEFAULT NULL,
  `vip_level` smallint DEFAULT NULL,
  `vip_points` int DEFAULT NULL,
  `first_login` datetime DEFAULT NULL,
  `last_login` datetime DEFAULT NULL,
  `last_logout` datetime DEFAULT NULL,
  `purchase_count` int DEFAULT NULL,
  `total_purchase` int DEFAULT NULL,
  `max_purchase` int DEFAULT NULL,
  `max_weekly_purchase` int DEFAULT NULL,
  `first_purchase` datetime DEFAULT NULL,
  `last_purchase` datetime DEFAULT NULL,
  `spin_count` int DEFAULT NULL,
  `store_disconnection_data` varchar(1024) DEFAULT NULL,
  `user_flag` int DEFAULT '0',
  `apple_sub_id` varchar(64) DEFAULT '',
  `hms_id` varchar(180) DEFAULT '',
  `login_in_cn` int DEFAULT '0',
  `extra_int_1` int DEFAULT '0',
  `extra_int_2` int DEFAULT '0',
  `extra_int_3` int DEFAULT '0',
  `pactsafe` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `index_hms_id` (`hms_id`) USING BTREE,
  KEY `index_facebook_id` (`facebook_id`) USING BTREE,
  KEY `index_apple_sub_id` (`apple_sub_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4718 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `coins_str` varchar(64) DEFAULT NULL,
  `nick_name` varchar(64) DEFAULT NULL,
  `user_icon1` varchar(64) DEFAULT NULL,
  `user_icon2` varchar(64) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `xx1` int DEFAULT NULL,
  `xx2` int DEFAULT NULL,
  `xx3` int DEFAULT NULL,
  `yy1` varchar(16) DEFAULT NULL,
  `yy2` varchar(16) DEFAULT NULL,
  `yy3` varchar(32) DEFAULT NULL,
  `yy4` varchar(64) DEFAULT NULL,
  `yy5` varchar(256) DEFAULT NULL,
  `yy6` varchar(1024) DEFAULT NULL,
  `yy7` varchar(1024) DEFAULT NULL,
  `yy8` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_user_info_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1125 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for user_persona
-- ----------------------------
DROP TABLE IF EXISTS `user_persona`;
CREATE TABLE `user_persona` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `platform_type` int DEFAULT NULL,
  `user_source` int DEFAULT NULL,
  `purchase_type` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_user_persona_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5746 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
