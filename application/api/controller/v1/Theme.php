<?php

namespace app\api\controller\v1;

use app\api\validate\IDCollection;

class Theme extends Base
{
    /**
     * @url /theme?ids=id1,id2,id3...
     * @return [type] [description]
     */
    public function getSimpleList($ids='')
    {
        (new IDCollection)->_check();
        echo 'list';
    }
}
