<?php
namespace app\api\validate;

use think\Validate;
use think\facade\Request;
use app\lib\exception\ParameterException;

/**
 * 
 * @authors xuyunlong (416319808@qq.com)
 * @date    2018-05-12 15:28:45
 * @version $Id$
 */
class BaseValidate extends Validate
{
	public $request;

	protected $rule = [
        'id'  =>  ['require', 'checkId'],
    ];

    protected $message  =   [
        'id.require' => 'id必须',
        'id.checkId'   => 'id必须是正整数',
    ];

    // 验证
	public function _check()
	{
		$this->request = Request::param();

		if (!$this->check($this->request)) {
			throw new ParameterException([
				'msg' => $this->error
			]);
			die();
		} else {
			return true;
		}
	}

	/**
	 * 验证主键id
	 * @param  integer $var id
	 * @return boolean
	 */
	protected function checkId($var)
	{
		return isPositiveInt($var);
	}
}
