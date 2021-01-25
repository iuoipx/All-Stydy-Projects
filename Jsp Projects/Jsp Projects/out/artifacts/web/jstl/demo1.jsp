<%@ page import="com.iuoip.domain.User" %><%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/27 0027
  Time: 下午 2:35
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
    <title>Title</title>
</head>
<body>

<%
    request.setAttribute("msg", "<<hello>>");
%>
<%--el表达式不解析尖括号--%>
el:${msg}<br> <%--el:<>--%>
jstl:<c:out value="${msg}"></c:out><br>

<%--在某个域中存值--%>
<c:set var="test" value="shinobu" scope="request"></c:set>
变量set赋值:<c:out value="${test}"></c:out><br>

<%
    User user = new User();
    request.setAttribute("user", user);
%>
<%--为某个对象属性赋值--%>
<c:set target="${user}" property="username" value="gxsb"></c:set>
user.username set赋值:<c:out value="${user.username}"></c:out><br>

<c:remove var="msg"></c:remove>
remove msg 后：<c:out value="${msg}"></c:out><br>

</body>
</html>
