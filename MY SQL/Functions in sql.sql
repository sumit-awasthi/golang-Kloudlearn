select min(age)from demo;
select max(age) from demo;
select count(*) from demo where Gender ="Male";
select sum(age) from demo ;
select avg(age) from demo;
select "             Sumit";
select ltrim("         Sumit");
select "HI EVERYONE THIS IS SUMIT";
select lower("HI EVERYONE THIS IS SUMIT");
select"hello world";
select upper("hello world");
select"sumit";
select reverse("sumit");
select"this is Sumit";
select substring("this is Sumit",9,5);
select * from demo order by age;
select *from demo order by agr desc;




select avg(age),Gender from demo GROUP BY Gender;