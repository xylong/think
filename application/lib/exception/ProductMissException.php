<?php
namespace app\lib\exception;

/**
* 
*/
class ProductMissException extends BaseException
{
	
	public $code = 404;

	public $errorCode = 40000;
	
	public $msg = '请求的Product不存在';
}