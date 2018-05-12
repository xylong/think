<?php
namespace app\lib\exception;

use think\Exception;
/**
* restful
*/
class BaseException extends Exception
{
	// http 状态码
	public $code = 400;

	// 自定义错误码
	public $errorCode = 10000;
	
	// 错误信息
	public $msg = '参数错误';

	public function __construct($params=[])
	{
		if (!is_array($params)) {
			return;
		}

		if (array_key_exists('code', $params) && $params['code']) {
			$this->code = $params['code'];
		}

		if (array_key_exists('errorCode', $params) && $params['errorCode']) {
			$this->errorCode = $params['errorCode'];
		}

		if (array_key_exists('msg', $params) && $params['msg']) {
			$this->msg = $params['msg'];
		}
	}
}