package tiktok

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func UserVideos(secUserId string, videoCount string) string {
	// https://www.douyin.com/user/MS4wLjABAAAA7xbdm1QfWD8Um6rFnrm0wVpnOI1uEHhbth1XDud_tWRxG5ZI6YUbNu9ES4uMjF0D?is_search=0&list_name=follow&nt=3
	// 用正则表达式提取user/后面到第一个?之间的字符串
	if strings.HasPrefix(secUserId, "https://") {
		reg := regexp.MustCompile(`(?m)user/(.*?)\?`)
		secUserId = reg.FindStringSubmatch(secUserId)[1]
	}
	return NewXBogusReq(secUserId, videoCount)
}

type UserVideoResp struct {
	AwemeList []struct {
		Anchors             interface{} `json:"anchors"`
		AuthenticationToken string      `json:"authentication_token"`
		Author              struct {
			AcceptPrivatePolicy bool   `json:"accept_private_policy"`
			AccountRegion       string `json:"account_region"`
			AppleAccount        int    `json:"apple_account"`
			AvatarThumb         struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			AvatarUri    string `json:"avatar_uri"`
			AwemeControl struct {
				CanComment     bool `json:"can_comment"`
				CanForward     bool `json:"can_forward"`
				CanShare       bool `json:"can_share"`
				CanShowComment bool `json:"can_show_comment"`
			} `json:"aweme_control"`
			AwemeCount               int           `json:"aweme_count"`
			AwemeHotsoonAuth         int           `json:"aweme_hotsoon_auth"`
			AwemeHotsoonAuthRelation int           `json:"aweme_hotsoon_auth_relation"`
			BanUserFunctions         []interface{} `json:"ban_user_functions"`
			BindPhone                string        `json:"bind_phone"`
			CanSetGeofencing         interface{}   `json:"can_set_geofencing"`
			CardEntries              interface{}   `json:"card_entries"`
			CardEntriesNotDisplay    interface{}   `json:"card_entries_not_display"`
			CardSortPriority         interface{}   `json:"card_sort_priority"`
			CfList                   interface{}   `json:"cf_list"`
			ChaList                  interface{}   `json:"cha_list"`
			CloseFriendType          int           `json:"close_friend_type"`
			Constellation            int           `json:"constellation"`
			ContactsStatus           int           `json:"contacts_status"`
			ContrailList             interface{}   `json:"contrail_list"`
			CoverUrl                 []struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_url"`
			CreateTime                             int           `json:"create_time"`
			CustomVerify                           string        `json:"custom_verify"`
			CvLevel                                string        `json:"cv_level"`
			DataLabelList                          interface{}   `json:"data_label_list"`
			DisplayInfo                            interface{}   `json:"display_info"`
			DownloadPromptTs                       int           `json:"download_prompt_ts"`
			EnableNearbyVisible                    bool          `json:"enable_nearby_visible"`
			EndorsementInfoList                    interface{}   `json:"endorsement_info_list"`
			EnterpriseVerifyReason                 string        `json:"enterprise_verify_reason"`
			FamiliarVisitorUser                    interface{}   `json:"familiar_visitor_user"`
			FavoritingCount                        int           `json:"favoriting_count"`
			FbExpireTime                           int           `json:"fb_expire_time"`
			FollowStatus                           int           `json:"follow_status"`
			FollowerCount                          int           `json:"follower_count"`
			FollowerListSecondaryInformationStruct interface{}   `json:"follower_list_secondary_information_struct"`
			FollowerRequestStatus                  int           `json:"follower_request_status"`
			FollowerStatus                         int           `json:"follower_status"`
			FollowingCount                         int           `json:"following_count"`
			Geofencing                             []interface{} `json:"geofencing"`
			GoogleAccount                          string        `json:"google_account"`
			HasEmail                               bool          `json:"has_email"`
			HasFacebookToken                       bool          `json:"has_facebook_token"`
			HasInsights                            bool          `json:"has_insights"`
			HasOrders                              bool          `json:"has_orders"`
			HasTwitterToken                        bool          `json:"has_twitter_token"`
			HasYoutubeToken                        bool          `json:"has_youtube_token"`
			HideSearch                             bool          `json:"hide_search"`
			HomepageBottomToast                    interface{}   `json:"homepage_bottom_toast"`
			ImRoleIds                              interface{}   `json:"im_role_ids"`
			InsId                                  string        `json:"ins_id"`
			InterestTags                           interface{}   `json:"interest_tags"`
			IsBindedWeibo                          bool          `json:"is_binded_weibo"`
			IsBlockedV2                            bool          `json:"is_blocked_v2"`
			IsBlockingV2                           bool          `json:"is_blocking_v2"`
			IsCf                                   int           `json:"is_cf"`
			IsNotShow                              bool          `json:"is_not_show"`
			IsPhoneBinded                          bool          `json:"is_phone_binded"`
			ItemList                               interface{}   `json:"item_list"`
			KyOnlyPredict                          int           `json:"ky_only_predict"`
			LinkItemList                           interface{}   `json:"link_item_list"`
			LiveAgreement                          int           `json:"live_agreement"`
			LiveAgreementTime                      int           `json:"live_agreement_time"`
			LiveCommerce                           bool          `json:"live_commerce"`
			LiveHighValue                          int           `json:"live_high_value"`
			LiveVerify                             int           `json:"live_verify"`
			MaxFollowerCount                       int           `json:"max_follower_count"`
			NeedPoints                             interface{}   `json:"need_points"`
			NeedRecommend                          int           `json:"need_recommend"`
			NeiguangShield                         int           `json:"neiguang_shield"`
			NewStoryCover                          interface{}   `json:"new_story_cover"`
			Nickname                               string        `json:"nickname"`
			NotSeenItemIdList                      interface{}   `json:"not_seen_item_id_list"`
			NotSeenItemIdListV2                    interface{}   `json:"not_seen_item_id_list_v2"`
			OfflineInfoList                        interface{}   `json:"offline_info_list"`
			PersonalTagList                        interface{}   `json:"personal_tag_list"`
			PlatformSyncInfo                       interface{}   `json:"platform_sync_info"`
			PreventDownload                        bool          `json:"prevent_download"`
			ReactSetting                           int           `json:"react_setting"`
			ReflowPageGid                          int           `json:"reflow_page_gid"`
			ReflowPageUid                          int           `json:"reflow_page_uid"`
			RiskNoticeText                         string        `json:"risk_notice_text"`
			SchoolCategory                         int           `json:"school_category"`
			SchoolId                               string        `json:"school_id"`
			SearchImpr                             struct {
				EntityId string `json:"entity_id"`
			} `json:"search_impr"`
			SecUid    string `json:"sec_uid"`
			Secret    int    `json:"secret"`
			ShareInfo struct {
				ShareDesc      string `json:"share_desc"`
				ShareDescInfo  string `json:"share_desc_info"`
				ShareQrcodeUrl struct {
					Height  int      `json:"height"`
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
				} `json:"share_qrcode_url"`
				ShareTitle       string `json:"share_title"`
				ShareTitleMyself string `json:"share_title_myself"`
				ShareTitleOther  string `json:"share_title_other"`
				ShareUrl         string `json:"share_url"`
				ShareWeiboDesc   string `json:"share_weibo_desc"`
			} `json:"share_info"`
			ShareQrcodeUri        string      `json:"share_qrcode_uri"`
			ShieldCommentNotice   int         `json:"shield_comment_notice"`
			ShieldDiggNotice      int         `json:"shield_digg_notice"`
			ShieldFollowNotice    int         `json:"shield_follow_notice"`
			ShortId               string      `json:"short_id"`
			ShowImageBubble       bool        `json:"show_image_bubble"`
			ShowNearbyActive      bool        `json:"show_nearby_active"`
			Signature             string      `json:"signature"`
			SignatureDisplayLines int         `json:"signature_display_lines"`
			SignatureExtra        interface{} `json:"signature_extra"`
			SpecialFollowStatus   int         `json:"special_follow_status"`
			SpecialLock           int         `json:"special_lock"`
			SpecialPeopleLabels   interface{} `json:"special_people_labels"`
			Status                int         `json:"status"`
			StoryOpen             bool        `json:"story_open"`
			TextExtra             interface{} `json:"text_extra"`
			TotalFavorited        int         `json:"total_favorited"`
			TwExpireTime          int         `json:"tw_expire_time"`
			TwitterId             string      `json:"twitter_id"`
			TwitterName           string      `json:"twitter_name"`
			TypeLabel             interface{} `json:"type_label"`
			Uid                   string      `json:"uid"`
			UniqueId              string      `json:"unique_id"`
			UniqueIdModifyTime    int         `json:"unique_id_modify_time"`
			UserAge               int         `json:"user_age"`
			UserCanceled          bool        `json:"user_canceled"`
			UserMode              int         `json:"user_mode"`
			UserNotSee            int         `json:"user_not_see"`
			UserNotShow           int         `json:"user_not_show"`
			UserPeriod            int         `json:"user_period"`
			UserPermissions       interface{} `json:"user_permissions"`
			UserRate              int         `json:"user_rate"`
			UserTags              interface{} `json:"user_tags"`
			VerificationType      int         `json:"verification_type"`
			WeiboName             string      `json:"weibo_name"`
			WeiboSchema           string      `json:"weibo_schema"`
			WeiboUrl              string      `json:"weibo_url"`
			WeiboVerify           string      `json:"weibo_verify"`
			WhiteCoverUrl         interface{} `json:"white_cover_url"`
			WithDouEntry          bool        `json:"with_dou_entry"`
			WithFusionShopEntry   bool        `json:"with_fusion_shop_entry"`
			WithShopEntry         bool        `json:"with_shop_entry"`
			YoutubeChannelId      string      `json:"youtube_channel_id"`
			YoutubeChannelTitle   string      `json:"youtube_channel_title"`
			YoutubeExpireTime     int         `json:"youtube_expire_time"`
		} `json:"author"`
		AuthorMaskTag int   `json:"author_mask_tag"`
		AuthorUserId  int64 `json:"author_user_id"`
		AwemeControl  struct {
			CanComment     bool `json:"can_comment"`
			CanForward     bool `json:"can_forward"`
			CanShare       bool `json:"can_share"`
			CanShowComment bool `json:"can_show_comment"`
		} `json:"aweme_control"`
		AwemeId   string `json:"aweme_id"`
		AwemeType int    `json:"aweme_type"`
		BookBar   struct {
		} `json:"book_bar"`
		ChallengePosition     interface{} `json:"challenge_position"`
		ChapterList           interface{} `json:"chapter_list"`
		CollectStat           int         `json:"collect_stat"`
		CollectionCornerMark  int         `json:"collection_corner_mark"`
		CommentGid            int64       `json:"comment_gid"`
		CommentList           interface{} `json:"comment_list"`
		CommentPermissionInfo struct {
			CanComment              bool `json:"can_comment"`
			CommentPermissionStatus int  `json:"comment_permission_status"`
			ItemDetailEntry         bool `json:"item_detail_entry"`
			PressEntry              bool `json:"press_entry"`
			ToastGuide              bool `json:"toast_guide"`
		} `json:"comment_permission_info"`
		CommerceConfigData interface{} `json:"commerce_config_data"`
		CommonBarInfo      string      `json:"common_bar_info"`
		ComponentInfoV2    string      `json:"component_info_v2"`
		CoverLabels        interface{} `json:"cover_labels"`
		CreateTime         int         `json:"create_time"`
		Desc               string      `json:"desc"`
		DiggLottie         struct {
			CanBomb  int    `json:"can_bomb"`
			LottieId string `json:"lottie_id"`
		} `json:"digg_lottie"`
		DisableRelationBar     int         `json:"disable_relation_bar"`
		DislikeDimensionList   interface{} `json:"dislike_dimension_list"`
		DislikeDimensionListV2 interface{} `json:"dislike_dimension_list_v2"`
		DistributeCircle       struct {
			CampusBlockInteraction bool `json:"campus_block_interaction"`
			DistributeType         int  `json:"distribute_type"`
		} `json:"distribute_circle"`
		DuetAggregateInMusicTab bool          `json:"duet_aggregate_in_music_tab"`
		Duration                int           `json:"duration"`
		Geofencing              []interface{} `json:"geofencing"`
		GeofencingRegions       interface{}   `json:"geofencing_regions"`
		GroupId                 string        `json:"group_id"`
		GuideBtnType            int           `json:"guide_btn_type"`
		GuideSceneInfo          struct {
			DiamondExposeInfoStr string `json:"diamond_expose_info_str"`
			FeedOriginGidInfoStr string `json:"feed_origin_gid_info_str"`
			GuideSceneType       int    `json:"guide_scene_type"`
		} `json:"guide_scene_info"`
		HybridLabel         interface{} `json:"hybrid_label"`
		ImageAlbumMusicInfo struct {
			BeginTime int `json:"begin_time"`
			EndTime   int `json:"end_time"`
			Volume    int `json:"volume"`
		} `json:"image_album_music_info"`
		ImageComment struct {
		} `json:"image_comment"`
		ImageInfos     interface{} `json:"image_infos"`
		ImageList      interface{} `json:"image_list"`
		Images         interface{} `json:"images"`
		ImgBitrate     interface{} `json:"img_bitrate"`
		ImpressionData struct {
			GroupIdListA   []interface{} `json:"group_id_list_a"`
			GroupIdListB   []interface{} `json:"group_id_list_b"`
			GroupIdListC   []int64       `json:"group_id_list_c"`
			SimilarIdListA interface{}   `json:"similar_id_list_a"`
			SimilarIdListB interface{}   `json:"similar_id_list_b"`
		} `json:"impression_data"`
		InteractionStickers  interface{} `json:"interaction_stickers"`
		IsCollectsSelected   int         `json:"is_collects_selected"`
		IsDuetSing           bool        `json:"is_duet_sing"`
		IsImageBeat          bool        `json:"is_image_beat"`
		IsLifeItem           bool        `json:"is_life_item"`
		IsSharePost          bool        `json:"is_share_post"`
		IsStory              int         `json:"is_story"`
		IsTop                int         `json:"is_top"`
		ItemWarnNotification struct {
			Content string `json:"content"`
			Show    bool   `json:"show"`
			Type    int    `json:"type"`
		} `json:"item_warn_notification"`
		LabelTopText interface{} `json:"label_top_text"`
		LongVideo    interface{} `json:"long_video"`
		Music        struct {
			Album            string        `json:"album"`
			ArtistUserInfos  interface{}   `json:"artist_user_infos"`
			Artists          []interface{} `json:"artists"`
			AuditionDuration int           `json:"audition_duration"`
			Author           string        `json:"author"`
			AuthorDeleted    bool          `json:"author_deleted"`
			AuthorPosition   interface{}   `json:"author_position"`
			AuthorStatus     int           `json:"author_status"`
			AvatarLarge      struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_large"`
			AvatarMedium struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_medium"`
			AvatarThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			BindedChallengeId int  `json:"binded_challenge_id"`
			CanBackgroundPlay bool `json:"can_background_play"`
			CollectStat       int  `json:"collect_stat"`
			CoverHd           struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_hd"`
			CoverLarge struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_large"`
			CoverMedium struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_medium"`
			CoverThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_thumb"`
			DmvAutoShow          bool          `json:"dmv_auto_show"`
			DspStatus            int           `json:"dsp_status"`
			Duration             int           `json:"duration"`
			EndTime              int           `json:"end_time"`
			ExternalSongInfo     []interface{} `json:"external_song_info"`
			Extra                string        `json:"extra"`
			Id                   int64         `json:"id"`
			IdStr                string        `json:"id_str"`
			IsAudioUrlWithCookie bool          `json:"is_audio_url_with_cookie"`
			IsCommerceMusic      bool          `json:"is_commerce_music"`
			IsDelVideo           bool          `json:"is_del_video"`
			IsMatchedMetadata    bool          `json:"is_matched_metadata"`
			IsOriginal           bool          `json:"is_original"`
			IsOriginalSound      bool          `json:"is_original_sound"`
			IsPgc                bool          `json:"is_pgc"`
			IsRestricted         bool          `json:"is_restricted"`
			IsVideoSelfSee       bool          `json:"is_video_self_see"`
			LunaInfo             struct {
				HasCopyright bool `json:"has_copyright,omitempty"`
				IsLunaUser   bool `json:"is_luna_user"`
			} `json:"luna_info"`
			LyricShortPosition interface{} `json:"lyric_short_position"`
			MatchedPgcSound    struct {
				Author      string `json:"author"`
				CoverMedium struct {
					Height  int      `json:"height"`
					Uri     string   `json:"uri"`
					UrlList []string `json:"url_list"`
					Width   int      `json:"width"`
				} `json:"cover_medium"`
				MixedAuthor string `json:"mixed_author"`
				MixedTitle  string `json:"mixed_title"`
				Title       string `json:"title"`
			} `json:"matched_pgc_sound,omitempty"`
			Mid                            string      `json:"mid"`
			MusicChartRanks                interface{} `json:"music_chart_ranks"`
			MusicCollectCount              int         `json:"music_collect_count"`
			MusicCoverAtmosphereColorValue string      `json:"music_cover_atmosphere_color_value"`
			MusicStatus                    int         `json:"music_status"`
			MusicianUserInfos              interface{} `json:"musician_user_infos"`
			MuteShare                      bool        `json:"mute_share"`
			OfflineDesc                    string      `json:"offline_desc"`
			OwnerHandle                    string      `json:"owner_handle"`
			OwnerId                        string      `json:"owner_id"`
			OwnerNickname                  string      `json:"owner_nickname"`
			PgcMusicType                   int         `json:"pgc_music_type"`
			PlayUrl                        struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlKey  string   `json:"url_key"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"play_url"`
			Position                  interface{} `json:"position"`
			PreventDownload           bool        `json:"prevent_download"`
			PreventItemDownloadStatus int         `json:"prevent_item_download_status"`
			PreviewEndTime            int         `json:"preview_end_time"`
			PreviewStartTime          int         `json:"preview_start_time"`
			ReasonType                int         `json:"reason_type"`
			Redirect                  bool        `json:"redirect"`
			SchemaUrl                 string      `json:"schema_url"`
			SearchImpr                struct {
				EntityId string `json:"entity_id"`
			} `json:"search_impr"`
			SecUid        string `json:"sec_uid"`
			ShootDuration int    `json:"shoot_duration"`
			Song          struct {
				Artists       interface{} `json:"artists"`
				ChorusV3Infos interface{} `json:"chorus_v3_infos"`
				Id            int64       `json:"id"`
				IdStr         string      `json:"id_str"`
			} `json:"song,omitempty"`
			SourcePlatform int `json:"source_platform"`
			StartTime      int `json:"start_time"`
			Status         int `json:"status"`
			StrongBeatUrl  struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"strong_beat_url"`
			TagList           interface{} `json:"tag_list"`
			Title             string      `json:"title"`
			UnshelveCountries interface{} `json:"unshelve_countries"`
			UserCount         int         `json:"user_count"`
			VideoDuration     int         `json:"video_duration"`
		} `json:"music"`
		NicknamePosition    interface{}   `json:"nickname_position"`
		OriginCommentIds    interface{}   `json:"origin_comment_ids"`
		OriginTextExtra     []interface{} `json:"origin_text_extra"`
		OriginalImages      interface{}   `json:"original_images"`
		PackedClips         interface{}   `json:"packed_clips"`
		PhotoSearchEntrance struct {
			EcomType int `json:"ecom_type"`
		} `json:"photo_search_entrance"`
		Position             interface{}   `json:"position"`
		PreventDownload      bool          `json:"prevent_download"`
		PreviewTitle         string        `json:"preview_title"`
		PreviewVideoStatus   int           `json:"preview_video_status"`
		Promotions           []interface{} `json:"promotions"`
		RefTtsIdList         interface{}   `json:"ref_tts_id_list"`
		RefVoiceModifyIdList interface{}   `json:"ref_voice_modify_id_list"`
		Region               string        `json:"region"`
		RelationLabels       interface{}   `json:"relation_labels"`
		ReportAction         bool          `json:"report_action"`
		SearchImpr           struct {
			EntityId   string `json:"entity_id"`
			EntityType string `json:"entity_type"`
		} `json:"search_impr"`
		SeoInfo struct {
		} `json:"seo_info"`
		SeriesPaidInfo struct {
			ItemPrice        int `json:"item_price"`
			SeriesPaidStatus int `json:"series_paid_status"`
		} `json:"series_paid_info"`
		ShareInfo struct {
			ShareLinkDesc string `json:"share_link_desc"`
			ShareUrl      string `json:"share_url"`
		} `json:"share_info"`
		ShareUrl           string `json:"share_url"`
		ShouldOpenAdReport bool   `json:"should_open_ad_report"`
		ShowFollowButton   struct {
		} `json:"show_follow_button"`
		SocialTagList       interface{} `json:"social_tag_list"`
		StandardBarInfoList interface{} `json:"standard_bar_info_list"`
		Statistics          struct {
			AdmireCount  int    `json:"admire_count"`
			AwemeId      string `json:"aweme_id"`
			CollectCount int    `json:"collect_count"`
			CommentCount int    `json:"comment_count"`
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
		} `json:"statistics"`
		Status struct {
			AllowShare        bool   `json:"allow_share"`
			AwemeId           string `json:"aweme_id"`
			InReviewing       bool   `json:"in_reviewing"`
			IsDelete          bool   `json:"is_delete"`
			IsProhibited      bool   `json:"is_prohibited"`
			ListenVideoStatus int    `json:"listen_video_status"`
			PartSee           int    `json:"part_see"`
			PrivateStatus     int    `json:"private_status"`
			ReviewResult      struct {
				ReviewStatus int `json:"review_status"`
			} `json:"review_result"`
		} `json:"status"`
		SuggestWords struct {
			SuggestWords []struct {
				ExtraInfo string `json:"extra_info"`
				HintText  string `json:"hint_text"`
				IconUrl   string `json:"icon_url"`
				Scene     string `json:"scene"`
				Words     []struct {
					Info   string `json:"info"`
					Word   string `json:"word"`
					WordId string `json:"word_id"`
				} `json:"words"`
			} `json:"suggest_words"`
		} `json:"suggest_words,omitempty"`
		TextExtra []struct {
			End         int    `json:"end"`
			HashtagId   string `json:"hashtag_id,omitempty"`
			HashtagName string `json:"hashtag_name,omitempty"`
			IsCommerce  bool   `json:"is_commerce,omitempty"`
			Start       int    `json:"start"`
			Type        int    `json:"type"`
			SecUid      string `json:"sec_uid,omitempty"`
			UserId      string `json:"user_id,omitempty"`
		} `json:"text_extra"`
		TtsIdList           interface{} `json:"tts_id_list"`
		UniqidPosition      interface{} `json:"uniqid_position"`
		UserDigged          int         `json:"user_digged"`
		UserRecommendStatus int         `json:"user_recommend_status"`
		Video               struct {
			AnimatedCover struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"animated_cover"`
			BigThumbs interface{} `json:"big_thumbs"`
			BitRate   []struct {
				FPS       int    `json:"FPS"`
				HDRBit    string `json:"HDR_bit"`
				HDRType   string `json:"HDR_type"`
				BitRate   int    `json:"bit_rate"`
				GearName  string `json:"gear_name"`
				IsBytevc1 int    `json:"is_bytevc1"`
				IsH265    int    `json:"is_h265"`
				PlayAddr  struct {
					DataSize int      `json:"data_size"`
					FileCs   string   `json:"file_cs"`
					FileHash string   `json:"file_hash"`
					Height   int      `json:"height"`
					Uri      string   `json:"uri"`
					UrlKey   string   `json:"url_key"`
					UrlList  []string `json:"url_list"`
					Width    int      `json:"width"`
				} `json:"play_addr"`
				QualityType int    `json:"quality_type"`
				VideoExtra  string `json:"video_extra"`
			} `json:"bit_rate"`
			BitRateAudio interface{} `json:"bit_rate_audio"`
			Cover        struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover"`
			DownloadAddr struct {
				Height   int      `json:"height"`
				Uri      string   `json:"uri"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
				DataSize int      `json:"data_size,omitempty"`
			} `json:"download_addr"`
			Duration     int `json:"duration"`
			DynamicCover struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"dynamic_cover"`
			Height      int    `json:"height"`
			IsH265      int    `json:"is_h265"`
			IsSourceHDR int    `json:"is_source_HDR"`
			Meta        string `json:"meta"`
			OriginCover struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"origin_cover"`
			PlayAddr struct {
				DataSize int      `json:"data_size"`
				FileCs   string   `json:"file_cs"`
				FileHash string   `json:"file_hash"`
				Height   int      `json:"height"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
			} `json:"play_addr"`
			PlayAddr265 struct {
				DataSize int      `json:"data_size"`
				FileCs   string   `json:"file_cs"`
				FileHash string   `json:"file_hash"`
				Height   int      `json:"height"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
			} `json:"play_addr_265"`
			PlayAddrH264 struct {
				DataSize int      `json:"data_size"`
				FileCs   string   `json:"file_cs"`
				FileHash string   `json:"file_hash"`
				Height   int      `json:"height"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
			} `json:"play_addr_h264"`
			Ratio      string `json:"ratio"`
			VideoModel string `json:"video_model"`
			Width      int    `json:"width"`
		} `json:"video"`
		VideoGameDataChannelConfig struct {
		} `json:"video_game_data_channel_config"`
		VideoLabels []interface{} `json:"video_labels"`
		VideoTag    []struct {
			Level   int    `json:"level"`
			TagId   int    `json:"tag_id"`
			TagName string `json:"tag_name"`
		} `json:"video_tag"`
		VideoText         []interface{} `json:"video_text"`
		VoiceModifyIdList interface{}   `json:"voice_modify_id_list"`
		AwemeAcl          struct {
			DownloadMaskPanel struct {
				Code     int `json:"code"`
				ShowType int `json:"show_type"`
			} `json:"download_mask_panel"`
		} `json:"aweme_acl,omitempty"`
	} `json:"aweme_list"`
	HasMore int `json:"has_more"`
	LogPb   struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	MaxCursor          int64       `json:"max_cursor"`
	MinCursor          int64       `json:"min_cursor"`
	PostSerial         int         `json:"post_serial"`
	ReplaceSeriesCover int         `json:"replace_series_cover"`
	RequestItemCursor  int64       `json:"request_item_cursor"`
	StatusCode         int         `json:"status_code"`
	TimeList           interface{} `json:"time_list"`
}

// GetAllVideoWithName 获取作者视频的名字和下载地址
func (u *UserVideoResp) GetAllVideoWithName() map[string]string {
	videoSlice := u.AwemeList
	descVideoplayaddr := make(map[string]string, 100)
	for _, videoInfo := range videoSlice {
		desc := videoInfo.Desc
		var videoPlayAddr string
		if videoInfo.Video.PlayAddrH264.UrlList != nil {
			videoPlayAddr = videoInfo.Video.PlayAddrH264.UrlList[0]
		} else if videoInfo.Video.PlayAddr265.UrlList != nil {
			videoPlayAddr = videoInfo.Video.PlayAddr265.UrlList[0]
		} else if videoInfo.Video.PlayAddr.UrlList != nil {
			videoPlayAddr = videoInfo.Video.PlayAddr.UrlList[0]
		} else {
			fmt.Fprintf(os.Stdout, "视频:%s 无法下载", desc)
		}
		descVideoplayaddr[desc] = videoPlayAddr
	}
	return descVideoplayaddr
}
