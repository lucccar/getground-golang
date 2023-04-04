CREATE TABLE IF NOT EXISTS `tables` (
  `id` INT NOT NULL auto_increment PRIMARY KEY,
  `capacity` INT NOT NULL,
  `freeseats` INT NOT NULL
);


CREATE TABLE IF NOT EXISTS `expectedguests` (
  `id` INT NOT NULL auto_increment PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `table` INT NOT NULL,
  `accompanyingGuests` INT
);

CREATE TABLE IF NOT EXISTS `guests` (
  `id` INT NOT NULL auto_increment PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `table` INT NOT NULL,
  `accompanyingGuests` INT,
  `timeArrived` TIMESTAMP NOT NULL
);