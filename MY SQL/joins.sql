create table employee(
e_id int not null,
e_name varchar(50),
e_salary int,
e_age int,
e_gender varchar (50),
e_dept varchar(20),
primary key(e_id)
);


insert into employee values(
1,'Sumit',95000,45,'Male','Backend');
insert into employee values(2,'Rohan',80000,'Male','Sales');
insert into employee values(3,'Matt',75000,'Male','analytic');
insert into employee values(4,'Rashmi',80000,"Female",'Backend');
insert into employee values(5,'Ram',95000,'Male','Sales');


create table department(
e_id int not null,
e_dept varchar(50),
dept_location varchar(50)
 );
 
insert into department values(
1,'Backend','Uttar Pradesh');
insert into department values(
2,'Sales','Maharastra');
insert into department values(
3,'analystic','Bihar');
insert into department values(
4,'Backend','Gujarat');
insert into department values(
5,'Sales','Uttar Pradesh');


select *from employee;
select *from department;
select * from employee inner join department on  employee.e_id=department.e_id;
select * from employee left join department on  employee.e_id=department.e_id;
select * from employee full join department on  employee.e_id=department.e_id;
