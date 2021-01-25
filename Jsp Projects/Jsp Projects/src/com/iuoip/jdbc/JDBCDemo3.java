package com.iuoip.jdbc;

import com.alibaba.druid.pool.DruidDataSourceFactory;
import com.iuoip.domain.User;
import com.iuoip.utils.JDBCUtils;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.sql.DataSource;
import java.io.IOException;
import java.io.InputStream;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

@WebServlet("/JDBCDemo3")
public class JDBCDemo3 extends HttpServlet {
    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {

    }

    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        Properties properties = new Properties();
        InputStream resourceAsStream = JDBCDemo3.class.getClassLoader().getResourceAsStream("druid.properties");
        properties.load(resourceAsStream);
        Connection connection = null;
        PreparedStatement ps = null;
        ResultSet resultSet = null;
        try {
            DataSource dataSource = DruidDataSourceFactory.createDataSource(properties);
            connection = dataSource.getConnection();
            List users = new ArrayList<User>();
            try {
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
            } catch (Exception e) {
                e.printStackTrace();
            }
        } catch (SQLException e) {
            e.printStackTrace();
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            JDBCUtils.close(resultSet, ps, connection);
        }
    }
}