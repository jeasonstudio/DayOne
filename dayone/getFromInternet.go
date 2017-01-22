package dayone

import (
	"io/ioutil"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

// ERRMSG err
const ERRMSG = "未知错误"

// Spy 抓取每日一句
func Spy() (string, error) {
	resp, err := http.Get("http://sentence.iciba.com/index.php?c=dailysentence&m=getTodaySentence")
	if err != nil {
		return ERRMSG, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ERRMSG, err
	}

	// fmt.Println(string(body))

	return makeHTML(contents)
}

func makeHTML(contents []byte) (string, error) {
	json, jsErr := simplejson.NewJson(contents)
	if jsErr != nil {
		return ERRMSG, jsErr
	}

	pic, _ := json.Get("picture").String()
	yw, _ := json.Get("content").String()
	fy, _ := json.Get("note").String()
	pl, _ := json.Get("translation").String()

	tagHTML := `<!DOCTYPE html>
<html lang="en">

<head>
    <title>每日一句</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
    <style>
        .yw {
            font-size: 24px;
            font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif
        }
        
        .fy {
            font-size: 18px;
            font-family: monospace, '微软雅黑';
        }
        
        .pl {
            font-size: 14px;
            font-family: monospace, '微软雅黑';
            color: #aaa;
        }
    </style>
    <img alt="" style="width: 100%" src="` + pic + `">
    <div class="content">
        <div class="yw">` + yw + `</div>
        <div class="fy">` + fy + `</div>
        <div class="pl">` + pl + `</div>
    </div>
</body>

</html>`

	return tagHTML, nil
}
