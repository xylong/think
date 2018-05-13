<?php
namespace app\api\controller;

use think\Controller;
use app\common\controller\Common;
use app\common\model\Product as ProductModel;
use app\api\validate\Product as ProductValidate;
use app\lib\exception\ProductMissException;

class Product extends Base
{
	protected function initialize()
    {
        (new ProductValidate)->_check();
        parent::initialize();
    }

    public function index()
    {
        return 'index';
    }

    public function getProduct($id)
    {
    	$product = ProductModel::get($id, true);
    	if (!$product) {
        	throw new ProductMissException();
        }
		return $product;
    }
}