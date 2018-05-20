<?php
namespace app\api\validate;

class Product extends BaseValidate
{
	protected $scene = [
		// 最新
        'recent'  =>  ['count'],
        // 分类
        'category' => ['category_id'],
    ];

	public function __construct($scene='')
	{
		parent::__construct();

		$this->rule = [
			'category_id'	=> ['require', 'checkId'],
			'count' => 'checkId|between:1,15'
		];

		$this->message = [
			'category_id.require'	=> 'id必须',
			'category_id.checkId'	=> 'category_id必须是正整数',
			'count.checkId' => 'count必须是正整数',
			'count.between' => 'count必须在1-15之间'
		];

    	$this->scene_name = $scene;
	}
}