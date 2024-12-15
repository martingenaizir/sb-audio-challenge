-- audio_challenge_db.users definition
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
     `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
     `user_level` tinyint(3) unsigned NOT NULL,
     `active` tinyint(1) NOT NULL,
     PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- audio_challenge_db.phrases definition
DROP TABLE IF EXISTS `phrases`;
CREATE TABLE `phrases` (
   `phrase_id` bigint(20) NOT NULL AUTO_INCREMENT,
   `level` tinyint(3) unsigned NOT NULL,
   `active` tinyint(1) NOT NULL,
   PRIMARY KEY (`phrase_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- audio_challenge_db.user_phrases definition

DROP TABLE IF EXISTS `user_phrases`;
CREATE TABLE `user_phrases` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `phrase_id` bigint(20) NOT NULL,
    `filepath` varchar(255) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT current_timestamp(),
    `last_update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`),
    KEY `user_phrase_idx` (`user_id`,`phrase_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- insert some default users
INSERT INTO audio_challenge_db.users (user_level, active)
VALUES (10, 1),(1,1);

-- insert some phrases
INSERT INTO audio_challenge_db.phrases(`level`, active)
    VALUES (1, 1), (11, 1);
