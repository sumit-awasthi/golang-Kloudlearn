/*use joins;
select* from table1 LEFT OUTER JOIN table2 on table1.column1=table2.column2;
use company*/
select name,jobtitle,title from employees LEFT OUTER JOIN projects on employees.employeeid=projects.employeeid;