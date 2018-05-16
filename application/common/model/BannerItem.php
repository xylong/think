<?php
namespace app\common\model;

use think\Model;

class BannerItem extends Model
{
	protected static function init()
    {
        //TODO:初始化内容
    }

    public function img()
    {
    	return $this->belongsTo('Image', 'img_id', 'id');
    }

}
