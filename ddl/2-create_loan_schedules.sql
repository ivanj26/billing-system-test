-- Table `loan_schedules`
CREATE TABLE `loan_schedules` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `loan_id` BIGINT unsigned NOT NULL,
    `amount` float(32, 2) DEFAULT 0,
    `due_date`datetime NOT NULL,

    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT `loan_id_foreign_key_ref_loans_id` FOREIGN KEY (`loan_id`) REFERENCES `loans` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);