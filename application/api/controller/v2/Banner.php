<?php
namespace app\api\controller\v2;

use app\common\model\Banner as BannerModel;
use app\api\validate\Banner as BannerValidate;
use app\lib\exception\BannerMissException;
/**
* Banner
*/
class Banner extends Base
{
	protected function initialize()
    {
        (new BannerValidate)->_check();
    }
	
	/**
	 * 获取banner信息
	 * @url 	/banner/:id/[:type]
	 * @http	get
	 * @param  integer $id id
	 * @param  string $type 类型
	 * @return json
	 */
	public function getBanner($id, $type='')
	{
        echo 'test version2';
	}
}
