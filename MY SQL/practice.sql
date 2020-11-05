/*create database School;

use School

create table Students(
studentid INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
firstname varchar(50)NOT NULL,
lastname varchar(50) NOT NULL,
class varchar(20),
age int
);
create table Teachers(
teacherid int not null primary key auto_increment,
name varchar(100),
phone varchar(10)
);

create table Subjects(
subjectid int not null primary key auto_increment,
 title varchar (50)
 );

use company

create table Employees(
employeeid int not null primary key auto_increment,
name varchar(50) not null,
jobtitle varchar(50) not null,
salary float
);
create table Clients(
clientid int not null primary key auto_increment,
name varchar(50) not null,
phone varchar(50),
address varchar(150)
);
create table Projects(
projectid int not null primary key auto_increment,
title varchar(50) not null,
clientid int,
employeeid int,
startdate datetime,
enddate datetime
);

use onlineshop

create table items(
itemid int not null auto_increment primary key,
categoryid int,
name varchar(50)not null,
price float
);
create table categories (
categoryid int not null primary key auto_increment,
title varchar(50) not null,
status tinyint(1) not null
);
create table customers(
userid  int not null primary key auto_increment,
name varchar(50)not null,
phone varchar(50),
address varchar(50)
);

create table orders(
orderid int not null auto_increment primary key,
userid int,
items text,
total float,
orderdate datetime
);


use School

insert into students(firstname,lastname,class,age) values
 ('John','Smith','First',5),
 ('Mary','Doe','Third',8),
 ('James','Walker','Six',12),
 ('Andrew','Grawfild','four',8),
 ('Riche','Smith','First',6),
 ('Emma','Stone','Third',6),
 ('Adan','Mall','Six',7),
 ('Jeff','carry','four',10);
 
 
 select * from Students
 
 insert into teachers(name,phone)values
 ('SpiderMan','9119119111'),
 ('SuperMan','8944198415'),
 ('BatMan','1234567890'),
 ('IronMan','9876543210');
 
 insert into subjects(title)values ('English'),('Maths'),('Science'),('Computer'),('History');
 
 use Company;
 insert into employees (name,jobtitle,salary)values
 ('Andrew Grawfield','Backend developer',8400),
 ('Emma Stone','Frontend developer',8977),
 ('Tonny Stark','Backend developer',9999),
 ('Nick Fury','System Admin',15000),
 ('Steve Roggers','System Admin',15000),
 ('Black Panter','Frontend developer',8914),
 ('Hulk','IT developer',4875),
 ('josh Bailly','Database Admin',9845),
 ('Jerry Midsten','Database Admin',8987),
 ('Jose Matt','IT develop',7894);
 */
 
 insert into clients(name,phone,address)values
 ('Jimmy Jones','9874563210','Victoria Palace 16 High Street '),
 ('Merry Clerk','7789413492','477/B Rozi Street'),
 ('')