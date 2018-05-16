<?php
namespace app\common\model;

use think\Model;

class BannerItem extends Model
{
	protected $hidden = ['id', 'img_id', 'banner_id', 'update_time', 'delete_time'];

	protected static function init()
    {
        //TODO:初始化内容
    }

    public function img()
    {
    	return $this->belongsTo('Image', 'img_id', 'id');
    }

}
