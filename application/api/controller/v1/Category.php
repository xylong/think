<?php

namespace app\api\controller\v1;

use app\common\model\Category as CategoryModel;
use app\lib\exception\CategoryException;

/**
 * 分类
 */
class Category extends Base
{
    public function getAllCategories()
    {
        $categories = CategoryModel::all([], 'img');

        if ($categories->isEmpty()) {
            throw new CategoryException;
        }

        return $categories;
    }
}
