/*create database joins;
use joins;

create table table1(column1 int);
create table table2(column2 int);
show tables 
insert into table1 values(11),(12),(13),(14),(15);
insert into table2 values(12),(15),(50);
select*from table1 INNER JOIN table2 on table1.column1=table2.column2;/*use of inner join 

use company;

select* from employees;
select * from projects;
select name,jobtitle,title from employees INNER JOIN projects ON employees.employeeid = projects.employeeid;*/

use onlineshop;
select* from customers;
select * from orders;
select customers.userid,name ,phone,items,total from customers INNER JOIN orders ON customers.userid=orders.userid;