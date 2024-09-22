-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 22, 2024 at 09:27 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `academia_db_main`
--

-- --------------------------------------------------------

--
-- Table structure for table `student`
--

CREATE TABLE `student` (
  `id` int(11) NOT NULL,
  `student_id` varchar(50) NOT NULL,
  `full_name` varchar(50) NOT NULL,
  `user_name` varchar(20) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `profile_picture` int(255) NOT NULL,
  `registration_year` date NOT NULL,
  `registration_semester` varchar(20) NOT NULL,
  `department` varchar(100) NOT NULL,
  `batch_number` int(11) NOT NULL,
  `current_section` varchar(20) NOT NULL,
  `advisor` varchar(100) NOT NULL,
  `super_advisor` varchar(100) NOT NULL,
  `current_cgpa` varchar(50) NOT NULL,
  `final_project_title` varchar(255) NOT NULL,
  `final_cgpa` varchar(50) DEFAULT NULL,
  `refresh_token` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `student`
--

INSERT INTO `student` (`id`, `student_id`, `full_name`, `user_name`, `email`, `password`, `profile_picture`, `registration_year`, `registration_semester`, `department`, `batch_number`, `current_section`, `advisor`, `super_advisor`, `current_cgpa`, `final_project_title`, `final_cgpa`, `refresh_token`) VALUES
(1, '203-15-3914', 'B.M. Rahat Almas', 'purplehand', 'almas15-3914@diu.edu.bd\r\n', '123456', 0, '2020-09-16', 'fall-2020', 'Computer Science And Engineering', 57, '57_D', '123456', '123456', '3.00', 'Academia', '', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `student`
--
ALTER TABLE `student`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
