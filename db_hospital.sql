CREATE TABLE IF NO EXISTS `tb_patients` (
    `id` int NOT NULL AUTO_INCREMENT,
    `identity_number` varchar(30) DEFAULT NULL,
    `family_card_number` varchar(30) DEFAULT NULL,
    `first_name` varchar(60) NOT NULL,
    `last_name` varchar(60) DEFAULT NULL,
    `birth_date` date DEFAULT NULL,
    `address` text DEFAULT NULL,
    `bpjs_number` varchar(30) DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
) ENGINE = InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET = latin1;


CREATE TABLE IF NOT EXISTS `tb_hospitals` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(60) NOT NULL,
    `address` text DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
) ENGINE = InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET = latin1;


CREATE TABLE IF NOT EXISTS `tb_hospitals_patients` (
    `id` int NOT NULL AUTO_INCREMENT,
    `hospital_id` int NOT NULL,
    `patient_id` int DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
) ENGINE = InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET = latin1;
