<?php
namespace app\common\model;

class Banner extends BaseModel
{
	protected $hidden = ['update_time', 'delete_time'];

	protected static function init()
    {
        //TODO:初始化内容
    }

    public function items()
    {
    	return $this->hasMany('BannerItem', 'banner_id', 'id');
    }

    public static function getBannerById($id)
    {
    	return self::with(['items', 'items.img'])->find($id);
    }

}
