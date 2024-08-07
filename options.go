package douban

var DBMovieSortOptions = []Option{
	Option{Value: "", Name: "默认"},
	Option{Value: "U", Name: "综合排序"},
	Option{Value: "T", Name: "近期热度"},
	Option{Value: "R", Name: "首播时间"},
	Option{Value: "S", Name: "高分优先"},
}

var DBMovieTagsOptions = []Option{
	Option{Value: "", Name: "全部"},
	Option{Value: "喜剧", Name: "喜剧"},
	Option{Value: "爱情", Name: "爱情"},
	Option{Value: "动作", Name: "动作"},
	Option{Value: "科幻", Name: "科幻"},
	Option{Value: "动画", Name: "动画"},
	Option{Value: "悬疑", Name: "悬疑"},
	Option{Value: "犯罪", Name: "犯罪"},
	Option{Value: "惊悚", Name: "惊悚"},
	Option{Value: "冒险", Name: "冒险"},
	Option{Value: "奇幻", Name: "奇幻"},
	Option{Value: "恐怖", Name: "恐怖"},
	Option{Value: "战争", Name: "战争"},
	Option{Value: "武侠", Name: "武侠"},
	Option{Value: "灾难", Name: "灾难"},
}

var DBMovieCountryOptions = []Option{
	Option{Value: "", Name: "全部地区"},
	Option{Value: "华语", Name: "华语"},
	Option{Value: "韩国", Name: "韩国"},
	Option{Value: "日本", Name: "日本"},
	Option{Value: "中国大陆", Name: "中国大陆"},
	Option{Value: "美国", Name: "美国"},
	Option{Value: "中国香港", Name: "中国香港"},
	Option{Value: "中国台湾", Name: "中国台湾"},
	Option{Value: "英国", Name: "英国"},
	Option{Value: "法国", Name: "法国"},
	Option{Value: "德国", Name: "德国"},
	Option{Value: "意大利", Name: "意大利"},
	Option{Value: "西班牙", Name: "西班牙"},
	Option{Value: "印度", Name: "印度"},
	Option{Value: "泰国", Name: "泰国"},
	Option{Value: "俄罗斯", Name: "俄罗斯"},
	Option{Value: "加拿大", Name: "加拿大"},
	Option{Value: "澳大利亚", Name: "澳大利亚"},
	Option{Value: "爱尔兰", Name: "爱尔兰"},
	Option{Value: "瑞典", Name: "瑞典"},
	Option{Value: "巴西", Name: "巴西"},
	Option{Value: "丹麦", Name: "丹麦"},
}

var DBYearOptions = []Option{
	{Value: "", Name: "全部年代"},
	{Value: "2024", Name: "2024"},
	{Value: "2023", Name: "2023"},
	{Value: "2022", Name: "2022"},
	{Value: "2021", Name: "2021"},
	{Value: "2020", Name: "2020"},
	{Value: "2019", Name: "2019"},
	{Value: "2010年代", Name: "2010年代"},
	{Value: "2000年代", Name: "2000年代"},
	{Value: "90年代", Name: "90年代"},
	{Value: "80年代", Name: "80年代"},
	{Value: "70年代", Name: "70年代"},
	{Value: "60年代", Name: "60年代"},
	{Value: "更早", Name: "更早"},
}

var DBTvTagOptions = []Option{
	{Value: "", Name: "全部"},
	{Value: "华语", Name: "华语"},
	{Value: "中国大陆", Name: "中国大陆"},
	{Value: "中国香港", Name: "中国香港"},
	{Value: "中国台湾", Name: "中国台湾"},
	{Value: "欧美", Name: "欧美"},
	{Value: "韩国", Name: "韩国"},
	{Value: "日本", Name: "日本"},
	{Value: "印度", Name: "印度"},
	{Value: "泰国", Name: "泰国"},
}

var DBTvTypeOptions = []Option{
	{Value: "", Name: "全部类型"},
	{Value: "电视剧", Name: "全部剧集"},
	{Value: "综艺", Name: "全部综艺"},
	{Value: "动画", Name: "动画"},
	{Value: "爱情", Name: "爱情"},
	{Value: "喜剧", Name: "喜剧"},
	{Value: "悬疑", Name: "悬疑"},
	{Value: "武侠", Name: "武侠"},
	{Value: "古装", Name: "古装"},
	{Value: "家庭", Name: "家庭"},
	{Value: "犯罪", Name: "犯罪"},
	{Value: "科幻", Name: "科幻"},
	{Value: "恐怖", Name: "恐怖"},
	{Value: "历史", Name: "历史"},
	{Value: "战争", Name: "战争"},
	{Value: "动作", Name: "动作"},
	{Value: "冒险", Name: "冒险"},
	{Value: "传记", Name: "传记"},
	{Value: "剧情", Name: "剧情"},
	{Value: "奇幻", Name: "奇幻"},
	{Value: "惊悚", Name: "惊悚"},
	{Value: "灾难", Name: "灾难"},
}

var DBTvPlatforms = []Option{
	{Value: "", Name: "全部平台"},
	{Value: "腾讯视频", Name: "腾讯视频"},
	{Value: "爱奇艺", Name: "爱奇艺"},
	{Value: "优酷", Name: "优酷"},
	{Value: "湖南卫视", Name: "湖南卫视"},
}

var TMDBMovieSortByOptions = []Option{
	{Value: "", Name: "默认"},
	{Value: "popularity.desc", Name: "近期热度"},
	{Value: "vote_average.desc", Name: "高分优先"},
	{Value: "release_date.desc", Name: "首播时间"},
}

var TMDBMovieWithGenresOptions = []Option{
	{Value: "", Name: "全部"},
	{Value: "12", Name: "冒险"},
	{Value: "16", Name: "动画"},
	{Value: "35", Name: "喜剧"},
	{Value: "80", Name: "犯罪"},
	{Value: "18", Name: "剧情"},
	{Value: "14", Name: "奇幻"},
	{Value: "27", Name: "恐怖"},
	{Value: "9648", Name: "悬疑"},
	{Value: "10749", Name: "爱情"},
	{Value: "878", Name: "科幻"},
	{Value: "53", Name: "惊悚"},
	{Value: "10752", Name: "战争"},
}

var TMDBMovieWithOriginalLanguageOptions = []Option{
	{Value: "", Name: "全部"},
	{Value: "zh", Name: "中文"},
	{Value: "en", Name: "英语"},
	{Value: "ja", Name: "日语"},
	{Value: "ko", Name: "韩语"},
	{Value: "fr", Name: "法语"},
	{Value: "de", Name: "德语"},
	{Value: "ru", Name: "俄语"},
	{Value: "hi", Name: "印地语"},
}

var TMDBTvSortByOptions = []Option{
	{Value: "", Name: "默认"},
	{Value: "popularity.desc", Name: "近期热度"},
	{Value: "vote_average.desc", Name: "高分优先"},
	{Value: "first_air_date.desc", Name: "首播时间"},
}

var TMDBTvWithGenresOptions = []Option{
	{Value: "", Name: "全部"},
	{Value: "16", Name: "动画"},
	{Value: "35", Name: "喜剧"},
	{Value: "80", Name: "犯罪"},
	{Value: "18", Name: "剧情"},
	{Value: "9648", Name: "悬疑"},
	{Value: "10764", Name: "真人秀"},
	{Value: "10765", Name: "科幻"},
}

var TMDBTvWithOriginalLanguageOptions = []Option{
	{Value: "", Name: "全部"},
	{Value: "zh", Name: "中文"},
	{Value: "en", Name: "英语"},
	{Value: "ja", Name: "日语"},
	{Value: "ko", Name: "韩语"},
	{Value: "fr", Name: "法语"},
	{Value: "de", Name: "德语"},
	{Value: "ru", Name: "俄语"},
	{Value: "hi", Name: "印地语"},
}

var BLDocumentaryOrderOptions = []Option{
	{Value: "0", Name: "更新时间"},
	{Value: "6", Name: "上映时间"},
	{Value: "4", Name: "最高评分"},
	{Value: "2", Name: "播放数量"},
	{Value: "1", Name: "弹幕数量"},
}

var BLDocumentaryStyleIdOptions = []Option{
	{Value: "-1", Name: "全部"},
	{Value: "10033", Name: "历史"},
	{Value: "10045", Name: "美食"},
	{Value: "10065", Name: "人文"},
	{Value: "10066", Name: "科技"},
	{Value: "10067", Name: "探险"},
	{Value: "10068", Name: "宇宙"},
	{Value: "10069", Name: "萌宠"},
	{Value: "10070", Name: "社会"},
	{Value: "10071", Name: "动物"},
	{Value: "10072", Name: "自然"},
	{Value: "10073", Name: "医疗"},
	{Value: "10074", Name: "军事"},
	{Value: "10064", Name: "灾难"},
	{Value: "100075", Name: "罪案"},
	{Value: "10076", Name: "神秘"},
	{Value: "10077", Name: "旅行"},
	{Value: "10038", Name: "运动"},
	{Value: "-10", Name: "电影"},
}

var BLVarietyOrderOptions = []Option{
	{Value: "0", Name: "更新时间"},
	{Value: "6", Name: "上映时间"},
	{Value: "4", Name: "最高评分"},
	{Value: "2", Name: "播放数量"},
	{Value: "1", Name: "弹幕数量"},
}

var BLVarietyStyleIdOptions = []Option{
	{Value: "-1", Name: "全部"},
	{Value: "10040", Name: "音乐"},
	{Value: "10090", Name: "访谈"},
	{Value: "10091", Name: "脱口秀"},
	{Value: "10092", Name: "真人秀"},
	{Value: "10094", Name: "选秀"},
	{Value: "10045", Name: "美食"},
	{Value: "10095", Name: "旅游"},
	{Value: "10098", Name: "晚会"},
	{Value: "10096", Name: "演唱会"},
	{Value: "10084", Name: "情感"},
	{Value: "10051", Name: "喜剧"},
	{Value: "10097", Name: "亲子"},
	{Value: "10100", Name: "文化"},
	{Value: "10048", Name: "职场"},
	{Value: "10069", Name: "萌宠"},
	{Value: "10099", Name: "养成"},
}

var BLGuoManOrderOptions = []Option{
	{Value: "0", Name: "更新时间"},
	{Value: "6", Name: "上映时间"},
	{Value: "4", Name: "最高评分"},
	{Value: "2", Name: "播放数量"},
	{Value: "1", Name: "弹幕数量"},
}

var BLGuoManStyleIdOptions = []Option{
	{Value: "-1", Name: "全部"},
	{Value: "10010", Name: "原创"},
	{Value: "10011", Name: "漫画改"},
	{Value: "10012", Name: "小说改"},
	{Value: "10013", Name: "游戏改"},
	{Value: "10014", Name: "动态漫"},
	{Value: "10015", Name: "布袋戏"},
	{Value: "10016", Name: "热血"},
	{Value: "10018", Name: "奇幻"},
	{Value: "10019", Name: "玄幻"},
	{Value: "10020", Name: "战斗"},
	{Value: "10021", Name: "搞笑"},
	{Value: "10078", Name: "武侠"},
	{Value: "10022", Name: "日常"},
	{Value: "10023", Name: "科幻"},
	{Value: "10024", Name: "萌系"},
	{Value: "10025", Name: "治愈"},
	{Value: "10057", Name: "悬疑"},
	{Value: "10026", Name: "校园"},
	{Value: "10027", Name: "少儿"},
	{Value: "10028", Name: "泡面"},
	{Value: "10029", Name: "恋爱"},
	{Value: "10030", Name: "少女"},
	{Value: "10031", Name: "魔法"},
	{Value: "10033", Name: "历史"},
	{Value: "10035", Name: "机战"},
	{Value: "10036", Name: "神魔"},
	{Value: "10037", Name: "声控"},
	{Value: "10038", Name: "运动"},
	{Value: "10039", Name: "励志"},
	{Value: "10040", Name: "音乐"},
	{Value: "10041", Name: "推理"},
	{Value: "10042", Name: "社团"},
	{Value: "10043", Name: "智斗"},
	{Value: "10044", Name: "催泪"},
	{Value: "10045", Name: "美食"},
	{Value: "10046", Name: "偶像"},
	{Value: "10047", Name: "乙女"},
	{Value: "10048", Name: "职场"},
	{Value: "10049", Name: "古风"},
}

var SkyNetNewPlaylistsOptions = []Option{
	{"all", "全部"},
	{"official", "豆瓣片单"},
	{"selected", "精选"},
	{"classical", "经典"},
	{"prize", "获奖"},
	{"high_score", "高分"},
	{"movie_list", "榜单"},
	{"dark_horse", "冷门佳作"},
	{"topic", "主题"},
	{"director", "导演"},
	{"actor", "演员"},
	{"series", "系列"},
	{"chinese", "华语"},
	{"western", "欧美"},
	{"japanese", "日本"},
	{"korea", "韩国"},
	{"comedy", "喜剧"},
	{"action", "动作"},
	{"love", "爱情"},
	{"science_fiction", "科幻"},
	{"cartoon", "动画"},
	{"mystery", "悬疑"},
	{"panic", "惊悚"},
	{"horrible", "恐怖"},
	{"criminal", "犯罪"},
	{"lgbt", "同性"},
	{"war", "战争"},
	{"fantasy", "奇幻"},
	{"erotica", "情色"},
	{"music", "音乐"},
	{"documentary", "纪录片"},
	{"cure", "治愈"},
	{"art", "艺术"},
	{"dark_humor", "黑色幽默"},
	{"youth", "青春"},
	{"female", "女性"},
	{"real_event", "真实事件改编"},
	{"violence", "暴力"},
	{"black_white", "黑白"},
	{"food", "美食"},
	{"travel", "旅行"},
	{"child", "儿童"},
	{"humanity", "人性"},
	{"family", "家庭"},
	{"literary_art", "文艺"},
	{"novel", "小说改编"},
	{"moving", "感人"},
	{"inspiration", "励志"},
}