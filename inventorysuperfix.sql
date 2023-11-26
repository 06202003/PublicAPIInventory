-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 26, 2023 at 08:35 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `inventorysuperfix`
--

-- --------------------------------------------------------

--
-- Table structure for table `history_pemakaian`
--

CREATE TABLE `history_pemakaian` (
  `id` bigint(20) NOT NULL,
  `nomor_induk_old` varchar(45) DEFAULT NULL,
  `nomor_induk_new` varchar(45) DEFAULT NULL,
  `tanggal` longtext DEFAULT NULL,
  `ruangan_old` varchar(20) DEFAULT NULL,
  `ruangan_new` varchar(20) DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL,
  `id_pemakaian` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `history_pemakaian`
--

INSERT INTO `history_pemakaian` (`id`, `nomor_induk_old`, `nomor_induk_new`, `tanggal`, `ruangan_old`, `ruangan_new`, `created_at`, `updated_at`, `id_pemakaian`) VALUES
(10, '', '123123', '2023-10-26 08:49:58', '', 'INT2', '2023-11-22 15:46:15', '2023-11-22 15:46:15', 'ORGPK-1'),
(11, '', '123123', '2023-10-26 08:49:58', '', 'INT2', '2023-11-22 15:46:15', '2023-11-22 15:46:15', 'ORGPK-1'),
(12, '', '123123', '2023-10-26 08:49:58', '', 'INT2', '2023-11-22 15:46:39', '2023-11-22 15:46:39', 'ORGPK-1'),
(13, '', '123123', '2023-10-26 08:49:58', '', 'INT2', '2023-11-22 15:46:39', '2023-11-22 15:46:39', 'ORGPK-1'),
(14, '', '123123', '2023-10-26 08:49:58', '', 'DBS', '2023-11-22 15:59:07', '2023-11-22 15:59:07', 'ORGPK-1'),
(15, '', '123123', '2023-10-26 08:49:58', '', 'INT2', '2023-11-22 15:59:07', '2023-11-22 15:59:07', 'ORGPK-1'),
(16, '123123', '123123', '2023-10-26 08:49:58', 'INT2', 'DBS', '2023-11-22 15:59:28', '2023-11-22 15:59:28', 'ORGPK-1'),
(17, '123123', '123123', '2023-10-27 08:49:58', 'DBS', 'INT2', '2023-11-22 16:00:07', '2023-11-22 16:00:07', 'ORGPK-1');

-- --------------------------------------------------------

--
-- Table structure for table `history_perbaikan`
--

CREATE TABLE `history_perbaikan` (
  `id` bigint(20) NOT NULL,
  `tanggal_perbaikan` longtext DEFAULT NULL,
  `biaya` bigint(20) DEFAULT NULL,
  `deskripsi` varchar(255) DEFAULT NULL,
  `tanggal_kerusakan` longtext DEFAULT NULL,
  `tanggal_selesai_perbaikan` longtext DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL,
  `id_pemakaian` varchar(20) DEFAULT NULL,
  `tempat_perbaikan` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `history_perbaikan`
--

INSERT INTO `history_perbaikan` (`id`, `tanggal_perbaikan`, `biaya`, `deskripsi`, `tanggal_kerusakan`, `tanggal_selesai_perbaikan`, `created_at`, `updated_at`, `id_pemakaian`, `tempat_perbaikan`) VALUES
(6, '2023-10-26 08:49:58', 100000, 'Rusak Ringan ges', '2023-10-25 08:49:58', '2023-10-27 08:49:58', '2023-10-27 08:49:58', '2023-10-27 08:49:58', 'ORGPK-1', 'BENGKEL LAPTOP');

-- --------------------------------------------------------

--
-- Table structure for table `inventory`
--

CREATE TABLE `inventory` (
  `kode_aset` varchar(20) NOT NULL,
  `merk` varchar(45) DEFAULT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `tanggal` longtext DEFAULT NULL,
  `harga` bigint(20) DEFAULT NULL,
  `nilai_residu` bigint(20) DEFAULT NULL,
  `masa_manfaat` bigint(20) DEFAULT NULL,
  `depresiasi` bigint(20) DEFAULT NULL,
  `deskripsi` varchar(255) DEFAULT NULL,
  `status` varchar(20) DEFAULT NULL,
  `id_kategori` varchar(3) DEFAULT NULL,
  `tahun_1` bigint(20) DEFAULT NULL,
  `tahun_2` bigint(20) DEFAULT NULL,
  `tahun_3` bigint(20) DEFAULT NULL,
  `tahun_4` bigint(20) DEFAULT NULL,
  `img_url` varchar(255) DEFAULT NULL,
  `vendor` varchar(100) DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `inventory`
--

INSERT INTO `inventory` (`kode_aset`, `merk`, `nama`, `tanggal`, `harga`, `nilai_residu`, `masa_manfaat`, `depresiasi`, `deskripsi`, `status`, `id_kategori`, `tahun_1`, `tahun_2`, `tahun_3`, `tahun_4`, `img_url`, `vendor`, `created_at`, `updated_at`) VALUES
('STV-001', 'Sony', 'Smart TV', '2023-10-20', 500000, 125000, 4, 93750, 'Smart TV 42inch', 'normal', 'SPL', 406250, 312500, 218750, 125000, '', '', '2023-10-26 04:17:32', '2023-10-26 08:49:58');

-- --------------------------------------------------------

--
-- Table structure for table `karyawan`
--

CREATE TABLE `karyawan` (
  `nomor_induk` varchar(20) NOT NULL,
  `gambar` varchar(255) DEFAULT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `gender` tinyint(1) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `telepon` varchar(20) DEFAULT NULL,
  `jabatan` varchar(20) DEFAULT NULL,
  `divisi` varchar(20) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `karyawan`
--

INSERT INTO `karyawan` (`nomor_induk`, `gambar`, `nama`, `gender`, `email`, `telepon`, `jabatan`, `divisi`, `alamat`, `created_at`, `updated_at`) VALUES
('123123', '', 'David Updated', 0, '2172015@maranatha.ac.id', '08755225', 'CTO', 'IT', 'Cimahi', '2023-10-24 13:43:32', '2023-10-26 06:27:04');

-- --------------------------------------------------------

--
-- Table structure for table `kategori`
--

CREATE TABLE `kategori` (
  `id_kategori` varchar(3) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `kategori`
--

INSERT INTO `kategori` (`id_kategori`, `nama`, `created_at`, `updated_at`) VALUES
('SPL', 'Sample Category Updated', '2023-11-19T12:00:00Z', '2023-11-19T12:30:00Z');

-- --------------------------------------------------------

--
-- Table structure for table `lokasi`
--

CREATE TABLE `lokasi` (
  `id_lokasi` varchar(5) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `alamat` varchar(100) DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `lokasi`
--

INSERT INTO `lokasi` (`id_lokasi`, `nama`, `alamat`, `created_at`, `updated_at`) VALUES
('BDG', 'Bandung', 'Bandung No 18', '2023-10-26 04:12:43', '2023-10-26 04:12:43'),
('SBY', 'Surabaya', 'Surabaya No 18', '2023-10-26 04:12:43', '2023-10-26 04:12:43');

-- --------------------------------------------------------

--
-- Table structure for table `pemakaian`
--

CREATE TABLE `pemakaian` (
  `id_pemakaian` varchar(20) NOT NULL,
  `kode_aset` varchar(20) DEFAULT NULL,
  `nomor_induk` varchar(20) DEFAULT NULL,
  `id_ruangan` varchar(5) DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pemakaian`
--

INSERT INTO `pemakaian` (`id_pemakaian`, `kode_aset`, `nomor_induk`, `id_ruangan`, `created_at`, `updated_at`) VALUES
('ORGPK-1', 'STV-001', '123123', 'INT2', '2023-10-26 04:17:32', '2023-10-27 08:49:58');

-- --------------------------------------------------------

--
-- Table structure for table `ruangan`
--

CREATE TABLE `ruangan` (
  `id_ruangan` varchar(5) NOT NULL,
  `nama` longtext DEFAULT NULL,
  `created_at` longtext DEFAULT NULL,
  `updated_at` longtext DEFAULT NULL,
  `id_lokasi` varchar(5) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ruangan`
--

INSERT INTO `ruangan` (`id_ruangan`, `nama`, `created_at`, `updated_at`, `id_lokasi`) VALUES
('DBS', 'Laboratorium Database', '2023-10-26 04:17:32', '2023-10-26 04:17:32', 'SBY'),
('INT2', 'Laboratorium Internet 2', '2023-10-26 04:12:43', '2023-10-26 04:12:43', 'SBY');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  `nama` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `status`, `nama`) VALUES
(1, '2172003@maranatha.ac.id', '$2a$10$WSyQLeaPOslYWdUcb3E9oe9PaGkLEhO6TH9AY12n3pQYqI9c8ST7K', 'Aktif', 'Yehezkiel David Setiawan'),
(2, '2172003@maranatha.ac.id', '$2a$10$dBnAgqvqSR2A0f2YD4Vphu.61SdY3thcJSYRFUXemuz59n8c0PAI6', 'Aktif', 'Yehezkiel David Setiawan'),
(3, '2172028@maranatha.ac.id', '$2a$10$JJZmgfcGDR3wHet94bqJQe/U4ZyYpog9tUYWN.YY2uh5OYBJMvvL2', 'Aktif', 'Laurentius Ontoseno'),
(4, '2172028@maranatha.ac.id', '$2a$10$YC2TE27bHZ7wlyKFhfDqUe3ZdCO093IoDr122i.sk6/WvHYTY2OnW', 'Aktif', 'Laurentius Ontoseno');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `history_pemakaian`
--
ALTER TABLE `history_pemakaian`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_historyPemakaian_pemakaian` (`id_pemakaian`);

--
-- Indexes for table `history_perbaikan`
--
ALTER TABLE `history_perbaikan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_historyPerbaikan_pemakaian` (`id_pemakaian`);

--
-- Indexes for table `inventory`
--
ALTER TABLE `inventory`
  ADD PRIMARY KEY (`kode_aset`),
  ADD KEY `fk_inventory_category` (`id_kategori`);

--
-- Indexes for table `karyawan`
--
ALTER TABLE `karyawan`
  ADD PRIMARY KEY (`nomor_induk`);

--
-- Indexes for table `kategori`
--
ALTER TABLE `kategori`
  ADD PRIMARY KEY (`id_kategori`);

--
-- Indexes for table `lokasi`
--
ALTER TABLE `lokasi`
  ADD PRIMARY KEY (`id_lokasi`);

--
-- Indexes for table `pemakaian`
--
ALTER TABLE `pemakaian`
  ADD PRIMARY KEY (`id_pemakaian`),
  ADD KEY `fk_pemakaian_ruangan` (`id_ruangan`),
  ADD KEY `fk_pemakaian_karyawan` (`nomor_induk`),
  ADD KEY `fk_pemakaian_inventaris` (`kode_aset`);

--
-- Indexes for table `ruangan`
--
ALTER TABLE `ruangan`
  ADD PRIMARY KEY (`id_ruangan`),
  ADD KEY `fk_ruangan_lokasi` (`id_lokasi`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `history_pemakaian`
--
ALTER TABLE `history_pemakaian`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- AUTO_INCREMENT for table `history_perbaikan`
--
ALTER TABLE `history_perbaikan`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `history_pemakaian`
--
ALTER TABLE `history_pemakaian`
  ADD CONSTRAINT `fk_historyPemakaian_pemakaian` FOREIGN KEY (`id_pemakaian`) REFERENCES `pemakaian` (`id_pemakaian`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `history_perbaikan`
--
ALTER TABLE `history_perbaikan`
  ADD CONSTRAINT `fk_historyPerbaikan_pemakaian` FOREIGN KEY (`id_pemakaian`) REFERENCES `pemakaian` (`id_pemakaian`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `inventory`
--
ALTER TABLE `inventory`
  ADD CONSTRAINT `fk_inventory_category` FOREIGN KEY (`id_kategori`) REFERENCES `kategori` (`id_kategori`);

--
-- Constraints for table `pemakaian`
--
ALTER TABLE `pemakaian`
  ADD CONSTRAINT `fk_pemakaian_employee` FOREIGN KEY (`nomor_induk`) REFERENCES `karyawan` (`nomor_induk`),
  ADD CONSTRAINT `fk_pemakaian_inventaris` FOREIGN KEY (`kode_aset`) REFERENCES `inventory` (`kode_aset`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_pemakaian_inventory` FOREIGN KEY (`kode_aset`) REFERENCES `inventory` (`kode_aset`),
  ADD CONSTRAINT `fk_pemakaian_karyawan` FOREIGN KEY (`nomor_induk`) REFERENCES `karyawan` (`nomor_induk`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_pemakaian_room` FOREIGN KEY (`id_ruangan`) REFERENCES `ruangan` (`id_ruangan`),
  ADD CONSTRAINT `fk_pemakaian_ruangan` FOREIGN KEY (`id_ruangan`) REFERENCES `ruangan` (`id_ruangan`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ruangan`
--
ALTER TABLE `ruangan`
  ADD CONSTRAINT `fk_ruangan_lokasi` FOREIGN KEY (`id_lokasi`) REFERENCES `lokasi` (`id_lokasi`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
