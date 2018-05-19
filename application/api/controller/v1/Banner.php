<?php
namespace app\api\controller\v1;

use app\common\model\Banner as BannerModel;
use app\api\validate\Banner as BannerValidate;
use app\lib\exception\BannerMissException;
/**
* Banner
*/
class Banner extends Base
{
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
		(new BannerValidate)->_check();
		
        $banner = BannerModel::getBannerById($id);
        if (!$banner) {
        	throw new BannerMissException();
        }
		return $banner;
	}
}
