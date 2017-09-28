package tick

import (
	"github.com/influxdata/kapacitor/pipeline"
	"github.com/influxdata/kapacitor/tick/ast"
)

// Alert converts the Alert pipeline node into the TICKScript AST
type Alert struct {
	Function
}

// NewAlert creates an Alert function builder
func NewAlert(parents []ast.Node) *Alert {
	return &Alert{
		Function{
			Parents: parents,
		},
	}
}

// Build creates a Alert ast.Node
func (n *Alert) Build(a *pipeline.AlertNode) (ast.Node, error) {
	n.Pipe("alert").
		Dot("topic", a.Topic).
		Dot("id", a.Id).
		Dot("message", a.Message).
		Dot("details", a.Details).
		Dot("info", a.Info).
		Dot("warn", a.Warn).
		Dot("crit", a.Crit).
		Dot("infoReset", a.InfoReset).
		Dot("warnReset", a.WarnReset).
		Dot("critReset", a.CritReset).
		Dot("history", a.History).
		Dot("levelTag", a.LevelTag).
		Dot("levelField", a.LevelField).
		Dot("messageField", a.MessageField).
		Dot("durationField", a.DurationField).
		Dot("idTag", a.IdTag).
		Dot("idField", a.IdField).
		DotIf("all", a.AllFlag).
		DotIf("noRecoveries", a.NoRecoveriesFlag).
		Dot("", a.IsStateChangesOnly).
		Dot("idField", a.IdField).
		Dot("idField", a.IdField)

	if a.IsStateChangesOnly {
		if a.StateChangesOnlyDuration == 0 {
			n.Dot("stateChangesOnly")
		} else {
			n.Dot("stateChangesOnly", a.StateChangesOnlyDuration)
		}
	}

	if a.UseFlapping {
		n.DotZeroValueOK("flapping", a.FlapLow, a.FlapHigh)
	}

	for _, h := range a.HTTPPostHandlers {
		n.Dot("post", h.URL).
			Dot("endpoint", h.Endpoint)
		for k, v := range h.Headers {
			n.Dot("header", k, v)
		}
	}

	for _, h := range a.TcpHandlers {
		n.Dot("tcp").
			Dot("address", h.Address)
	}

	for _, h := range a.EmailHandlers {
		n.Dot("email")
		for _, to := range h.ToList {
			n.Dot("to", to)
		}
	}

	for _, h := range a.ExecHandlers {
		n.Dot("exec", args(h.Command))
	}

	for _, h := range a.LogHandlers {
		n.Dot("log", h.FilePath)
		mode := ast.NumberNode{
			IsInt: true,
			Int64: h.Mode,
			Base:  8,
		}
		n.Dot("mode", mode)
	}

	for _, h := range a.VictorOpsHandlers {
		n.Dot("victorOps").
			Dot("routingKey", h.RoutingKey)
	}

	for _, h := range a.PagerDutyHandlers {
		n.Dot("pagerDuty").
			Dot("serviceKey", h.ServiceKey)
	}

	for _, h := range a.PushoverHandlers {
		n.Dot("pushover").
			Dot("userKey", h.UserKey).
			Dot("device", h.Device).
			Dot("title", h.Title).
			Dot("url", h.URL).
			Dot("urlTitle", h.URLTitle).
			Dot("sound", h.Sound)
	}

	for _, h := range a.SensuHandlers {
		n.Dot("sensu").
			Dot("source", h.Source).
			Dot("handlers", args(h.HandlersList))
	}

	for _, h := range a.SlackHandlers {
		n.Dot("slack").
			Dot("channel", h.Channel).
			Dot("username", h.Username).
			Dot("iconEmoji", h.IconEmoji)
	}

	for _, h := range a.TelegramHandlers {
		n.Dot("telegram").
			Dot("chatId", h.ChatId).
			Dot("parseMode", h.ParseMode).
			DotIf("disableWebPagePreview", h.IsDisableWebPagePreview).
			DotIf("disableNotification", h.IsDisableNotification)
	}

	for _, h := range a.HipChatHandlers {
		n.Dot("hipChat").
			Dot("room", h.Room).
			Dot("token", h.Token)
	}

	for _, h := range a.AlertaHandlers {
		n.Dot("alerta").
			Dot("token", h.Token).
			Dot("resource", h.Resource).
			Dot("event", h.Event).
			Dot("environment", h.Environment).
			Dot("group", h.Group).
			Dot("value", h.Value).
			Dot("origin", h.Origin).
			Dot("services", args(h.Service))
	}

	for _, h := range a.OpsGenieHandlers {
		n.Dot("opsGenie").
			Dot("teams", args(h.TeamsList)).
			Dot("recipients", args(h.RecipientsList))
	}

	for _ = range a.TalkHandlers {
		n.Dot("talk")
	}

	for _, h := range a.MQTTHandlers {
		n.Dot("mqtt").
			Dot("brokerName", h.BrokerName).
			Dot("topic", h.Topic).
			Dot("qos", h.Qos).
			Dot("retained", h.Retained)
	}

	for _, h := range a.SNMPTrapHandlers {
		n.Dot("snmpTrap", h.TrapOid)
		for _, d := range h.DataList {
			n.Dot("data", d.Oid, d.Type, d.Value)
		}
	}

	return n.prev, n.err
}

func args(a []string) []interface{} {
	r := make([]interface{}, len(a))
	for i := range a {
		r[i] = a[i]
	}
	return r
}

func largs(a []*ast.LambdaNode) []interface{} {
	r := make([]interface{}, len(a))
	for i := range a {
		r[i] = a[i]
	}
	return r
}