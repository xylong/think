<?php
namespace app\lib\exception;

/**
* 
*/
class WeChatException extends BaseException
{
	
	public $code = 401;

	public $errorCode = 10001;
	
	public $msg = '微信异常';
}