package email

const (
	resetSubject = "Instructions for resetting your password."
)
const resetTextTmpl = `Hi there! 
It appears that you have requested a password reset. If this was you, please follow the link below to update your password: 
%s 
If you are asked for a token, please use the following value: 
%s 
If you didn't request a password reset you can safely ignore this email and your account will not be changed. 

Best, 
LensLocked Support 
`
const resetHTMLTmpl = `Hi there!<br/>
<br/> 
It appears that you have requested a password reset. If this was you, please follow the link below to update your password:<br/> 
<br/> 
<a href="%s">%s</a><br/> 
<br/> 
If you are asked for a token, please use the following value:<br/> 
<br/> 
%s<br/> 
<br/> 
If you didn't request a password reset you can safely ignore this email and your account will not be changed.<br/> 
<br/> 
Best,<br/> 
LensLocked Support<br/> 
`

const resetBaseURL = "https://www.lenslocked. com/reset"

func (c *Client) ResetPw(toEmail, token string) error {
	v := url.Values{}
	v.Set("token", token)
	resetUrl := resetBaseURL + "?" + v.Encode()
	resetText := fmt.Sprintf(resetTextTmpl, resetUrl, token)
	message := mailgun.NewMessage(c.from, resetSubject, resetText, toEmail)
	resetHTML := fmt.Sprintf(resetHTMLTmpl, resetUrl, resetUrl, token)
	message.SetHtml(resetHTML)
	_, _, err := c.mg.Send(message)
	return err
}
