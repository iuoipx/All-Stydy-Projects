package com.iuoip.servlet;

import com.iuoip.domain.User;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

@WebServlet("/LoginServlet")
public class LoginServlet extends HttpServlet {
    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        request.setCharacterEncoding("utf-8");
        String username = request.getParameter("username");
        String password = request.getParameter("password");
        String checkCode = request.getParameter("checkCode"); //用户输入的验证码
        String checkCode_session = (String) request.getSession().getAttribute("checkCode_session");

        if (checkCode_session.equalsIgnoreCase(checkCode)) {
            User loginUser = new User();
            loginUser.setUsername(username);
            loginUser.setPassword(password);

            User initUser = new User("tom", "123456");

            boolean flag = initUser.equals(loginUser);
            if (flag) {
                Cookie cookie = new Cookie(username, password);
                response.addCookie(cookie);
                request.getRequestDispatcher("SuccessServlet").forward(request, response);
            } else {
                request.getRequestDispatcher("FailedServlet").forward(request, response);
            }
        } else {
            System.out.println("验证码错误");
            response.sendRedirect("login.html");
        }

    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

    }
}
