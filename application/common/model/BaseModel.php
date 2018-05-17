<?php

namespace app\common\model;

use think\Model;
use think\facade\Config;


class BaseModel extends Model
{
    protected function prefixImgUrl($value, $data)
    {
    	return $data['from'] === 1 ? Config::get('storage.img_prefix') . $value : $value;
    }
}
