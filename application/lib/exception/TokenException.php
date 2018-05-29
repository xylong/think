<?php
namespace app\lib\exception;

/**
* 
*/
class TokenException extends BaseException
{
	
	public $code = 401;

	public $errorCode = 10001;
	
	public $msg = 'token过期或无效token';
}