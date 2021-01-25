import sys
#递归函数
def print_lol(_list,indent=False,level=0,fn=sys.stdout):
    for item in _list:
        if isinstance(item,list):
            print_lol(item,indent,level+1,fn)  #如果所处理的列表项本身是一个列表，则调用函数
        else:
            if indent:
                # for tab_stop in range(level):
                #     print('\t',end='')
                print('\t'*level,end='',file=fn)
            print(item,file=fn)