create database test;
use test;
create table demo(
Sno int,
Name varchar(50),
Age int,
Gender varchar(10),
Phone varchar(10),
Address varchar(50));
insert into demo values(1,"ram",20,'Male',789456123,'0078/oj');
insert into demo values(2,"rajesh",28,'Male',889456123,'7b/oj');
insert into demo values(3,"rashi",20,'Female',897456123,'07e/oj');
insert into demo values(4,"rohan",15,'Male',7789456123,'71/oj');
insert into demo values(5,"muuskan",86,'Female',7979456123,'7p/oj');

select* from demo
set sql_safe_updates=0;
update demo
set age='25'
where Name='rohan';

select *from demo
delete from demo where Name='muuskan';

select* from demo

drop database test






  