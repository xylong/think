<?php
namespace app\api\controller;

use think\Controller;
use think\Validate;
use think\facade\Request;
use app\common\model\Banner as BannerModel;
use app\api\validate\Banner as BannerValidate;
use app\lib\exception\BannerMissException;
/**
* Banner
*/
class Banner extends Controller
{
	protected function initialize()
    {
        (new BannerValidate)->_check();
    }
	
	/**
	 * 获取banner信息
	 * @url 	/banner/:id
	 * @http	get
	 * @param  integer $id id
	 * @return json
	 */
	public function getBanner($id=0)
	{
        $banner = BannerModel::get($id, true);
        if (!$banner) {
        	throw new BannerMissException();
        }
		return $banner;
	}
}
