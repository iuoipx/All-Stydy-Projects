import zipfile,os

#打包文件夹内的文件成zip文件
def makeDirZip(get_Dir_path, set_zips_path):
   if(os.path.exists(get_Dir_path)):
       list = set_zips_path.split('\\')
       list.pop()
       str_array = "\\"
       tempZipsPath = str_array.join(list)
       if(os.path.exists(tempZipsPath)):
           f = zipfile.ZipFile(set_zips_path, 'w', zipfile.ZIP_DEFLATED )
           for dirpath, dirnames, filenames in os.walk( get_Dir_path ):
               fpath = dirpath.replace(get_Dir_path, '')
           fpath = fpath and fpath + os.sep or ''
           for filename in filenames:
               f.write(os.path.join(dirpath, filename), fpath + filename)
           f.close()
           print("成功打包目录下的所有文件")
           return True
       else:
           print("打包目录:zip参数错误")
           return False
   else:
       print("打包目录:文件夹参数错误")
       return False

#打包单个文件成zip文件
def makeFileZip(get_File_path, set_zip_path):
    if(os.path.exists(get_File_path)):
        list = set_zip_path.split('\\')
        list.pop()
        str_array = "\\"
        tempZipsPath = str_array.join(list)
        if(os.path.exists(tempZipsPath)):
            f = zipfile.ZipFile(set_zip_path, 'w', zipfile.ZIP_DEFLATED)
            for dirpath, dirnames, filenames in os.walk(get_File_path):
                for filename in filenames:
                    f.write(os.path.join(dirpath, filename), filename)
            f.close()
            print("成功打包指定文件")
            return True
        else:
            print("打包文件:zip参数错误")
            return False
    else:
        print("打包文件:文件参数错误")
        return False


if __name__=='__main__':
    get_Dir_path = "D:\\文档\\dir" #需要打包的文件夹路径
    set_zips_path = "D:\\文档\\outputDir\\Dir.zip" #存放打包后的压缩文件(注意:不能与上述压缩文件夹一样)
    get_File_path = "D:\\文档\\file" #需要打包的文件路径
    set_zip_path = "D:\\文档\\outputFile\\File.zip" #存放打包后的压缩文件
    makeDirZip(get_Dir_path, set_zips_path)
    makeFileZip(get_File_path,set_zip_path)
    

