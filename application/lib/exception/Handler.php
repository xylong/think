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
        switch (true) {
            // 请求异常
            case $e instanceof HttpException:
                $this->code = $e->getStatusCode();
                $this->errorCode = 400;
                $this->msg = $e->getMessage();
                break;

            // 自定义异常
            case $e instanceof BaseException:
                $this->code         = $e->code;
                $this->errorCode    = $e->errorCode;
                $this->msg          = $e->msg;
                break;
            
            // 系统异常
            default:
                // 判断异常是展示还是记录
                if (Config::get('app_debug')) {
                    return parent::render($e);
                }
                $this->recodeErrorLog($e);
                break;
        }
        
        $result = [
            'error_code'    => $this->errorCode,
            'msg'           => $this->msg,
            'request_url'   => Request::url()
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
