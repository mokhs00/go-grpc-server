CREATE DATABASE IF NOT EXISTS `go_grpc_server` DEFAULT CHARACTER SET utf8mb4;

USE `go_grpc_server`;

CREATE TABLE `content`
(
    `content_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '컨텐츠 ID',
    `title`      varchar(255) COMMENT '컨텐츠 제목',
    `body`       text COMMENT '컨텐츠 본문',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '생성 일시',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '수정 일시',
    PRIMARY KEY (`content_id`),
    KEY `idx_content_m1` (`created_at`),
    KEY `idx_content_m2` (`updated_at`)
) ENGINE = InnoDB DEFAULT CHAR SET = `utf8mb4` COMMENT '컨텐츠';