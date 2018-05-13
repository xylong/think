<?php
namespace app\common\controller;

use think\Controller;
/**
 * 
 * @authors xuyunlong (416319808@qq.com)
 * @date    2018-05-13 20:54:20
 * @version $Id$
 */
class Common extends Controller
{
    protected function initialize()
    {
    	// $this->cross();
    }

    // 跨域
    public function cross()
    {
    	header('Access-Control-Allow-Origin: '.$_SERVER['HTTP_ORIGIN']);
        header('Access-Control-Allow-Credentials: true');
        header('Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS');
        header("Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept, authKey, sessionId");
    }
}