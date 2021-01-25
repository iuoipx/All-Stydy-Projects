package com.iuoip.controller.servlet;

import com.iuoip.domain.User;
import com.iuoip.service.impl.UserServiceImpl;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.net.URLEncoder;

@WebServlet("/RegisterServlet")
public class RegisterServlet extends HttpServlet {
    static int id = 2;

    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        request.setCharacterEncoding("utf-8");
        response.setContentType("text/html;charset=utf-8");
        String username = request.getParameter("username");
        String password = request.getParameter("password");
        String sex = request.getParameter("sex");
        int age = Integer.parseInt(request.getParameter("age"));
        String phone = request.getParameter("phone");
        String email = request.getParameter("email");

        String pid = String.valueOf(++id);

        User registerUser = new User(pid, username, password, sex, age, phone, email);
        UserServiceImpl userService = new UserServiceImpl();
        int register = userService.register(registerUser);

        if (register == 1) {
            response.getWriter().print("<script language='javascript'>alert('注册成功');window.location.href='index.jsp'</script>");
//            request.getRequestDispatcher("index.jsp").forward(request, response);
        } else {
            response.getWriter().print("<script language='javascript'>alert('注册失败');window.location.href='index.jsp'</script>");
        }

    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

    }
}
