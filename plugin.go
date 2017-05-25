package kong

type RequestTransformer struct {
	Name               string `url:"name"`
	ConsumerId         string `url:"consumer_id,omitempty"`
	RemoveHeaders      string `url:"config.remove.headers,omitempty"`
	RemoveQuerystring  string `url:"config.remove.querystring,omitempty"`
	RemoveBody         string `url:"config.remove.body,omitempty"`
	ReplaceHeaders     string `url:"config.replace.headers,omitempty"`
	ReplaceQuerystring string `url:"config.replace.querystring,omitempty"`
	ReplaceBody        string `url:"config.replace.body,omitempty"`
	AddHeaders         string `url:"config.add.headers,omitempty"`
	AddQuerystring     string `url:"config.add.querystring,omitempty"`
	AddBody            string `url:"config.add.body,omitempty"`
	AppendHeaders      string `url:"config.append.headers,omitempty"`
	AppendQuerystring  string `url:"config.append.querystring,omitempty"`
	AppendBody         string `url:"config.append.body,omitempty"`
}

type RateLimiting struct {
	Name                string `url:"name"`
	ConsumerId          string `url:"consumer_id,omitempty"`
	ConfigSecond        string `url:"config.second,omitempty"`
	ConfigMinute        string `url:"config.minute,omitempty"`
	ConfigHour          string `url:"config.hour,omitempty"`
	ConfigDay           string `url:"config.day,omitempty"`
	ConfigMonth         string `url:"config.month,omitempty"`
	ConfigYear          string `url:"config.year,omitempty"`
	ConfigLimitBy       string `url:"config.limit_by,omitempty"`
	ConfigPolicy        string `url:"config.policy,omitempty"`
	ConfigFaultTolerant string `url:"config.fault_tolerant,omitempty"`
	ConfigRedisHost     string `url:"config.redis_host,omitempty"`
	ConfigRedisPort     string `url:"config.redis_port,omitempty"`
	ConfigRedisPassword string `url:"config.redis_password,omitempty"`
	ConfigRedisTimeout  string `url:"config.redis_timeout,omitempty"`
	ConfigRedisDatabase string `url:"config.redis_database,omitempty"`
}
