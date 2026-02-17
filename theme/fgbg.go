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

func (th *Theme) FgBgTxt(kind Kind, component string) (color.NRGBA, color.NRGBA, color.NRGBA) {
	cm := th.GetComponent(component)

	switch kind {
	case KindPrimary:
		return cm.Primary, cm.Surface, cm.Text
	case KindSecondary:
		return cm.Secondary, cm.Surface, cm.Text
	case KindInfo:
		return cm.Info, cm.Surface, cm.Text
	case KindSuccess:
		return cm.Success, cm.Surface, cm.Text
	case KindNotice:
		return cm.Notice, cm.Surface, cm.Text
	case KindWarning:
		return cm.Warning, cm.Surface, cm.Text
	case KindDanger:
		return cm.Danger, cm.Surface, cm.Text
	}
	return color.NRGBA{}, color.NRGBA{}, color.NRGBA{}
}
