package com.iuoip.dao.impl;

import com.iuoip.dao.UserDao;
import com.iuoip.domain.User;
import com.iuoip.utils.DBUtils;
import org.springframework.jdbc.core.BeanPropertyRowMapper;
import org.springframework.jdbc.core.JdbcTemplate;

public class UserDaoImpl implements UserDao {
    private JdbcTemplate jdbcTemplate = new JdbcTemplate(DBUtils.getDataSource());

    @Override
    public User findUser(User user) {
        try {
            String sql = "select * from user where username=? and password=?";
            User result = jdbcTemplate.queryForObject(sql, new BeanPropertyRowMapper<User>(User.class), user.getUsername(), user.getPassword());
            return result;
        } catch (Exception e) {
            return null;
        }
    }

    @Override
    public int addUser(User user) {
        try {
            String sql = "insert user value(?,?,?,?,?,?,?)";
            int result = jdbcTemplate.update(sql, user.getId(),
                    user.getUsername(), user.getPassword(), user.getSex(), user.getAge(), user.getPhone(), user.getEmail());
            return result;
        } catch (Exception e) {
            return 0;
        }
    }
}
