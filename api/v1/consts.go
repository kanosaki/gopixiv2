package v1

type RankingMode string

const (
	RankingModeDay          RankingMode = "day"
	RankingModeWeek         RankingMode = "week"
	RankingModeMonth        RankingMode = "month"
	RankingModeDayMale      RankingMode = "male"
	RankingModeDayFemale    RankingMode = "female"
	RankingModeWeekOriginal RankingMode = "week_original"
	RankingModeWeekRookie   RankingMode = "week_rookie"
	RankingModeDayR18       RankingMode = "day_r18"
	RankingModeDayMaleR18   RankingMode = "day_male_r18"
	RankingModeDayFemaleR18 RankingMode = "day_female_r18"
	RankingModeWeekR18      RankingMode = "week_r18"
	RankingModeWeekR18G     RankingMode = "week_r18g"
)

type SearchTarget string

const (
	SearchTargetPartialMatchForTags SearchTarget = "partial_match_for_tags"
	SearchTargetExactMatchForTags   SearchTarget = "exact_match_for_tags"
	SearchTargetTitleAndCaption     SearchTarget = "title_and_caption"
)

type SearchSortOrder string

const (
	SearchSortOrderDateDesc          SearchSortOrder = "date_desc"
	SearchSortOrderDateAsc           SearchSortOrder = "date_asc"
	SearchSortOrderPopularDesc       SearchSortOrder = "popular_desc"
	SearchSortOrderPopularMaleDesc   SearchSortOrder = "popular_male_desc"
	SearchSortOrderPopularFemaleDesc SearchSortOrder = "popular_female_desc"
)
