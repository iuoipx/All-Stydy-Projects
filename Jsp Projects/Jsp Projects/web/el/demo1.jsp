<%@ page import="com.iuoip.domain.User" %>
<%@ page import="java.util.ArrayList" %>
<%@ page import="java.util.List" %>
<%@ page import="java.util.HashMap" %>
<%@ page import="java.util.Map" %><%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/11/24 0024
  Time: 下午 4:32
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<html>
<head>
    <title>Title</title>
</head>
<body>
<%
    User user = null;
    user = new User("gxsb", "123456");
    session.setAttribute("user1", user);
%>

<%
    user = (User) session.getAttribute("user1");
    out.print(user.getUsername() + ":" + user.getPassword());
%>
<br>
<h2>EL表达式取值</h2>
<h3>对象取值</h3>
${user1.username}:${user1.password}

<%
    request.setAttribute("str", null);
    request.setAttribute("str1", "");

    List arrayList = new ArrayList();
    arrayList.add("");
    arrayList.add(null);
    arrayList.add("listValue");
    arrayList.add(user);
    request.setAttribute("list", arrayList);

    Map hashMap = new HashMap();
    hashMap.put("str", "mapValue");
    hashMap.put("user", user);
    request.setAttribute("map", hashMap);
%>

<h3>字符串取值</h3>
<%--依次从最小的域中查找是否有该键对应的值--%>
${str}<br>
${str1}<br>

<%--通过域名.键名取值--%>
${requestScope.str}
${requestScope.str1}

${empty str}<br>
${empty str}<br>

<h3>数组取值</h3>
${list[0]}<br>
${list[1]}<br>
${list[2]}<br>
${list[3].username}:${list[3].password}<br>

通过域名.键名取值<br>
${requestScope.list[2]}

<h3>map取值</h3>
${map.str}<br>
${map.user.username}:${map.user.password}<br>
<%--本质上通过User类中的setter getter方法名获取--%>
${map.user.username1}<br>

</body>
</html>
