<%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/20 0020
  Time: 下午 2:36
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
<head>
    <title>Title</title>
</head>
<body>

<%
    pageContext.setAttribute("pageContextKey", "pageContextValue");
    request.setAttribute("requestKey", "requestValue");
    session.setAttribute("sessionKey", "sessionValue");
    application.setAttribute("applicationKey", "applicationValue");
%>

pageContextvalue: <%=pageContext.getAttribute("pageContextKey")%> <br>
requestKey: <%=request.getAttribute("requestKey")%> <br>
sessionKey: <%=session.getAttribute("sessionKey")%> <br>
applicationKey: <%=application.getAttribute("applicationKey")%> <br>

<%--<a href="demo3.jsp">to demo3</a>--%>
<%--<%--%>
<%--    request.getRequestDispatcher("demo3.jsp").forward(request, response);--%>
<%--%>--%>
<%--<%--%>
<%--    request.getRequestDispatcher("../JspServlet").forward(request, response);--%>
<%--%>--%>

<%
    out.println("222222222");
%>

<%
    response.getWriter().write("111111111");
%>

</body>
</html>
