<?php

namespace app\common\model;

use think\Model;
use think\facade\Config;


class Image extends Model
{
    protected $hidden = ['id', 'from', 'update_time', 'delete_time'];

    public function getUrlAttr($value, $data)
    {
    	return $data['from'] === 1 ? Config::get('storage.img_prefix') . $value : $value;
    }
}
