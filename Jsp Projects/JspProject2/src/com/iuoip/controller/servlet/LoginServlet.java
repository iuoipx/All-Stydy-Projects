package com.iuoip.controller.servlet;

import com.iuoip.domain.User;
import com.iuoip.service.impl.UserServiceImpl;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
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

            UserServiceImpl userService = new UserServiceImpl();
            User login = userService.login(loginUser);

            if (login != null) {
                request.getSession().setAttribute("username", username);
                request.getRequestDispatcher("jsp/success.jsp").forward(request, response);
            } else {
                request.getSession().setAttribute("loginflag", "用户名或密码错误");
                request.getRequestDispatcher("index.jsp").forward(request, response);
            }
        } else {
            System.out.println("验证码错误");
            request.getSession().setAttribute("loginflag", "验证码错误");
            request.getRequestDispatcher("index.jsp").forward(request, response);

        }
    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {


    }
}
