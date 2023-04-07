package formatters

import (
	"fmt"
	"io"
	"math"

	"github.com/alecthomas/chroma"
)

type ttyTable struct {
	foreground map[chroma.Colour]string
	background map[chroma.Colour]string
}

var c = chroma.MustParseColour

var ttyTables = map[int]*ttyTable{
	8: {
		foreground: map[chroma.Colour]string{
			c("#000000"): "\033[30m", c("#7f0000"): "\033[31m", c("#007f00"): "\033[32m", c("#7f7fe0"): "\033[33m",
			c("#00007f"): "\033[34m", c("#7f007f"): "\033[35m", c("#007f7f"): "\033[36m", c("#e5e5e5"): "\033[37m",
			c("#555555"): "\033[1m\033[30m", c("#ff0000"): "\033[1m\033[31m", c("#00ff00"): "\033[1m\033[32m", c("#ffff00"): "\033[1m\033[33m",
			c("#0000ff"): "\033[1m\033[34m", c("#ff00ff"): "\033[1m\033[35m", c("#00ffff"): "\033[1m\033[36m", c("#ffffff"): "\033[1m\033[37m",
		},
		background: map[chroma.Colour]string{
			c("#000000"): "\033[40m", c("#7f0000"): "\033[41m", c("#007f00"): "\033[42m", c("#7f7fe0"): "\033[43m",
			c("#00007f"): "\033[44m", c("#7f007f"): "\033[45m", c("#007f7f"): "\033[46m", c("#e5e5e5"): "\033[47m",
			c("#555555"): "\033[1m\033[40m", c("#ff0000"): "\033[1m\033[41m", c("#00ff00"): "\033[1m\033[42m", c("#ffff00"): "\033[1m\033[43m",
			c("#0000ff"): "\033[1m\033[44m", c("#ff00ff"): "\033[1m\033[45m", c("#00ffff"): "\033[1m\033[46m", c("#ffffff"): "\033[1m\033[47m",
		},
	},
	16: {
		foreground: map[chroma.Colour]string{
			c("#000000"): "\033[30m", c("#7f0000"): "\033[31m", c("#007f00"): "\033[32m", c("#7f7fe0"): "\033[33m",
			c("#00007f"): "\033[34m", c("#7f007f"): "\033[35m", c("#007f7f"): "\033[36m", c("#e5e5e5"): "\033[37m",
			c("#555555"): "\033[90m", c("#ff0000"): "\033[91m", c("#00ff00"): "\033[92m", c("#ffff00"): "\033[93m",
			c("#0000ff"): "\033[94m", c("#ff00ff"): "\033[95m", c("#00ffff"): "\033[96m", c("#ffffff"): "\033[97m",
		},
		background: map[chroma.Colour]string{
			c("#000000"): "\033[40m", c("#7f0000"): "\033[41m", c("#007f00"): "\033[42m", c("#7f7fe0"): "\033[43m",
			c("#00007f"): "\033[44m", c("#7f007f"): "\033[45m", c("#007f7f"): "\033[46m", c("#e5e5e5"): "\033[47m",
			c("#555555"): "\033[100m", c("#ff0000"): "\033[101m", c("#00ff00"): "\033[102m", c("#ffff00"): "\033[103m",
			c("#0000ff"): "\033[104m", c("#ff00ff"): "\033[105m", c("#00ffff"): "\033[106m", c("#ffffff"): "\033[107m",
		},
	},
	256: {
		foreground: map[chroma.Colour]string{
			c("#000000"): "\033[38;5;0m", c("#800000"): "\033[38;5;1m", c("#008000"): "\033[38;5;2m", c("#808000"): "\033[38;5;3m",
			c("#000080"): "\033[38;5;4m", c("#800080"): "\033[38;5;5m", c("#008080"): "\033[38;5;6m", c("#c0c0c0"): "\033[38;5;7m",
			c("#808080"): "\033[38;5;8m", c("#ff0000"): "\033[38;5;9m", c("#00ff00"): "\033[38;5;10m", c("#ffff00"): "\033[38;5;11m",
			c("#0000ff"): "\033[38;5;12m", c("#ff00ff"): "\033[38;5;13m", c("#00ffff"): "\033[38;5;14m", c("#ffffff"): "\033[38;5;15m",
			c("#000000"): "\033[38;5;16m", c("#00005f"): "\033[38;5;17m", c("#000087"): "\033[38;5;18m", c("#0000af"): "\033[38;5;19m",
			c("#0000d7"): "\033[38;5;20m", c("#0000ff"): "\033[38;5;21m", c("#005f00"): "\033[38;5;22m", c("#005f5f"): "\033[38;5;23m",
			c("#005f87"): "\033[38;5;24m", c("#005faf"): "\033[38;5;25m", c("#005fd7"): "\033[38;5;26m", c("#005fff"): "\033[38;5;27m",
			c("#008700"): "\033[38;5;28m", c("#00875f"): "\033[38;5;29m", c("#008787"): "\033[38;5;30m", c("#0087af"): "\033[38;5;31m",
			c("#0087d7"): "\033[38;5;32m", c("#0087ff"): "\033[38;5;33m", c("#00af00"): "\033[38;5;34m", c("#00af5f"): "\033[38;5;35m",
			c("#00af87"): "\033[38;5;36m", c("#00afaf"): "\033[38;5;37m", c("#00afd7"): "\033[38;5;38m", c("#00afff"): "\033[38;5;39m",
			c("#00d700"): "\033[38;5;40m", c("#00d75f"): "\033[38;5;41m", c("#00d787"): "\033[38;5;42m", c("#00d7af"): "\033[38;5;43m",
			c("#00d7d7"): "\033[38;5;44m", c("#00d7ff"): "\033[38;5;45m", c("#00ff00"): "\033[38;5;46m", c("#00ff5f"): "\033[38;5;47m",
			c("#00ff87"): "\033[38;5;48m", c("#00ffaf"): "\033[38;5;49m", c("#00ffd7"): "\033[38;5;50m", c("#00ffff"): "\033[38;5;51m",
			c("#5f0000"): "\033[38;5;52m", c("#5f005f"): "\033[38;5;53m", c("#5f0087"): "\033[38;5;54m", c("#5f00af"): "\033[38;5;55m",
			c("#5f00d7"): "\033[38;5;56m", c("#5f00ff"): "\033[38;5;57m", c("#5f5f00"): "\033[38;5;58m", c("#5f5f5f"): "\033[38;5;59m",
			c("#5f5f87"): "\033[38;5;60m", c("#5f5faf"): "\033[38;5;61m", c("#5f5fd7"): "\033[38;5;62m", c("#5f5fff"): "\033[38;5;63m",
			c("#5f8700"): "\033[38;5;64m", c("#5f875f"): "\033[38;5;65m", c("#5f8787"): "\033[38;5;66m", c("#5f87af"): "\033[38;5;67m",
			c("#5f87d7"): "\033[38;5;68m", c("#5f87ff"): "\033[38;5;69m", c("#5faf00"): "\033[38;5;70m", c("#5faf5f"): "\033[38;5;71m",
			c("#5faf87"): "\033[38;5;72m", c("#5fafaf"): "\033[38;5;73m", c("#5fafd7"): "\033[38;5;74m", c("#5fafff"): "\033[38;5;75m",
			c("#5fd700"): "\033[38;5;76m", c("#5fd75f"): "\033[38;5;77m", c("#5fd787"): "\033[38;5;78m", c("#5fd7af"): "\033[38;5;79m",
			c("#5fd7d7"): "\033[38;5;80m", c("#5fd7ff"): "\033[38;5;81m", c("#5fff00"): "\033[38;5;82m", c("#5fff5f"): "\033[38;5;83m",
			c("#5fff87"): "\033[38;5;84m", c("#5fffaf"): "\033[38;5;85m", c("#5fffd7"): "\033[38;5;86m", c("#5fffff"): "\033[38;5;87m",
			c("#870000"): "\033[38;5;88m", c("#87005f"): "\033[38;5;89m", c("#870087"): "\033[38;5;90m", c("#8700af"): "\033[38;5;91m",
			c("#8700d7"): "\033[38;5;92m", c("#8700ff"): "\033[38;5;93m", c("#875f00"): "\033[38;5;94m", c("#875f5f"): "\033[38;5;95m",
			c("#875f87"): "\033[38;5;96m", c("#875faf"): "\033[38;5;97m", c("#875fd7"): "\033[38;5;98m", c("#875fff"): "\033[38;5;99m",
			c("#878700"): "\033[38;5;100m", c("#87875f"): "\033[38;5;101m", c("#878787"): "\033[38;5;102m", c("#8787af"): "\033[38;5;103m",
			c("#8787d7"): "\033[38;5;104m", c("#8787ff"): "\033[38;5;105m", c("#87af00"): "\033[38;5;106m", c("#87af5f"): "\033[38;5;107m",
			c("#87af87"): "\033[38;5;108m", c("#87afaf"): "\033[38;5;109m", c("#87afd7"): "\033[38;5;110m", c("#87afff"): "\033[38;5;111m",
			c("#87d700"): "\033[38;5;112m", c("#87d75f"): "\033[38;5;113m", c("#87d787"): "\033[38;5;114m", c("#87d7af"): "\033[38;5;115m",
			c("#87d7d7"): "\033[38;5;116m", c("#87d7ff"): "\033[38;5;117m", c("#87ff00"): "\033[38;5;118m", c("#87ff5f"): "\033[38;5;119m",
			c("#87ff87"): "\033[38;5;120m", c("#87ffaf"): "\033[38;5;121m", c("#87ffd7"): "\033[38;5;122m", c("#87ffff"): "\033[38;5;123m",
			c("#af0000"): "\033[38;5;124m", c("#af005f"): "\033[38;5;125m", c("#af0087"): "\033[38;5;126m", c("#af00af"): "\033[38;5;127m",
			c("#af00d7"): "\033[38;5;128m", c("#af00ff"): "\033[38;5;129m", c("#af5f00"): "\033[38;5;130m", c("#af5f5f"): "\033[38;5;131m",
			c("#af5f87"): "\033[38;5;132m", c("#af5faf"): "\033[38;5;133m", c("#af5fd7"): "\033[38;5;134m", c("#af5fff"): "\033[38;5;135m",
			c("#af8700"): "\033[38;5;136m", c("#af875f"): "\033[38;5;137m", c("#af8787"): "\033[38;5;138m", c("#af87af"): "\033[38;5;139m",
			c("#af87d7"): "\033[38;5;140m", c("#af87ff"): "\033[38;5;141m", c("#afaf00"): "\033[38;5;142m", c("#afaf5f"): "\033[38;5;143m",
			c("#afaf87"): "\033[38;5;144m", c("#afafaf"): "\033[38;5;145m", c("#afafd7"): "\033[38;5;146m", c("#afafff"): "\033[38;5;147m",
			c("#afd700"): "\033[38;5;148m", c("#afd75f"): "\033[38;5;149m", c("#afd787"): "\033[38;5;150m", c("#afd7af"): "\033[38;5;151m",
			c("#afd7d7"): "\033[38;5;152m", c("#afd7ff"): "\033[38;5;153m", c("#afff00"): "\033[38;5;154m", c("#afff5f"): "\033[38;5;155m",
			c("#afff87"): "\033[38;5;156m", c("#afffaf"): "\033[38;5;157m", c("#afffd7"): "\033[38;5;158m", c("#afffff"): "\033[38;5;159m",
			c("#d70000"): "\033[38;5;160m", c("#d7005f"): "\033[38;5;161m", c("#d70087"): "\033[38;5;162m", c("#d700af"): "\033[38;5;163m",
			c("#d700d7"): "\033[38;5;164m", c("#d700ff"): "\033[38;5;165m", c("#d75f00"): "\033[38;5;166m", c("#d75f5f"): "\033[38;5;167m",
			c("#d75f87"): "\033[38;5;168m", c("#d75faf"): "\033[38;5;169m", c("#d75fd7"): "\033[38;5;170m", c("#d75fff"): "\033[38;5;171m",
			c("#d78700"): "\033[38;5;172m", c("#d7875f"): "\033[38;5;173m", c("#d78787"): "\033[38;5;174m", c("#d787af"): "\033[38;5;175m",
			c("#d787d7"): "\033[38;5;176m", c("#d787ff"): "\033[38;5;177m", c("#d7af00"): "\033[38;5;178m", c("#d7af5f"): "\033[38;5;179m",
			c("#d7af87"): "\033[38;5;180m", c("#d7afaf"): "\033[38;5;181m", c("#d7afd7"): "\033[38;5;182m", c("#d7afff"): "\033[38;5;183m",
			c("#d7d700"): "\033[38;5;184m", c("#d7d75f"): "\033[38;5;185m", c("#d7d787"): "\033[38;5;186m", c("#d7d7af"): "\033[38;5;187m",
			c("#d7d7d7"): "\033[38;5;188m", c("#d7d7ff"): "\033[38;5;189m", c("#d7ff00"): "\033[38;5;190m", c("#d7ff5f"): "\033[38;5;191m",
			c("#d7ff87"): "\033[38;5;192m", c("#d7ffaf"): "\033[38;5;193m", c("#d7ffd7"): "\033[38;5;194m", c("#d7ffff"): "\033[38;5;195m",
			c("#ff0000"): "\033[38;5;196m", c("#ff005f"): "\033[38;5;197m", c("#ff0087"): "\033[38;5;198m", c("#ff00af"): "\033[38;5;199m",
			c("#ff00d7"): "\033[38;5;200m", c("#ff00ff"): "\033[38;5;201m", c("#ff5f00"): "\033[38;5;202m", c("#ff5f5f"): "\033[38;5;203m",
			c("#ff5f87"): "\033[38;5;204m", c("#ff5faf"): "\033[38;5;205m", c("#ff5fd7"): "\033[38;5;206m", c("#ff5fff"): "\033[38;5;207m",
			c("#ff8700"): "\033[38;5;208m", c("#ff875f"): "\033[38;5;209m", c("#ff8787"): "\033[38;5;210m", c("#ff87af"): "\033[38;5;211m",
			c("#ff87d7"): "\033[38;5;212m", c("#ff87ff"): "\033[38;5;213m", c("#ffaf00"): "\033[38;5;214m", c("#ffaf5f"): "\033[38;5;215m",
			c("#ffaf87"): "\033[38;5;216m", c("#ffafaf"): "\033[38;5;217m", c("#ffafd7"): "\033[38;5;218m", c("#ffafff"): "\033[38;5;219m",
			c("#ffd700"): "\033[38;5;220m", c("#ffd75f"): "\033[38;5;221m", c("#ffd787"): "\033[38;5;222m", c("#ffd7af"): "\033[38;5;223m",
			c("#ffd7d7"): "\033[38;5;224m", c("#ffd7ff"): "\033[38;5;225m", c("#ffff00"): "\033[38;5;226m", c("#ffff5f"): "\033[38;5;227m",
			c("#ffff87"): "\033[38;5;228m", c("#ffffaf"): "\033[38;5;229m", c("#ffffd7"): "\033[38;5;230m", c("#ffffff"): "\033[38;5;231m",
			c("#080808"): "\033[38;5;232m", c("#121212"): "\033[38;5;233m", c("#1c1c1c"): "\033[38;5;234m", c("#262626"): "\033[38;5;235m",
			c("#303030"): "\033[38;5;236m", c("#3a3a3a"): "\033[38;5;237m", c("#444444"): "\033[38;5;238m", c("#4e4e4e"): "\033[38;5;239m",
			c("#585858"): "\033[38;5;240m", c("#626262"): "\033[38;5;241m", c("#6c6c6c"): "\033[38;5;242m", c("#767676"): "\033[38;5;243m",
			c("#808080"): "\033[38;5;244m", c("#8a8a8a"): "\033[38;5;245m", c("#949494"): "\033[38;5;246m", c("#9e9e9e"): "\033[38;5;247m",
			c("#a8a8a8"): "\033[38;5;248m", c("#b2b2b2"): "\033[38;5;249m", c("#bcbcbc"): "\033[38;5;250m", c("#c6c6c6"): "\033[38;5;251m",
			c("#d0d0d0"): "\033[38;5;252m", c("#dadada"): "\033[38;5;253m", c("#e4e4e4"): "\033[38;5;254m", c("#eeeeee"): "\033[38;5;255m",
		},
		background: map[chroma.Colour]string{
			c("#000000"): "\033[48;5;0m", c("#800000"): "\033[48;5;1m", c("#008000"): "\033[48;5;2m", c("#808000"): "\033[48;5;3m",
			c("#000080"): "\033[48;5;4m", c("#800080"): "\033[48;5;5m", c("#008080"): "\033[48;5;6m", c("#c0c0c0"): "\033[48;5;7m",
			c("#808080"): "\033[48;5;8m", c("#ff0000"): "\033[48;5;9m", c("#00ff00"): "\033[48;5;10m", c("#ffff00"): "\033[48;5;11m",
			c("#0000ff"): "\033[48;5;12m", c("#ff00ff"): "\033[48;5;13m", c("#00ffff"): "\033[48;5;14m", c("#ffffff"): "\033[48;5;15m",
			c("#000000"): "\033[48;5;16m", c("#00005f"): "\033[48;5;17m", c("#000087"): "\033[48;5;18m", c("#0000af"): "\033[48;5;19m",
			c("#0000d7"): "\033[48;5;20m", c("#0000ff"): "\033[48;5;21m", c("#005f00"): "\033[48;5;22m", c("#005f5f"): "\033[48;5;23m",
			c("#005f87"): "\033[48;5;24m", c("#005faf"): "\033[48;5;25m", c("#005fd7"): "\033[48;5;26m", c("#005fff"): "\033[48;5;27m",
			c("#008700"): "\033[48;5;28m", c("#00875f"): "\033[48;5;29m", c("#008787"): "\033[48;5;30m", c("#0087af"): "\033[48;5;31m",
			c("#0087d7"): "\033[48;5;32m", c("#0087ff"): "\033[48;5;33m", c("#00af00"): "\033[48;5;34m", c("#00af5f"): "\033[48;5;35m",
			c("#00af87"): "\033[48;5;36m", c("#00afaf"): "\033[48;5;37m", c("#00afd7"): "\033[48;5;38m", c("#00afff"): "\033[48;5;39m",
			c("#00d700"): "\033[48;5;40m", c("#00d75f"): "\033[48;5;41m", c("#00d787"): "\033[48;5;42m", c("#00d7af"): "\033[48;5;43m",
			c("#00d7d7"): "\033[48;5;44m", c("#00d7ff"): "\033[48;5;45m", c("#00ff00"): "\033[48;5;46m", c("#00ff5f"): "\033[48;5;47m",
			c("#00ff87"): "\033[48;5;48m", c("#00ffaf"): "\033[48;5;49m", c("#00ffd7"): "\033[48;5;50m", c("#00ffff"): "\033[48;5;51m",
			c("#5f0000"): "\033[48;5;52m", c("#5f005f"): "\033[48;5;53m", c("#5f0087"): "\033[48;5;54m", c("#5f00af"): "\033[48;5;55m",
			c("#5f00d7"): "\033[48;5;56m", c("#5f00ff"): "\033[48;5;57m", c("#5f5f00"): "\033[48;5;58m", c("#5f5f5f"): "\033[48;5;59m",
			c("#5f5f87"): "\033[48;5;60m", c("#5f5faf"): "\033[48;5;61m", c("#5f5fd7"): "\033[48;5;62m", c("#5f5fff"): "\033[48;5;63m",
			c("#5f8700"): "\033[48;5;64m", c("#5f875f"): "\033[48;5;65m", c("#5f8787"): "\033[48;5;66m", c("#5f87af"): "\033[48;5;67m",
			c("#5f87d7"): "\033[48;5;68m", c("#5f87ff"): "\033[48;5;69m", c("#5faf00"): "\033[48;5;70m", c("#5faf5f"): "\033[48;5;71m",
			c("#5faf87"): "\033[48;5;72m", c("#5fafaf"): "\033[48;5;73m", c("#5fafd7"): "\033[48;5;74m", c("#5fafff"): "\033[48;5;75m",
			c("#5fd700"): "\033[48;5;76m", c("#5fd75f"): "\033[48;5;77m", c("#5fd787"): "\033[48;5;78m", c("#5fd7af"): "\033[48;5;79m",
			c("#5fd7d7"): "\033[48;5;80m", c("#5fd7ff"): "\033[48;5;81m", c("#5fff00"): "\033[48;5;82m", c("#5fff5f"): "\033[48;5;83m",
			c("#5fff87"): "\033[48;5;84m", c("#5fffaf"): "\033[48;5;85m", c("#5fffd7"): "\033[48;5;86m", c("#5fffff"): "\033[48;5;87m",
			c("#870000"): "\033[48;5;88m", c("#87005f"): "\033[48;5;89m", c("#870087"): "\033[48;5;90m", c("#8700af"): "\033[48;5;91m",
			c("#8700d7"): "\033[48;5;92m", c("#8700ff"): "\033[48;5;93m", c("#875f00"): "\033[48;5;94m", c("#875f5f"): "\033[48;5;95m",
			c("#875f87"): "\033[48;5;96m", c("#875faf"): "\033[48;5;97m", c("#875fd7"): "\033[48;5;98m", c("#875fff"): "\033[48;5;99m",
			c("#878700"): "\033[48;5;100m", c("#87875f"): "\033[48;5;101m", c("#878787"): "\033[48;5;102m", c("#8787af"): "\033[48;5;103m",
			c("#8787d7"): "\033[48;5;104m", c("#8787ff"): "\033[48;5;105m", c("#87af00"): "\033[48;5;106m", c("#87af5f"): "\033[48;5;107m",
			c("#87af87"): "\033[48;5;108m", c("#87afaf"): "\033[48;5;109m", c("#87afd7"): "\033[48;5;110m", c("#87afff"): "\033[48;5;111m",
			c("#87d700"): "\033[48;5;112m", c("#87d75f"): "\033[48;5;113m", c("#87d787"): "\033[48;5;114m", c("#87d7af"): "\033[48;5;115m",
			c("#87d7d7"): "\033[48;5;116m", c("#87d7ff"): "\033[48;5;117m", c("#87ff00"): "\033[48;5;118m", c("#87ff5f"): "\033[48;5;119m",
			c("#87ff87"): "\033[48;5;120m", c("#87ffaf"): "\033[48;5;121m", c("#87ffd7"): "\033[48;5;122m", c("#87ffff"): "\033[48;5;123m",
			c("#af0000"): "\033[48;5;124m", c("#af005f"): "\033[48;5;125m", c("#af0087"): "\033[48;5;126m", c("#af00af"): "\033[48;5;127m",
			c("#af00d7"): "\033[48;5;128m", c("#af00ff"): "\033[48;5;129m", c("#af5f00"): "\033[48;5;130m", c("#af5f5f"): "\033[48;5;131m",
			c("#af5f87"): "\033[48;5;132m", c("#af5faf"): "\033[48;5;133m", c("#af5fd7"): "\033[48;5;134m", c("#af5fff"): "\033[48;5;135m",
			c("#af8700"): "\033[48;5;136m", c("#af875f"): "\033[48;5;137m", c("#af8787"): "\033[48;5;138m", c("#af87af"): "\033[48;5;139m",
			c("#af87d7"): "\033[48;5;140m", c("#af87ff"): "\033[48;5;141m", c("#afaf00"): "\033[48;5;142m", c("#afaf5f"): "\033[48;5;143m",
			c("#afaf87"): "\033[48;5;144m", c("#afafaf"): "\033[48;5;145m", c("#afafd7"): "\033[48;5;146m", c("#afafff"): "\033[48;5;147m",
			c("#afd700"): "\033[48;5;148m", c("#afd75f"): "\033[48;5;149m", c("#afd787"): "\033[48;5;150m", c("#afd7af"): "\033[48;5;151m",
			c("#afd7d7"): "\033[48;5;152m", c("#afd7ff"): "\033[48;5;153m", c("#afff00"): "\033[48;5;154m", c("#afff5f"): "\033[48;5;155m",
			c("#afff87"): "\033[48;5;156m", c("#afffaf"): "\033[48;5;157m", c("#afffd7"): "\033[48;5;158m", c("#afffff"): "\033[48;5;159m",
			c("#d70000"): "\033[48;5;160m", c("#d7005f"): "\033[48;5;161m", c("#d70087"): "\033[48;5;162m", c("#d700af"): "\033[48;5;163m",
			c("#d700d7"): "\033[48;5;164m", c("#d700ff"): "\033[48;5;165m", c("#d75f00"): "\033[48;5;166m", c("#d75f5f"): "\033[48;5;167m",
			c("#d75f87"): "\033[48;5;168m", c("#d75faf"): "\033[48;5;169m", c("#d75fd7"): "\033[48;5;170m", c("#d75fff"): "\033[48;5;171m",
			c("#d78700"): "\033[48;5;172m", c("#d7875f"): "\033[48;5;173m", c("#d78787"): "\033[48;5;174m", c("#d787af"): "\033[48;5;175m",
			c("#d787d7"): "\033[48;5;176m", c("#d787ff"): "\033[48;5;177m", c("#d7af00"): "\033[48;5;178m", c("#d7af5f"): "\033[48;5;179m",
			c("#d7af87"): "\033[48;5;180m", c("#d7afaf"): "\033[48;5;181m", c("#d7afd7"): "\033[48;5;182m", c("#d7afff"): "\033[48;5;183m",
			c("#d7d700"): "\033[48;5;184m", c("#d7d75f"): "\033[48;5;185m", c("#d7d787"): "\033[48;5;186m", c("#d7d7af"): "\033[48;5;187m",
			c("#d7d7d7"): "\033[48;5;188m", c("#d7d7ff"): "\033[48;5;189m", c("#d7ff00"): "\033[48;5;190m", c("#d7ff5f"): "\033[48;5;191m",
			c("#d7ff87"): "\033[48;5;192m", c("#d7ffaf"): "\033[48;5;193m", c("#d7ffd7"): "\033[48;5;194m", c("#d7ffff"): "\033[48;5;195m",
			c("#ff0000"): "\033[48;5;196m", c("#ff005f"): "\033[48;5;197m", c("#ff0087"): "\033[48;5;198m", c("#ff00af"): "\033[48;5;199m",
			c("#ff00d7"): "\033[48;5;200m", c("#ff00ff"): "\033[48;5;201m", c("#ff5f00"): "\033[48;5;202m", c("#ff5f5f"): "\033[48;5;203m",
			c("#ff5f87"): "\033[48;5;204m", c("#ff5faf"): "\033[48;5;205m", c("#ff5fd7"): "\033[48;5;206m", c("#ff5fff"): "\033[48;5;207m",
			c("#ff8700"): "\033[48;5;208m", c("#ff875f"): "\033[48;5;209m", c("#ff8787"): "\033[48;5;210m", c("#ff87af"): "\033[48;5;211m",
			c("#ff87d7"): "\033[48;5;212m", c("#ff87ff"): "\033[48;5;213m", c("#ffaf00"): "\033[48;5;214m", c("#ffaf5f"): "\033[48;5;215m",
			c("#ffaf87"): "\033[48;5;216m", c("#ffafaf"): "\033[48;5;217m", c("#ffafd7"): "\033[48;5;218m", c("#ffafff"): "\033[48;5;219m",
			c("#ffd700"): "\033[48;5;220m", c("#ffd75f"): "\033[48;5;221m", c("#ffd787"): "\033[48;5;222m", c("#ffd7af"): "\033[48;5;223m",
			c("#ffd7d7"): "\033[48;5;224m", c("#ffd7ff"): "\033[48;5;225m", c("#ffff00"): "\033[48;5;226m", c("#ffff5f"): "\033[48;5;227m",
			c("#ffff87"): "\033[48;5;228m", c("#ffffaf"): "\033[48;5;229m", c("#ffffd7"): "\033[48;5;230m", c("#ffffff"): "\033[48;5;231m",
			c("#080808"): "\033[48;5;232m", c("#121212"): "\033[48;5;233m", c("#1c1c1c"): "\033[48;5;234m", c("#262626"): "\033[48;5;235m",
			c("#303030"): "\033[48;5;236m", c("#3a3a3a"): "\033[48;5;237m", c("#444444"): "\033[48;5;238m", c("#4e4e4e"): "\033[48;5;239m",
			c("#585858"): "\033[48;5;240m", c("#626262"): "\033[48;5;241m", c("#6c6c6c"): "\033[48;5;242m", c("#767676"): "\033[48;5;243m",
			c("#808080"): "\033[48;5;244m", c("#8a8a8a"): "\033[48;5;245m", c("#949494"): "\033[48;5;246m", c("#9e9e9e"): "\033[48;5;247m",
			c("#a8a8a8"): "\033[48;5;248m", c("#b2b2b2"): "\033[48;5;249m", c("#bcbcbc"): "\033[48;5;250m", c("#c6c6c6"): "\033[48;5;251m",
			c("#d0d0d0"): "\033[48;5;252m", c("#dadada"): "\033[48;5;253m", c("#e4e4e4"): "\033[48;5;254m", c("#eeeeee"): "\033[48;5;255m",
		},
	},
}

func entryToEscapeSequence(table *ttyTable, entry chroma.StyleEntry) string {
	out := ""
	if entry.Bold == chroma.Yes {
		out += "\033[1m"
	}
	if entry.Underline == chroma.Yes {
		out += "\033[4m"
	}
	if entry.Italic == chroma.Yes {
		out += "\033[3m"
	}
	if entry.Colour.IsSet() {
		out += table.foreground[findClosest(table, entry.Colour)]
	}
	if entry.Background.IsSet() {
		out += table.background[findClosest(table, entry.Background)]
	}
	return out
}

func findClosest(table *ttyTable, seeking chroma.Colour) chroma.Colour {
	closestColour := chroma.Colour(0)
	closest := float64(math.MaxFloat64)
	for colour := range table.foreground {
		distance := colour.Distance(seeking)
		if distance < closest {
			closest = distance
			closestColour = colour
		}
	}
	return closestColour
}

func styleToEscapeSequence(table *ttyTable, style *chroma.Style) map[chroma.TokenType]string {
	style = clearBackground(style)
	out := map[chroma.TokenType]string{}
	for _, ttype := range style.Types() {
		entry := style.Get(ttype)
		out[ttype] = entryToEscapeSequence(table, entry)
	}
	return out
}

// Clear the background colour.
func clearBackground(style *chroma.Style) *chroma.Style {
	builder := style.Builder()
	bg := builder.Get(chroma.Background)
	bg.Background = 0
	bg.NoInherit = true
	builder.AddEntry(chroma.Background, bg)
	style, _ = builder.Build()
	return style
}

type indexedTTYFormatter struct {
	table *ttyTable
}

func (c *indexedTTYFormatter) Format(w io.Writer, style *chroma.Style, it chroma.Iterator) (err error) {
	theme := styleToEscapeSequence(c.table, style)
	for token := it(); token != chroma.EOF; token = it() {
		clr, ok := theme[token.Type]
		if !ok {
			clr, ok = theme[token.Type.SubCategory()]
			if !ok {
				clr = theme[token.Type.Category()]
			}
		}
		if clr != "" {
			fmt.Fprint(w, clr)
		}
		fmt.Fprint(w, token.Value)
		if clr != "" {
			fmt.Fprintf(w, "\033[0m")
		}
	}
	return nil
}

// TTY is an 8-colour terminal formatter.
//
// The Lab colour space is used to map RGB values to the most appropriate index colour.
var TTY = Register("terminal", &indexedTTYFormatter{ttyTables[8]})

// TTY8 is an 8-colour terminal formatter.
//
// The Lab colour space is used to map RGB values to the most appropriate index colour.
var TTY8 = Register("terminal8", &indexedTTYFormatter{ttyTables[8]})

// TTY16 is a 16-colour terminal formatter.
//
// It uses \033[3xm for normal colours and \033[90Xm for bright colours.
//
// The Lab colour space is used to map RGB values to the most appropriate index colour.
var TTY16 = Register("terminal16", &indexedTTYFormatter{ttyTables[16]})

// TTY256 is a 256-colour terminal formatter.
//
// The Lab colour space is used to map RGB values to the most appropriate index colour.
var TTY256 = Register("terminal256", &indexedTTYFormatter{ttyTables[256]})
