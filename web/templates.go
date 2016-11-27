// Package web is generated with ftmpl {{{v0.3.1}}}, do not edit!!!! */
package web

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// TMPLERRbase evaluates a template base.tmpl
func TMPLERRbase(title string, loggedIn bool, unread int) (string, error) {
	_template := "base.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
    `)
	_w(`/* http://meyerweb.com/eric/tools/css/reset/
   v2.0 | 20110126
   License: none (public domain)
*/

html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
b, u, i, center,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td,
article, aside, canvas, details, embed,
figure, figcaption, footer, header, hgroup,
menu, nav, output, ruby, section, summary,
time, mark, audio, video {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 100%;
	font: inherit;
	vertical-align: baseline;
}
/* HTML5 display-role reset for older browsers */
article, aside, details, figcaption, figure,
footer, header, hgroup, menu, nav, section {
	display: block;
}
body {
	line-height: 1;
}
ol, ul {
	list-style: none;
}
blockquote, q {
	quotes: none;
}
blockquote:before, blockquote:after,
q:before, q:after {
	content: '';
	content: none;
}
table {
	border-collapse: collapse;
	border-spacing: 0;
}
`)
	_w(`
    `)
	_w(`body {
  font-family: Sans-Serif;
}

ul {
	list-style-type: none;
	margin: 0;
	padding: 0;
	overflow: hidden;
	background-color: #333;
}

li {
	float: left;
}

li a {
	display: block;
	color: white;
	text-align: center;
	padding: 14px 16px;
	text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
	background-color: #111;
}
`)
	_w(`
    </style>

</head>
<body>
	<ul id="menu">
		<li><a href="/">Kiasu</a></li>
`)
	if loggedIn {
		_w(`    <li>
        <a href="/feeds">Feeds `)
		_w(fmt.Sprintf(`%d`, unread))
		_w(`</a>
    </li>
`)
	} else {
		_w(`    <li><a href="/login">Login</a></li>
    <li><a href="/register">Register</a></li>
`)
	}
	_w(`	</ul>

`)
	_w(`
`)
	_w(`</body>
</html>
`)

	return _ftmpl.String(), nil
}

// TMPLbase evaluates a template base.tmpl
func TMPLbase(title string, loggedIn bool, unread int) string {
	html, err := TMPLERRbase(title, loggedIn, unread)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template base.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRbaseStyle evaluates a template baseStyle.tmpl
func TMPLERRbaseStyle() (string, error) {
	_template := "baseStyle.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`body {
  font-family: Sans-Serif;
}

ul {
	list-style-type: none;
	margin: 0;
	padding: 0;
	overflow: hidden;
	background-color: #333;
}

li {
	float: left;
}

li a {
	display: block;
	color: white;
	text-align: center;
	padding: 14px 16px;
	text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
	background-color: #111;
}
`)

	return _ftmpl.String(), nil
}

// TMPLbaseStyle evaluates a template baseStyle.tmpl
func TMPLbaseStyle() string {
	html, err := TMPLERRbaseStyle()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template baseStyle.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRhome evaluates a template home.tmpl
func TMPLERRhome(title string, loggedIn bool, unread int) (string, error) {
	_template := "home.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`
`)
	_w(`
<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
    `)
	_w(`/* http://meyerweb.com/eric/tools/css/reset/
   v2.0 | 20110126
   License: none (public domain)
*/

html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
b, u, i, center,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td,
article, aside, canvas, details, embed,
figure, figcaption, footer, header, hgroup,
menu, nav, output, ruby, section, summary,
time, mark, audio, video {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 100%;
	font: inherit;
	vertical-align: baseline;
}
/* HTML5 display-role reset for older browsers */
article, aside, details, figcaption, figure,
footer, header, hgroup, menu, nav, section {
	display: block;
}
body {
	line-height: 1;
}
ol, ul {
	list-style: none;
}
blockquote, q {
	quotes: none;
}
blockquote:before, blockquote:after,
q:before, q:after {
	content: '';
	content: none;
}
table {
	border-collapse: collapse;
	border-spacing: 0;
}
`)
	_w(`
    `)
	_w(`body {
  font-family: Sans-Serif;
}

ul {
	list-style-type: none;
	margin: 0;
	padding: 0;
	overflow: hidden;
	background-color: #333;
}

li {
	float: left;
}

li a {
	display: block;
	color: white;
	text-align: center;
	padding: 14px 16px;
	text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
	background-color: #111;
}
`)
	_w(`
    </style>

</head>
<body>
	<ul id="menu">
		<li><a href="/">Kiasu</a></li>
`)
	if loggedIn {
		_w(`    <li>
        <a href="/feeds">Feeds `)
		_w(fmt.Sprintf(`%d`, unread))
		_w(`</a>
    </li>
`)
	} else {
		_w(`    <li><a href="/login">Login</a></li>
    <li><a href="/register">Register</a></li>
`)
	}
	_w(`	</ul>

`)
	_w(`
<h1>Body!</h1>

`)
	_w(`
`)
	_w(`
<p> hi im footer</p>
`)
	_w(`</body>
</html>
`)

	return _ftmpl.String(), nil
}

// TMPLhome evaluates a template home.tmpl
func TMPLhome(title string, loggedIn bool, unread int) string {
	html, err := TMPLERRhome(title, loggedIn, unread)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template home.tmpl:" + err.Error())
	}
	return html
}

// TMPLERRresetStyle evaluates a template resetStyle.tmpl
func TMPLERRresetStyle() (string, error) {
	_template := "resetStyle.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`/* http://meyerweb.com/eric/tools/css/reset/
   v2.0 | 20110126
   License: none (public domain)
*/

html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
b, u, i, center,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td,
article, aside, canvas, details, embed,
figure, figcaption, footer, header, hgroup,
menu, nav, output, ruby, section, summary,
time, mark, audio, video {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 100%;
	font: inherit;
	vertical-align: baseline;
}
/* HTML5 display-role reset for older browsers */
article, aside, details, figcaption, figure,
footer, header, hgroup, menu, nav, section {
	display: block;
}
body {
	line-height: 1;
}
ol, ul {
	list-style: none;
}
blockquote, q {
	quotes: none;
}
blockquote:before, blockquote:after,
q:before, q:after {
	content: '';
	content: none;
}
table {
	border-collapse: collapse;
	border-spacing: 0;
}
`)

	return _ftmpl.String(), nil
}

// TMPLresetStyle evaluates a template resetStyle.tmpl
func TMPLresetStyle() string {
	html, err := TMPLERRresetStyle()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template resetStyle.tmpl:" + err.Error())
	}
	return html
}
