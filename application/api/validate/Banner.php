<?php
namespace app\api\validate;

use think\Validate;

class Banner extends BaseValidate
{
	public function __construct()
	{
		parent::__construct();

		// 重载规则
		$this->rule = [
			'id' => ['require', 'checkId'],
			'type' => ['in' => ['new', 'hot']],
		];

		$this->message = [
			'type.in' => 'type类型错误',
			'id.require' => 'id必须',
        	'id.checkId'   => 'id必须是正整数',
		];
	}
}