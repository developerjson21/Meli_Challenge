/*
Se deberá contar con una base de datos PostgreSQL ó MySQL (según lo que se considere
mejor para persistir la información del resultado de los escaneos). En esta base de datos debe
poder verse: configuraciones, historial de escaneos, y la clasificación actualizada a partir del
último escaneo de cada repositorio.

La tabla de clasificación debería contener los siguientes campos:  
*/
DROP TABLE IF EXISTS `db_meli_scanner`.`CLASSIFICATION`;
CREATE SCHEMA `db_meli_scanner` ;
CREATE TABLE `db_meli_scanner`.`CLASSIFICATION` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `REPOSITORY_URL` VARCHAR(200) NOT NULL,
  `FILE` VARCHAR(45) NOT NULL,
  `INFORMATION_TYPE` VARCHAR(45) NOT NULL,
  `NUMBER_OF_LINES` VARCHAR(45),
  `DETECTION_DATE` DATETIME NOT NULL,
  `JOB_ID` VARCHAR(45) NOT NULL,
  `REPOSITORY_OWNER` VARCHAR(45) NOT NULL,
  `AMOUNT_DETECTIONS` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`ID`)
  );
  
  
  