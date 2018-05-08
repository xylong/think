<?php
namespace app\lib\exception;

use think\exception\Handle;
use think\exception\HttpException;
use think\exception\ValidateException;
use think\facade\Request;
/**
* 
*/
class ExceptionHandler extends Handle
{
	private $code = 500;

	private $errorCode = 999;

	private $msg = '服务器内部错误';

	public function render(Exception $e)
	{
		if ($e instanceof BaseException) {
			// 自定义异常
			$this->code = $e->code;
			$this->errorCode = $e->errorCode;
			$this->msg = $e->msg;
		}

		return json([
			'msg' => $this->msg,
			'error_code' => $this->errorCode,
			'request_url' => Request::url()
		], $this->code);
	}
}