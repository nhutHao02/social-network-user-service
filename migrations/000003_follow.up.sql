CREATE TABLE IF NOT EXISTS `follow` (
  `ID` INT(11) NOT NULL AUTO_INCREMENT,
  `FollowerID` INT(11) NOT NULL ,
  `FollowingID` INT(11) NOT NULL,
  PRIMARY KEY (`ID`, `FollowerID`,`FollowingID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
