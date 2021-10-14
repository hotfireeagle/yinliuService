create database yinliu character set utf8mb4 collate utf8mb4_unicode_ci;

use yinliu;

create table banners
(
    id char(36) NOT NULL,
    src varchar(255) NOT NULL,
    redirect_url varchar(255) NOT NULL,
    created_at datetime,
    deleted_at datetime,
    updated_at datetime,
    PRIMARY KEY (id)
);

create table buttons
(
    id char(36) NOT NULL,
    icon varchar(255) NOT NULL,
    title varchar(20) NOT NULL,
    desc_txt varchar(20) NOT NULL,
    btn_txt varchar(20) NOT NULL,
    redirect_url varchar(255) NOT NULL,
    created_at datetime,
    deleted_at datetime,
    updated_at datetime,
    PRIMARY KEY (id)
);

create table menus
(
    id char(36) NOT NULL,
    icon varchar(255) NOT NULL,
    text varchar(15) NOT NULL,
    redirect_url varchar(255) NOT NULL,
    created_at datetime,
    deleted_at datetime,
    updated_at datetime,
    PRIMARY KEY (id)
);

e