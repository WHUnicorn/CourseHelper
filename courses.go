package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testCourse/utils"
)

type resultSet struct {
	CurrentPage   int  `json:"currentPage"`
	CurrentResult int  `json:"currentResult"`
	EntityOrField bool `json:"entityOrField"`
	Items         []struct {
		Date               string `json:"date"`
		DateDigit          string `json:"dateDigit"`
		DateDigitSeparator string `json:"dateDigitSeparator"`
		Day                string `json:"day"`
		Jgpxzd             string `json:"jgpxzd"`
		Jsxm               string `json:"jsxm"`
		JxbId              string `json:"jxb_id"`
		Jxbmc              string `json:"jxbmc"`
		Jxdd               string `json:"jxdd,omitempty"`
		Kch                string `json:"kch"`
		Kcmc               string `json:"kcmc"`
		Kkxy               string `json:"kkxy"`
		Listnav            string `json:"listnav"`
		LocaleKey          string `json:"localeKey"`
		Month              string `json:"month"`
		PageTotal          int    `json:"pageTotal"`
		Pageable           bool   `json:"pageable"`
		QueryModel         struct {
			CurrentPage   int           `json:"currentPage"`
			CurrentResult int           `json:"currentResult"`
			EntityOrField bool          `json:"entityOrField"`
			Limit         int           `json:"limit"`
			Offset        int           `json:"offset"`
			PageNo        int           `json:"pageNo"`
			PageSize      int           `json:"pageSize"`
			ShowCount     int           `json:"showCount"`
			Sorts         []interface{} `json:"sorts"`
			TotalCount    int           `json:"totalCount"`
			TotalPage     int           `json:"totalPage"`
			TotalResult   int           `json:"totalResult"`
		} `json:"queryModel"`
		Rangeable   bool   `json:"rangeable"`
		RowId       string `json:"row_id"`
		Sksj        string `json:"sksj,omitempty"`
		TotalResult string `json:"totalResult"`
		UserModel   struct {
			Monitor    bool   `json:"monitor"`
			RoleCount  int    `json:"roleCount"`
			RoleKeys   string `json:"roleKeys"`
			RoleValues string `json:"roleValues"`
			Status     int    `json:"status"`
			Usable     bool   `json:"usable"`
		} `json:"userModel"`
		Xf     string `json:"xf"`
		Xnm    string `json:"xnm"`
		Xnmc   string `json:"xnmc"`
		Xqm    string `json:"xqm"`
		Xqmmc  string `json:"xqmmc"`
		Year   string `json:"year"`
		Kclbmc string `json:"kclbmc,omitempty"`
	} `json:"items"`
	Limit       int           `json:"limit"`
	Offset      int           `json:"offset"`
	PageNo      int           `json:"pageNo"`
	PageSize    int           `json:"pageSize"`
	ShowCount   int           `json:"showCount"`
	Sorts       []interface{} `json:"sorts"`
	TotalCount  int           `json:"totalCount"`
	TotalPage   int           `json:"totalPage"`
	TotalResult int           `json:"totalResult"`
}

type PersonalCourse struct {
	CourseIdentifier string `json:"courseIdentifier"`
	CourseName       string `json:"courseName"`
}

func getPersonalCourses(cookie string) (res []PersonalCourse) {
	urlToConn, _ := url.Parse("https://jwgl.whu.edu.cn/xsxxxggl/xsxxwh_cxXsxkxx.html")
	param := url.Values{}
	param.Set("gnmkdm", "N100801")
	param.Set("queryModel.currentPage", "1")
	param.Set("queryModel.showCount", "2000")
	urlToConn.RawQuery = param.Encode()
	// body 为空，无需关闭
	req, _ := http.NewRequest("GET", urlToConn.String(), nil)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "deflate")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/116.0.0.0 Mobile Safari/537.36 "+
		"Edg/116.0.1938.76")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Error(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	responseBody, _ := io.ReadAll(resp.Body)
	var reply resultSet
	err = json.Unmarshal(responseBody, &reply)
	if err != nil {
		utils.Error(err)
		return nil
	}

	res = make([]PersonalCourse, 0)
	for _, item := range reply.Items {
		res = append(res, PersonalCourse{
			CourseIdentifier: item.Kch,
			CourseName:       item.Kcmc,
		})
	}

	return
}
