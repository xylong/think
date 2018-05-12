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
}