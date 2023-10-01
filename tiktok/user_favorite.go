package tiktok

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"tiktok/utils"
	"time"
)

type UserFavoriteVideoXBogusReq struct {
	DevicePlatform string `json:"device_platform"`
	Aid            string `json:"aid"`
	SecUserId      string `json:"sec_user_id"`
	UserUid        string `json:"user_id"`
	MaxCursor      string `json:"max_cursor"`
	Count          string `json:"count"`
}

func DownloadFavoriteVideos(secUserId string, userUid string, videoNum int, filePath string) {
	XBogusReq := &UserFavoriteVideoXBogusReq{
		DevicePlatform: "android",
		Aid:            "6383",
		SecUserId:      secUserId,
		UserUid:        userUid,
		MaxCursor:      "0",
		Count:          "18",
	}

	url := utils.ConcatXBogusUrlString(XBogusReq, USERFAVORITEPREFIX)
	url = NewXBogusReq(url)

	req, _ := utils.HttpNewRequest("GET", url, nil)

	h := &http.Client{}
	respStruct := &UserFavoriteResp{}
	var resp *http.Response

	// 需要多次获取到不同时间的视频列表然后合并
	var nameVideos = make(map[string][]string)
	for {
		resp, _ = h.Do(req)
		if resp.StatusCode == http.StatusOK && resp.ContentLength == 0 {
			fmt.Println("请求失败 2 秒后重试")
			resp.Body.Close()
			time.Sleep(2 * time.Second)
			continue
		} else if resp.StatusCode != http.StatusOK {
			fmt.Printf("请求失败 %d", resp.StatusCode)
			resp.Body.Close()
		}

		_ = json.NewDecoder(resp.Body).Decode(respStruct)

		if nameVideo, ok := respStruct.getAllVideoAdd(); ok {
			for k, v := range nameVideo {
				nameVideos[k] = v
			}
		} else if len(nameVideos) == 0 {
			panic("获取视频下载地址数量为空")
		}

		if len(nameVideos) >= videoNum || respStruct.HasMore == 0 {
			resp.Body.Close()
			break
		} else {
			XBogusReq.MaxCursor = strconv.Itoa(int(respStruct.MaxCursor))
			url := utils.ConcatXBogusUrlString(XBogusReq, USERFAVORITEPREFIX)
			url = NewXBogusReq(url)
			req, _ = utils.HttpNewRequest("GET", url, nil)
			resp.Body.Close()
		}
	}

	p := path.Join(filePath, "favorite")
	fmt.Println("开始下载喜欢的视频")
	for n, adds := range nameVideos {
		wg.Add(1)
		go download(n, adds, &wg, p)
	}

	wg.Wait()
	fmt.Println("共下载" + fmt.Sprintf("%d", len(nameVideos)) + "个视频")

	fmt.Println("下载完成")
}

type UserFavoriteResp struct {
	MaxCursor int64 `json:"max_cursor"`
	AwemeList []struct {
		CreateTime           int `json:"create_time"`
		ItemWarnNotification struct {
			Content string `json:"content"`
			Show    bool   `json:"show"`
			Type    int    `json:"type"`
		} `json:"item_warn_notification"`
		GeofencingRegions interface{} `json:"geofencing_regions"`
		VideoTag          []struct {
			TagId   int    `json:"tag_id"`
			TagName string `json:"tag_name"`
			Level   int    `json:"level"`
		} `json:"video_tag"`
		DisableRelationBar int   `json:"disable_relation_bar"`
		Duration           int   `json:"duration"`
		IsStory            int   `json:"is_story"`
		CommentGid         int64 `json:"comment_gid"`
		Statistics         struct {
			PlayCount    int `json:"play_count"`
			DiggCount    int `json:"digg_count"`
			CollectCount int `json:"collect_count"`
			CommentCount int `json:"comment_count"`
			ShareCount   int `json:"share_count"`
		} `json:"statistics"`
		Region          string      `json:"region"`
		IsTop           int         `json:"is_top"`
		PackedClips     interface{} `json:"packed_clips"`
		VideoText       interface{} `json:"video_text"`
		GroupId         string      `json:"group_id"`
		PreventDownload bool        `json:"prevent_download"`
		Desc            string      `json:"desc"`
		SocialTagList   interface{} `json:"social_tag_list"`
		OriginalImages  interface{} `json:"original_images"`
		ImageList       interface{} `json:"image_list"`
		Status          struct {
			PrivateStatus int  `json:"private_status"`
			InReviewing   bool `json:"in_reviewing"`
			IsProhibited  bool `json:"is_prohibited"`
			ReviewResult  struct {
				ReviewStatus int `json:"review_status"`
			} `json:"review_result"`
			ListenVideoStatus int  `json:"listen_video_status"`
			AllowShare        bool `json:"allow_share"`
			IsDelete          bool `json:"is_delete"`
			PartSee           int  `json:"part_see"`
		} `json:"status"`
		IsImageBeat          bool `json:"is_image_beat"`
		CollectStat          int  `json:"collect_stat"`
		CollectionCornerMark int  `json:"collection_corner_mark"`
		UserDigged           int  `json:"user_digged"`
		TextExtra            []struct {
			Start       int    `json:"start"`
			HashtagName string `json:"hashtag_name"`
			HashtagId   string `json:"hashtag_id"`
			IsCommerce  bool   `json:"is_commerce"`
			Type        int    `json:"type"`
			End         int    `json:"end"`
		} `json:"text_extra"`
		IsLifeItem                 bool `json:"is_life_item"`
		VideoGameDataChannelConfig struct {
		} `json:"video_game_data_channel_config"`
		LongVideo         interface{} `json:"long_video"`
		ImgBitrate        interface{} `json:"img_bitrate"`
		ChallengePosition interface{} `json:"challenge_position"`
		IsDuetSing        bool        `json:"is_duet_sing"`
		VideoControl      struct {
			ShareGrayed              bool `json:"share_grayed"`
			AllowDownload            bool `json:"allow_download"`
			DraftProgressBar         int  `json:"draft_progress_bar"`
			AllowReact               bool `json:"allow_react"`
			TimerStatus              int  `json:"timer_status"`
			AllowStitch              bool `json:"allow_stitch"`
			ShareType                int  `json:"share_type"`
			AllowDuet                bool `json:"allow_duet"`
			AllowDynamicWallpaper    bool `json:"allow_dynamic_wallpaper"`
			AllowMusic               bool `json:"allow_music"`
			AllowShare               bool `json:"allow_share"`
			PreventDownloadType      int  `json:"prevent_download_type"`
			DownloadIgnoreVisibility bool `json:"download_ignore_visibility"`
			ShareIgnoreVisibility    bool `json:"share_ignore_visibility"`
			DownloadInfo             struct {
				Level    int `json:"level"`
				FailInfo struct {
					Code   int    `json:"code"`
					Reason string `json:"reason"`
					Msg    string `json:"msg"`
				} `json:"fail_info,omitempty"`
			} `json:"download_info"`
			DuetInfo struct {
				Level int `json:"level"`
			} `json:"duet_info"`
			ShowProgressBar      int  `json:"show_progress_bar"`
			AllowDouplus         bool `json:"allow_douplus"`
			DuetIgnoreVisibility bool `json:"duet_ignore_visibility"`
		} `json:"video_control"`
		LibfinsertTaskId string `json:"libfinsert_task_id"`
		AwemeType        int    `json:"aweme_type"`
		ShowFollowButton struct {
		} `json:"show_follow_button"`
		Video struct {
			Cover struct {
				Uri     string   `json:"uri"`
				Width   int      `json:"width"`
				UrlList []string `json:"url_list"`
				Height  int      `json:"height"`
			} `json:"cover"`
			Height       int `json:"height"`
			PlayAddrH264 struct {
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
				FileCs   string   `json:"file_cs"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				Height   int      `json:"height"`
				DataSize int      `json:"data_size"`
				FileHash string   `json:"file_hash"`
			} `json:"play_addr_h264"`
			BigThumbs []struct {
				ImgNum   int     `json:"img_num"`
				Uri      string  `json:"uri"`
				Duration float64 `json:"duration"`
				Interval int     `json:"interval"`
				Fext     string  `json:"fext"`
				ImgUrl   string  `json:"img_url"`
				ImgXSize int     `json:"img_x_size"`
				ImgYSize int     `json:"img_y_size"`
				ImgXLen  int     `json:"img_x_len"`
				ImgYLen  int     `json:"img_y_len"`
			} `json:"big_thumbs"`
			Meta          string `json:"meta"`
			GaussianCover struct {
				Uri     string   `json:"uri"`
				Height  int      `json:"height"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"gaussian_cover"`
			PlayAddr265 struct {
				Uri      string   `json:"uri"`
				Height   int      `json:"height"`
				DataSize int      `json:"data_size"`
				FileHash string   `json:"file_hash"`
				UrlList  []string `json:"url_list"`
				UrlKey   string   `json:"url_key"`
				Width    int      `json:"width"`
				FileCs   string   `json:"file_cs"`
			} `json:"play_addr_265"`
			IsSourceHDR int    `json:"is_source_HDR"`
			VideoModel  string `json:"video_model"`
			BitRate     []struct {
				IsH265      int    `json:"is_h265"`
				IsBytevc1   int    `json:"is_bytevc1"`
				HDRType     string `json:"HDR_type"`
				FPS         int    `json:"FPS"`
				VideoExtra  string `json:"video_extra"`
				GearName    string `json:"gear_name"`
				QualityType int    `json:"quality_type"`
				BitRate     int    `json:"bit_rate"`
				HDRBit      string `json:"HDR_bit"`
				PlayAddr    struct {
					Uri      string   `json:"uri"`
					Height   int      `json:"height"`
					Width    int      `json:"width"`
					FileHash string   `json:"file_hash"`
					UrlList  []string `json:"url_list"`
					UrlKey   string   `json:"url_key"`
					DataSize int      `json:"data_size"`
					FileCs   string   `json:"file_cs"`
				} `json:"play_addr"`
			} `json:"bit_rate"`
			DynamicCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"dynamic_cover"`
			Duration     int         `json:"duration"`
			Width        int         `json:"width"`
			BitRateAudio interface{} `json:"bit_rate_audio"`
			PlayAddr     struct {
				DataSize int      `json:"data_size"`
				FileHash string   `json:"file_hash"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				Height   int      `json:"height"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
				FileCs   string   `json:"file_cs"`
			} `json:"play_addr"`
			OriginCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Height  int      `json:"height"`
				Width   int      `json:"width"`
			} `json:"origin_cover"`
			Ratio          string `json:"ratio"`
			HorizontalType int    `json:"horizontal_type,omitempty"`
		} `json:"video"`
		Images              interface{} `json:"images"`
		ImageAlbumMusicInfo struct {
			BeginTime int `json:"begin_time"`
			EndTime   int `json:"end_time"`
			Volume    int `json:"volume"`
		} `json:"image_album_music_info"`
		SeriesPaidInfo struct {
			SeriesPaidStatus int `json:"series_paid_status"`
			ItemPrice        int `json:"item_price"`
		} `json:"series_paid_info"`
		VideoLabels         interface{} `json:"video_labels"`
		PhotoSearchEntrance struct {
			EcomType int `json:"ecom_type"`
		} `json:"photo_search_entrance"`
		ImageCropCtrl         int `json:"image_crop_ctrl"`
		AuthorMaskTag         int `json:"author_mask_tag"`
		CommentPermissionInfo struct {
			ItemDetailEntry         bool `json:"item_detail_entry"`
			PressEntry              bool `json:"press_entry"`
			ToastGuide              bool `json:"toast_guide"`
			CommentPermissionStatus int  `json:"comment_permission_status"`
			CanComment              bool `json:"can_comment"`
		} `json:"comment_permission_info"`
		DistributeCircle struct {
			DistributeType         int  `json:"distribute_type"`
			IsCampus               bool `json:"is_campus"`
			CampusBlockInteraction bool `json:"campus_block_interaction"`
		} `json:"distribute_circle"`
		OriginCommentIds interface{} `json:"origin_comment_ids"`
		Music            struct {
			Title             string `json:"title"`
			CanBackgroundPlay bool   `json:"can_background_play"`
			PgcMusicType      int    `json:"pgc_music_type"`
			IsRestricted      bool   `json:"is_restricted"`
			AvatarThumb       struct {
				Uri     string   `json:"uri"`
				Height  int      `json:"height"`
				Width   int      `json:"width"`
				UrlList []string `json:"url_list"`
			} `json:"avatar_thumb"`
			SecUid          string `json:"sec_uid"`
			Status          int    `json:"status"`
			ReasonType      int    `json:"reason_type"`
			IsOriginalSound bool   `json:"is_original_sound"`
			CollectStat     int    `json:"collect_stat"`
			AvatarMedium    struct {
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Uri     string   `json:"uri"`
				Height  int      `json:"height"`
			} `json:"avatar_medium"`
			MusicChartRanks                interface{} `json:"music_chart_ranks"`
			MusicStatus                    int         `json:"music_status"`
			VideoDuration                  int         `json:"video_duration"`
			PreventItemDownloadStatus      int         `json:"prevent_item_download_status"`
			Id                             int64       `json:"id"`
			DmvAutoShow                    bool        `json:"dmv_auto_show"`
			Position                       interface{} `json:"position"`
			AuthorStatus                   int         `json:"author_status"`
			IsVideoSelfSee                 bool        `json:"is_video_self_see"`
			UnshelveCountries              interface{} `json:"unshelve_countries"`
			MusicCoverAtmosphereColorValue string      `json:"music_cover_atmosphere_color_value"`
			PreviewStartTime               int         `json:"preview_start_time"`
			CoverThumb                     struct {
				Height  int      `json:"height"`
				Width   int      `json:"width"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"cover_thumb"`
			AuthorDeleted        bool        `json:"author_deleted"`
			OfflineDesc          string      `json:"offline_desc"`
			Redirect             bool        `json:"redirect"`
			PreventDownload      bool        `json:"prevent_download"`
			OwnerNickname        string      `json:"owner_nickname"`
			IsAudioUrlWithCookie bool        `json:"is_audio_url_with_cookie"`
			ShootDuration        int         `json:"shoot_duration"`
			IsDelVideo           bool        `json:"is_del_video"`
			ArtistUserInfos      interface{} `json:"artist_user_infos"`
			AuthorPosition       interface{} `json:"author_position"`
			LyricShortPosition   interface{} `json:"lyric_short_position"`
			Album                string      `json:"album"`
			Duration             int         `json:"duration"`
			IsOriginal           bool        `json:"is_original"`
			UserCount            int         `json:"user_count"`
			PreviewEndTime       int         `json:"preview_end_time"`
			PlayUrl              struct {
				UrlKey  string   `json:"url_key"`
				Uri     string   `json:"uri"`
				Height  int      `json:"height"`
				Width   int      `json:"width"`
				UrlList []string `json:"url_list"`
			} `json:"play_url"`
			OwnerId          string `json:"owner_id"`
			SourcePlatform   int    `json:"source_platform"`
			Author           string `json:"author"`
			Mid              string `json:"mid"`
			EndTime          int    `json:"end_time"`
			OwnerHandle      string `json:"owner_handle"`
			AuditionDuration int    `json:"audition_duration"`
			CoverLarge       struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Height  int      `json:"height"`
				Width   int      `json:"width"`
			} `json:"cover_large"`
			ExternalSongInfo []interface{} `json:"external_song_info"`
			DspStatus        int           `json:"dsp_status"`
			StartTime        int           `json:"start_time"`
			AvatarLarge      struct {
				UrlList []string `json:"url_list"`
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				Width   int      `json:"width"`
			} `json:"avatar_large"`
			Artists []struct {
				EnterType int    `json:"enter_type"`
				Uid       string `json:"uid"`
				SecUid    string `json:"sec_uid"`
				NickName  string `json:"nick_name"`
				Handle    string `json:"handle"`
				Avatar    struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"avatar"`
				IsVerified bool `json:"is_verified"`
			} `json:"artists"`
			SchemaUrl         string `json:"schema_url"`
			MusicCollectCount int    `json:"music_collect_count"`
			StrongBeatUrl     struct {
				UrlList []string `json:"url_list"`
				Height  int      `json:"height"`
				Width   int      `json:"width"`
				Uri     string   `json:"uri"`
			} `json:"strong_beat_url,omitempty"`
			IsPgc       bool        `json:"is_pgc"`
			TagList     interface{} `json:"tag_list"`
			CoverMedium struct {
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
				Uri     string   `json:"uri"`
				Height  int      `json:"height"`
			} `json:"cover_medium"`
			IsCommerceMusic   bool        `json:"is_commerce_music"`
			MuteShare         bool        `json:"mute_share"`
			MusicianUserInfos interface{} `json:"musician_user_infos"`
			Extra             string      `json:"extra"`
			BindedChallengeId int         `json:"binded_challenge_id"`
			IsMatchedMetadata bool        `json:"is_matched_metadata"`
			SearchImpr        struct {
				EntityId string `json:"entity_id"`
			} `json:"search_impr"`
			IdStr   string `json:"id_str"`
			CoverHd struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_hd"`
			Song struct {
				Id      int64       `json:"id"`
				IdStr   string      `json:"id_str"`
				Artists interface{} `json:"artists"`
				Chorus  struct {
					StartMs    int `json:"start_ms"`
					DurationMs int `json:"duration_ms"`
				} `json:"chorus,omitempty"`
				Title string `json:"title,omitempty"`
			} `json:"song,omitempty"`
			MatchedPgcSound struct {
				MixedTitle  string `json:"mixed_title"`
				MixedAuthor string `json:"mixed_author"`
				CoverMedium struct {
					Width   int      `json:"width"`
					Height  int      `json:"height"`
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
				} `json:"cover_medium"`
				Author string `json:"author"`
				Title  string `json:"title"`
			} `json:"matched_pgc_sound,omitempty"`
			CoverColorHsv struct {
				H int `json:"h"`
				S int `json:"s"`
				V int `json:"v"`
			} `json:"cover_color_hsv,omitempty"`
		} `json:"music"`
		IsAds                   bool          `json:"is_ads"`
		UserRecommendStatus     int           `json:"user_recommend_status"`
		Geofencing              []interface{} `json:"geofencing"`
		ChapterList             interface{}   `json:"chapter_list"`
		BoostStatus             int           `json:"boost_status"`
		Promotions              []interface{} `json:"promotions"`
		DuetAggregateInMusicTab bool          `json:"duet_aggregate_in_music_tab"`
		AuthorUserId            int64         `json:"author_user_id"`
		Anchors                 interface{}   `json:"anchors"`
		ActivityVideoType       int           `json:"activity_video_type"`
		AuthenticationToken     string        `json:"authentication_token"`
		ShareUrl                string        `json:"share_url"`
		Author                  struct {
			ShareInfo struct {
				ShareTitle       string `json:"share_title"`
				ShareTitleMyself string `json:"share_title_myself"`
				ShareDesc        string `json:"share_desc"`
				ShareTitleOther  string `json:"share_title_other"`
				ShareUrl         string `json:"share_url"`
				ShareDescInfo    string `json:"share_desc_info"`
				ShareQrcodeUrl   struct {
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Height  int      `json:"height"`
					Width   int      `json:"width"`
				} `json:"share_qrcode_url"`
				ShareWeiboDesc string `json:"share_weibo_desc"`
			} `json:"share_info"`
			Uid         string `json:"uid"`
			AvatarThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			PreventDownload        bool   `json:"prevent_download"`
			SecUid                 string `json:"sec_uid"`
			CustomVerify           string `json:"custom_verify"`
			EnterpriseVerifyReason string `json:"enterprise_verify_reason"`
			RiskNoticeText         string `json:"risk_notice_text"`
			Nickname               string `json:"nickname"`
			FollowerStatus         int    `json:"follower_status"`
			FollowStatus           int    `json:"follow_status"`
		} `json:"author"`
		ImageComment struct {
		} `json:"image_comment"`
		ComponentInfoV2 string      `json:"component_info_v2"`
		Position        interface{} `json:"position"`
		UniqidPosition  interface{} `json:"uniqid_position"`
		AwemeControl    struct {
			CanComment     bool `json:"can_comment"`
			CanShowComment bool `json:"can_show_comment"`
			CanForward     bool `json:"can_forward"`
			CanShare       bool `json:"can_share"`
		} `json:"aweme_control"`
		IsCollectsSelected  int         `json:"is_collects_selected"`
		InteractionStickers interface{} `json:"interaction_stickers"`
		CommerceConfigData  interface{} `json:"commerce_config_data"`
		ImageInfos          interface{} `json:"image_infos"`
		ImpressionData      struct {
			GroupIdListB   []interface{} `json:"group_id_list_b"`
			SimilarIdListA []int64       `json:"similar_id_list_a"`
			SimilarIdListB []int64       `json:"similar_id_list_b"`
			GroupIdListC   []interface{} `json:"group_id_list_c"`
			GroupIdListA   []int64       `json:"group_id_list_a"`
		} `json:"impression_data"`
		AwemeId    string `json:"aweme_id"`
		DiggLottie struct {
			LottieId string `json:"lottie_id"`
			CanBomb  int    `json:"can_bomb"`
		} `json:"digg_lottie"`
		LabelTopText      interface{} `json:"label_top_text"`
		FeedCommentConfig struct {
		} `json:"feed_comment_config"`
		ShareInfo struct {
			ShareUrl      string `json:"share_url"`
			ShareLinkDesc string `json:"share_link_desc"`
		} `json:"share_info"`
		NicknamePosition interface{} `json:"nickname_position"`
		MediaType        int         `json:"media_type"`
		CommentList      interface{} `json:"comment_list"`
		CommonBarInfo    string      `json:"common_bar_info"`
		CoverLabels      interface{} `json:"cover_labels"`
		HybridLabel      interface{} `json:"hybrid_label"`
		VtagSearch       struct {
			VtagEnable  bool `json:"vtag_enable"`
			VtagDelayTs int  `json:"vtag_delay_ts"`
		} `json:"vtag_search,omitempty"`
		HorizontalType int `json:"horizontal_type,omitempty"`
		DanmakuControl struct {
			EnableDanmaku      bool   `json:"enable_danmaku"`
			PostPrivilegeLevel int    `json:"post_privilege_level"`
			IsPostDenied       bool   `json:"is_post_denied"`
			PostDeniedReason   string `json:"post_denied_reason"`
			SkipDanmaku        bool   `json:"skip_danmaku"`
			DanmakuCnt         int    `json:"danmaku_cnt"`
			Activities         []struct {
				Id   int `json:"id"`
				Type int `json:"type"`
			} `json:"activities"`
		} `json:"danmaku_control,omitempty"`
		OpenPlatformName string `json:"open_platform_name,omitempty"`
		AnchorInfo       struct {
			TitleTag  string `json:"title_tag"`
			Extra     string `json:"extra"`
			LogExtra  string `json:"log_extra"`
			Content   string `json:"content"`
			StyleInfo struct {
				DefaultIcon string `json:"default_icon"`
				SceneIcon   string `json:"scene_icon"`
				Extra       string `json:"extra"`
			} `json:"style_info"`
			Type int    `json:"type"`
			Id   string `json:"id"`
			Icon struct {
				Height  int      `json:"height"`
				Width   int      `json:"width"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				UrlKey  string   `json:"url_key"`
			} `json:"icon"`
			Title string `json:"title"`
		} `json:"anchor_info,omitempty"`
		OpenPlatformInfo struct {
			Name string `json:"name"`
			Link string `json:"link"`
			Icon string `json:"icon"`
		} `json:"open_platform_info,omitempty"`
	} `json:"aweme_list"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	HasMore    int `json:"has_more"`
	StatusCode int `json:"status_code"`
}

func (u *UserFavoriteResp) getAllVideoAdd() (map[string][]string, bool) {
	if len(u.AwemeList) == 0 {
		return nil, false
	}
	videoSlice := u.AwemeList
	descVideoPlayAddr := make(map[string][]string)
	for _, videoInfo := range videoSlice {
		desc := videoInfo.Desc
		//deep copy the list
		var urlSlice = make([]string, len(videoInfo.Video.PlayAddr.UrlList))
		copy(urlSlice, videoInfo.Video.PlayAddr.UrlList)
		descVideoPlayAddr[desc] = urlSlice
	}
	return descVideoPlayAddr, true
}
