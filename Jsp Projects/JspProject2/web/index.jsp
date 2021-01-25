<%--
  Created by IntelliJ IDEA.
  User: Administrator
  Date: 2020/12/4 0004
  Time: 下午 3:37
  To change this template use File | Settings | File Templates.
--%>
<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%
    //    获取绝对路径
    String contextPath = request.getContextPath();
%>
<!DOCTYPE html>
<html lang="zh-CN">

<head>
    <meta charset="utf-8">
    <title>登录注册 </title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.8.2/css/all.min.css">
    <link rel="stylesheet" href="css/dmaku.css">
    <link rel="stylesheet" href="css/bootstrap.min.css">
    <script>
        window.onload = function () {
            document.getElementById('checkCode').onclick = function () {
                this.src = "<%=contextPath%>/CheckCodeServlet?time=" + new Date().getTime();
            }
        }
    </script>
</head>

<body>
<div class="dowebok" id="dowebok">
    <div class="form-container sign-up-container">
        <form action="<%=contextPath%>/RegisterServlet" method="post">
            <h1>注册</h1>
            <input type="text" name="username" placeholder="用户名">
            <input type="password" name="password" placeholder="密码">
            <input type="text" name="sex" placeholder="性别">
            <input type="text" name="age" placeholder="年龄" oninput="value=value.replace(/[^\d]/g,'')">
            <input type="text" name="phone" placeholder="手机" maxlength="11" oninput="value=value.replace(/[^\d]/g,'')">
            <input type="text" name="email" placeholder="email">
            <button>注册</button>
        </form>
    </div>
    <div class="form-container sign-in-container">
        <form action="<%=contextPath%>/LoginServlet" method="post">
            <h1>登录</h1>
            <input type="text" name="username" placeholder="用户名">
            <input type="password" name="password" placeholder="密码">
            <input type="text" name="checkCode" placeholder="验证码">
            <img src="<%=contextPath%>/CheckCodeServlet" title="看不清点击刷新" id="checkCode"/>
            <button>登录</button>
            <div class="alert alert-warning alert-dismissible">
                <strong>${loginflag}</strong>
            </div>
        </form>
    </div>
    <div class="overlay-container">
        <div class="overlay">
            <div class="overlay-panel overlay-left">
                <h1>已有帐号？</h1>
                <p>请使用您的帐号进行登录</p>
                <button class="ghost" id="signIn">登录</button>
            </div>
            <div class="overlay-panel overlay-right">
                <h1>没有帐号？</h1>
                <p>立即注册加入我们，和我们一起开始旅程吧</p>
                <button class="ghost" id="signUp">注册</button>
            </div>
        </div>
    </div>
</div>
<script src="js/dmaku.js"></script>
</body>
</html>
