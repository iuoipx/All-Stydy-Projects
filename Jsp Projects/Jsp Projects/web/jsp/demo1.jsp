<%@ page import="java.util.Date" %>
<%@ page contentType="text/html;charset=UTF-8" errorPage="500.jsp" language="java" %>
<%@include file="success.jsp" %>
<html>
<head>
    <title>Title</title>
</head>
<body>
<%
    //    int i = 3;
    String name = request.getParameter("name");
    out.print(name);
%>

<%!
    int i = 5;
%>

<%=i%>

<%--<%--%>
<%--    int j = 1 / 0;--%>
<%--%>--%>

<%
    Date date = new Date();
%>
</body>
</html>
