package routes

import (
	"dinifarb/codewars_readme_stats/stats"
	"fmt"
	"net/http"
	"strings"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	view := stats.GetStats()

	var rows strings.Builder
	for _, f := range view.Features {
		rows.WriteString(fmt.Sprintf("| %-30s | %6d |\n", f.Name, f.Count))
	}

	separator := "+" + strings.Repeat("-", 32) + "+" + strings.Repeat("-", 8) + "+"
	header := fmt.Sprintf("| %-30s | %6s |", "FEATURE", "COUNT")

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Codewars Stats - Usage</title>
<style>
  * { margin: 0; padding: 0; box-sizing: border-box; }
  body {
    background: #000;
    color: #fff;
    font-family: "Courier New", Courier, monospace;
    font-size: 14px;
    padding: 40px;
    line-height: 1.6;
  }
  .container { max-width: 600px; margin: 0 auto; }
  pre {
    white-space: pre-wrap;
    font-family: inherit;
  }
</style>
</head>
<body>
<div class="container">
<pre>
total_requests  %d
unique_users    %d

%s
%s
%s
%s
%s
</pre>
</div>
</body>
</html>`, view.TotalRequests, view.UniqueUserCount, separator, header, separator, rows.String(), separator)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
