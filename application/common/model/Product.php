<?php
namespace app\common\model;


class Product extends BaseModel
{
	protected $hidden = ['pivot', 'create_time', 'update_time', 'delete_time', 'from', 'category_id'];

	public function getMainImgUrlAttr($value, $data)
    {
    	return $this->prefixImgUrl($value, $data);
    }

    public static function getMostRecent($count)
    {
    	return self::limit($count)->order('create_time desc')->select();
    }

    public static function getProductsByCategoryId($category_id)
    {
    	return self::where('category_id', $category_id)->select();
    }
}