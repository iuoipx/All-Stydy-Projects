package com.iuoip.filter;

import javax.servlet.*;
import javax.servlet.annotation.WebFilter;
import java.io.IOException;

@WebFilter("/*")
//按字符串执行filter
public class FilterDemo2 implements Filter {
    public void destroy() {
    }

    public void doFilter(ServletRequest req, ServletResponse resp, FilterChain chain) throws ServletException, IOException {
        System.out.println("do filterDemo2...");
        chain.doFilter(req, resp);
        System.out.println("do filter2 after...");
    }

    public void init(FilterConfig config) throws ServletException {

    }

}
