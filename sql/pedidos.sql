-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 21-Dez-2020 às 20:36
-- Versão do servidor: 10.4.13-MariaDB
-- versão do PHP: 7.4.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `pedidos`
--
CREATE DATABASE IF NOT EXISTS `pedidos` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `pedidos`;

-- --------------------------------------------------------

--
-- Estrutura da tabela `pedidos`
--

CREATE TABLE `pedidos` (
  `id` int(11) UNSIGNED NOT NULL,
  `data` date NOT NULL,
  `valor_total` decimal(8,2) DEFAULT NULL,
  `pagamento` varchar(100) DEFAULT NULL,
  `idusuarios` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `pedidos`
--

INSERT INTO `pedidos` (`id`, `data`, `valor_total`, `pagamento`, `idusuarios`) VALUES
(1, '2020-12-18', '24.00', 'dinheiro', 1),
(2, '2020-12-21', '33.50', 'cartão de crédito', 2);

-- --------------------------------------------------------

--
-- Estrutura da tabela `pedidos_itens`
--

CREATE TABLE `pedidos_itens` (
  `id` int(11) NOT NULL,
  `idpedidos` int(11) UNSIGNED NOT NULL,
  `descricao` varchar(255) NOT NULL,
  `valor` decimal(8,2) UNSIGNED NOT NULL,
  `qtde` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `pedidos_itens`
--

INSERT INTO `pedidos_itens` (`id`, `idpedidos`, `descricao`, `valor`, `qtde`) VALUES
(1, 1, 'Sorvete de Chocolate', '6.50', 1),
(2, 1, 'Sorvete de Morango', '6.00', 2),
(3, 1, 'Sorvete de Uva', '5.50', 1),
(4, 2, 'Sorvete de Flocos', '5.50', 2),
(5, 2, 'Sorvete de Creme', '6.00', 2),
(6, 2, 'Sorvete de Napolitano', '5.50', 1),
(7, 2, 'Sorvete de Milho', '5.00', 1);

-- --------------------------------------------------------

--
-- Estrutura da tabela `usuarios`
--

CREATE TABLE `usuarios` (
  `id` int(11) UNSIGNED NOT NULL,
  `nome` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `usuarios`
--

INSERT INTO `usuarios` (`id`, `nome`, `email`) VALUES
(1, 'Lucas S. Rosa', 'lucas.tarta@hotmail.com'),
(2, 'Teste Completo e Sobrenome', 'teste@teste.com');

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `pedidos`
--
ALTER TABLE `pedidos`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pedidos_usuarios_id_fk` (`idusuarios`);

--
-- Índices para tabela `pedidos_itens`
--
ALTER TABLE `pedidos_itens`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pedidos_itens_pedidos_id_fk` (`idpedidos`);

--
-- Índices para tabela `usuarios`
--
ALTER TABLE `usuarios`
  ADD PRIMARY KEY (`id`);

--
-- Restrições para despejos de tabelas
--

--
-- Limitadores para a tabela `pedidos`
--
ALTER TABLE `pedidos`
  ADD CONSTRAINT `pedidos_usuarios_id_fk` FOREIGN KEY (`idusuarios`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Limitadores para a tabela `pedidos_itens`
--
ALTER TABLE `pedidos_itens`
  ADD CONSTRAINT `pedidos_itens_pedidos_id_fk` FOREIGN KEY (`idpedidos`) REFERENCES `pedidos` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
