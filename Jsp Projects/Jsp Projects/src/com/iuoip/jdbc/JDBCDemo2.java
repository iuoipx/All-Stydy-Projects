package com.iuoip.jdbc;

import com.iuoip.domain.User;
import com.mchange.v2.c3p0.ComboPooledDataSource;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;

@WebServlet("/JDBCDemo2")
public class JDBCDemo2 extends HttpServlet {
    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        ComboPooledDataSource cpds = new ComboPooledDataSource();

        Connection connection = null;
        PreparedStatement ps = null;
        ResultSet resultSet = null;

        List users = new ArrayList<User>();
        try {
            connection = cpds.getConnection();
            String sql = "select * from user";
            ps = connection.prepareStatement(sql);
            resultSet = ps.executeQuery();

            while (resultSet.next()) {
                User user = new User();
                user.setId(resultSet.getString("id"));
                user.setUsername(resultSet.getString("username"));
                user.setPassword(resultSet.getString("password"));
                user.setSex(resultSet.getString("sex"));
                user.setAge(resultSet.getInt("age"));
                user.setPhone(resultSet.getString("phone"));
                user.setEmail(resultSet.getString("email"));
                users.add(user);
            }
            request.setAttribute("userlist", users);
            request.getRequestDispatcher("jdbc/usertable.jsp").forward(request, response);
        } catch (
                SQLException e) {
            e.printStackTrace();
        }
    }
}
