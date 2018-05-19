<?php
namespace app\api\validate;

class IDCollection extends BaseValidate
{
	public function __construct()
	{
		parent::__construct();

		$this->rule = [
			'ids' => 'require|checkIDs'
		];

		$this->message = [
			'ids' => 'ids参数必须是以逗号分隔的多个正整数'
		];
	}

	protected function checkIDs($value)
	{
		$params = explode(',', $value);

		if (empty($params)) {
			return false;
		}

		foreach ($params as $id ) {
			if (!$this->checkId($id)) {
				return false;
			}
		}

		return true;
	}
}
