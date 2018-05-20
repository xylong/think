<?php
/**
 * 微信配置
 * @authors xuyunlong (416319808@qq.com)
 * @date    2018-05-20 21:41:28
 * @version $Id$
 */

return [
	'app_id' => '',
	'app_secret' => '',
	'login_url' => "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
];
