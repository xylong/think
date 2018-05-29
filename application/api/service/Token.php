<?php

namespace app\api\service;

use think\facade\Config;

/**
 * 
 * @authors xuyunlong (416319808@qq.com)
 * @date    2018-05-20 21:14:22
 * @version $Id$
 */

class Token
{
	// 生成token
	public static function generateToken()
	{
		$randChars = getRandChar();

		$timestamp = $_SERVER['REQUEST_TIME_FLOAT'];

		$salt = Config::get('secure.token_salt');

		return md5($randChars . $timestamp . $salt);
	}
}