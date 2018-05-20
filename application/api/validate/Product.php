<?php
namespace app\api\validate;

class Product extends BaseValidate
{
	public function __construct()
	{
		parent::__construct();

		$this->rule = [
			'count' => 'checkId|between:1,15'
		];

		$this->message = [
			'count.checkId' => 'count必须是正整数',
			'count.between' => 'count必须在1-15之间'
		];
	}
}