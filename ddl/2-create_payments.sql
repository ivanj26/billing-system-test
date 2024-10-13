-- Table `loans`
CREATE TABLE `payments` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `loan_schedule_id` BIGINT unsigned NOT NULL,
    `status` ENUM('pending', 'paid') DEFAULT 'pending', -- assume: the payment row will be created as a `pending payment` for every 1st day of month
    `amount` float(32, 2) DEFAULT NULL,
    `paid_date`datetime DEFAULT NULL,

    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT `loan_schedule_id_foreign_key_ref_loans_id` FOREIGN KEY (`loan_schedule_id`) REFERENCES `loan_schedules` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);