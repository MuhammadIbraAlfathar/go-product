CREATE TABLE users (
                       id INT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
                       name VARCHAR(255) NOT NULL ,
                       user_name VARCHAR(255),
                       email VARCHAR(255) NOT NULL ,
                       password VARCHAR(255) NOT NULL ,
                       gender enum ('female', 'male') NOT NULL,
                       address VARCHAR(500) NOT NULL ,
                       created_at TIMESTAMP DEFAULT current_timestamp,
                       updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp
) ENGINE = InnoDB;