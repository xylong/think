<?php

namespace app\common\model;

/**
 * 专题
 */
class Theme extends BaseModel
{
    public function topicImg()
    {
        return $this->belongsTo('Image', 'topic_img_id', 'id ');
    }

    public function headImg()
    {
    	return $this->belongsTo('Image', 'head_img_id', 'id ');
    }
}
