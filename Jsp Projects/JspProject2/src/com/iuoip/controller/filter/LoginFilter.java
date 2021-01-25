package com.iuoip.controller.filter;

import javax.servlet.*;
import javax.servlet.annotation.WebFilter;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

@WebFilter("/jsp/*")
public class LoginFilter implements Filter {
    public void destroy() {
    }

    public void doFilter(ServletRequest req, ServletResponse resp, FilterChain chain) throws ServletException, IOException {
        HttpServletRequest request = (HttpServletRequest) req;
        HttpServletResponse response = (HttpServletResponse) resp;
        response.setContentType("text/html;charset=utf-8");
        Object username = request.getSession().getAttribute("username");
//        String requestURI = request.getRequestURI();
        if (username != null) {
            chain.doFilter(req, resp);
        } else {
            resp.getWriter().print("<script language='javascript'>alert('请先进行登录');window.location.href='../index.jsp'</script>");
        }
    }

    public void init(FilterConfig config) throws ServletException {

    }

}
