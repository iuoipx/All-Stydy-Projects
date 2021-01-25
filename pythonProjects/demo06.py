import os
os.chdir('D:\文档\pythonDir\chapter06')

def sanitize(time_string):
    if '-' in time_string:
        splitter='-'
    elif ':' in time_string:
        splitter=':'
    else:
        return time_string
    (mins,secs)=time_string.split(splitter)
    return(mins+'.'+secs)

def openFile(flieName):
    try:
        with open(flieName) as f:
            data=f.readline()
        tempList=data.strip().split(',')
        return({
            'NAME':tempList.pop(0),
            'DOB':tempList.pop(0),
            'TIME':str(sorted(set([sanitize(t) for t in tempList]))[0:3])
        })
    except IOError as err:
        print('File Error'+str(err))
        return None


#sarah=openFile('sarah.txt')
# (sarah_name,sarah_dob)=sarah.pop(0),sarah.pop(0)
# print(sarah_name+"'s fastest time are:"+str(sorted(set([sanitize(t) for t in sarah]))[0:3]))


#sarah=openFile('sarah.txt')
# sarah_data={}
# sarah_data['NAME']=sarah.pop(0)
# sarah_data['DOB']=sarah.pop(0)
# sarah_data['TIME']=sarah
# print(sarah_data['NAME']+"'s fastest time are:"+str(sorted(set([sanitize(t) for t in sarah_data['TIME']]))[0:3]))

sarah=openFile('sarah.txt')
print(sarah['NAME']+"'s fastest time are:"+sarah['TIME'])