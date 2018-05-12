<?php
namespace app\lib\exception;

/**
 * 参数异常
 */
class ParameterException extends BaseException
{
	public $code = 400;

	public $errorCode = 30000;
	
	public $msg = '参数错误';
}