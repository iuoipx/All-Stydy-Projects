package com.iuoip.jdbc;

import com.iuoip.domain.User;
import com.iuoip.utils.DBUtils;
import org.springframework.jdbc.core.BeanPropertyRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.List;

@WebServlet("/SpringJDBC")
public class SpringJDBC extends HttpServlet {
    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        JdbcTemplate jdbcTemplate = new JdbcTemplate(DBUtils.getDataSource());
        String sql = "select * from user";
        List<User> userList = jdbcTemplate.query(sql,new BeanPropertyRowMapper());
        request.setAttribute("userlist", userList);
        request.getRequestDispatcher("jdbc/usertable.jsp").forward(request, response);
    }
}
