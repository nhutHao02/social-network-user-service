CREATE TABLE IF NOT EXISTS `follow` (
  `FollowerID` INT(11) NOT NULL ,
  `FollowingID` INT(11) NOT NULL,
  `CreatedAt` DATETIME NOT NULL DEFAULT current_timestamp(),
  `UpdatedAt` DATETIME NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `DeletedAt` DATETIME,
  PRIMARY KEY (`FollowerID`,`FollowingID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
