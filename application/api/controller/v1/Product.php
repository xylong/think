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
        (new ProductValidate('recent'))->_check();

        $products = ProductModel::getMostRecent($count);

        if ($products->isEmpty()) {
            throw new ProductException;
        }

        $products = $products->hidden(['summary']);

        return $products;
    }

    public function getAllInCategory($category_id)
    {
    	(new ProductValidate('category'))->_check();
    	
    	$products = ProductModel::getProductsByCategoryId($category_id);

    	if ($products->isEmpty()) {
    		throw new ProductException([
    			'code' => 404,
    			'errorCode' => 20000,
    			'msg' => '指定类目商品不存在'
    		]);
    	}

    	$products = $products->hidden(['summary']);

    	return $products;
    }
}
