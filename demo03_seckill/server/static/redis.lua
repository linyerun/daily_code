function seckill(userId)
    local userId = "userId_" .. userId;

    local prodKey = "product_apples";

    local usersKey = "usersInProductApples";

    local userExists = redis.call("sismember", usersKey, userid);

    if tonumber(userExists) == 1 then
        -- 重复抢购
        return 2;
    end

    local num = redis.call("get", prodKey);

    if tonumber(num) == 0 then
        -- 商品数量为0
        return 0;
    else
        redis.call("decr", prodKey);
        redis.call("sadd", usersKey, userId);
    end

    -- 请购成功
    return 1;
end