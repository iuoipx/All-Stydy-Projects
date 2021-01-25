<%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/20 0020
  Time: 下午 2:42
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
<head>
    <title>Title</title>
</head>
<body>

pageContextvalue: <%=pageContext.getAttribute("pageContextKey")%> <br>
requestKey: <%=request.getAttribute("requestKey")%> <br>
sessionKey: <%=session.getAttribute("sessionKey")%> <br>
applicationKey: <%=application.getAttribute("applicationKey")%> <br>

</body>
</html>
