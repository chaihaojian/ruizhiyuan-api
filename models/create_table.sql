CREATE TABLE `admin` (
    `id` bigint(20) NOT NULL ,
    `name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `password` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `last_active` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    `status` int(16) NOT NULL DEFAULT '0' ,
    PRIMARY KEY (`id`) ,
    KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `video` (
    `id` bigint(20) NOT NULL ,
    `partition` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `interviewee` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `title` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `link` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    `status` int(16) NOT NULL DEFAULT '0' ,
    PRIMARY KEY (`id`) ,
    KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE `article` (
    `id` bigint(20) NOT NULL ,
    `partition` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `author` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `source` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `title` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
    `outline` varchar(512) COLLATE utf8mb4_general_ci NOT NULL ,
    `content` text COLLATE utf8mb4_general_ci NOT NULL ,
    `cover` varchar(256) COLLATE utf8mb4_general_ci ,
    `is_show` bool COLLATE utf8mb4_general_ci NOT NULL default false,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    `status` int(16) NOT NULL DEFAULT '0' ,
    PRIMARY KEY (`id`) ,
    KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;