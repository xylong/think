<?php

namespace app\api\service;

use think\facade\Config;

/**
 * 
 * @authors xuyunlong (416319808@qq.com)
 * @date    2018-05-20 21:14:22
 * @version $Id$
 */

class UserToken
{
	protected $code;

	protected $wxAppId;

	protected $wxAppSecret;

	protected $wxLoginUrl;

	public function __construct($code)
	{
		$this->code = $code;
		$this->wxAppId = Config::get('wx.app_id');
		$this->wxAppSecret = Config::get('wx.app_secret');
		$this->wxLoginUrl = sprintf(Config::get('wx.login_url'), $this->wxAppId, $this->wxAppSecret) ;
	}

	public function get($code)
	{
		$response = curl_get($this->wxLoginUrl);
		$wx = json_decode($response, true);

		if (empty($wx)) {
			throw new Exception("there is an error in get session_key adn openID, internal error of weixin");
		}
	}
}