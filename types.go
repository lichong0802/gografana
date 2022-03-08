package gografana




import (
	"encoding/json"
	"time"
)

// Board represents Grafana dashboard.
type Board struct {
	ID           uint     `json:"id,omitempty"`
	UID          string   `json:"uid"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	Editable     bool     `json:"editable"`
	HideControls bool     `json:"hideControls"`
	Style        string   `json:"style"`
	Timezone     string   `json:"timezone"`
	Version      uint     `json:"version"`
	IsStarred    bool     `json:"isStarred"`
	FolderId     uint     `json:"folderId"`
	FolderUid    string   `json:"folderUid"`
	FolderTitle  string   `json:"folderTitle"`
	FolderUrl    string   `json:"folderUrl"`
	Url          string   `json:"url,omitempty"` //TODO: ??
	Rows         []*Row   `json:"rows"`
	Annotations Annotations `json:"annotations"`
	Description string `json:"description"`
	FiscalYearStartMonth int `json:"fiscalYearStartMonth"`
	GnetID int `json:"gnetId"`
	GraphTooltip int `json:"graphTooltip"`
	Iteration int64 `json:"iteration"`
	Links []interface{} `json:"links"`
	LiveNow bool `json:"liveNow"`
	Panels []*Panel `json:"panels"`
	Refresh string `json:"refresh"`
	SchemaVersion int `json:"schemaVersion"`
	Templating Templating `json:"templating"`
	WeekStart string `json:"weekStart"`
}
type Annotations struct {
	List []struct {
		BuiltIn int `json:"builtIn"`
		Datasource string `json:"datasource"`
		Enable bool `json:"enable"`
		Hide bool `json:"hide"`
		IconColor string `json:"iconColor"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"list"`
}

type Templating struct {
	List []struct {
		Current struct {
			IsNone bool `json:"isNone"`
			Selected bool `json:"selected"`
			Text string `json:"text"`
			Value string `json:"value"`
		} `json:"current,omitempty"`
		Hide int `json:"hide"`
		IncludeAll bool `json:"includeAll"`
		Label string `json:"label"`
		Multi bool `json:"multi"`
		Name string `json:"name"`
		Options []interface{} `json:"options"`
		Query string `json:"query"`
		Refresh int `json:"refresh"`
		Regex string `json:"regex"`
		SkipURLSync bool `json:"skipUrlSync"`
		Type string `json:"type"`
		AllValue interface{} `json:"allValue,omitempty"`
		Datasource string `json:"datasource,omitempty"`
		Definition string `json:"definition,omitempty"`
		Sort int `json:"sort,omitempty"`
		TagValuesQuery string `json:"tagValuesQuery,omitempty"`
		Tags []interface{} `json:"tags,omitempty"`
		TagsQuery string `json:"tagsQuery,omitempty"`
		UseTags bool `json:"useTags,omitempty"`
	} `json:"list"`
}

type Custom struct {
}
type Defaults struct {
	Custom Custom `json:"custom"`
}
type FieldConfig struct {
	Defaults Defaults `json:"defaults"`
	Overrides []interface{} `json:"overrides"`
}
type Panel struct {
	CacheTimeout interface{} `json:"cacheTimeout,omitempty"`
	ColorBackground bool `json:"colorBackground,omitempty"`
	ColorValue bool `json:"colorValue,omitempty"`
	Colors []string `json:"colors,omitempty"`
	Datasource interface{} `json:"datasource"`
	Editable bool `json:"editable"`
	Error bool `json:"error"`
	FieldConfig *FieldConfig`json:"fieldConfig,omitempty"`
	Format string `json:"format,omitempty"`
	Gauge struct {
		MaxValue int `json:"maxValue"`
		MinValue int `json:"minValue"`
		Show bool `json:"show"`
		ThresholdLabels bool `json:"thresholdLabels"`
		ThresholdMarkers bool `json:"thresholdMarkers"`
	} `json:"gauge,omitempty"`
	GridPos struct {
		H int `json:"h"`
		W int `json:"w"`
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gridPos"`
	Height string `json:"height,omitempty"`
	ID int `json:"id"`
	Interval interface{} `json:"interval,omitempty"`
	IsNew bool `json:"isNew"`
	Links []interface{} `json:"links"`
	MappingType int `json:"mappingType,omitempty"`
	MappingTypes []struct {
		Name string `json:"name"`
		Value int `json:"value"`
	} `json:"mappingTypes,omitempty"`
	MaxDataPoints int `json:"maxDataPoints,omitempty"`
	NullPointMode string `json:"nullPointMode"`
	NullText interface{} `json:"nullText,omitempty"`
	Postfix string `json:"postfix,omitempty"`
	PostfixFontSize string `json:"postfixFontSize,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	PrefixFontSize string `json:"prefixFontSize,omitempty"`
	RangeMaps []struct {
		From string `json:"from"`
		Text string `json:"text"`
		To string `json:"to"`
	} `json:"rangeMaps,omitempty"`
	Sparkline struct {
		FillColor string `json:"fillColor"`
		Full bool `json:"full"`
		LineColor string `json:"lineColor"`
		Show bool `json:"show"`
	} `json:"sparkline,omitempty"`
	TableColumn string `json:"tableColumn,omitempty"`
	Targets []struct {
		Expr string `json:"expr"`
		IntervalFactor int `json:"intervalFactor"`
		LegendFormat string `json:"legendFormat"`
		Metric string `json:"metric"`
		RefID string `json:"refId"`
		Step int `json:"step"`
	} `json:"targets"`
	Thresholds interface{} `json:"thresholds"`
	Title string `json:"title"`
	Transparent bool `json:"transparent,omitempty"`
	Type string `json:"type"`
	ValueFontSize string `json:"valueFontSize,omitempty"`
	ValueMaps []struct {
		Op string `json:"op"`
		Text string `json:"text"`
		Value string `json:"value"`
	} `json:"valueMaps,omitempty"`
	ValueName string `json:"valueName,omitempty"`
	AliasColors struct {
	} `json:"aliasColors,omitempty"`
	Bars bool `json:"bars,omitempty"`
	DashLength int `json:"dashLength,omitempty"`
	Dashes bool `json:"dashes,omitempty"`
	Decimals int `json:"decimals,omitempty"`
	Fill int `json:"fill,omitempty"`
	FillGradient int `json:"fillGradient,omitempty"`
	Grid struct {
	} `json:"grid,omitempty"`
	HiddenSeries bool `json:"hiddenSeries,omitempty"`
	Legend struct {
		AlignAsTable bool `json:"alignAsTable"`
		Avg bool `json:"avg"`
		Current bool `json:"current"`
		Max bool `json:"max"`
		Min bool `json:"min"`
		RightSide bool `json:"rightSide"`
		Show bool `json:"show"`
		Total bool `json:"total"`
		Values bool `json:"values"`
	} `json:"legend,omitempty"`
	Lines bool `json:"lines,omitempty"`
	Linewidth int `json:"linewidth,omitempty"`
	Options struct {
		DataLinks []interface{} `json:"dataLinks"`
	} `json:"options,omitempty"`
	Percentage bool `json:"percentage,omitempty"`
	Pointradius int `json:"pointradius,omitempty"`
	Points bool `json:"points,omitempty"`
	Renderer string `json:"renderer,omitempty"`
	SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`
	SpaceLength int `json:"spaceLength,omitempty"`
	Stack bool `json:"stack,omitempty"`
	SteppedLine bool `json:"steppedLine,omitempty"`
	TimeFrom interface{} `json:"timeFrom,omitempty"`
	TimeRegions []interface{} `json:"timeRegions,omitempty"`
	TimeShift interface{} `json:"timeShift,omitempty"`
	Tooltip struct {
		MsResolution bool `json:"msResolution"`
		Shared bool `json:"shared"`
		Sort int `json:"sort"`
		ValueType string `json:"value_type"`
	} `json:"tooltip,omitempty"`
	Xaxis struct {
		Buckets interface{} `json:"buckets"`
		Mode string `json:"mode"`
		Name interface{} `json:"name"`
		Show bool `json:"show"`
		Values []interface{} `json:"values"`
	} `json:"xaxis,omitempty"`
	Yaxes []struct {
		Format string `json:"format"`
		Label interface{} `json:"label"`
		LogBase int `json:"logBase"`
		Max interface{} `json:"max"`
		Min interface{} `json:"min"`
		Show bool `json:"show"`
	} `json:"yaxes,omitempty"`
	Yaxis struct {
		Align bool `json:"align"`
		AlignLevel interface{} `json:"alignLevel"`
	} `json:"yaxis,omitempty"`
}





type CreateDashboardRequest struct {
	Board Board `json:"dashboard"`
	//The id of the folder to save the dashboard in.
	FolderId uint `json:"folderId"`
	//Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite bool `json:"overwrite,omitempty"`
	//Set a commit message for the version history.
	Message string `json:"message"`
}

type CreateFolderRequest struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type CreateFolderResponse struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasAcl    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

type CreateAPIKeyResponse struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

/*
Status Codes:
---------------
200 – Created
400 – Errors (invalid json, missing or invalid fields, etc)
401 – Unauthorized
403 – Access denied
412 – Precondition failed

The 412 status code is used for explaing that you cannot create the dashboard and why. There can be different reasons for this:

The dashboard has been changed by someone else, status=version-mismatch
A dashboard with the same name in the folder already exists, status=name-exists
A dashboard with the same uid already exists, status=name-exists
The dashboard belongs to plugin <plugin title>, status=plugin-dashboard
*/
type CreateDashboardResponse struct {
	ID      uint   `json:"id,omitempty"`
	UID     string `json:"uid"`
	Url     string `json:"url"`
	Status  string `json:"status"`
	Version uint   `json:"version"`
	Message string `json:"message,omitempty"`
}

type Panel_5_0 struct {
	AliasColors map[string]string `json:"aliasColors"`
	Alert       *struct {
		AlertRuleTags struct {
		} `json:"alertRuleTags"`
		Conditions []struct {
			Evaluator struct {
				Params []int64 `json:"params"`
				Type   string  `json:"type"`
			} `json:"evaluator"`
			Operator struct {
				Type string `json:"type"`
			} `json:"operator"`
			Query struct {
				Params []string `json:"params"`
			} `json:"query"`
			Reducer struct {
				Params []interface{} `json:"params"`
				Type   string        `json:"type"`
			} `json:"reducer"`
			Type string `json:"type"`
		} `json:"conditions"`
		ExecutionErrorState string `json:"executionErrorState"`
		For                 string `json:"for"`
		Frequency           string `json:"frequency"`
		Handler             int    `json:"handler"`
		Message             string `json:"message"`
		Name                string `json:"name"`
		NoDataState         string `json:"noDataState"`
		Notifications       []struct {
			ID int `json:"id"`
		} `json:"notifications"`
	} `json:"alert,omitempty"`
	Bars       bool   `json:"bars"`
	DashLength int    `json:"dashLength"`
	Dashes     bool   `json:"dashes"`
	Datasource string `json:"datasource"`
	Fill       int    `json:"fill"`
	GridPos    struct {
		H int `json:"h"`
		W int `json:"w"`
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gridPos,omitempty"`
	ID     int `json:"id"`
	Legend struct {
		Avg          bool `json:"avg"`
		Current      bool `json:"current"`
		Max          bool `json:"max"`
		Min          bool `json:"min"`
		Show         bool `json:"show"`
		Total        bool `json:"total"`
		Values       bool `json:"values"`
		AlignAsTable bool `json:"alignAsTable"`
	} `json:"legend,omitempty"`
	Lines           bool          `json:"lines"`
	Linewidth       int           `json:"linewidth"`
	Links           []interface{} `json:"links"`
	NullPointMode   string        `json:"nullPointMode"`
	Percentage      bool          `json:"percentage"`
	Pointradius     int           `json:"pointradius"`
	Points          bool          `json:"points"`
	Renderer        string        `json:"renderer"`
	SeriesOverrides []interface{} `json:"seriesOverrides"`
	SpaceLength     int           `json:"spaceLength"`
	Stack           bool          `json:"stack"`
	SteppedLine     bool          `json:"steppedLine"`
	Targets         []struct {
		Expr           string `json:"expr"`
		Format         string `json:"format"`
		Instant        bool   `json:"instant"`
		IntervalFactor int    `json:"intervalFactor"`
		LegendFormat   string `json:"legendFormat"`
		RefID          string `json:"refId"`
	} `json:"targets"`
	Thresholds []interface{} `json:"thresholds"`
	TimeFrom   interface{}   `json:"timeFrom"`
	TimeShift  interface{}   `json:"timeShift"`
	Title      string        `json:"title"`
	Tooltip    struct {
		Shared    bool   `json:"shared"`
		Sort      int    `json:"sort"`
		ValueType string `json:"value_type"`
	} `json:"tooltip,omitempty"`
	Transparent bool   `json:"transparent"`
	Type        string `json:"type"`
	Xaxis       struct {
		Buckets interface{}   `json:"buckets"`
		Mode    string        `json:"mode"`
		Name    interface{}   `json:"name"`
		Show    bool          `json:"show"`
		Values  []interface{} `json:"values"`
	} `json:"xaxis,omitempty"`
	Yaxes []struct {
		Format  string      `json:"format"`
		Label   interface{} `json:"label"`
		LogBase int         `json:"logBase"`
		Max     interface{} `json:"max"`
		Min     interface{} `json:"min"`
		Show    bool        `json:"show"`
	} `json:"yaxes,omitempty"`
}

func (p *Panel_5_0) MarshalJSON() ([]byte, error) {
	type Alias Panel_5_0
	aliasColors := map[string]string{}
	if p.AliasColors != nil {
		aliasColors = p.AliasColors
	}
	return json.Marshal(&struct {
		AliasColors map[string]string `json:"aliasColors"`
		*Alias
	}{
		AliasColors: aliasColors,
		Alias:       (*Alias)(p),
	})
}

type Row struct {
	Title     string      `json:"title"`
	ShowTitle bool        `json:"showTitle"`
	Collapse  bool        `json:"collapse"`
	Editable  bool        `json:"editable"`
	Height    string      `json:"height"`
	Panels    []Panel_5_0 `json:"panels"`
}

type GetDashboardByUIdResponse struct {
	Meta struct {
		Type        string    `json:"type"`
		CanSave     bool      `json:"canSave"`
		CanEdit     bool      `json:"canEdit"`
		CanAdmin    bool      `json:"canAdmin"`
		CanStar     bool      `json:"canStar"`
		Slug        string    `json:"slug"`
		URL         string    `json:"url"`
		Expires     time.Time `json:"expires"`
		Created     time.Time `json:"created"`
		Updated     time.Time `json:"updated"`
		UpdatedBy   string    `json:"updatedBy"`
		CreatedBy   string    `json:"createdBy"`
		Version     int       `json:"version"`
		HasAcl      bool      `json:"hasAcl"`
		IsFolder    bool      `json:"isFolder"`
		FolderID    int       `json:"folderId"`
		FolderTitle string    `json:"folderTitle"`
		FolderURL   string    `json:"folderUrl"`
	} `json:"meta"`
	Dashboard Board `json:"dashboard"`
}

type DataSource struct {
	ID          int                    `json:"id"`
	OrgID       int                    `json:"orgId"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	TypeLogoURL string                 `json:"typeLogoUrl"`
	Access      string                 `json:"access"`
	URL         string                 `json:"url"`
	Password    string                 `json:"password"`
	User        string                 `json:"user"`
	Database    string                 `json:"database"`
	BasicAuth   bool                   `json:"basicAuth"`
	IsDefault   bool                   `json:"isDefault"`
	JSONData    map[string]interface{} `json:"jsonData"`
	ReadOnly    bool                   `json:"readOnly"`
}

type CreateDataSourceResponse struct {
	DataSource *DataSource `json:"datasource"`
	ID         int         `json:"id"`
	Message    string      `json:"message"`
	Name       string      `json:"name"`
}

type CreateNotificationChannelResponse struct {
	ID           int    `json:"id"`
	UID          string `json:"uid"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	IsDefault    bool   `json:"isDefault"`
	SendReminder bool   `json:"sendReminder"`
	Settings     struct {
		Addresses string `json:"addresses"`
	} `json:"settings"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type APIKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Folder struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasACL    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

type NotificationChannel struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	Type                  string    `json:"type"`
	IsDefault             bool      `json:"isDefault"`
	SendReminder          bool      `json:"sendReminder"`
	DisableResolveMessage bool      `json:"disableResolveMessage"`
	Frequency             string    `json:"frequency"`
	Created               time.Time `json:"created"`
	Updated               time.Time `json:"updated"`
	Settings              struct {
		Addresses   string `json:"addresses"`
		AutoResolve bool   `json:"autoResolve"`
		HTTPMethod  string `json:"httpMethod"`
		UploadImage bool   `json:"uploadImage"`
	} `json:"settings"`
}
