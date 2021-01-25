package com.iuoip.servlet;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

@WebServlet("/demo3")
public class ServletDemo3 extends HttpServlet {
    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
        System.out.println("demo3 do get...");
        String username = req.getParameter("username");
        String password = req.getParameter("password");
        System.out.println(username+":"+password);
//        System.out.println(req.getContextPath()); // /web
//        System.out.println(req.getRequestURI()); // /web/demo3
        req.getRequestDispatcher("ServletDemo4").forward(req,resp);
    }

    @Override
    protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
        req.setCharacterEncoding("utf-8"); //post乱码设置编码
        System.out.println("demo3 do post...");
        String username = req.getParameter("username");
        String password = req.getParameter("password");
        System.out.println(username+":"+password);

        //转发
        req.getRequestDispatcher("ServletDemo4").forward(req,resp);
    }
}
