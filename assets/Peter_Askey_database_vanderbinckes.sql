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
-- Table structure for table `Bicycle`
--

DROP TABLE IF EXISTS `Bicycle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Bicycle` (
  `bicycleId` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `type` varchar(10) NOT NULL,
  `price` decimal(4,2) NOT NULL,
  `amount` int(11) NOT NULL,
  PRIMARY KEY (`bicycleId`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Bicycle`
--

LOCK TABLES `Bicycle` WRITE;
/*!40000 ALTER TABLE `Bicycle` DISABLE KEYS */;
INSERT INTO `Bicycle` VALUES (3,'Cargo Bike Green','CBG1',20.00,10),(4,'Cargo Bike Electric','CBE1',40.00,10);
/*!40000 ALTER TABLE `Bicycle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Customer`
--

DROP TABLE IF EXISTS `Customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Customer` (
  `customerId` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `surname` varchar(15) NOT NULL,
  `postalCode` varchar(6) NOT NULL,
  `houseNumber` int(11) NOT NULL,
  `houseNumberSuffix` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`customerId`),
  UNIQUE KEY `unique_customer` (`postalCode`,`houseNumber`,`surname`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Customer`
--

LOCK TABLES `Customer` WRITE;
/*!40000 ALTER TABLE `Customer` DISABLE KEYS */;
INSERT INTO `Customer` VALUES (1,'Leo','Sharp','1999XB',10,'a'),(2,'George','Long','1999XB',84,NULL),(3,'Coby','Caouette','1019BY',60,NULL),(4,'Sean','Donatelli','1019BY',97,NULL),(5,'Rachael','Guyer','1019BY',3,NULL),(6,'Ike','Perillous','6931NX',76,NULL),(7,'Ciara','Linhart','6931NX',13,NULL),(8,'Oscar','Francis','2691UQ',27,NULL),(9,'Gillian','Stannard','6591FR',77,NULL),(10,'Rasmus','Zapetis','6591FR',94,NULL),(11,'Mariska','Anderson','6591FR',54,NULL),(12,'Jurre','Kuehn','2532VL',26,NULL),(13,'Luis','Gonzalez','2532VM',96,NULL),(14,'Maja','Stevens','7176NU',66,NULL),(15,'Scotty','Sterrett','7176NU',25,NULL),(16,'Charlotte','Schubert','7668WA',16,NULL),(17,'Sophia','Beckbau','7668WA',64,NULL),(18,'Rachael','Goodman','6173XD',37,NULL),(19,'Luis','Love','6173XD',84,NULL),(20,'Sophia','Hedgecock','6173XD',28,NULL),(21,'Lara','Brennan','3440JV',64,NULL),(22,'Jurre','Pengilly','3440JV',52,NULL),(23,'Vanessa','Noteboom','7290ZN',35,NULL),(24,'Barbara','Daley','1952FB',93,NULL),(25,'Cloe','Bruno','1952FB',79,NULL),(27,'Peter','Askey','5632LK',269,NULL),(28,'Henk','Waters','5503XG',52,'A'),(29,'Sander','Merks','9979AB',997,'A'),(30,'Henkie','Penkie','5598AS',59,NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Employee`
--

LOCK TABLES `Employee` WRITE;
/*!40000 ALTER TABLE `Employee` DISABLE KEYS */;
INSERT INTO `Employee` VALUES (1,'Bas','2018-05-21','Jansen'),(2,'Vincent','2017-09-01','Rademakers'),(3,'Karel','2018-05-30','Heiden, van der'),(4,'Irene','2017-09-01','Kraymans'),(5,'Francine','2018-07-01','Boer, de'),(6,'Jaap','2017-11-01','Velzenmaker'),(7,'Peter','1995-01-07','Askey');
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderAccessory`
--

LOCK TABLES `OrderAccessory` WRITE;
/*!40000 ALTER TABLE `OrderAccessory` DISABLE KEYS */;
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
  `bicycleId` int(11) NOT NULL,
  `customerId` int(11) NOT NULL,
  `employeeId` int(11) NOT NULL,
  `startDate` date NOT NULL,
  `days` int(11) NOT NULL,
  `totalPrice` decimal(5,2) DEFAULT NULL,
  PRIMARY KEY (`orderLineId`),
  KEY `FK_verhuur_klant` (`customerId`),
  KEY `FK_verhuur_medewerker` (`employeeId`),
  KEY `FK_verhuur_bakfiets` (`bicycleId`),
  CONSTRAINT `FK_verhuur_fiets` FOREIGN KEY (`bicycleId`) REFERENCES `Bicycle` (`bicycleId`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_verhuur_klant` FOREIGN KEY (`customerId`) REFERENCES `Customer` (`customerId`),
  CONSTRAINT `FK_verhuur_medewerker` FOREIGN KEY (`employeeId`) REFERENCES `Employee` (`employeeId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OrderLine`
--

LOCK TABLES `OrderLine` WRITE;
/*!40000 ALTER TABLE `OrderLine` DISABLE KEYS */;
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

-- Dump completed on 2020-01-06 18:49:51
