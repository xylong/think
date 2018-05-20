<?php
namespace app\common\model;


class Product extends BaseModel
{
	protected $hidden = ['pivot', 'create_time', 'update_time', 'delete_time', 'from', 'category_id'];

	public function getMainImgUrlAttr($value, $data)
    {
    	return $this->prefixImgUrl($value, $data);
    }
}