<?php

namespace app\api\controller\v1;

use app\common\model\Theme as ThemeModel;
use app\api\validate\Theme as ThemeValidate;
use app\api\validate\IDCollection;
use app\lib\exception\ThemeException;

class Theme extends Base
{
    /**
     * @url /theme?ids=id1,id2,id3...
     * @return object
     */
    public function getSimpleList($ids='')
    {
        (new IDCollection)->_check();

        $res = ThemeModel::with(['topicImg', 'headImg'])->select($ids);

        if ($res->isEmpty()) {
            throw new ThemeException;
        }

        return $res;
    }

    /**
     * @url /theme/:id
     * @param  integer $id 主题🆔
     * @return object
     */
    public function getComplexOne($id)
    {
        (new ThemeValidate)->_check();
        
        $theme = ThemeModel::getThemeWithProducts($id);

        if (!$theme) {
            throw new ThemeException;
        }

        return $theme;
    }
}
