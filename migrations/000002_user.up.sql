CREATE TABLE IF NOT EXISTS `user` (
  `ID` INT(11) NOT NULL AUTO_INCREMENT,
  `Email` VARCHAR(255) NOT NULL ,
  `Password` VARCHAR(255) NOT NULL,
  `FullName` VARCHAR(255),
  `Sex` BIT,
  `Bio` TEXT,
  `UrlAvt` VARCHAR(255),
  `UrlBackground` VARCHAR(255),
  `CreatedAt` DATETIME NOT NULL DEFAULT current_timestamp(),
  `UpdatedAt` DATETIME NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `DeletedAt` DATETIME,
  `LocationID` INT(11),
  PRIMARY KEY (`ID`),
  UNIQUE KEY (`Email`),
  FOREIGN KEY (`LocationID`) REFERENCES `location`(`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
