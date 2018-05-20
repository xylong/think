<?php

namespace app\common\model;


/**
 * 分类
 */
class Category extends BaseModel
{
	protected $hidden = ['delete_time', 'update_time'];

    public function img()
    {
    	return $this->belongsTo('Image', 'topic_img_id', 'id');
    }
}
