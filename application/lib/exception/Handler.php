<?php
namespace app\lib\exception;

use Exception;
use think\exception\Handle;
use think\exception\HttpException;
use think\exception\ValidateException;
use think\facade\Request;
use think\facade\Log;
/**
 * 异常处理
 */
class Handler extends Handle
{
	public $code = 500;

	public $errorCode = 10000;

	public $msg = '服务器内部错误';

	public function render(Exception $e)
    {
        // 参数验证错误
        if ($e instanceof ValidateException) {
            return json($e->getError(), 422);
        }

        // 请求异常
        if ($e instanceof HttpException && request()->isAjax()) {
            return response($e->getMessage(), $e->getStatusCode());
        }

        // 其他错误交给系统处理
        // return parent::render($e);

        if ($e instanceof BaseException) {
        	$this->code 		= $e->code;
        	$this->errorCode 	= $e->errorCode;
        	$this->msg 			= $e->msg;
        } else {
        	$this->recodeErrorLog($e);
        }
        
        $result = [
        	'error_code'	=> $this->errorCode,
        	'msg'			=> $this->msg,
        	'request_url'	=> Request::url()
        ];

        return json($result, $this->code);
    }

    /**
     * 记录异常日志
     */
    private function recodeErrorLog(Exception $e)
    {echo $e->getMessage();exit();
    	Log::record($e->getMessage(), 'error');
    }

}
