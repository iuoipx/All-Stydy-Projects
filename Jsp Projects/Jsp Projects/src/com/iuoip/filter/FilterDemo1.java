package com.iuoip.filter;

import javax.servlet.*;
import javax.servlet.annotation.WebFilter;
import java.io.IOException;

//@WebFilter("/*")
//@WebFilter("/index.jsp")
//@WebFilter("*.jsp")
@WebFilter("/jstl/*")
public class FilterDemo1 implements Filter {
    public void destroy() {
        System.out.println("do filter1 destroy...");
    }

    public void doFilter(ServletRequest req, ServletResponse resp, FilterChain chain) throws ServletException, IOException {
        System.out.println("do filterDemo1...");
        chain.doFilter(req, resp);
        System.out.println("do filter1 after...");
    }

    public void init(FilterConfig config) throws ServletException {
        System.out.println("do filter1 init...");
    }

}
