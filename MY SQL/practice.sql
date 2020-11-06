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

 
 insert into clients(name,phone,address)values
 ('Jimmy Jones','9874563210','Victoria Palace 16 High Street '),
 ('Merry Clerk','7789413492','477/B Rozi Street'),
 ('Jeff Hardy','8989868620','8/w champ road'),
 ('Smith Rosey','8896368620','dvv champ road'); 
 select * from clients
 
 use company;
 
 insert into projects(title,clientid,employeeid,startdate,enddate)values
 ('Golang course from scratch',1,3,now(),date_add(now(),interval 30 day)),
 ('AWS course for beginners',1,2,now(),date_add(now(),interval 45 day)),
 ('Java course for beginners',2,7,now(),date_add(now(),interval 45 day)),
 ('Python course for beginners',3,8,now(),date_add(now(),interval 15 day)),
 ('Microservices course for beginners',4,null,now(),date_add(now(),interval 20 day));
 
 
 use onlineshop;
 
 insert into categories(title,status)values('Electronics',1),('Books',1),('cloths',1);
 
 insert into items(categoryid,name,price)values
 (1,'Smart phone',2500.00),
 (1,'Laptops',10000.00),
 (2,'How To Train Your Dog',50.00),
 (2,'Healthy Dogs Food',150.00),
 (3,'Classy tshirt combo',520.00),
 (3,'Cool Denim shirt',1500.00);
 
 insert into customers(name,phone,address)values
 ('James Jones',7794949892,'victoria street,las vegas'),
 ('Henry Hellen',8794561230,'415/G5 sydney place,Maimi'),
 ('Ruby Young',7784569120,'22/E Washington DC' ),
 ('Michel Pal',55134441446,'22-H Canada');
 
 
 insert into Orders(userid,items,total,orderdate) values
 (2,'laptop',10000,now()),
 (1,'Healthy Dogs Food',150.00,now()),
 (25,'Food to be eaten',258,now()),
 (3,'Cool Denim Shirt',1500,now()),
 (5,'Classy tshirt combo',520.00,now());
 
 */
 
 /*Select clause in mysql*/
 
use school;
select* from students;/*to selecting everything from a table*/
select firstname,age from students;/*to select specific fields from the table */
select distinct firstname from students;
select* from students where studentid =6;
