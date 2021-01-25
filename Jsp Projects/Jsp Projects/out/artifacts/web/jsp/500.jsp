<%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/17 0017
  Time: 下午 5:39
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" isErrorPage="true" language="java" %>
<html>
<head>
    <title>Title</title>
</head>
<body>
服务器异常
<%
    String message = exception.getMessage();
    out.write(message);
%>
</body>
</html>
