<?php
namespace app\common\model;

use think\Model;

class Banner extends Model
{
	protected static function init()
    {
        //TODO:初始化内容
    }

    public static function test($id)
    {
    	try {
    		1/0;
    	} catch (\Exception $e) {
    		throw $e;
    	}
    	return 'this is a test';
    }
}