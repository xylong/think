<?php

namespace app\api\controller\v1;

use app\api\validate\IDCollection;
use app\common\model\Theme as ThemeModel;
use app\lib\exception\ThemeException;

class Theme extends Base
{
    /**
     * @url /theme?ids=id1,id2,id3...
     * @return [type] [description]
     */
    public function getSimpleList($ids='')
    {
        (new IDCollection)->_check();

        $res = ThemeModel::with(['topicImg', 'headImg'])->select($ids);

        // debug 直接判断!$res判断不了(空数组或对象都返回true)
        if (!count($res)) {
            throw new ThemeException;
        }

        return $res;
    }

    public function getComplexOne($id)
    {
        echo $id;
    }
}
