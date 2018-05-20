<?php

namespace app\api\controller\v1;

use think\Controller;
use think\Request;
use app\api\validate\TokenGet;
use app\api\service\UserToken;

class Token extends Controller
{
    public function getToken($code='')
    {
        (new TokenGet)->_check();

        $ut = new UserToken;
        $token = $ut->get($code);
        return $token;
    }
}
