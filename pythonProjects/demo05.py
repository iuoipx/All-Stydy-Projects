import os
os.chdir('D:\文档\pythonDir\chapter05')

# with open('james.txt') as jaf:
#     data=jaf.readline()
# james=data.strip().split(',')

# with open('julie.txt') as juf:
#     data=juf.readline()
# julie=data.strip().split(',')

# with open('mikey.txt') as mif:
#     data=mif.readline()
# mikey=data.strip().split(',')

# with open('sarah.txt') as saf:
#     data=saf.readline()
# sarah=data.strip().split(',')

    #等于

def openFile(flieName):
    try:
        with open(flieName) as f:
            data=f.readline()
        return(data.strip().split(','))
    except IOError as err:
        print('File Error'+str(err))
        return None

james=openFile('james.txt')
julie=openFile('julie.txt')
mikey=openFile('mikey.txt')
sarah=openFile('sarah.txt')



def sanitize(time_string):
    #格式化时间数据
    if '-' in time_string:
        splitter='-'
    elif ':' in time_string:
        splitter=':'
    else:
        return time_string
    (mins,secs)=time_string.split(splitter)
    return(mins+'.'+secs)
    

print(sorted(set([sanitize(item) for item in james]))[0:3])
print(sorted(set([sanitize(item) for item in julie]))[0:3])
print(sorted(set([sanitize(item) for item in mikey]))[0:3])
print(sorted(set([sanitize(item) for item in sarah]))[0:3])


# unipue_jame=[]

# for item in clean_james:
#     if item not in unipue_jame:
#         unipue_jame.append(item)

# print(unipue_jame[0:3])        


"""
    sort()方法可以在原地改变列表的顺序

    sorted()通过提供复制排序可以对几乎任何数据结构排序

    向sort()或sorted()传入reverse=True可以按降序排列数据

    列表推导   new_list=[fn(t) for t in old_list]

    分片可以访问列表中的多个数据项
    [3,6]会访问列表中从索引3直到但不包括索引位置6的列表项

    使用set()工厂方法可以创建一个集合
"""