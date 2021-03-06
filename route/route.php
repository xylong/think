<?php
// +----------------------------------------------------------------------
// | ThinkPHP [ WE CAN DO IT JUST THINK ]
// +----------------------------------------------------------------------
// | Copyright (c) 2006~2018 http://thinkphp.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: liu21st <liu21st@gmail.com>
// +----------------------------------------------------------------------

Route::get('think', function () {
    return 'hello,ThinkPHP5!';
});

Route::get('hello/:name', 'index/hello');

// 获取banner
Route::get('api/:version/banner/:id/[:type]', 'api/:version.Banner/getBanner');

// 获取主题列表
Route::get('api/:version/theme', 'api/:version.Theme/getSimpleList');

// 获取完整一个主题
Route::get('api/:version/theme/:id', 'api/:version.Theme/getComplexOne');

// 获取最新商品
Route::get('api/:version/new_product/[:count]', 'api/:version.Product/getRecent');

// 获取分类商品
Route::get('api/:version/category_product/:category_id', 'api/:version.Product/getAllInCategory');

// 获取分类
Route::get('api/:version/category/all', 'api/:version.Category/getAllCategories');

// 获取token
Route::post('api/:version/token/user', 'api/:version.Token/getToken');
