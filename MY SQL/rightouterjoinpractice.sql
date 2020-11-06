/*use joins

select * from table1 RIGHT OUTER JOIN table2 ON table1.column1=table2.column2;

use company
select name,jobtitle,title from employees RIGHT OUTER JOIN projects ON employees.employeeid=projects.employeeid;

use onlineshop;
select c.userid,c.name,c.phone,o.items,o.total from customers as c right outer join orders as o on c.userid=o.userid;*/
use joins;
select * from table1 left outer join table2 on table1.column1=table2.column2 UNION select * from table1 RIGHT OUTER JOIN table2 on table1.column1=table2.column2;