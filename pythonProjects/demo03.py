import os
#print(os.getcwd())  #d:\Projects\pythonProjects 当前工作目录
os.chdir('D:\文档\pythonDir')
#print(os.getcwd())  #D:\文档\pythonDir
#data=open('sketch.txt')
#readline()方法从文件获取一个数据行
#print(data.readline(),end='')  
#print(data.readline(),end='')

#seek()方法返回到文件起始位置，对python文件可以使用tell()
#print(data.seek(0))   #0
# if os.path.exists('sketch.txt'):
#     data=open('sketch.txt')
#     for itemLine in data:
#         if not itemLine.find(':')==-1:
#             (role,line_spoken)=itemLine.split(':',1)
#             print(role,end='')
#             print('said:',end='')
#             print(line_spoken,end='')
    
#     data.close()
# else:
#     print('文件不存在')

try:
    #open()打开一个磁盘文件、创建一个迭代器从文件读取数据，一次读取一个数据行
    data=open('sketch.txt')
    for itemLine in data:
        try:
            (role,line_spoken)=itemLine.split(':',1)
            print(role,end='')
            print('said:',end='')
            print(line_spoken,end='')
        except ValueError:
            pass  #忽略错误
    data.close()  #colse()关闭一个之前打开的文件
except IOError:
    print('文件不存在')
    
"""
   使用open()打开一个磁盘文件，创建一个迭代器从文件读取数据
   一次读取一个数据行

   readline()方法从一个打开的文件读取一行数据

   seek()方法可以用来将文件"退回"到起始位置

   close()方法关闭一个之前打开的文件

   split()方法可以将一个字符串分解为一个子串列表

   python中不可改变的常量列表称为元组(tuple)
   一旦将列表数据赋至一个元组，就不能改变 
   
   数据不符合期望的格式时会出现ValueError

   数据无法正常访问时会出现IOError(例如，
   可能弄得数据文件已经被移走或者重命名)

   find()方法会在一个字符串中查找一个特定子串

   not关键字将一个条件取反

   try/except语句提供一个异常处理机制，
   从而保护可能导致运行时错误的某些代码行

   pass语句就是python的空语句或null语句，它什么也不做

"""