package home

const (
	English = "en"
	Chinese = "zh-CN"

	repositoryURL = "https://github.com/daqing/iNeovim"
)

type featureContent struct {
	Number      string
	Title       string
	Description string
	Shortcut    string
}

type pageContent struct {
	Language              string
	Title                 string
	Description           string
	NavigationLabel       string
	HomeLabel             string
	NavFeatures           string
	LanguageAction        string
	HeroKicker            string
	HeroTitle             string
	HeroDescription       string
	PrimaryAction         string
	SecondaryAction       string
	ScreenshotAlt         string
	ScreenshotCaption     string
	FeaturesKicker        string
	FeaturesTitle         string
	FeaturesDescription   string
	Features              []featureContent
	OpenSourceKicker      string
	OpenSourceTitle       string
	OpenSourceDescription string
	OpenSourceAction      string
	FooterTagline         string
}

func contentFor(language string) pageContent {
	if language == Chinese {
		return pageContent{
			Language:            Chinese,
			Title:               "iNeovim — macOS 原生 Neovim 客户端",
			Description:         "简洁、快速的 macOS 原生 Neovim 客户端，集成独立终端面板与直观的 LSP 诊断。",
			NavigationLabel:     "主导航",
			HomeLabel:           "iNeovim 首页",
			NavFeatures:         "功能",
			LanguageAction:      "English",
			HeroKicker:          "macOS 原生开源 Neovim 客户端",
			HeroTitle:           "Neovim，融入 macOS。",
			HeroDescription:     "一款简洁、快速的 Mac 原生 Neovim 客户端，保留你熟悉的编辑体验，自然融入日常工作流。",
			PrimaryAction:       "在 GitHub 上查看",
			SecondaryAction:     "了解功能",
			ScreenshotAlt:       "iNeovim 在 macOS 上展示 Go 代码、LSP 诊断弹窗、Problems 侧栏和独立终端面板。",
			ScreenshotCaption:   "代码、终端与诊断，在同一个原生窗口中清晰呈现。",
			FeaturesKicker:      "为什么选择 iNeovim",
			FeaturesTitle:       "专注编辑，也照顾整个工作流。",
			FeaturesDescription: "Neovim 仍然是核心，iNeovim 则为它补上贴合 Mac 使用习惯的原生体验。",
			Features: []featureContent{
				{
					Number:      "01",
					Title:       "原生、克制、快速",
					Description: "为 macOS 打造的简洁界面，遵循熟悉的 Mac 交互习惯，没有多余干扰，让你专注于代码。",
				},
				{
					Number:      "02",
					Title:       "独立终端面板",
					Description: "基于开源 libghostty 库构建，在独立面板中运行命令。代码与终端同窗协作，无需来回切换。",
				},
				{
					Number:      "03",
					Title:       "点击查看诊断详情",
					Description: "语言服务器发现的问题不再难以阅读。点击代码中的诊断，即可通过弹窗查看详细信息。",
				},
				{
					Number:      "04",
					Title:       "所有问题，一栏尽览",
					Description: "打开 Problems 侧栏，集中查看所有 LSP 诊断，快速找到需要处理的问题。",
					Shortcut:    "Cmd+P",
				},
			},
			OpenSourceKicker:      "开源项目",
			OpenSourceTitle:       "代码开放，欢迎参与。",
			OpenSourceDescription: "iNeovim 完全开源。你可以在 GitHub 上阅读源码、提交问题，或参与项目建设。",
			OpenSourceAction:      "在 GitHub 上探索 iNeovim",
			FooterTagline:         "为 Neovim 而构建，为 macOS 而打磨。",
		}
	}

	return pageContent{
		Language:            English,
		Title:               "iNeovim — Native Neovim for macOS",
		Description:         "A fast, native macOS client for Neovim with an integrated terminal and simple, visual LSP diagnostics.",
		NavigationLabel:     "Primary navigation",
		HomeLabel:           "iNeovim home",
		NavFeatures:         "Features",
		LanguageAction:      "中文",
		HeroKicker:          "Open-source Neovim client for macOS",
		HeroTitle:           "Neovim, at home on macOS.",
		HeroDescription:     "A native Mac client that stays out of your way. Fast, focused, and built around the way you already use Neovim.",
		PrimaryAction:       "View on GitHub",
		SecondaryAction:     "Explore features",
		ScreenshotAlt:       "iNeovim on macOS showing Go code, an LSP diagnostic popover, the Problems sidebar, and a terminal panel.",
		ScreenshotCaption:   "Native editing, terminal, and diagnostics in one window.",
		FeaturesKicker:      "Why iNeovim",
		FeaturesTitle:       "The editor you know. The workflow you need.",
		FeaturesDescription: "Neovim is still the editor. iNeovim makes the surrounding workflow feel at home on a Mac.",
		Features: []featureContent{
			{
				Number:      "01",
				Title:       "Native by design",
				Description: "A clean, responsive interface built for macOS. Familiar Mac conventions and a focused layout keep the experience fast and uncomplicated.",
			},
			{
				Number:      "02",
				Title:       "A terminal that belongs",
				Description: "A dedicated terminal panel powered by the open-source libghostty library. Run commands alongside your code without leaving the window.",
			},
			{
				Number:      "03",
				Title:       "Diagnostics at a click",
				Description: "See what your language server is telling you. Click a diagnostic to read the details in a popover, right where you are working.",
			},
			{
				Number:      "04",
				Title:       "Every problem, one place",
				Description: "Open the Problems sidebar to review all LSP diagnostics in one place and quickly find what needs your attention.",
				Shortcut:    "Cmd+P",
			},
		},
		OpenSourceKicker:      "Open source",
		OpenSourceTitle:       "Built in the open.",
		OpenSourceDescription: "iNeovim is fully open source. Explore the code, report an issue, or help shape the project on GitHub.",
		OpenSourceAction:      "Explore the project on GitHub",
		FooterTagline:         "Built for Neovim. Made for macOS.",
	}
}

func alternateLanguage(language string) string {
	if language == Chinese {
		return English
	}
	return Chinese
}

func homePath(language string) string {
	if language == Chinese {
		return "/zh/"
	}
	return "/"
}
