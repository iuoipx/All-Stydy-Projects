<%@ page import="java.util.ArrayList" %>
<%@ page import="java.util.List" %>
<%@ page import="com.iuoip.domain.User" %><%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/27 0027
  Time: 下午 3:24
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
    <title>Title</title>
</head>
<body>

<c:forEach var="i" begin="1" end="10" step="2">
    <c:out value="gxsb${i}"></c:out>
</c:forEach>
<br>

<%
    List list = new ArrayList();
    list.add(new User("gxsb", "1SB"));
    list.add(new User("gxsb", "2SB"));
    list.add(new User("gxsb", "3SB"));
    list.add(new User("gxsb", "4SB"));
    list.add(new User("gxsb", "5SB"));
    request.setAttribute("list", list);
%>
<%--遍历arraylist--%>
<c:forEach items="${list}" var="user" step="1">
    ${user.username}：${user.password}<br>
</c:forEach>

<%--切割字符串--%>
<c:forTokens items="g,x,s,b" delims="," var="i">
    ${i}
</c:forTokens>

</body>
</html>
