<%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/27 0027
  Time: 下午 2:58
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
    int age = 20;
    if (age > 18)
        out.print("gxsb");
    else
        out.print("isn't");
%>

<%
    if (age > 18) {
%>
<h1>gxsb</h1>
<%
} else {
%>
<h1>isn't</h1>
<%
    }
%>

<c:set var="age1" value="21" scope="request"></c:set>
<c:if test="${age1>18}">
    <h2>gxsb</h2>
</c:if>

<c:set var="score" value="50" scope="request"></c:set>
<c:choose>
    <c:when test="${score>90}">
        <h3>优秀</h3>
    </c:when>
    <c:when test="${score>70}">
        <h3>良好</h3>
    </c:when>
    <c:when test="${score>=60}">
        <h3>及格</h3>
    </c:when>
    <c:otherwise>
        <h3>gxsb</h3>
    </c:otherwise>
</c:choose>

</body>
</html>
