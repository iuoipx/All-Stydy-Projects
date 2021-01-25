package com.iuoip.service.impl;

import com.iuoip.dao.UserDao;
import com.iuoip.dao.impl.UserDaoImpl;
import com.iuoip.domain.User;
import com.iuoip.service.UserService;

public class UserServiceImpl implements UserService {

    private UserDao userDao = new UserDaoImpl();

    @Override
    public User login(User user) {
        return userDao.findUser(user);
    }

    @Override
    public int register(User user) {
        return userDao.addUser(user);
    }
}
