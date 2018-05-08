<?php
namespace app\api\validate;

use think\Validate;

class Banner extends Validate
{
    protected $rule = [
        'id'  =>  'require|number',
    ];

    protected $message  =   [
        'name.require' => 'id必须',
        'id.number'   => 'id必须是数字',
    ];

}