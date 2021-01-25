<%@ page import="java.util.List" %>
<%@ page import="java.util.ArrayList" %>
<%@ page import="com.iuoip.domain.User" %><%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/27 0027
  Time: 下午 3:42
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<html>
<head>
    <title>Title</title>
    <style>
        table {
            border: 1px;
            text-align: center;
            width: 300px;
            margin: 0 auto;
        }
        tr th {
            padding: 4px 10px;
        }

        td {

            color: whitesmoke;
            padding: 4px 10px;
        }
    </style>
</head>
<body>
<%
    List<User> list = new ArrayList();
    list.add(new User("gxsb", "1SB"));
    list.add(new User("gxsb", "2SB"));
    list.add(new User("gxsb", "3SB"));
    list.add(new User("gxsb", "4SB"));
    list.add(new User("gxsb", "5SB"));
    request.setAttribute("list", list);
%>

<table>
    <tr>
        <th>编号</th>
        <th>用户名</th>
        <th>密码</th>
    </tr>
    <c:forEach items="${list}" var="user" varStatus="s">
        <c:if test="${s.count%2==0}">
            <tr bgcolor="#87ceeb">
                <td>${s.count}</td>
                <td>${user.username}</td>
                <td>${user.password}</td>
            </tr>
        </c:if>
        <c:if test="${s.count%2!=0}">
            <tr bgcolor="#ff4500">
                <td>${s.count}</td>
                <td>${user.username}</td>
                <td>${user.password}</td>
            </tr>
        </c:if>
    </c:forEach>
</table>
</body>
</html>
