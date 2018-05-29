<?php

namespace app\common\model;


class User extends BaseModel
{
    public static function getOpenID($openid)
    {
    	return self::where('openid', $openid)->find();
    }
}
