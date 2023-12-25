import {
	WS_CONNECTION,
	constructMessage,
	UI_THEME,
	setCookie
} from "../index.js"
import {
	menubar_text
} from "./menubar_text.js"
import {
	menubar_img
} from "./menubar_img.js"
import { 
	playButton 
} from "./playButton.js"

export class menubar extends React.Component {
	constructor(props) {
		super(props)	
	}
	render() {
		return React.createElement("div", {
				id : "menubar",
				style : {
					display : "flex",
					flexDirection : "row",
				}
			},
			React.createElement("div", {
					style: {
						width : "44vmax",
						borderRadius: ".2vmax",
						border: "solid " + (this.props.loginData.uiTheme == UI_THEME.LIGHT ? "black " : "white ") + "1px",
						justifyContent: "space-evenly",
						display: "flex",
						height : "1.6vmax",
						padding: ".4vmax",
						letterSpacing: ".1vmax",
						flexDirection: "row",
						backgroundColor: this.props.loginData.uiTheme == UI_THEME.LIGHT ? "gainsboro" : "#555555",
					},
				},
				React.createElement(menubar_text, {
						name : "home",
						content : this.props.content,
						clickEvent : () => this.props.setStateRoot({content : "home"}),
						uiTheme : this.props.loginData.uiTheme,
					},
				),
				React.createElement(menubar_text, {
						name : "lobby",
						content : this.props.content,
						clickEvent : () => {
							this.props.setStateRoot({content : "lobby"})
						},
						uiTheme : this.props.loginData.uiTheme,
					},
				),
				React.createElement(menubar_text, {
						name : "profile",
						content : this.props.content,
						clickEvent : () => this.props.setStateRoot({content : "profile"}),
						uiTheme : this.props.loginData.uiTheme,
					},
				),
				React.createElement(menubar_text, {
						name : "leaderboard",
						content : this.props.content,
						clickEvent : () => WS_CONNECTION.send(constructMessage("leaderboard", )),
						uiTheme : this.props.loginData.uiTheme,
					},
				),
				React.createElement(menubar_text, {
						name : "editor",
						content : this.props.content,
						clickEvent : () => this.props.setStateRoot({content : "editor"}),
						uiTheme : this.props.loginData.uiTheme,
					},
				),
				React.createElement(menubar_text, {
						disabled : !this.props.game,
						name : "game",
						content : this.props.content,
						clickEvent : () => this.props.setStateRoot({content : "game"}),
						uiTheme : this.props.loginData.uiTheme,
					},
				),
			),

			React.createElement(playButton, {
				uiTheme : this.props.loginData.uiTheme,
			},),
			
			React.createElement("div", {
					style: {
						width : "44vmax",
						borderRadius: ".2vmax",
						border: "solid " + (this.props.loginData.uiTheme == UI_THEME.LIGHT ? "black " : "white ") + "1px",
						justifyContent: "space-evenly",
						display: "flex",
						height : "1.6vmax",
						padding: ".4vmax",
						letterSpacing: ".07vmax",
						flexDirection: "row",
						backgroundColor: this.props.loginData.uiTheme == UI_THEME.LIGHT ? "gainsboro" : "#555555",
					},
				},
				React.createElement("div", {
						style: {
							display: "flex",
							flexDirection: "row",
							gap : "1vmax"
						}
					},
					React.createElement("div", {
							style: {
								display: "flex",
								flexDirection: "row",
								alignItems: "center",
								justifyContent: "center",
								padding: "0.1vmax",
							}
						},
						React.createElement("img", {
								src : this.props.loginData.uiTheme == UI_THEME.LIGHT ? "https://nodge.s3.eu-central-1.amazonaws.com/user_b.svg" : "https://nodge.s3.eu-central-1.amazonaws.com/user_w.svg",
								onClick :() => {
									WS_CONNECTION.send(constructMessage("profile", this.props.loginData.username))
								},
								style: {
									cursor: "pointer",
									width: "1.3vmax",
									height: "1.3vmax",
								}
							}, 
						),
					),

					React.createElement("div", {
							style: {
								display: "flex",
								alignItems : "center",
								justifyContent : "right",
							}
						},
						React.createElement("b", {
								id : "menubarName",
								style : {
									fontSize : "1.3vmax",
									cursor: "pointer",
								},
							}, 
							this.props.loginData.username,
						)
					),

					this.props.loginData.userPower > 0 ? React.createElement(menubar_img, {
						src : this.props.loginData.uiTheme == UI_THEME.LIGHT ? "https://nodge.s3.eu-central-1.amazonaws.com/logout_svg.svg" : "https://nodge.s3.eu-central-1.amazonaws.com/logout_white.svg.svg",
						clickEvent : () => {
							if(confirm("Are you sure you want to logout?")) {
								setCookie("username", "", 1)
								setCookie("password", "", 1)
								location.reload()
							}
						},
					}, ) : null,
				),

				this.props.loginData.userPower === 0 ? React.createElement(menubar_text, {
						name : "login",
						content : this.props.content,
						clickEvent : () => this.props.setStateRoot({content : "login"}),
						uiTheme : this.props.loginData.uiTheme,
					},
				):"" ,

				this.props.loginData.userPower === 0 ? React.createElement(menubar_text, {
						name : "register",
						content : this.props.content,
						clickEvent : () => this.props.setStateRoot({content : "register"}),
						uiTheme : this.props.loginData.uiTheme,
					},
				):"" ,

				
				React.createElement(menubar_img, {
					src : this.props.loginData.uiTheme == UI_THEME.LIGHT ? this.props.loginData.showCommunicator ? "https://nodge.s3.eu-central-1.amazonaws.com/com-white.png" : "https://nodge.s3.eu-central-1.amazonaws.com/com-black.png" : this.props.loginData.showCommunicator ? "https://nodge.s3.eu-central-1.amazonaws.com/com-black.png" : "https://nodge.s3.eu-central-1.amazonaws.com/com-white.png",
					backgroundColor : this.props.loginData.uiTheme == UI_THEME.LIGHT ? this.props.loginData.showCommunicator ? "#5d5d5d" : "" : this.props.loginData.showCommunicator ? "lightGrey" : "",
					clickEvent : () => {
						WS_CONNECTION.send(constructMessage("showCommunicator", this.props.loginData.showCommunicator ? "false" : "true"))
					}
				}),
				React.createElement(menubar_img, {
					src : this.props.loginData.uiTheme == UI_THEME.LIGHT ? "https://nodge.s3.eu-central-1.amazonaws.com/sun-b.png" : "https://nodge.s3.eu-central-1.amazonaws.com/moon-white.svg",
					clickEvent : () => {
						WS_CONNECTION.send(constructMessage("uiTheme", this.props.loginData.uiTheme == UI_THEME.LIGHT ? UI_THEME.DARK : UI_THEME.LIGHT))
					}
				}),
				React.createElement(menubar_img, {
					src :  this.props.loginData.uiTheme == UI_THEME.LIGHT ? "https://nodge.s3.eu-central-1.amazonaws.com/discord.svg" : "https://nodge.s3.eu-central-1.amazonaws.com/discord-w.svg",
					clickEvent : () => {
						window.open("https://discord.gg/tghbWmdtgB")
					}
				}),
			),
		)
	}
}