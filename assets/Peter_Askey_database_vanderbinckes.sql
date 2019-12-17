CREATE DATABASE  IF NOT EXISTS `vanderbinckes` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `vanderbinckes`;
-- MySQL dump 10.13  Distrib 5.7.25, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: vanderbinckes
-- ------------------------------------------------------
-- Server version	5.7.28-0ubuntu0.18.04.4

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Accessory`
--

DROP TABLE IF EXISTS `Accessory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Accessory` (
  `accessoryId` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `price` decimal(3,2) NOT NULL,
  PRIMARY KEY (`accessoryId`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Accessory`
--

LOCK TABLES `Accessory` WRITE;
/*!40000 ALTER TABLE `Accessory` DISABLE KEYS */;
INSERT INTO `Accessory` VALUES (1,'regendak',2.50),(2,'zonnedak',2.00),(3,'babystoeltje',3.00),(4,'smart phone houder',1.00),(5,'kaarthouder',1.00),(6,'helm',1.50);
/*!40000 ALTER TABLE `Accessory` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Bycicle`
--

DROP TABLE IF EXISTS `Bycicle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Bycicle` (
  `bycicleId` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `type` varchar(10) NOT NULL,
  `price` decimal(4,2) NOT NULL,
  `amount` int(11) NOT NULL,
  PRIMARY KEY (`bycicleId`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Bycicle`
--

LOCK TABLES `Bycicle` WRITE;
/*!40000 ALTER TABLE `Bycicle` DISABLE KEYS */;
INSERT INTO `Bycicle` VALUES (3,'Cargo Bike Green','CBG1',20.00,10),(4,'Cargo Bike Electric','CBE1',40.00,10);
/*!40000 ALTER TABLE `Bycicle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Customer`
--

DROP TABLE IF EXISTS `Customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Customer` (
  `customerId` int(11) NOT NULL AUTO_INCREMENT,
  `postalCode` varchar(6) NOT NULL,
  `houseNumber` int(11) NOT NULL,
  `surname` varchar(15) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `houseNumberSuffix` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`customerId`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Customer`
--

LOCK TABLES `Customer` WRITE;
/*!40000 ALTER TABLE `Customer` DISABLE KEYS */;
INSERT INTO `Customer` VALUES (1,'1999XB',10,'Sharp','Leo','a'),(2,'1999XB',84,'Long','George',NULL),(3,'1019BY',60,'Caouette','Coby',NULL),(4,'1019BY',97,'Donatelli','Sean',NULL),(5,'1019BY',3,'Guyer','Rachael',NULL),(6,'6931NX',76,'Perillous','Ike',NULL),(7,'6931NX',13,'Linhart','Ciara',NULL),(8,'2691UQ',27,'Francis','Oscar',NULL),(9,'6591FR',77,'Stannard','Gillian',NULL),(10,'6591FR',94,'Zapetis','Rasmus',NULL),(11,'6591FR',54,'Anderson','Mariska',NULL),(12,'2532VL',26,'Kuehn','Jurre',NULL),(13,'2532VM',96,'Gonzalez','Luis',NULL),(14,'7176NU',66,'Stevens','Maja',NULL),(15,'7176NU',25,'Sterrett','Scotty',NULL),(16,'7668WA',16,'Schubert','Charlotte',NULL),(17,'7668WA',64,'Beckbau','Sophia',NULL),(18,'6173XD',37,'Goodman','Rachael',NULL),(19,'6173XD',84,'Love','Luis',NULL),(20,'6173XD',28,'Hedgecock','Sophia',NULL),(21,'3440JV',64,'Brennan','Lara',NULL),(22,'3440JV',52,'Pengilly','Jurre',NULL),(23,'7290ZN',35,'Noteboom','Vanessa',NULL),(24,'1952FB',93,'Daley','Barbara',NULL),(25,'1952FB',79,'Bruno','Cloe',NULL);
/*!40000 ALTER TABLE `Customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Employee`
--

DROP TABLE IF EXISTS `Employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Employee` (
  `employeeId` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `employee_since` date NOT NULL,
  `surname` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`employeeId`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Employee`
--

LOCK TABLES `Employee` WRITE;
/*!40000 ALTER TABLE `Employee` DISABLE KEYS */;
INSERT INTO `Employee` VALUES (1,'Bas','2018-05-21','Jansen'),(2,'Vincent','2017-09-01','Rademakers'),(3,'Karel','2018-05-30','Heiden, van der'),(4,'Irene','2017-09-01','Kraymans'),(5,'Francine','2018-07-01','Boer, de'),(6,'Jaap','2017-11-01','Velzenmaker');
/*!40000 ALTER TABLE `Employee` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OrderAccessory`
--

DROP TABLE IF EXISTS `OrderAccessory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OrderAccessory` (
  `orderAccessoryId` int(11) NOT NULL AUTO_INCREMENT,
  `orderLineId` int(11) NOT NULL,
  `accessoryId` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  PRIMARY KEY (`orderAccessoryId`),
  KEY `FK_verhuuracc_accessoire` (`accessoryId`),
  KEY `FK_verhuuracc_verhuur` (`orderLineId`),
  CONSTRAINT `FK_verhuuracc_accessoire` FOREIGN KEY (`accessoryId`) REFERENCES `Accessory` (`accessoryId`),
  CONSTRAINT `FK_verhuuracc_verhuur` FOREIGN KEY (`orderLineId`) REFERENCES `OrderLine` (`orderLineId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderAccessory`
--

LOCK TABLES `OrderAccessory` WRITE;
/*!40000 ALTER TABLE `OrderAccessory` DISABLE KEYS */;
INSERT INTO `OrderAccessory` VALUES (1,1,1,1),(2,1,6,2);
/*!40000 ALTER TABLE `OrderAccessory` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OrderLine`
--

DROP TABLE IF EXISTS `OrderLine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OrderLine` (
  `orderLineId` int(11) NOT NULL AUTO_INCREMENT,
  `bycicleId` int(11) NOT NULL,
  `customerId` int(11) NOT NULL,
  `employeeId` int(11) NOT NULL,
  `startDate` date NOT NULL,
  `days` int(11) NOT NULL,
  `totalPrice` decimal(5,2) DEFAULT NULL,
  PRIMARY KEY (`orderLineId`),
  KEY `FK_verhuur_klant` (`customerId`),
  KEY `FK_verhuur_medewerker` (`employeeId`),
  KEY `FK_verhuur_bakfiets` (`bycicleId`),
  CONSTRAINT `FK_verhuur_bakfiets` FOREIGN KEY (`bycicleId`) REFERENCES `Bycicle` (`bycicleId`),
  CONSTRAINT `FK_verhuur_klant` FOREIGN KEY (`customerId`) REFERENCES `Customer` (`customerId`),
  CONSTRAINT `FK_verhuur_medewerker` FOREIGN KEY (`employeeId`) REFERENCES `Employee` (`employeeId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderLine`
--

LOCK TABLES `OrderLine` WRITE;
/*!40000 ALTER TABLE `OrderLine` DISABLE KEYS */;
INSERT INTO `OrderLine` VALUES (1,1,2,1,'2019-12-02',5,127.50),(2,1,4,3,'2019-12-01',3,80.00),(3,2,10,3,'2019-11-30',7,140.00);
/*!40000 ALTER TABLE `OrderLine` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-12-17  9:14:01
