<?php

namespace app\api\controller\v1;

use app\common\model\Product as ProductModel;
use app\api\validate\Product as ProductValidate;
use app\lib\exception\ProductException;

/**
 * 商品
 */
class Product extends Base
{
    public function getRecent($count=15)
    {
        (new ProductValidate)->_check();

        $products = ProductModel::getMostRecent($count);

        if ($products->isEmpty()) {
            throw new ProductException;
        }

        return $products;
    }
}
