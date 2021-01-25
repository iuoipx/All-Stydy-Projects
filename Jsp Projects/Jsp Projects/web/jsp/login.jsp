<%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/17 0017
  Time: 下午 4:52
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%
//    获取绝对路径
    String contextPath = request.getContextPath();
    System.out.println(contextPath);
%>
<html>
<head>
    <title>Title</title>
    <script>
        window.onload = function () {
            document.getElementById('checkCodeImg').onclick = function () {
                this.src = "<%=contextPath%>/checkCodeServlet?time=" + new Date().getTime();
            }
        }
    </script>
</head>
<body>
<%--<%=request.getAttribute("checkCodeMessage") == null ? "" : request.getAttribute("checkCodeMessage")%>--%>
<%--<%=request.getAttribute("failMessage") == null ? "" : request.getAttribute("failMessage")%>--%>

${requestScope.checkCodeMessage}
${requestScope.failMessage}

<form action="<%=contextPath%>/JspLoginServlet" method="POST">
    username：&nbsp;&nbsp;<input type="text" name="username">
    <br>
    password： &nbsp;<input type="password" name="password">
    <br>
    checkcode：&nbsp;<input type="text" name="checkCode"><br>
    <img src="<%=contextPath%>/checkCodeServlet" id="checkCodeImg"><br>
    <input type="submit" value="submit">
</form>

</body>
</html>
