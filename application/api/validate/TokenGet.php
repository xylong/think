<?php
namespace app\api\validate;

class TokenGet extends BaseValidate
{
	public function __construct()
	{
		parent::__construct();

		$this->rule = [
			'code' => ['require', 'isNotEmpty'],
		];

		$this->message = [
			'code.require' => 'code必须',
			'code.isNotEmpty' => 'code不能为空'
		];
	}
}