import nester;
movies=[
    'The Holy Grail',
    1975,
    'Terry Jones & Terry Gilliam',
    91,
    [
        'Graham Chapman',
        [
            'Michael Palin',
            'John Cleese',
            'Terry Cilliam',
            'Eric Idle',
            'Terry Jones'
        ]
    ]
]
#print(isinstance(movies,list))  #True
#print(isinstance(len(movies),int))  #True
# for item in movies:
#     if isinstance(item,list):
#         for subItem in item:
#             if isinstance(subItem,list):
#                 for minItem in subItem:
#                     print(minItem)
#             else:
#                 print(subItem)
#     else:
#         print(item)

# #递归函数
# def print_lol(_list):
#     for item in _list:
#         if isinstance(item,list):
#             print_lol(item)  #如果所处理的列表项本身是一个列表，则调用函数
#         else:
#             print(item)

nester.print_lol(movies,True)

"""
    模块是一个包含Python代码的文本文件

    发布工具允许将模块转换为可共享的包

    setup.py程序提供了模块的元数据，用来构建、安装和
    上传打包的发布

    使用import语句可以将模块导入到其他程序中

    python中的各个模块提供了自己的命名空间，使用module.function()
    形式调用模块的函数时，要用命名空间名限定函数

    使用import语句的from module import function形式可以从一个
    模块将函数专门导入到当前命名空间

    使用#可以注释一行代码，或者为程序增加一个简短的单行注释

    内置函数(built-in functions,BIF)有自己的命名空间，名为__builtins__
    这会自动包含在每一个python程序中

    range()可以与for结合使用，从而迭代固定次数

    包含end=''作为print()的一个参数会关闭其默认行为(即在输入中
    自动包含换行)

    如果为函数参数提供一个缺省值，这个函数参数就是可选的
"""