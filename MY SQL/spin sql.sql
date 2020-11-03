/*1st procedure call without passing paramter*/
CREATE PROCEDURE Sal_Emplyee
AS
	select Emp_Sal from tblemplyee
GO


exec Sal_Emplyee
/*2st procedure call without passing paramter*/
Create procedure employee_details
as
select*from tblemplyee
go

exec employee_details

/*1st procedure call with passing paramter*/
create procedure employee_gender @gender varchar(50)
as
select *from employee
where e_gender=@gender
go
exec employee_gender @gender='Male'
exec employee_gender @gender='Female'