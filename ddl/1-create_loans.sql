-- Table `loans`
CREATE TABLE `loans` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `borrower_id` BIGINT unsigned NOT NULL,
    `amount` float(32, 2) DEFAULT 0,
    `status` ENUM('inprogress', 'completed') DEFAULT 'inprogress',
    -- Assume 1: Status always updated to `completed` if the all loan schedules are paid
    -- Assume 2: A borrower can only borrow 1x at the moment. If the borrower want to borrow again, we will reject it if they have an inprogress loan. 
    
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT `borrower_id_foreign_key_ref_borrowers_id` FOREIGN KEY (`borrower_id`) REFERENCES `borrowers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);