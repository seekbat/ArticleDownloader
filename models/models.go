package models

/*
 * ==========================================
 *  Title:
 *  Project:   ArticleDownloader
 * Author:    Olivier Eggimann
 * Date:      03.Mai.2019
 * ==========================================
 */

type Site struct {
	ID        int    `json:"_id"`       //ID for Internal Use
	Name      string `json:"name"`      //Name of the Site
	URL       string `json:"url"`       //to the Site
	LinkRegex string `json:"linkregex"` //to find te Links
	IDRegex   string `json:"idregex"`   //to extract the ID
}

type LinkList struct {
	SiteName string        `json:"site_name"`
	Links    []ArticleLink `json:"links"`
}
type ArticleLink struct {
	ArticleId int    `json:"article_id" bson:"_id"`
	URL       string `json:"url"        bson:"url"`
	Timestamp int64  `json:"timestamp"  bson:"timestamp"`
}
