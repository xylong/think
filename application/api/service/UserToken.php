<?php

namespace app\api\service;

use think\facade\Config;
use think\facade\Cache;
use app\lib\exception\WeChatException;
use app\lib\exception\TokenException;
use app\common\model\User;

/**
 * 
 * @authors xuyunlong (416319808@qq.com)
 * @date    2018-05-20 21:14:22
 * @version $Id$
 */

class UserToken extends Token
{
	protected $code;

	protected $wxAppId;

	protected $wxAppSecret;

	protected $wxLoginUrl;

	public function __construct($code='')
	{
		$this->code = $code;
		$this->wxAppId = Config::get('wx.app_id');
		$this->wxAppSecret = Config::get('wx.app_secret');
		$this->wxLoginUrl = sprintf(Config::get('wx.login_url'), $this->wxAppId, $this->wxAppSecret, $this->code);
	}

	public function get()
	{
		$response = curl_get($this->wxLoginUrl);
		$wx = json_decode($response, true);

		if (empty($wx)) {
			throw new Exception("there is an error in get session_key adn openID, internal error of weixin");
		} else {
			$loginFail = array_key_exists('errcode', $wx);

			if ($loginFail) {
				$this->processLoginError($wx);
			} else {
				return $this->grantToken($wx);
			}
		}
	}

	private function grantToken($wx)
	{
		$openid = $wx['openid'];
		$user = User::getOpenID($openid);

		if ($user) {
			$uid = $user->id;
		} else {
			$uid = $this->createUser($openid);
		}

		$cacheValue = $this->prepareCacheValue($wx, $uid);
		$token = $this->saveCache($cacheValue);
		return $token;
	}

	/**
	 * 缓存用户信息
	 * @param  array $cacheValue 缓存值
	 * @return string            token令牌
	 */
	private function saveCache($cacheValue)
	{
		$key = self::generateToken();
		$value = json_encode(cacheValue);
		$expire_in = Config::get('storage.token_expire_in');

		$request = Cache::set($key, $value, $expire_in);
		if (!$request) {
			throw TokenException([
				'msg' => '服务器缓存异常',
				'errorCode' => 10005
			]);
		}
		
		return $key;
	}

	/**
	 * 准备缓存值
	 * @param  string $wx  微信返回值
	 * @param  integer $uid 用户主键🆔
	 * @return array      要缓存的数组
	 */
	private function prepareCacheValue($wx, $uid)
	{
		$cacheValue = $wx;
		$cacheValue['uid'] = $uid; 
		$cacheValue['scope'] = 16;
		return $cacheValue;
	}

	/**
	 * 创建用户
	 * @param  string $openid 微信openid
	 * @return integer        用户主键🆔
	 */
	private function createUser($openid)
	{
		$user = User::create([
			'openid' => $openid
		]);

		return $user->id;
	}

	private function processLoginError($wx)
	{
		throw new WeChatException([
			'msg' => $wx['errmsg'],
			'errorCode' => $wx['errcode']
		]);
		
	}
}