<?php
namespace app\lib\exception;

/**
* 
*/
class ThemeException extends BaseException
{
	
	public $code = 404;

	public $errorCode = 30000;
	
	public $msg = '请求的主题不存在';
}