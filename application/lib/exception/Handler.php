<?php
namespace app\lib\exception;

use Exception;
use think\exception\Handle;
use think\exception\HttpException;
use think\exception\ValidateException;
use think\facade\Request;
use think\facade\Log;
use think\facade\Config;
/**
 * 异常处理
 */
class Handler extends Handle
{
	public $code = 500;

	public $errorCode = 999;

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

        // 判断是否是自定义的异常
        if ($e instanceof BaseException) {
        	$this->code 		= $e->code;
        	$this->errorCode 	= $e->errorCode;
        	$this->msg 			= $e->msg;
        } else {
        	// 判断是默认显示还是json
        	if (Config::get('app_debug')) {
        		return parent::render($e);
        	}
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
     * 手动记录日志
     */
    private function recodeErrorLog(Exception $e)
    {
    	Log::init([
    		'type' => 'File',
    		'path' => LOG_PATH,
    		'level'=> ['error']
    	]);

    	Log::record($e->getMessage(), 'error');
    }

}
