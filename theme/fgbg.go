package theme

import (
	"image/color"
)

type Kind string

const (
	KindPrimary   Kind = "primary"
	KindSecondary Kind = "secondary"
	KindInfo      Kind = "info"
	KindSuccess   Kind = "success"
	KindNotice    Kind = "notice"
	KindWarning   Kind = "warning"
	KindDanger    Kind = "danger"
)

func (th *Theme) FgBg(kind Kind, component string) (color.NRGBA, color.NRGBA) {
	cm := th.GetComponent(component)

	switch kind {
	case KindPrimary:
		return cm.Primary, cm.Surface
	case KindSecondary:
		return cm.Secondary, cm.Surface
	case KindInfo:
		return cm.Info, cm.Surface
	case KindSuccess:
		return cm.Success, cm.Surface
	case KindNotice:
		return cm.Notice, cm.Surface
	case KindWarning:
		return cm.Warning, cm.Surface
	case KindDanger:
		return cm.Danger, cm.Surface
	}
	return color.NRGBA{}, color.NRGBA{}
}
