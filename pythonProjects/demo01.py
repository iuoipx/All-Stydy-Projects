
cast=['Cleese','Palin','Jones','Idle']
print(len(cast))  #4
cast.append('Gilliam') #在列表末尾增加一个数据项
print(cast) #['Cleese', 'Palin', 'Jones', 'Idle', 'Gilliam']
cast.pop() #'Gilliam'  从列表末尾删除数据
print(cast) #['Cleese', 'Palin', 'Jones', 'Idle']
cast.extend(['Filliam',"Chapman"]) #在列表末尾增加一个数据项集合
print(cast) #['Cleese', 'Palin', 'Jones', 'Idle', 'Filliam', 'Chapman']
cast.remove('Chapman') #在列表中找到并删除一个特定的数据项
print(cast) #['Cleese', 'Palin', 'Jones', 'Idle', 'Filliam']
cast.insert(0,'chapman') #在特定位置前面增加一个数据项
print(cast) #['chapman', 'Cleese', 'Palin', 'Jones', 'Idle', 'Filliam']
movies=[
    'The Holy Grail',
    'teh Life of Brian',
    'The Meaning of Life'
]
movies.insert(1,1975)
movies.insert(3,1979)
movies.append(1983)
print(movies)

# for item in movies:
#     print(item)

# count=0
# while count<len(movies):
#     print(movies[count])
#     count=count+1


#chcp 65001编码  中文乱码