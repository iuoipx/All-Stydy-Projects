import os
import nester
import pickle

os.chdir('D:\文档\pythonDir')
man=[]
other=[]
try:
    data=open('sketch.txt')
    for itemLine in data:
        try:
            (role,line_spoken)=itemLine.split(':',1)
            line_spoken=line_spoken.strip()
            if role=='Man':
                man.append(line_spoken)
            elif role=='Other Man':
                other.append(line_spoken)
        except ValueError:
            pass  
    data.close()  
except IOError:
    print('文件不存在')

try:
    man_file=open('man_file.txt','w')
    other_file=open('other_file.txt','w')
    print(man,file=man_file)
    print(other,file=other_file)
except IOError:
    print('文件不存在')
finally: #无论出现什么错误都会执行
    man_file.close()
    other_file.close()

# try:
#     data=open('missing.txt')
#     print(data.readline(),end='')
# except IOError as err:
#     #No such file or directory: 'missing.txt'
#     print('File error'+str(err)) 
# finally:
#     if 'data' in locals():
#         data.close()


# try:
#     data=open('its.txt','w')
#     print('It\'s...',file=data)
# except IOError as err:
#     print('File error'+str(err))
# finally:
#     if 'data' in locals():
#         data.close()

        #等价于下面

try:
    with open('its.txt','w') as data:
        print('It\'s...',file=data)
except IOError as err:
    print('File error'+str(err))
"""使用with时，不需要操心关闭打开的文件，因为Python解释器会
   自动考虑这一点
"""

try:
    # with open('man_data.txt','w') as man_file:
    #     print(man,file=man_file)
    # with open('other_data.txt','w') as other_file:
    #     print(other,file=other_file)
    with open('man_data.txt','wb') as man_file,open('other_data.txt','wb') as other_file:
        pickle.dump(man,man_file)
        pickle.dump(other,other_file)
except IOError as err:
    print('File error:'+str(err))
except pickle.PickleError as perr:
    print('Picking error:'+str(perr))

# #使用dump()保存数据
# with open('mydata.pickle','wb') as mySaveData:
#     pickle.dump([1,2,'three'],mySaveData)
# # b 告诉python已二进制模式打开数据文件
# #使用load()从文件恢复数据
# with open('mydata.pickle','rb') as myRestoreData:
#     a_list=pickle.load(myRestoreData)

# print(a_list)


# with open('other_data.txt') as mdf:
#     print(mdf.readline())

"""
    strip()方法可以从字符串去除不想要的空白符

    print()的file参数控制将数据发送/保存到哪里

    finally组总会执行，而不论try/except语句中出现什么异常

    会向except组传入一个异常对象，并使用as关键字赋至一个标识符

    str()可以用来访问任何数据对象(支持串转换)的串表示

    locals()返回当前作用域中的变量集合

    in操作符用于检查成员关系

    "+"操作符用于字符串时将联接两个字符串，用于数字时则会将两个数相加

    with语句会自动处理所有以打开文件的关闭工作，即出现异常
    也不例外。with语句也可以使用as关键字

    sys.stdout是python中所谓的"标准输出",可以从标准库的sys模块访问

    标准库的pickle模块允许你容易而高效地将python数据对象保存
    到磁盘以及从磁盘恢复

    pickle.dump()函数将数据保存到磁盘

    pickle.load()函数从磁盘恢复数据
"""