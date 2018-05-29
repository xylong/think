<?php
// +----------------------------------------------------------------------
// | ThinkPHP [ WE CAN DO IT JUST THINK ]
// +----------------------------------------------------------------------
// | Copyright (c) 2006-2016 http://thinkphp.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: 流年 <liu21st@gmail.com>
// +----------------------------------------------------------------------

// 应用公共文件

/**
 * 检测是否是正整数
 * @param  integer  $var
 * @return boolean
 */
function isPositiveInt($var)
{
	return is_numeric($var) && is_int(($var + 0)) && ($var + 0) > 0;
}

/**
 * curl
 * @param  string  $url       get请求地址
 * @param  integer &$httpCode 返回状态码
 * @return mixed
 */
function curl_get($url, &$httpCode=0)
{
	$ch = curl_init();
	curl_setopt($ch, CURLOPT_URL, $url);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

	// 不做证书校验
	curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, true);	// linux->true,windows->false
	curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 10);
	$file_contents = curl_exec($ch);
	$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
	curl_close($ch);

	return $file_contents;
}

/**
 * 生成随机字符串
 * @param  integer $length 指定长度
 * @return string
 */
function getRandChar($length=32)
{
	$str = '';
	$strPol = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
	$max = strlen($strPol) - 1;

	for ($i=0; $i < $length; $i++) { 
		$str .= $strPol[rand(0, $max)];
	}

	return $str;
}