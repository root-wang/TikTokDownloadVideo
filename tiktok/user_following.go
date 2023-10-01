package tiktok

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"tiktok/utils"
	"time"
)

type FollowingXBogusReq struct {
	DevicePlatform string `json:"device_platform"`
	Aid            string `json:"aid"`
	UserId         string `json:"user_id"`
	SecUserId      string `json:"sec_user_id"`
	MaxTime        string `json:"max_time"`
	MinTime        string `json:"min_time"`
	Offset         string `json:"offset"`
	Count          string `json:"count"`
	SourceType     string `json:"source_type"`
}

// 获取用户关注列表 返回用户的secId
func FollowingUsersSecId(secUserId string, UserId string, count int) []string {
	req := &FollowingXBogusReq{
		DevicePlatform: "android",
		Aid:            "6383",
		UserId:         UserId,
		SecUserId:      secUserId,
		Offset:         "0",
		MinTime:        "0",
		MaxTime:        "0",
		Count:          "20",
		SourceType:     "4",
	}

	xBogusUrl := utils.ConcatXBogusUrlString(req, FOLLOWINGURLPREFIX)
	followingUserListUrl := NewXBogusReq(xBogusUrl)

	followingUserReq, _ := utils.HttpNewRequest("GET", followingUserListUrl, nil)

	h := &http.Client{}
	respStruct := &UserFollowingResp{}

	var userSecIds = []string{}
	var resp *http.Response

	buff := bytes.NewBuffer(make([]byte, 0, 512))
	for {
		resp, _ = h.Do(followingUserReq)
		_, _ = io.Copy(buff, resp.Body)

		if resp.StatusCode == http.StatusOK && resp.ContentLength == 0 {
			fmt.Println("请求失败 2 秒后重试")
			resp.Body.Close()
			time.Sleep(2 * time.Second)
			continue
		} else if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Printf("请求失败 %d", resp.StatusCode)
		}
		_ = json.NewDecoder(buff).Decode(respStruct)

		for _, u := range respStruct.Followings {
			userSecIds = append(userSecIds, u.SecUid)
		}
		if len(userSecIds) >= count || !respStruct.HasMore {
			userSecIds = userSecIds[:count+1]
			break
		} else {
			offset, _ := strconv.Atoi(req.Offset)
			req.Offset = strconv.Itoa(offset + 20)
			followingUserListUrl = NewXBogusReq(utils.ConcatXBogusUrlString(req, FOLLOWINGURLPREFIX))
			followingUserReq, _ = utils.HttpNewRequest("GET", followingUserListUrl, nil)
			resp.Body.Close()
			buff.Reset()
		}
	}

	fmt.Printf("共获取到%d个关注用户的ID 即将开始下载这些用户中%d个的主页视频\n", len(userSecIds)-1, count)
	if count >= len(userSecIds) {
		count = len(userSecIds)
	}
	return userSecIds[:count]
}

type UserFollowingResp struct {
	Extra struct {
		FatalItemIds []interface{} `json:"fatal_item_ids"`
		LogId        string        `json:"logid"`
		Now          int64         `json:"now"`
	} `json:"extra"`
	Followings []struct {
		AcceptPrivatePolicy    bool        `json:"accept_private_policy"`
		AccountRegion          string      `json:"account_region"`
		AccountType            interface{} `json:"account_type"`
		Activity               interface{} `json:"activity"`
		ActivityLabel          interface{} `json:"activity_label"`
		AdCoverTitle           interface{} `json:"ad_cover_title"`
		AdCoverUrl             interface{} `json:"ad_cover_url"`
		AdOrderId              interface{} `json:"ad_order_id"`
		AgeGateAction          interface{} `json:"age_gate_action"`
		AgeGatePostAction      interface{} `json:"age_gate_post_action"`
		AgeGateTime            interface{} `json:"age_gate_time"`
		AllowStatus            interface{} `json:"allow_status"`
		AnchorInfo             interface{} `json:"anchor_info"`
		AnchorScheduleGuideTxt interface{} `json:"anchor_schedule_guide_txt"`
		AppleAccount           int         `json:"apple_account"`
		AuthorityStatus        int         `json:"authority_status"`
		Avatar168X168          struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_168x168"`
		Avatar300X300 struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_300x300"`
		AvatarDecoration   interface{} `json:"avatar_decoration"`
		AvatarDecorationId interface{} `json:"avatar_decoration_id"`
		AvatarLarger       struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_larger"`
		AvatarMedium struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_medium"`
		AvatarPendantLarger interface{} `json:"avatar_pendant_larger"`
		AvatarPendantMedium interface{} `json:"avatar_pendant_medium"`
		AvatarPendantThumb  interface{} `json:"avatar_pendant_thumb"`
		AvatarThumb         struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_thumb"`
		AvatarUpdateReminder interface{} `json:"avatar_update_reminder"`
		AvatarUri            string      `json:"avatar_uri"`
		AwemeControl         struct {
			CanComment     bool `json:"can_comment"`
			CanForward     bool `json:"can_forward"`
			CanShare       bool `json:"can_share"`
			CanShowComment bool `json:"can_show_comment"`
		} `json:"aweme_control"`
		AwemeCount                  int         `json:"aweme_count"`
		AwemeCover                  interface{} `json:"aweme_cover"`
		AwemeHotsoonAuth            *int        `json:"aweme_hotsoon_auth"`
		AwemeHotsoonAuthRelation    *int        `json:"aweme_hotsoon_auth_relation"`
		BanUserFunctions            []int       `json:"ban_user_functions"`
		BioEmail                    interface{} `json:"bio_email"`
		BioLocation                 interface{} `json:"bio_location"`
		BioPermission               interface{} `json:"bio_permission"`
		BioPhone                    interface{} `json:"bio_phone"`
		BioSecureUrl                interface{} `json:"bio_secure_url"`
		BioUrl                      interface{} `json:"bio_url"`
		BirthdayHideLevel           interface{} `json:"birthday_hide_level"`
		BizAccountInfo              interface{} `json:"biz_account_info"`
		BrandInfo                   interface{} `json:"brand_info"`
		CanModifyHometownInfo       interface{} `json:"can_modify_hometown_info"`
		CanModifySchoolInfo         interface{} `json:"can_modify_school_info"`
		CanSetGeofencing            interface{} `json:"can_set_geofencing"`
		CanShowGroupCard            interface{} `json:"can_show_group_card"`
		CancelType                  interface{} `json:"cancel_type"`
		CardEntries                 interface{} `json:"card_entries"`
		CardEntriesInfo             interface{} `json:"card_entries_info"`
		CardEntriesNotDisplay       interface{} `json:"card_entries_not_display"`
		CardSortPriority            interface{} `json:"card_sort_priority"`
		Category                    interface{} `json:"category"`
		ChaList                     interface{} `json:"cha_list"`
		CleanFollowingReason        interface{} `json:"clean_following_reason"`
		CollectCount                interface{} `json:"collect_count"`
		CommentFilterStatus         int         `json:"comment_filter_status"`
		CommentSetting              int         `json:"comment_setting"`
		CommerceBubble              interface{} `json:"commerce_bubble"`
		CommerceInfo                interface{} `json:"commerce_info"`
		CommercePermissions         interface{} `json:"commerce_permissions"`
		CommerceUserInfo            interface{} `json:"commerce_user_info"`
		CommerceUserLevel           int         `json:"commerce_user_level"`
		Constellation               int         `json:"constellation"`
		ContactName                 interface{} `json:"contact_name"`
		ContentLanguageAlreadyPopup interface{} `json:"content_language_already_popup"`
		CountStatus                 interface{} `json:"count_status"`
		CoverColour                 interface{} `json:"cover_colour"`
		CoverJumpUrl                interface{} `json:"cover_jump_url"`
		CoverUrl                    []struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_url"`
		CreateTime                              int         `json:"create_time"`
		CreatorLevel                            interface{} `json:"creator_level"`
		CustomVerify                            string      `json:"custom_verify"`
		CvLevel                                 string      `json:"cv_level"`
		DisplayInfo                             interface{} `json:"display_info"`
		DisplayWvalantineActivityEntry          interface{} `json:"display_wvalantine_activity_entry"`
		DogCardInfo                             interface{} `json:"dog_card_info"`
		DongtaiCount                            interface{} `json:"dongtai_count"`
		DormerGroup                             interface{} `json:"dormer_group"`
		DouPlusShareLocation                    interface{} `json:"dou_plus_share_location"`
		DouplusOldUser                          interface{} `json:"douplus_old_user"`
		DouplusToast                            interface{} `json:"douplus_toast"`
		DownloadPromptTs                        int         `json:"download_prompt_ts"`
		DownloadSetting                         int         `json:"download_setting"`
		DpLevel                                 interface{} `json:"dp_level"`
		DuetSetting                             int         `json:"duet_setting"`
		EffectDetail                            interface{} `json:"effect_detail"`
		EnableNearbyVisible                     bool        `json:"enable_nearby_visible"`
		EnableWish                              interface{} `json:"enable_wish"`
		EnterpriseUserInfo                      interface{} `json:"enterprise_user_info"`
		EnterpriseVerifyReason                  string      `json:"enterprise_verify_reason"`
		EverOver1KFollower                      interface{} `json:"ever_over_1k_follower"`
		FastCommentTexts                        interface{} `json:"fast_comment_texts"`
		FavoritingCount                         int         `json:"favoriting_count"`
		FbExpireTime                            int         `json:"fb_expire_time"`
		FollowAsSubscription                    interface{} `json:"follow_as_subscription"`
		FollowGuide                             interface{} `json:"follow_guide"`
		FollowStatus                            int         `json:"follow_status"`
		FollowVerifyStatus                      interface{} `json:"follow_verify_status"`
		FollowerCount                           int         `json:"follower_count"`
		FollowerRequestStatus                   int         `json:"follower_request_status"`
		FollowerStatus                          int         `json:"follower_status"`
		FollowingCount                          int         `json:"following_count"`
		FollowingListSecondaryInformationStruct struct {
			SecondaryInformationPriority int    `json:"secondary_information_priority"`
			SecondaryInformationText     string `json:"secondary_information_text"`
			SecondaryInformationTextType int    `json:"secondary_information_text_type"`
		} `json:"following_list_secondary_information_struct"`
		ForcePrivateAccount          interface{}   `json:"force_private_account"`
		ForwardCount                 interface{}   `json:"forward_count"`
		FriendCount                  interface{}   `json:"friend_count"`
		GeneralPermission            interface{}   `json:"general_permission"`
		Geofencing                   []interface{} `json:"geofencing"`
		GoogleAccount                string        `json:"google_account"`
		HasActivityMedal             interface{}   `json:"has_activity_medal"`
		HasCardEditPageEntrance      interface{}   `json:"has_card_edit_page_entrance"`
		HasEmail                     bool          `json:"has_email"`
		HasFacebookToken             bool          `json:"has_facebook_token"`
		HasHelpDeskEntrance          interface{}   `json:"has_help_desk_entrance"`
		HasInsights                  bool          `json:"has_insights"`
		HasOrders                    bool          `json:"has_orders"`
		HasStory                     interface{}   `json:"has_story"`
		HasSubscription              interface{}   `json:"has_subscription"`
		HasTwitterToken              bool          `json:"has_twitter_token"`
		HasUnreadStory               bool          `json:"has_unread_story"`
		HasYoutubeToken              bool          `json:"has_youtube_token"`
		HideFollowingFollowerList    interface{}   `json:"hide_following_follower_list"`
		HideLocation                 bool          `json:"hide_location"`
		HideSearch                   bool          `json:"hide_search"`
		HideShootButton              interface{}   `json:"hide_shoot_button"`
		HomepageBottomToast          interface{}   `json:"homepage_bottom_toast"`
		Hometown                     interface{}   `json:"hometown"`
		HometownFellowship           interface{}   `json:"hometown_fellowship"`
		HometownVisible              interface{}   `json:"hometown_visible"`
		HonorInfo                    interface{}   `json:"honor_info"`
		HotList                      interface{}   `json:"hot_list"`
		ImAgeStage                   interface{}   `json:"im_age_stage"`
		ImExaminationInfo            interface{}   `json:"im_examination_info"`
		ImSubscriptionPublisher      interface{}   `json:"im_subscription_publisher"`
		InfringementReportRemindInfo interface{}   `json:"infringement_report_remind_info"`
		InsId                        string        `json:"ins_id"`
		InterestTags                 interface{}   `json:"interest_tags"`
		IsActivityUser               interface{}   `json:"is_activity_user"`
		IsAdFake                     bool          `json:"is_ad_fake"`
		IsBindedWeibo                bool          `json:"is_binded_weibo"`
		IsBlock                      bool          `json:"is_block"`
		IsBlocked                    interface{}   `json:"is_blocked"`
		IsDisciplineMember           bool          `json:"is_discipline_member"`
		IsDouManager                 interface{}   `json:"is_dou_manager"`
		IsEffectArtist               interface{}   `json:"is_effect_artist"`
		IsEmailVerified              interface{}   `json:"is_email_verified"`
		IsEqualQuery                 interface{}   `json:"is_equal_query"`
		IsFlowcardMember             interface{}   `json:"is_flowcard_member"`
		IsGovMediaVip                bool          `json:"is_gov_media_vip"`
		IsLifeStyle                  interface{}   `json:"is_life_style"`
		IsMinor                      interface{}   `json:"is_minor"`
		IsMirror                     interface{}   `json:"is_mirror"`
		IsMixUser                    bool          `json:"is_mix_user"`
		IsNotShow                    bool          `json:"is_not_show"`
		IsPhoneBinded                bool          `json:"is_phone_binded"`
		IsProAccount                 interface{}   `json:"is_pro_account"`
		IsSeriesUser                 interface{}   `json:"is_series_user"`
		IsStar                       bool          `json:"is_star"`
		IsTop                        interface{}   `json:"is_top"`
		IsVerified                   bool          `json:"is_verified"`
		IsoCountryCode               interface{}   `json:"iso_country_code"`
		ItemList                     interface{}   `json:"item_list"`
		KyOnlyPredict                float64       `json:"ky_only_predict"`
		Language                     string        `json:"language"`
		LatestOrderTime              interface{}   `json:"latest_order_time"`
		LifeStoryBlock               interface{}   `json:"life_story_block"`
		LiveAgreement                int           `json:"live_agreement"`
		LiveAgreementTime            int           `json:"live_agreement_time"`
		LiveCommerce                 bool          `json:"live_commerce"`
		LiveStatus                   int           `json:"live_status"`
		LiveVerify                   int           `json:"live_verify"`
		LoginPlatform                interface{}   `json:"login_platform"`
		MessageChatEntry             interface{}   `json:"message_chat_entry"`
		MinorMode                    interface{}   `json:"minor_mode"`
		MplatformFollowersCount      interface{}   `json:"mplatform_followers_count"`
		MusicComplianceAccount       interface{}   `json:"music_compliance_account"`
		NameField                    interface{}   `json:"name_field"`
		NeedAddrCard                 interface{}   `json:"need_addr_card"`
		NeedPoints                   interface{}   `json:"need_points"`
		NeedRecommend                int           `json:"need_recommend"`
		NeiguangShield               int           `json:"neiguang_shield"`
		NewFriendType                int           `json:"new_friend_type"`
		NewStoryCover                interface{}   `json:"new_story_cover"`
		NewVisitorCount              interface{}   `json:"new_visitor_count"`
		Nickname                     string        `json:"nickname"`
		NicknameUpdateReminder       interface{}   `json:"nickname_update_reminder"`
		NormalTopCommentPermission   interface{}   `json:"normal_top_comment_permission"`
		NotSeenItemIdListV2          []string      `json:"not_seen_item_id_list_v2"`
		NotifyPrivateAccount         interface{}   `json:"notify_private_account"`
		OpenInsightTime              interface{}   `json:"open_insight_time"`
		OriginalMusician             struct {
			DiggCount      int `json:"digg_count"`
			MusicCount     int `json:"music_count"`
			MusicUsedCount int `json:"music_used_count"`
		} `json:"original_musician"`
		PersonalizedTag              interface{} `json:"personalized_tag"`
		PlatformSyncInfo             interface{} `json:"platform_sync_info"`
		PlayCount                    interface{} `json:"play_count"`
		PostDefaultDownloadSetting   interface{} `json:"post_default_download_setting"`
		PrExempt                     interface{} `json:"pr_exempt"`
		PreventDownload              bool        `json:"prevent_download"`
		PrivateAccountReviewReminder interface{} `json:"private_account_review_reminder"`
		PrivateAwemeCount            interface{} `json:"private_aweme_count"`
		ProAccountTcmRedDot          interface{} `json:"pro_account_tcm_red_dot"`
		ProfileCompletion            interface{} `json:"profile_completion"`
		ProfilePv                    interface{} `json:"profile_pv"`
		ProfileStory                 interface{} `json:"profile_story"`
		ProfileTabType               interface{} `json:"profile_tab_type"`
		PublishLandingTab            interface{} `json:"publish_landing_tab"`
		PunishRemindInfo             interface{} `json:"punish_remind_info"`
		QuickShopInfo                interface{} `json:"quick_shop_info"`
		RFansGroupInfo               interface{} `json:"r_fans_group_info"`
		ReactSetting                 int         `json:"react_setting"`
		RealnameVerifyStatus         interface{} `json:"realname_verify_status"`
		RecAgeStage                  interface{} `json:"rec_age_stage"`
		RecommendReason              string      `json:"recommend_reason"`
		RecommendReasonRelation      interface{} `json:"recommend_reason_relation"`
		RecommendScore               interface{} `json:"recommend_score"`
		RecommendTemplate            interface{} `json:"recommend_template"`
		RecommendUserReasonSource    interface{} `json:"recommend_user_reason_source"`
		ReflowPageGid                int         `json:"reflow_page_gid"`
		ReflowPageUid                int         `json:"reflow_page_uid"`
		RegisterFrom                 interface{} `json:"register_from"`
		RegisterTime                 interface{} `json:"register_time"`
		RelationLabel                interface{} `json:"relation_label"`
		RelationShip                 interface{} `json:"relation_ship"`
		RelativeUsers                interface{} `json:"relative_users"`
		RemarkName                   interface{} `json:"remark_name"`
		RoomCover                    interface{} `json:"room_cover"`
		RoomData                     *string     `json:"room_data"`
		RoomId                       int64       `json:"room_id"`
		RoomIdStr                    interface{} `json:"room_id_str"`
		RoomTypeTag                  interface{} `json:"room_type_tag"`
		SchoolAuth                   interface{} `json:"school_auth"`
		SchoolCategory               int         `json:"school_category"`
		SchoolId                     string      `json:"school_id"`
		SchoolVisible                interface{} `json:"school_visible"`
		SearchImpr                   struct {
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
		ShareQrcodeUri               string      `json:"share_qrcode_uri"`
		ShieldCommentNotice          int         `json:"shield_comment_notice"`
		ShieldDiggNotice             int         `json:"shield_digg_notice"`
		ShieldFollowNotice           int         `json:"shield_follow_notice"`
		ShopMicroApp                 interface{} `json:"shop_micro_app"`
		ShortId                      string      `json:"short_id"`
		ShowArtistPlaylist           interface{} `json:"show_artist_playlist"`
		ShowAvatarDecorationEntrance interface{} `json:"show_avatar_decoration_entrance"`
		ShowEffectList               interface{} `json:"show_effect_list"`
		ShowFavoriteList             interface{} `json:"show_favorite_list"`
		ShowFavoriteListOnItem       interface{} `json:"show_favorite_list_on_item"`
		ShowFirstAvatarDecoration    interface{} `json:"show_first_avatar_decoration"`
		ShowFollowingFollowerBanner  interface{} `json:"show_following_follower_banner"`
		ShowGenderStrategy           int         `json:"show_gender_strategy"`
		ShowImageBubble              bool        `json:"show_image_bubble"`
		ShowLocatedBanner            interface{} `json:"show_located_banner"`
		ShowMusicianCard             interface{} `json:"show_musician_card"`
		ShowNearbyActive             bool        `json:"show_nearby_active"`
		ShowPrivacyBanner            interface{} `json:"show_privacy_banner"`
		ShowPrivateTab               interface{} `json:"show_private_tab"`
		ShowRelationBanner           interface{} `json:"show_relation_banner"`
		ShowSecretBanner             interface{} `json:"show_secret_banner"`
		ShowSubscription             interface{} `json:"show_subscription"`
		ShowTelBookBanner            interface{} `json:"show_tel_book_banner"`
		ShowUserBanDialog            interface{} `json:"show_user_ban_dialog"`
		Signature                    string      `json:"signature"`
		SignatureDisplayLines        int         `json:"signature_display_lines"`
		SignatureLanguage            interface{} `json:"signature_language"`
		SpecialLock                  int         `json:"special_lock"`
		SpecialStateInfo             interface{} `json:"special_state_info"`
		SprintSupportUserInfo        interface{} `json:"sprint_support_user_info"`
		StarActivityEntrance         interface{} `json:"star_activity_entrance"`
		StarBillboardInfo            interface{} `json:"star_billboard_info"`
		StarBillboardRank            interface{} `json:"star_billboard_rank"`
		StarUseNewDownload           interface{} `json:"star_use_new_download"`
		Status                       int         `json:"status"`
		StitchSetting                int         `json:"stitch_setting"`
		StoryCount                   int         `json:"story_count"`
		StoryExpiredGuide            interface{} `json:"story_expired_guide"`
		StoryOpen                    bool        `json:"story_open"`
		String                       interface{} `json:"string"`
		SyncToToutiao                int         `json:"sync_to_toutiao"`
		TabSettings                  interface{} `json:"tab_settings"`
		ThirdName                    interface{} `json:"third_name"`
		TotalFavorited               int         `json:"total_favorited"`
		TwExpireTime                 int         `json:"tw_expire_time"`
		TwitterId                    string      `json:"twitter_id"`
		TwitterName                  string      `json:"twitter_name"`
		TypeLabel                    interface{} `json:"type_label"`
		Uid                          string      `json:"uid"`
		UniqueId                     string      `json:"unique_id"`
		UniqueIdModifyTime           int         `json:"unique_id_modify_time"`
		UniqueIdUpdateReminder       interface{} `json:"unique_id_update_reminder"`
		UrgeDetail                   struct {
			UserUrged int `json:"user_urged"`
		} `json:"urge_detail"`
		UserCanceled          bool        `json:"user_canceled"`
		UserDeleted           interface{} `json:"user_deleted"`
		UserMode              int         `json:"user_mode"`
		UserNotSee            int         `json:"user_not_see"`
		UserNotShow           int         `json:"user_not_show"`
		UserPeriod            int         `json:"user_period"`
		UserRate              int         `json:"user_rate"`
		UserRateRemindInfo    interface{} `json:"user_rate_remind_info"`
		UserRipEntry          interface{} `json:"user_rip_entry"`
		UserStoryCount        interface{} `json:"user_story_count"`
		UserTags              interface{} `json:"user_tags"`
		VcdAuthBlock          interface{} `json:"vcd_auth_block"`
		VerificationBadgeType interface{} `json:"verification_badge_type"`
		VerificationType      int         `json:"verification_type"`
		VerifyInfo            string      `json:"verify_info"`
		VersatileDisplay      interface{} `json:"versatile_display"`
		VideoCover            interface{} `json:"video_cover"`
		VideoIcon             struct {
			Height  int           `json:"height"`
			Uri     string        `json:"uri"`
			UrlList []interface{} `json:"url_list"`
			Width   int           `json:"width"`
		} `json:"video_icon"`
		VideoIconVirtualURI            interface{} `json:"video_icon_virtual_URI"`
		VideoUnreadInfo                interface{} `json:"video_unread_info"`
		VsPersonal                     interface{} `json:"vs_personal"`
		VxeTag                         interface{} `json:"vxe_tag"`
		WatchStatus                    interface{} `json:"watch_status"`
		WeiboName                      string      `json:"weibo_name"`
		WeiboSchema                    string      `json:"weibo_schema"`
		WeiboUrl                       string      `json:"weibo_url"`
		WeiboVerify                    string      `json:"weibo_verify"`
		WhiteCoverUrl                  interface{} `json:"white_cover_url"`
		WithCommerceEnterpriseTabEntry interface{} `json:"with_commerce_enterprise_tab_entry"`
		WithCommerceEntry              bool        `json:"with_commerce_entry"`
		WithCommerceNewbieTask         interface{} `json:"with_commerce_newbie_task"`
		WithDouEntry                   bool        `json:"with_dou_entry"`
		WithDouplusEntry               interface{} `json:"with_douplus_entry"`
		WithEcpEntry                   interface{} `json:"with_ecp_entry"`
		WithFusionShopEntry            bool        `json:"with_fusion_shop_entry"`
		WithItemCommerceEntry          interface{} `json:"with_item_commerce_entry"`
		WithLubanEntry                 interface{} `json:"with_luban_entry"`
		WithNewGoods                   interface{} `json:"with_new_goods"`
		WithShopEntry                  bool        `json:"with_shop_entry"`
		WithStarAtlasEntry             interface{} `json:"with_star_atlas_entry"`
		WithStickEntry                 interface{} `json:"with_stick_entry"`
		WithVisitorShopEntry           interface{} `json:"with_visitor_shop_entry"`
		WxInfo                         interface{} `json:"wx_info"`
		WxTag                          interface{} `json:"wx_tag"`
		YoutubeChannelId               string      `json:"youtube_channel_id"`
		YoutubeChannelTitle            string      `json:"youtube_channel_title"`
		YoutubeExpireTime              int         `json:"youtube_expire_time"`
		YoutubeLastRefreshTime         interface{} `json:"youtube_last_refresh_time"`
		YoutubeRefreshToken            interface{} `json:"youtube_refresh_token"`
		YtRawToken                     interface{} `json:"yt_raw_token"`
		ZeroPostUserTask               interface{} `json:"zero_post_user_task"`
	} `json:"followings"`
	HasMore        bool   `json:"has_more"`
	HotsoonHasMore int    `json:"hotsoon_has_more"`
	HotsoonText    string `json:"hotsoon_text"`
	LogPb          struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	MaxTime      int    `json:"max_time"`
	MinTime      int    `json:"min_time"`
	MyselfUserId string `json:"myself_user_id"`
	Offset       int    `json:"offset"`
	RecHasMore   bool   `json:"rec_has_more"`
	StatusCode   int    `json:"status_code"`
	StorePage    string `json:"store_page"`
	Total        int    `json:"total"`
	VcdCount     int    `json:"vcd_count"`
}
