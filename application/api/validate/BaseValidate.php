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

    // 场景
    protected $scene_name = '';

    // 验证
	public function _check()
	{
		$this->request = Request::param();

		if (!$this->scene($this->scene_name)->batch()->check($this->request)) {
			throw new ParameterException([
				'msg' => $this->error
			]);
			exit;
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

	protected function isNotEmpty($value, $rule, $data='', $field='')
	{
		return !empty($value);
	}
}
