-- create table merchants
DROP TABLE IF EXISTS `merchants`;
CREATE TABLE IF NOT EXISTS `merchants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(32) DEFAULT NULL,
  `username` varchar(50) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- add merchant
INSERT INTO Merchants 
(nama, status, username, password)
VALUES
('merchant1', '1', 'merchantone', 'lupalagi');

-- edit merchant
UPDATE Merchants
SET nama = 'new name', status = 'new status', username='new username', password = 'new password'
WHERE id = n;


-- get all merchant
SELECT * FROM Merchants;

-- get detail merchants
SELECT * FROM Merchants WHERE id = n;

-- delete merchant
DELETE FROM Merchants WHERE id = n;

-- =====================================
DROP TABLE IF EXISTS `products`;
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sku` varchar(50) DEFAULT NULL,
  `id_merchant` int(11) DEFAULT NULL,
  `image` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sku` (`sku`),
  KEY `FK_products_merchants` (`id_merchant`),
  CONSTRAINT `FK_products_merchants` FOREIGN KEY (`id_merchant`) REFERENCES `merchants` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- get all products
SELECT * FROM products;

-- get detail products
SELECT * FROM products WHERE id = n;

-- insert into products
INSERT INTO products 
(sku, image, id_merchant)
VALUES
('ctz-001', 'default/image.png', 1);

-- update products
UPDATE products
SET sku= 'ctz-002', image = 'image/ctz.png'
WHERE id_merchant=1;

-- DELETE PRODUCTS
DELETE FROM products WHERE sku = 'ctz-002';

-- =====================================

