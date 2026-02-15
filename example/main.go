package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/chapar-rest/uikit/actionbar"
	"github.com/chapar-rest/uikit/button"
	"github.com/chapar-rest/uikit/divider"
	"github.com/chapar-rest/uikit/icons"
	"github.com/chapar-rest/uikit/sidebar"
	"github.com/chapar-rest/uikit/split"
	"github.com/chapar-rest/uikit/tabs"
	"github.com/chapar-rest/uikit/theme"
	"github.com/chapar-rest/uikit/theme/themes"
	"github.com/chapar-rest/uikit/toggle"
	"github.com/chapar-rest/uikit/treeview"
	"github.com/oligo/gvcode"
	"github.com/oligo/gvcode/addons/completion"
	gvcolor "github.com/oligo/gvcode/color"
	"github.com/oligo/gvcode/textstyle/syntax"
	wg "github.com/oligo/gvcode/widget"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Size(unit.Dp(800), unit.Dp(700)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := themes.Light()
	fonts := theme.BuiltinFonts()
	th.WithFonts(fonts)

	state := &appState{
		split: &split.Split{
			Axis:  layout.Horizontal,
			Ratio: 0.3,
			HandleStyle: split.HandleStyle{
				Color:      th.Base.Border,
				Width:      unit.Dp(3),
				HoverColor: th.Base.Secondary,
			},
		},
		sidebar:      sidebar.New(),
		actionbar:    actionbar.NewActionBar(layout.Horizontal, layout.Start, layout.SpaceAround),
		appBar:       actionbar.NewActionBar(layout.Horizontal, layout.Start, layout.SpaceBetween),
		theme:        th,
		openFiles:    make(map[string]fileView),
		openTabs:     make(map[string]*tabs.Tab),
		openPaths:    make([]string, 0),
		tabToPath:    make(map[*tabs.Tab]string),
		projectIndex: make([]string, 0),
		memberIndex:  make(map[string][]string),
	}
	state.tree = state.buildFileTree(th)
	state.tabitems = tabs.NewTabs()
	state.buildProjectIndex()

	state.actionbar.AddItem(button.IconButton(state.theme, &state.NewFileClickable, icons.FileAdd, theme.KindPrimary))
	state.actionbar.AddItem(button.IconButton(state.theme, &state.SearchClickable, icons.Search, theme.KindPrimary))
	state.actionbar.AddItem(button.IconButton(state.theme, &state.OpenFileClickable, icons.FileInput, theme.KindPrimary))
	state.actionbar.AddItem(button.IconButton(state.theme, &state.HistoryClickable, icons.History, theme.KindPrimary))

	state.themeToggleClickable = toggle.NewToggleButton(state.theme, []*toggle.State{
		{
			Tag:  "dark",
			Icon: icons.Moon,
		},
		{
			Tag:  "light",
			Icon: icons.Sun,
		},
	})

	state.appBar.AddItem(actionbar.ActionBarItemFunc(func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
		return material.Label(th.Material(), unit.Sp(14), "VOID editor").Layout(gtx)
	}))

	state.appBar.AddItem(actionbar.ActionBarItemFunc(func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
		return state.themeToggleClickable.Layout(gtx, th)
	}))

	state.sidebar.AddNavItem(sidebar.Item{
		Tag:  "files",
		Name: "Files",
		Icon: icons.Files,
	})

	state.sidebar.AddNavItem(sidebar.Item{
		Tag:  "history",
		Name: "History",
		Icon: icons.History,
	})

	state.sidebar.AddNavItem(sidebar.Item{
		Tag:  "search",
		Name: "Search",
		Icon: icons.Search,
	})

	state.sidebar.AddNavItem(sidebar.Item{
		Tag:  "settings",
		Name: "Settings",
		Icon: icons.Settings,
	})

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			state.appLayout(gtx)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}

type appState struct {
	SearchClickable   widget.Clickable
	NewFileClickable  widget.Clickable
	OpenFileClickable widget.Clickable
	HistoryClickable  widget.Clickable

	sidebar *sidebar.Sidebar

	currentTheme         string
	themeToggleClickable *toggle.ToggleButton

	split *split.Split

	tree *treeview.Tree

	theme *theme.Theme

	tabitems *tabs.Tabs

	actionbar *actionbar.ActionBar

	appBar *actionbar.ActionBar

	openFiles    map[string]fileView
	openTabs     map[string]*tabs.Tab
	openPaths    []string             // path order matching tab order
	tabToPath    map[*tabs.Tab]string // tab -> path for close callback
	projectIndex []string             // unique identifiers from project files (for completion)
	memberIndex  map[string][]string  // receiver -> members seen after "receiver." in project
}

type fileView struct {
	Title           string
	Path            string
	Editor          *gvcode.Editor
	OriginalContent string                      // content when file was opened (from disk)
	OnChange        func(currentContent string) // called when editor content changes; pass ed.Text() to update dirty state
	Layout          func(gtx layout.Context, th *theme.Theme) layout.Dimensions
}

func (s *appState) onFileNodeClick(node *treeview.Node) {
	path := node.ID
	if _, ok := s.openFiles[path]; ok {
		if tab := s.openTabs[path]; tab != nil {
			s.tabitems.SelectTab(tab)
		}
		return
	}

	s.openFiles[path] = s.buildFileView(s.theme, path)

	t := tabs.NewTab(func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
		lb := material.Label(th.Material(), unit.Sp(14), s.openFiles[path].Title)
		return lb.Layout(gtx)
	})

	t.OnCloseFunc = func(tab *tabs.Tab) bool {
		path := s.tabToPath[tab]
		delete(s.openFiles, path)
		delete(s.openTabs, path)
		delete(s.tabToPath, tab)
		if i := slices.Index(s.openPaths, path); i >= 0 {
			s.openPaths = slices.Delete(s.openPaths, i, i+1)
		}
		return true
	}

	t.State = tabs.TabStateClean
	s.tabitems.AddTab(t)
	s.openTabs[path] = t
	s.openPaths = append(s.openPaths, path)
	s.tabToPath[t] = path
}

func (s *appState) appLayout(gtx layout.Context) {

	th := s.theme
	if s.currentTheme != s.themeToggleClickable.StateTag() {
		s.currentTheme = s.themeToggleClickable.StateTag()
		if s.currentTheme == "dark" {
			th = themes.Dark()
		} else {
			th = themes.Light()
		}
		s.theme = th
	}

	paint.Fill(gtx.Ops, th.Base.Surface)

	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top:    unit.Dp(12),
				Left:   unit.Dp(8),
				Right:  unit.Dp(8),
				Bottom: unit.Dp(12),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return s.appBar.Layout(gtx, th)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return divider.NewDivider(layout.Horizontal, unit.Dp(1)).Layout(gtx, th)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.X = gtx.Dp(60)
					return s.sidebar.Layout(gtx, th)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return divider.NewDivider(layout.Vertical, unit.Dp(1)).Layout(gtx, th)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return s.split.Layout(gtx,
						func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis: layout.Vertical,
							}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return s.actionbar.Layout(gtx, s.theme)
									})
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return divider.NewDivider(layout.Horizontal, unit.Dp(1)).Layout(gtx, s.theme)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return s.tree.Layout(gtx, s.theme)
								}),
							)
						},
						func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis: layout.Vertical,
							}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return s.tabitems.Layout(gtx, s.theme)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									if s.tabitems.CurrentView() < 0 || s.tabitems.CurrentView() >= len(s.openPaths) {
										return layout.Dimensions{}
									}
									path := s.openPaths[s.tabitems.CurrentView()]
									fv, ok := s.openFiles[path]
									if !ok {
										return layout.Dimensions{}
									}
									return fv.Layout(gtx, s.theme)
								}),
							)
						},
					)
				}),
			)
		}),
	)
}

func (s *appState) buildFileTree(th *theme.Theme) *treeview.Tree {
	tree := treeview.NewTree()

	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return tree
	}

	var ignoreList = []string{".git", ".idea", ".vscode", ".DS_Store", ".env"}

	for _, entry := range entries {
		if slices.Contains(ignoreList, entry.Name()) {
			continue
		}

		node := s.buildFileNode(th, entry, ".")
		if node != nil {
			tree.Insert(node)
		}
	}

	return tree
}

func (s *appState) buildFileNode(th *theme.Theme, entry os.DirEntry, parentPath string) *treeview.Node {
	name := entry.Name()
	fullPath := filepath.Join(parentPath, name)

	node := treeview.NewNode(fullPath, func(gtx layout.Context) layout.Dimensions {
		return material.Label(th.Material(), unit.Sp(14), name).Layout(gtx)
	})

	node.OnClickFunc = func(node *treeview.Node) {
		s.onFileNodeClick(node)
	}

	if entry.IsDir() {
		dirEntries, err := os.ReadDir(fullPath)
		if err == nil {
			for _, childEntry := range dirEntries {
				childNode := s.buildFileNode(th, childEntry, fullPath)
				if childNode != nil {
					node.AddChild(childNode)
				}
			}
		}
	}

	return node
}

var projectIndexIgnoreDirs = []string{".git", ".idea", ".vscode", "node_modules", "vendor"}

// buildProjectIndex walks the project directory and indexes identifiers and "receiver.member" pairs for completion.
func (s *appState) buildProjectIndex() {
	seen := make(map[string]bool)
	members := make(map[string]map[string]bool) // receiver -> set of members
	filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if slices.Contains(projectIndexIgnoreDirs, d.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		switch ext {
		case ".go", ".mod", ".sum", ".txt", ".md", ".yaml", ".yml", ".json", ".toml":
			// ok
		default:
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		if len(content) > 512*1024 {
			return nil
		}
		text := string(content)
		for _, word := range extractIdentifiers(text) {
			if len(word) >= 2 && len(word) <= 64 {
				seen[word] = true
			}
		}
		for receiver, list := range extractMemberPairs(text) {
			if members[receiver] == nil {
				members[receiver] = make(map[string]bool)
			}
			for _, m := range list {
				if len(m) >= 1 && len(m) <= 64 {
					members[receiver][m] = true
				}
			}
		}
		return nil
	})
	s.projectIndex = make([]string, 0, len(seen))
	for w := range seen {
		s.projectIndex = append(s.projectIndex, w)
	}
	slices.Sort(s.projectIndex)
	for rec, set := range members {
		list := make([]string, 0, len(set))
		for m := range set {
			list = append(list, m)
		}
		slices.Sort(list)
		s.memberIndex[rec] = list
	}
}

// extractMemberPairs returns receiver -> members seen after "receiver." in content.
func extractMemberPairs(content string) map[string][]string {
	out := make(map[string][]string)
	runes := []rune(content)
	i := 0
	for i < len(runes) {
		// find start of identifier
		if !(unicode.IsLetter(runes[i]) || runes[i] == '_') {
			i++
			continue
		}
		start := i
		for i < len(runes) && (unicode.IsLetter(runes[i]) || runes[i] == '_' || unicode.IsDigit(runes[i])) {
			i++
		}
		receiver := string(runes[start:i])
		if i < len(runes) && runes[i] == '.' {
			i++ // consume '.'
			if i < len(runes) && (unicode.IsLetter(runes[i]) || runes[i] == '_') {
				mStart := i
				for i < len(runes) && (unicode.IsLetter(runes[i]) || runes[i] == '_' || unicode.IsDigit(runes[i])) {
					i++
				}
				member := string(runes[mStart:i])
				if len(receiver) <= 64 && len(member) <= 64 {
					out[receiver] = append(out[receiver], member)
				}
			}
		}
	}
	return out
}

func extractIdentifiers(content string) []string {
	var words []string
	var buf []rune
	for _, r := range content {
		if unicode.IsLetter(r) || r == '_' || (len(buf) > 0 && (unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_')) {
			buf = append(buf, r)
		} else {
			if len(buf) > 0 {
				words = append(words, string(buf))
				buf = buf[:0]
			}
		}
	}
	if len(buf) > 0 {
		words = append(words, string(buf))
	}
	return words
}

// projectCompletor suggests completions from the project index and member index (after ".").
type projectCompletor struct {
	editor      *gvcode.Editor
	index       []string
	memberIndex map[string][]string
}

func isSymbolSeparator(ch rune) bool {
	return !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_')
}

func (c *projectCompletor) Trigger() gvcode.Trigger {
	return gvcode.Trigger{
		Characters: []string{"."},
		KeyBinding: struct {
			Name      key.Name
			Modifiers key.Modifiers
		}{
			Name: key.NameSpace, Modifiers: key.ModShortcut,
		},
	}
}

// textBeforeCaret returns runes before the given rune position (for "receiver.member" parsing).
func (c *projectCompletor) textBeforeCaret(runePos int) []rune {
	text := c.editor.Text()
	runes := []rune(text)
	if runePos <= 0 {
		return nil
	}
	if runePos > len(runes) {
		runePos = len(runes)
	}
	return runes[:runePos]
}

func (c *projectCompletor) Suggest(ctx gvcode.CompletionContext) []gvcode.CompletionCandidate {
	before := c.textBeforeCaret(ctx.Position.Runes)
	if len(before) == 0 {
		return nil
	}
	// Check if we're after "receiver." or "receiver.memberPrefix"
	lastDot := -1
	for i := len(before) - 1; i >= 0; i-- {
		if before[i] == '.' {
			lastDot = i
			break
		}
	}
	if lastDot >= 0 && c.memberIndex != nil {
		receiver := string(trimIdentifierRight(before[:lastDot]))
		memberPrefix := string(trimIdentifierLeft(before[lastDot+1:]))
		if list, ok := c.memberIndex[receiver]; ok {
			candidates := make([]gvcode.CompletionCandidate, 0)
			for _, m := range list {
				if strings.HasPrefix(m, memberPrefix) {
					candidates = append(candidates, gvcode.CompletionCandidate{
						Label: m,
						TextEdit: gvcode.TextEdit{
							NewText: m,
						},
						Description: receiver + " member",
						Kind:        "property",
						TextFormat:  "PlainText",
					})
				}
			}
			if len(candidates) > 0 {
				return candidates
			}
		}
	}
	// Default: prefix match on full project index
	prefix := c.editor.ReadUntil(-1, isSymbolSeparator)
	candidates := make([]gvcode.CompletionCandidate, 0)
	for _, w := range c.index {
		if strings.HasPrefix(w, prefix) {
			candidates = append(candidates, gvcode.CompletionCandidate{
				Label: w,
				TextEdit: gvcode.TextEdit{
					NewText: w,
				},
				Description: "project",
				Kind:        "text",
				TextFormat:  "PlainText",
			})
		}
	}
	return candidates
}

func trimIdentifierRight(r []rune) []rune {
	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsLetter(r[i]) || r[i] == '_' || unicode.IsDigit(r[i]) {
			return r[:i+1]
		}
	}
	return nil
}

func trimIdentifierLeft(r []rune) []rune {
	for i, x := range r {
		if unicode.IsLetter(x) || x == '_' || unicode.IsDigit(x) {
			return r[i:]
		}
	}
	return nil
}

func (c *projectCompletor) FilterAndRank(pattern string, candidates []gvcode.CompletionCandidate) []gvcode.CompletionCandidate {
	if pattern == "" {
		return candidates
	}
	filtered := make([]gvcode.CompletionCandidate, 0)
	for _, c := range candidates {
		if strings.HasPrefix(strings.ToLower(c.Label), strings.ToLower(pattern)) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (s *appState) buildFileView(th *theme.Theme, path string) fileView {
	content, err := os.ReadFile(path)
	if err != nil {
		content = []byte(fmt.Sprintf("// Error reading %s: %v", path, err))
	}

	ed := wg.NewEditor(th.Material())
	ed.WithOptions(
		gvcode.WithLineNumber(true),
		gvcode.WithLineNumberGutterGap(unit.Dp(12)),
		gvcode.WithTextSize(unit.Sp(14)),
		gvcode.WithLineHeight(0, 1.35),
		gvcode.WithTabWidth(4),
	)
	ed.SetText(string(content))

	// Auto-completion from project index
	cm := &completion.DefaultCompletion{Editor: ed}
	popup := completion.NewCompletionPopup(ed, cm)
	popup.Theme = th.Material()
	popup.TextSize = unit.Sp(12)
	_ = cm.AddCompletor(&projectCompletor{editor: ed, index: s.projectIndex, memberIndex: s.memberIndex}, popup)
	ed.WithOptions(gvcode.WithAutoCompletion(cm))

	// Build color scheme from chroma style and apply syntax highlighting
	chromaStyle := styles.Get("dracula")
	if chromaStyle == nil {
		chromaStyle = styles.Fallback
	}
	gvScheme := buildColorSchemeFromChroma(th.Material(), chromaStyle)
	ed.WithOptions(gvcode.WithColorScheme(gvScheme))

	originalContent := string(content)
	tokens := chromaTokensToGvcode(path, originalContent, chromaStyle)
	if len(tokens) > 0 {
		ed.SetSyntaxTokens(tokens...)
	}

	onChange := func(currentContent string) {
		if tab := s.openTabs[path]; tab != nil {
			if currentContent == originalContent {
				tab.State = tabs.TabStateClean
			} else {
				tab.State = tabs.TabStateDirty
			}
		}
	}

	return fileView{
		Title:           path,
		Path:            path,
		Editor:          ed,
		OriginalContent: originalContent,
		OnChange:        onChange,
		Layout: func(gtx layout.Context, th *theme.Theme) layout.Dimensions {
			for {
				evt, ok := ed.Update(gtx)
				if !ok {
					break
				}
				if _, isChange := evt.(gvcode.ChangeEvent); isChange {
					if onChange != nil {
						onChange(ed.Text())
					}
					ed.OnTextEdit()
					// Keep syntax highlighting in sync
					tokens := chromaTokensToGvcode(path, ed.Text(), chromaStyle)
					if len(tokens) > 0 {
						ed.SetSyntaxTokens(tokens...)
					}
				}
			}
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return ed.Layout(gtx, th.Material().Shaper)
			})
		},
	}
}

// buildColorSchemeFromChroma creates a gvcode ColorScheme from chroma style.
func buildColorSchemeFromChroma(mat *material.Theme, chromaStyle *chroma.Style) syntax.ColorScheme {
	cs := syntax.ColorScheme{}
	cs.Foreground = gvcolor.MakeColor(mat.Fg)
	cs.Background = gvcolor.MakeColor(mat.Bg)
	cs.SelectColor = gvcolor.MakeColor(mat.ContrastBg).MulAlpha(0x60)
	cs.LineColor = gvcolor.MakeColor(mat.ContrastBg).MulAlpha(0x30)
	cs.LineNumberColor = gvcolor.MakeColor(mat.Fg).MulAlpha(0xb6)

	// Register styles for common chroma token types so they are available when we set tokens.
	for _, tt := range []chroma.TokenType{
		chroma.Keyword, chroma.KeywordConstant, chroma.KeywordDeclaration, chroma.KeywordType,
		chroma.Name, chroma.NameBuiltin, chroma.NameFunction, chroma.NameVariable,
		chroma.LiteralString, chroma.LiteralStringChar, chroma.LiteralStringEscape,
		chroma.LiteralNumber, chroma.LiteralNumberInteger, chroma.LiteralNumberFloat,
		chroma.Comment, chroma.CommentSingle, chroma.CommentMultiline,
		chroma.Operator, chroma.Punctuation,
		chroma.Text, chroma.Whitespace,
	} {
		entry := chromaStyle.Get(tt)
		if entry.Colour.IsSet() {
			fg := gvcolor.MakeColor(color.NRGBA{
				R: entry.Colour.Red(),
				G: entry.Colour.Green(),
				B: entry.Colour.Blue(),
				A: 255,
			})
			cs.AddStyle(syntax.StyleScope(tt.String()), 0, fg, gvcolor.Color{})
		}
	}
	return cs
}

// chromaTokensToGvcode tokenizes content with chroma and returns gvcode syntax tokens (rune offsets).
func chromaTokensToGvcode(filename, content string, _ *chroma.Style) []syntax.Token {
	lexer := lexers.Match(filename)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	it, err := lexer.Tokenise(nil, content)
	if err != nil {
		return nil
	}

	var tokens []syntax.Token
	runeOffset := 0
	for t := it(); t != chroma.EOF; t = it() {
		if t.Value == "" {
			continue
		}
		start := runeOffset
		runeOffset += utf8.RuneCountInString(t.Value)
		end := runeOffset
		scope := syntax.StyleScope(t.Type.String())
		if scope.IsValid() {
			tokens = append(tokens, syntax.Token{Start: start, End: end, Scope: scope})
		}
	}
	return tokens
}
