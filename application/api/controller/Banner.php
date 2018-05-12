<?php
namespace app\api\controller;

use think\Controller;
use think\Validate;
use think\facade\Request;
use app\common\model\Banner as BannerModel;
use app\lib\exception\BannerMissException;
/**
* Banner
*/
class Banner extends Controller
{
	
	/**
	 * 获取banner信息
	 * @url 	/banner/:id
	 * @http	get
	 * @param  integer $id id
	 * @return json
	 */
	public function getBanner($id=0)
	{
		$result = $this->validate([
			'id' => $id
		], 'app\api\validate\Banner');

		if (true !== $result) {
            exit($result);
        }
 
        $banner = BannerModel::get($id, true);
        if (!$banner) {
        	// throw new BannerMissException();
        	throw new \Exception("aaaa", 1);
        	
        }
		return $banner;
	}
}
