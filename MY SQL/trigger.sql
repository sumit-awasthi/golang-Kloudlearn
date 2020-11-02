create table tblemplyee(
Emp_ID INT auto_increment,
Emp_name varchar(100),
Emp_Sal decimal(10.2)
);

insert into tblemplyee values(1,'Anand',48500.50);
insert into tblemplyee values(2,'Sumit',48500.50);
insert into tblemplyee values(3,'Raman',48500.50);
insert into tblemplyee values(4,'Nitin',38500.50);
insert into tblemplyee values(3,'Neha',51500.50);

create table emplyee_log(
Emp_ID int,
Emp_name varchar (50),
Emp_sal decimal(10,2),
Log_Active varchar(50),
log_Timestamp  datetime);
 create trigger tgrAfterInsert on tblemplyee
 After INSERT AS
  declare @empid int;
  declare @empname varchar(100);
  declare @empsal decimal(10,2);
  declare @log_action varchar(100);
  select @empid=i.Emp_ID,@empname=i.Emp_name,@empsal=i.Emp_sal from i;
  set@log_action ='Inserted record --After Insert Trigger';
  
  
  insert into Employee_log
	(Emp_ID,Emp_Name,Emp_sal,Log_Action,Log_Timwstamp)
	values(@empid, @empname,@empsal,@log_action,getdate());

    print'AFTER INSERT trigger fired'
GO
select* from tbtemplyee
select*from Emplyee_log
insert into tblemplyee values('Ram',10500)
   
create trigger thrAfterUpdate on tblemplyee
   After update
  as 
  declare @empid int;
  declare @empname varchar(50);
  declare @empsal decimal(10,2);
  declare@log_action varchar(100);
  
  
  select @empid=i.Emp_ID from inserted i;
  select @empname=i.Emp_Name from inserted i;
  select @empsal=i.Emp_sal from inserted i;
  
  
  if update (Emp_Name)
	set @log_action='updated Record--After update Trigger.';
  if update (Emp_Sal)
	set @log_action='updated Record--After update Trigger.';
  insert into emplyee_log(Emp_ID,Emp_Name,Emp_Sal,log_Action,Log_Timestamp)
  values(@empid,@empname,@empsal,@log_action,getdate());
  print'AFTER UPDATE Trigger Fired.'
go
update tblemployee set Emp_Sal=14500 where Emp_ID=6
select *from tblemployee
select * from employee_log

create trigger tgrAfterDelete On tblemployee
after delete
as
declare @empid int;
declare @empname varchar(100);
declare @empsal decimal(10,2);
declare@log_action varchar (100);

select @empid=d.Emp_ID from delected d;
select @empname=d.Emp_Name from delected d;
select @empsal=d.Emp_Sal from delected d;
set @log_action ='Deleted --After Delete Trigger';


insert into employee_log
(Emp_ID,Emp_Name,Emp_Sal,Log_Action,Log_Timestamp)
values(@empid,@empname,@empsal,@Log_action,getdate());
print'after delete trigger fired.'
go
delete from tblemployee where emp_id=6
select *from tblemployee 
select*from employee_log


create  trigger trgInsteadOfDelete ON tblemployee
INSTEAD OF  DELETE
AS
DECLARE@emp_id int;
declare @emp_name varchar(100);
declare@emp_sal int;


select @emp_id =d.Emp_ID from deleted d;
select @emp_name=d.Emp_Name from deleted d;
select @emp_sal=d.Emp_Sal from deleted d;

BEGIN
IF(@emp_sal>12000)
begin RAISERROR('Cannot delete  where salary>12000',16,1);
rollback;
end 
else
begin
	delete from tblemployee where Emp_ID=@emp_id;
	commit;
    insert into Employee_log(Emp_ID ,Emp_Name,Emp_Sal,Log_Action,Log_Timestamp)
    values(@emp_id,@emp_name,@emp_sal,'DELETED--Instead of delete trigger.',getdate());
    PRINT'Record deleted --Instead  of delete Trigger.'
    end
END
GO
