ALTER TABLE `location`
ADD COLUMN `CreatedAt` DATETIME NOT NULL DEFAULT current_timestamp(),
ADD COLUMN `UpdatedAt` DATETIME NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
ADD COLUMN `DeletedAt` DATETIME;
