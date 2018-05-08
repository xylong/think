<?php
namespace app\lib\exception;

use think\Exception;
/**
* restful
*/
class BaseException extends Exception
{
	
	// http 状态码
	public $code;

	// 自定义错误码
	public $errorCode;
	
	// 错误信息
	public $msg;
}