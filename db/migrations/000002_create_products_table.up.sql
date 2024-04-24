create table products (
    id int not null auto_increment primary key ,
    product_name varchar(255) not null,
    stock int,
    price int not null ,
    description longtext,
    category_product varchar(255) not null ,
    picture_url longtext,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp

) ENGINE = InnoDB;