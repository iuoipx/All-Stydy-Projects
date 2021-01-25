package com.iuoip.dao;

import com.iuoip.domain.User;

public interface UserDao {
    User findUser(User user);

    int addUser(User user);
}
