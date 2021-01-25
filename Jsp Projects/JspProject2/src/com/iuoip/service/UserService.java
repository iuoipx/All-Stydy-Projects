package com.iuoip.service;

import com.iuoip.domain.User;

public interface UserService {
    User login(User user);

    int register(User user);
}
